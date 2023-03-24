/** Author    : Rakshita Mathur
 *  StudentID : 300215340
 *  Course    : CSI 2120 Programming Paradigms
 *  Assignment 1 Question 2
 */

package main

import (
	"fmt"
	"sync"
)

// sort sorts the given array of float64s in ascending order
func sort(tab []float64) {
	for i := 0; i < len(tab); i++ {
		for j := i + 1; j < len(tab); j++ {
			if tab[i] > tab[j] {
				tab[i], tab[j] = tab[j], tab[i]
			}
		}
	}
}

// sortRows sorts the rows of the given 2D array of float64s in ascending order
func sortRows(tab [][]float64) {
	var wg sync.WaitGroup
	wg.Add(len(tab))
	for i := 0; i < len(tab); i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < len(tab[i]); j++ {
				for k := j + 1; k < len(tab[i]); k++ {
					if tab[i][j] > tab[i][k] {
						tab[i][j], tab[i][k] = tab[i][k], tab[i][j]
					}
				}
			}
		}(i)
	}
	wg.Wait()
}

// transpose transposes the given 2D array of float64s
func transpose(tab [][]float64) [][]float64 {
	var newTab [][]float64
	for i := 0; i < len(tab[0]); i++ {
		var row []float64
		for j := 0; j < len(tab); j++ {
			row = append(row, tab[j][i])
		}
		newTab = append(newTab, row)
	}
	return newTab
}

func main() {
	array := [][]float64{{1.1, 7.3, 3.2, 0.3, 3.1},
		{4.3, 5.6, 1.8, 5.3, 3.1},
		{1.3, 2.7, 3.5, 9.3, 1.1},
		{7.5, 5.1, 0.6, 2.3, 3.9}}

	// a. Print the input 2D array to console
	fmt.Println("Input array:")
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			fmt.Printf("%.1f ", array[i][j])
		}
		fmt.Println()
	}

	// b. Sort the rows of this array with sortRows
	sortRows(array)

	// c. Transpose the array with transpose
	array = transpose(array)

	// d. Sort the rows of this array with sortRows
	sortRows(array)

	// e. Transpose the array with transpose

	array = transpose(array)

	// f. Print the output 2D array to console
	fmt.Println("\nOutput array:")
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			fmt.Printf("%.1f ", array[i][j])
		}
		fmt.Println()
	}

}
