/** Author    : Rakshita Mathur
 *  StudentID : 300215340
 *  Course    : CSI 2120 Programming Paradigms
 *  Assignment 1 Question 1
 */

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

// MidPoint function calculates the midpoint and distance between two points
// and returns them as separate values.
func MidPoint(p1 Point, p2 Point) (midPoint Point, length float64) {
	midPoint = Point{(p1.x + p2.x) / 2, (p1.y + p2.y) / 2}
	length = math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
	return midPoint, length
}

func main() {
	// create a channel of type Point
	ch := make(chan Point)
	// create an array of points
	points := []Point{{8., 1.}, {3., 2.}, {7., 4.}, {6., 3.}}
	// create a goroutine that will call MidPoint for all combinations of two points
	// in the array of points
	go func() {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				// Calculate the midpoint and length between two points
				midPoint, length := MidPoint(points[i], points[j])
				// Print the points, midpoint, and length in the desired format
				fmt.Printf("Points: (%.0f, %.0f) (%.0f, %.0f)\nMidPoint= (%.1f, %.1f)\nLength= %.2f\n\n",
					points[i].x, points[i].y, points[j].x, points[j].y, midPoint.x, midPoint.y, length)
				// Send the midpoint to the channel
				ch <- midPoint
			}
		}
		// Close the channel after all midpoints are sent
		close(ch)
	}()

	// Read from the channel until it is closed
	for range ch {
	}
}
