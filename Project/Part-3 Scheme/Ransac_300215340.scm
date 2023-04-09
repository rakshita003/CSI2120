;lang scheme
; author: Rakshita Mathur
; Student ID: 300215340
; date: 31/03/2023
; Course: Csi 2120 Project-Part-3
; description: RANSAC algorithm to find the best fitting plane in a set of points

;Run the following commands to get the output attached in the zipfile
;(run-planeRANSAC "Point_Cloud_1_No_Road_Reduced.xyz" 0.99 0.5 0.8)
;(run-planeRANSAC "Point_Cloud_2_No_Road_Reduced.xyz" 0.99 0.5 0.8)
;(run-planeRANSAC "Point_Cloud_3_No_Road_Reduced.xyz" 0.99 0.5 0.8)



; Function used to run the RANSAC algorithm
(define (run-planeRANSAC filepath confidence percentage eps) 
  (let* ((full-path (string-append "" filepath))
         (result (planeRANSAC full-path confidence percentage eps))
         (best-plane (car result))
         (best-support (cdr result))
         (inliers (cdr (support best-plane (readXYZ full-path) eps)))
         (filename (string-split full-path "/"))
         (output-path (string-append (list-ref filename (- (length filename) 1))
                                     "-result.xyz")))
    (with-output-to-file output-path
      (lambda ()
        (write "x y z")(newline)
        (for-each (lambda (p)
                    (write (car p))
                    (write-char #\space)
                    (write (cadr p))
                    (write-char #\space)
                    (write (caddr p)) (write-char #\newline))
                  inliers)))
    result))

; Function used to write the output to a file
(define (write-lines lst port) 
  (for-each (lambda (x)
              (write (first x) port)
              (write-char #\space port)
              (write (second x) port)
              (write-char #\space port)
              (write (third x) port)
              (write-char #\newline port))
            lst))

; Function used to read the file and convert the string values to numbers
(define (readXYZ fileIn)
  (let ((sL (map (lambda (s)
                    (map (lambda (x)
                           (if (eqv? (string->number x) #f)
                               x
                               (string->number x)))
                         (string-split s))))
                 (cdr (file->lines fileIn)))))
    sL)

; Function used to pick three random points from the set of points
(define (pick-random-points Ps) 
  (list (list-ref Ps (random (length Ps)))
        (list-ref Ps (random (length Ps)))
        (list-ref Ps (random (length Ps)))))

; Function used to calculate the plane equation given three points
(define (plane P1 P2 P3) 
  (let ((v1 (list (- (car P2) (car P1))
                  (- (cadr P2) (cadr P1))
                  (- (caddr P2) (caddr P1))))
        (v2 (list (- (car P3) (car P1))
                  (- (cadr P3) (cadr P1))
                  (- (caddr P3) (caddr P1)))))
    (let ((a (- (* (cadr v1) (caddr v2))
                (* (caddr v1) (cadr v2))))
          (b (- (* (caddr v1) (car v2))
                (* (car v1) (caddr v2))))
          (c (- (* (car v1) (cadr v2))
                (* (cadr v1) (car v2)))))
      (let ((d (+ (* a (car P1))
                  (* b (cadr P1))
                  (* c (caddr P1)))))
        (list a b c d)))))

; Function used to calculate the support of a plane
(define (support plane points eps)
  (let ((count 0)
        (inliers '()))
    (let loop ((ps points))
      (cond ((null? ps)
             (cons count inliers))
            (else
             (let ((p (car ps))
                   (dist (distance-to-plane plane p)))
               (if (<= dist eps)
                   (begin
                     (set! count (+ count 1))
                     (set! inliers (cons p inliers))))
               (loop (cdr ps))))))))

; Function used to calculate the distance of a point from a plane
(define (distance-to-plane plane point)
  (let ((a (car plane))
        (b (cadr plane))
        (c (caddr plane))
        (d (cadddr plane))
        (x (car point))
        (y (cadr point))
        (z (caddr point)))
    (/ (abs (+ (* a x) (* b y) (* c z) d))
       (sqrt (+ (* a a) (* b b) (* c c))))))


; Function used to calculate the dot product of two vectors
(define (dot-product v1 v2) 
  (apply + (map * v1 v2)))

; Function used to calculate the cross product of two vectors
(define (ransacNumberOfIteration confidence percentage)
  (let ((n (length confidence)))
    (ceiling (/ (* (log (- 1 confidence))
                   (log (- 1 (expt percentage n))))
                (log (- 1 (expt (expt percentage n) confidence)))))))

; Function used to calculate the cross product of two vectors
(define (dominantPlane Ps k) 
  (let ((max-support 0)
        (dominant-plane '()))
    (do ((i 0 (+ i 1)))
        ((>= i k) dominant-plane)
      (let ((random-points (pick-random-points Ps))
            (current-plane '()))
        (set! current-plane (plane (first random-points)
                                   (second random-points)
                                   (third random-points)))
        (let ((support-count (support current-plane Ps)))
          (when (> (car support-count) max-support)
            (set! max-support (car support-count))
            (set! dominant-plane (cdr support-count))))))))


; Function used to calculate the cross product of two vectors
(define (planeRANSAC filepath confidence percentage eps)    
  (let ((Ps (readXYZ filepath))
        (k (ransacNumberOfIteration confidence percentage))
        (n (length (readXYZ filepath)))
        (t (* eps eps)))
    (let loop ((i 0) (best-support 0) (best-plane '()))
      (cond ((>= i k) (list best-plane best-support))
            (else
             (let* ((p1 (list-ref Ps (random n)))
                    (p2 (list-ref Ps (random n)))
                    (p3 (list-ref Ps (random n)))
                    (current-plane (plane p1 p2 p3))
                    (support-count (support current-plane Ps t)))
               (if (> (car support-count) best-support)
                   (loop (+ i 1) (car support-count) (cdr support-count))
                   (loop (+ i 1) best-support best-plane)))))))
  )