/**
*@author: Rakshita Mathur
*@student number: 300215340
*@date: 03/11/2023
*@Course: CSI 2120 - programming paradigms
*@description: This program implements the RANSAC algorithm to find the best plane that fits a set of points.
*              The program takes as input a file containing the points in the format x y z, one point per line.
*              The program outputs the plane coefficients A, B, C, D and the number of points that are on the plane.
*              The program also outputs the time taken to run the algorithm.
*              The program takes as input the following parameters:
*              -i: input file name (default: PointCloud1) // do not include the extension
*              -c: confidence level (default: 0.99)
*              -p: percentage of points that are on the plane (default: 20)
*              -e: maximum distance between a point and the plane (default: 0.2)
*              Example run:  go run RANSAC.go PointCloud2 .99 20 .01
* 			  The program will run the RANSAC algorithm on the points in PointCloud2.xyz with a confidence level of 0.99, a percentage of points on the plane of 20% and a maximum distance of 0.01.
*
 */
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Point3D struct {
	X float64
	Y float64
	Z float64
}
type Plane3D struct {
	A float64
	B float64
	C float64
	D float64
}
type Plane3DwSupport struct {
	Plane3D
	SupportSize int
}

// Implement a function to read the XYZ file and return a slice of Point3D.
func ReadXYZ(filename string) []Point3D {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the data from the file
	scanner := bufio.NewScanner(file)
	var points []Point3D
	var skippedFirstLine bool
	for scanner.Scan() {
		line := scanner.Text()
		if !skippedFirstLine {
			skippedFirstLine = true
			continue
		}
		fields := strings.Fields(line)
		if len(fields) != 3 {
			continue
		}
		x, err := strconv.ParseFloat(fields[0], 64)
		if err != nil {
			continue
		}
		y, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			continue
		}
		z, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			continue
		}
		points = append(points, Point3D{x, y, z})
	}
	return points
}

// Implement a function to save a slice of Point3D into an XYZ file.
func SaveXYZ(filename string, points []Point3D) {
	// Open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write each point to the file in XYZ format
	for _, point := range points {
		_, err = fmt.Fprintf(file, "%.6f %.6f %.6f\n", point.X, point.Y, point.Z)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Implement a function to compute the distance between two points in 3D space.
func (p1 *Point3D) GetDistance(p2 *Point3D) float64 {
	// compute the distance
	return math.Sqrt(math.Pow(p1.X-p2.X, 2) + math.Pow(p1.Y-p2.Y, 2) + math.Pow(p1.Z-p2.Z, 2))
}

// Implement a function to compute the plane defined by a set of three points.
// The plane is defined by the equation Ax + By + Cz + D = 0.
func GetPlane(points []Point3D) Plane3D {
	if len(points) != 3 {
		return Plane3D{}
	}
	// Compute the normal vector of the plane
	v1 := Point3D{X: points[1].X - points[0].X, Y: points[1].Y - points[0].Y, Z: points[1].Z - points[0].Z}
	v2 := Point3D{X: points[2].X - points[0].X, Y: points[2].Y - points[0].Y, Z: points[2].Z - points[0].Z}
	n := Point3D{X: v1.Y*v2.Z - v1.Z*v2.Y, Y: v1.Z*v2.X - v1.X*v2.Z, Z: v1.X*v2.Y - v1.Y*v2.X}
	norm := math.Sqrt(n.X*n.X + n.Y*n.Y + n.Z*n.Z)
	if norm == 0 {
		return Plane3D{}
	}
	n.X /= norm
	n.Y /= norm
	n.Z /= norm

	// Compute the plane coefficients
	plane := Plane3D{A: n.X, B: n.Y, C: n.Z}
	plane.D = -1 * (plane.A*points[0].X + plane.B*points[0].Y + plane.C*points[0].Z)

	return plane
}

// Implement a function to compute the number of required RANSAC iterations.
func GetNumberOfIterations(confidence float64, percentageOfPointsOnPlane float64) float64 {
	percentageOfPointsOnPlane = percentageOfPointsOnPlane / 100.0
	numberOfIterations := math.Ceil((math.Log(1-confidence) / math.Log(2)) / (math.Log(1-math.Pow(percentageOfPointsOnPlane, 3)) / math.Log(2)))
	return numberOfIterations
}

// Implement a function to compute the support of a plane in a set of points
func GetSupport(plane Plane3D, points []Point3D, eps float64) Plane3DwSupport {
	var supportSize int
	for _, point := range points {
		dist := point.GetDistance(&Point3D{X: -plane.D / plane.A, Y: -plane.D / plane.B, Z: -plane.D / plane.C})
		if dist <= eps {
			supportSize++
		}
	}
	return Plane3DwSupport{Plane3D: plane, SupportSize: supportSize}
}

// Implement a function to extract the points that support a given plane.
// Implement a function to extract the points that support a given plane.

func GetSupportingPoints(plane Plane3D, points []Point3D, eps float64) []Point3D {
	var supportingPoints []Point3D
	for _, point := range points {
		dist := math.Abs(plane.A*point.X+plane.B*point.Y+plane.C*point.Z+plane.D) / math.Sqrt(plane.A*plane.A+plane.B*plane.B+plane.C*plane.C)
		if dist <= eps {
			supportingPoints = append(supportingPoints, point)
		}
	}
	return supportingPoints
}

// Implement a function to create a new slice of points in which all points belonging to the plane have been removed.
func RemovePoints(points []Point3D, plane Plane3D, threshold float64) []Point3D {
	var newPoints []Point3D
	for _, point := range points {
		if math.Abs(plane.A*point.X+plane.B*point.Y+plane.C*point.Z+plane.D) > threshold {
			newPoints = append(newPoints, point)
		}
	}
	return newPoints
}

// Implement the pipeline, consisting of the following components:
// 1.Random point generator: randomly selects a point from the provided slice of Point3D.
func RandomPointGenerator(points []Point3D) <-chan Point3D {
	out := make(chan Point3D)
	go func() {
		for {
			out <- points[rand.Intn(len(points))]
		}
	}()
	return out
}

// 2.Triplet of points generator: reads Point3D instances and accumulates 3 points, then sends them as an array of Point3D.

func TripletOfPointsGenerator(in <-chan Point3D) <-chan []Point3D {
	out := make(chan []Point3D)
	go func() {
		var points []Point3D
		for {
			points = append(points, <-in)
			if len(points) == 3 {
				out <- points
				points = nil
			}
		}
	}()
	return out
}

// 3.TakeN: reads arrays of Point3D and resends them. It stops the pipeline after having received N arrays. Where N is the number of iterations float64.

func TakeN(in <-chan []Point3D, numberOfIterations float64) <-chan []Point3D {
	out := make(chan []Point3D)
	go func() {
		for i := 0; i < int(numberOfIterations); i++ {
			out <- <-in
		}
	}()
	return out
}

// 4.Plane estimator: reads arrays of three Point3D and computes the plane defined by these points. Using the function GetPlane.

func PlaneEstimator(in <-chan []Point3D) <-chan Plane3D {
	out := make(chan Plane3D)
	go func() {
		for {
			points := <-in
			out <- GetPlane(points)
		}
	}()
	return out
}

//5. Supporting point finder: counts the number of points in the provided slice of Point3D that supports the received 3D plane.

func SupportingPointFinder(in <-chan Plane3D, points []Point3D, threshold float64) <-chan Plane3DwSupport {
	out := make(chan Plane3DwSupport)
	go func() {
		for {
			plane := <-in
			support := GetSupport(plane, points, threshold)
			// out <- Plane3DwSupport{Plane3D, int}
			out <- Plane3DwSupport{plane, support.SupportSize}
		}
	}()
	return out
}

// 6.Fan In: multiplexes the results received from multiple channels into one output channel.

func FanIn(ins ...<-chan Plane3DwSupport) <-chan Plane3DwSupport {
	out := make(chan Plane3DwSupport)
	for _, in := range ins {
		go func(in <-chan Plane3DwSupport) {
			for {
				out <- <-in
			}
		}(in)
	}
	return out
}

// 7.Dominant plane identifier: receives Plane3DwSupport instances and keeps in memory the plane with the best support received so far.

func DominantPlaneIdentifier(in <-chan Plane3DwSupport) <-chan Plane3DwSupport {
	out := make(chan Plane3DwSupport)
	go func() {
		var best Plane3DwSupport
		for {
			plane := <-in
			if plane.SupportSize > best.SupportSize {
				best = plane
			}
			out <- best
		}
	}()
	return out
}

//In the main function, read the XYZ file, find the number of iterations required based on the specified confidence and percentage, create and start the RANSAC find dominant plane pipeline, save the supporting points of the identified dominant plane in a file named by appending _p to the input filename, and save the original point cloud without the supporting points of the dominant plane in a file named by appending _p0 to the input filename.

func main() {

	// read the input arguments
	if len(os.Args) != 5 {
		log.Fatal("Usage: RANSAC <input file> <confidence> <percentage> <eps>")
	}
	inputFilename := os.Args[1]
	confidence, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		log.Fatal(err)
	}
	percentage, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		log.Fatal(err)
	}
	eps, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		log.Fatal(err)
	}

	// read the input file
	points := ReadXYZ(inputFilename)

	// find the number of iterations required based on the specified confidence and percentage
	iterations := GetNumberOfIterations(confidence, percentage)

	randomPointGenerator := RandomPointGenerator(points)

	tripletOfPointsGenerator := TripletOfPointsGenerator(randomPointGenerator)

	takeN := TakeN(tripletOfPointsGenerator, iterations)

	planeEstimator := PlaneEstimator(takeN)

	supportingPointFinder := SupportingPointFinder(planeEstimator, points, eps)

	fanIn := FanIn(supportingPointFinder)

	dominantPlaneIdentifier := DominantPlaneIdentifier(fanIn)

	plane := <-dominantPlaneIdentifier

	supportingPoints := GetSupport(plane.Plane3D, points, eps)
	// save the inline points of the identified dominant plane in a file named by appending _p to the input filename
	SaveXYZ(inputFilename+"_p", GetSupportingPoints(supportingPoints.Plane3D, points, eps))
	// save the original point cloud without the supporting points of the dominant plane in a file named by appending _p0 to the input filename
	SaveXYZ(inputFilename+"_p0", RemovePoints(points, plane.Plane3D, eps))
}
