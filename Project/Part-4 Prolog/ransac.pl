% File: ransac.pl
% Author: Rakshita Mathur
% Student ID: 300215340
% Date: 2023-04-10
% Description: Predicates for implementing the RANSAC (RANdom SAmple Consensus) algorithm in Prolog.



% Predicate to randomly select 3 points from a list of points
% random3points(+Points, -Point3)
% This predicate is true if Point3 is a triplet of points randomly selected from the list of points Points.
% The triplet of points is of the form [[x1,y1,z1], [x2,y2,z2], [x3,y3,z3]].
random3points(Points, Point3) :-
    length(Points, N),
    N >= 3, % Ensure that the list has at least 3 points
    random_permutation(Points, RandomPoints), % Randomly permute the list of points
    take(RandomPoints, 3, Point3). % Take the first 3 points from the permuted list

% Helper predicate to take first N elements from a list
take(_, 0, []).
take([X|Xs], N, [X|Ys]) :-
    N > 0,
    N1 is N - 1,
    take(Xs, N1, Ys).

% Predicate to compute the coefficients of a plane equation given 3 points
% plane(+Point3, -Plane)
% This predicate is true if Plane is the equation of the plane defined by the three points of the list Point3.
% The plane is specified by the list [a, b, c, d] from the equation ax + by + cz = d.
% The list of points is of the form [[x1,y1,z1], [x2,y2,z2], [x3,y3,z3]].
plane(Point3, Plane) :-
    length(Point3, 3), % Ensure that Point3 contains exactly 3 points
    Point3 = [[X1, Y1, Z1], [X2, Y2, Z2], [X3, Y3, Z3]], % Extract the coordinates of the three points
    % Compute the coefficients of the plane equation
    A is (Y2 - Y1) * (Z3 - Z1) - (Z2 - Z1) * (Y3 - Y1),
    B is (Z2 - Z1) * (X3 - X1) - (X2 - X1) * (Z3 - Z1),
    C is (X2 - X1) * (Y3 - Y1) - (Y2 - Y1) * (X3 - X1),
    D is -(A * X1 + B * Y1 + C * Z1),
    Plane = [A, B, C, D]. % Construct the plane equation as a list [a, b, c, d]

% Predicate to find the support points of a plane given a list of points and an epsilon value
% support(+Plane, +Points, +Eps, -N)
% This predicate is true if the support of plane Plane is composed of N points from the list of points Point3
% when the distance Eps is used.
support(Plane, Points, Eps, N) :-
    length(Points, NumPoints), % Get the total number of points in the list
    NumPoints >= N, % Ensure that there are enough points to form a support of size N
    find_support_points(Plane, Points, Eps, SupportPoints), % Find the support points
    length(SupportPoints, N). % Check if the number of support points is N


% Helper predicate to find the support points for a given plane
find_support_points(_, [], _, []).
find_support_points(Plane, [Point|Points], Eps, SupportPoints) :-
    % Extract the coordinates of the current point
    Point = [X, Y, Z],
    % Compute the distance from the current point to the plane
    Dist is abs(Plane[0] * X + Plane[1] * Y + Plane[2] * Z + Plane[3]) / sqrt(Plane[0] * Plane[0] + Plane[1] * Plane[1] + Plane[2] * Plane[2]),
    % Check if the distance is within the given epsilon value
    Dist =< Eps,
    % Recursively find the support points in the remaining list of points
    find_support_points(Plane, Points, Eps, RemainingSupportPoints),
    % Construct the list of support points
    SupportPoints = [Point|RemainingSupportPoints].

% Predicate to calculate the number of iterations required by RANSAC with given confidence and percentage parameters
% ransac_number_of_iterations(+Confidence, +Percentage, -N)
% This predicate is true if N is the number of iterations required by RANSAC with parameters Confidence and Percentage
ransac_number_of_iterations(Confidence, Percentage, N) :-
    Confidence > 0, % Ensure that Confidence is a positive value
    Confidence < 1, % Ensure that Confidence is less than 1
    Percentage > 0, % Ensure that Percentage is a positive value
    Percentage < 1, % Ensure that Percentage is less than 1
    NumPoints is round(log(1 - Confidence) / log(1 - Percentage^3)), % Compute the number of iterations using the formula
    N is max(1, NumPoints). % Set N to be at least 1, as per the problem description

%test case 

% Test cases for random3points/2 predicate
test(random3points, 1) :-
    Points = [[0,0,0], [-5.1323336,-4.089636333,0.243960825], [-5.141415625,-4.020067234,0.242623445]],
    random3points(Points, Point3),
    length(Point3, 3),
    writeln('Randomly selected 3 points:'),
    writeln(Point3).

test(random3points, 2) :-
    Points = [[1,1,1], [2,2,2], [3,3,3], [4,4,4]],
    random3points(Points, Point3),
    length(Point3, 3),
    writeln('Randomly selected 3 points:'),
    writeln(Point3).

% Test cases for plane/2 predicate
test(plane, 1) :-
    Point3 = [[0,0,0], [1,1,1], [2,2,2]],
    plane(Point3, Plane),
    writeln('Plane equation:'),
    writeln(Plane).

test(plane, 2) :-
    Point3 = [[-1,-1,-1], [0,0,0], [1,1,1]],
    plane(Point3, Plane),
    writeln('Plane equation:'),
    writeln(Plane).

% Test cases for support/4 predicate
test(support, 1) :-
    Plane = [1,1,1,0],
    Points = [[0,0,0], [1,1,1], [2,2,2], [3,3,3]],
    Eps = 0.1,
    N = 3,
    support(Plane, Points, Eps, N),
    writeln('Support points for the plane:'),
    writeln(Points).

test(support, 2) :-
    Plane = [0,0,1,5],
    Points = [[0,0,5], [1,1,6], [2,2,7], [3,3,8]],
    Eps = 0.2,
    N = 2,
    support(Plane, Points, Eps, N),
    writeln('Support points for the plane:'),
    writeln(Points).

% Test cases for ransac_number_of_iterations/3 predicate
test(ransac_number_of_iterations, 1) :-
    Confidence = 0.99,
    Percentage = 0.3,
    ransac_number_of_iterations(Confidence, Percentage, N),
    writeln('Number of iterations for RANSAC:'),
    writeln(N).

test(ransac_number_of_iterations, 2) :-
    Confidence = 0.95,
    Percentage = 0.2,
    ransac_number_of_iterations(Confidence, Percentage, N),
    writeln('Number of iterations for RANSAC:'),
    writeln(N).

% Instruction to run the test cases
% ?- test(random3points, N).
% ?- test(plane, N).
% ?- test(support, N).
% ?- test(ransac_number_of_iterations, N).

