## RANSAC Plane Fitting

This program implements the RANSAC algorithm to find the best plane that fits a set of points in 3D space.

**Author:** Rakshita Mathur  
**Student Number:** 300215340  
**Date:** 03/11/2023  
**Course:** CSI 2120 - Programming Paradigms  

### Description

This program takes as input a file containing the points in the format x y z, one point per line. It then uses the RANSAC algorithm to fit a plane to these points. The program outputs the plane coefficients A, B, C, D, and the number of points that are on the plane. The program also outputs the time taken to run the algorithm.

The program takes the following parameters as input:
-i: Input file name (default: PointCloud1) // do not include the extension
-c: Confidence level (default: 0.99)
-p: Percentage of points that are on the plane (default: 20)
-e: Maximum distance between a point and the plane (default: 0.2)


### Example run:

```sh
go run RANSAC.go PointCloud2 .99 20 .01
```

The above command will run the RANSAC algorithm on the points in PointCloud2.xyz with a confidence level of 0.99, a percentage of points on the plane of 20%, and a maximum distance of 0.01.

### Usage
- Make sure you have Go installed on your machine.
- Compile the program using the following command: go build RANSAC.go.
- Run the program with the desired parameters using the following command: ./RANSAC -i <input_file_name> -c <confidence_level> -p <percentage_of_points_on_plane> -e <maximum_distance>.
- The program will output the plane coefficients, the number of points on the plane, and the time taken to run the algorithm.

