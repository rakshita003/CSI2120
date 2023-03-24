This archive contains this file (ReadMe),the main PlaneRANSAC.go file containing the code. There are the data files PointCloud1.xyz, PointCloud2.xyz and PointColid3.xyz and the 3 png file from which the data points are collected.

author: Rakshita Mathur

student number: 300215340

date: 03/11/2023

Course: CSI 2120 - programming paradigms

description: This program implements the RANSAC algorithm to find the best plane that fits a set of points.

              The program takes as input a file containing the points in the format x y z, one point per line.
              
              The program outputs the plane coefficients A, B, C, D and the number of points that are on the plane.
              
              The program also outputs the time taken to run the algorithm.
              
              The program takes as input the following parameters:
              
              -i: input file name (default: PointCloud1) // do not include the extension
              
              -c: confidence level (default: 0.99)
              
              -p: percentage of points that are on the plane (default: 20)
              
              -e: maximum distance between a point and the plane (default: 0.2)
              
             Example run:  go run RANSAC.go PointCloud2 .99 20 .01
             
 			       The program will run the RANSAC algorithm on the points in PointCloud2.xyz with a confidence level of 0.99,
             a percentage of points on the plane of 20% and a maximum distance of 0.01.

