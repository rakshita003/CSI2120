# RANSAC Algorithm for Plane Fitting in Point Cloud Data

This repository contains a Scheme implementation of the Random Sample Consensus (RANSAC) algorithm for fitting a plane in a point cloud data. The RANSAC algorithm is used to estimate the parameters of a plane model that best fits a set of points contaminated with outliers.

## Project Details
- Author: Rakshita Mathur
- Student ID: 300215340
- Date: 31/03/2023
- Course: CSI 2120 Project - Part 3

## How to Run
The RANSAC algorithm can be run by calling the `run-planeRANSAC` function with the following parameters:

```scheme
(run-planeRANSAC "filepath" confidence percentage eps)
```
## Parameters

- `"filepath"`: The path to the point cloud data file in XYZ format.
- `confidence`: The desired confidence level for the RANSAC algorithm, ranging from 0 to 1.
- `percentage`: The percentage of points to be used for the initial sample, ranging from 0 to 1.
- `eps`: The maximum allowed error (distance) for a point to be considered an inlier.

The output of the algorithm, which includes the estimated plane parameters and the inliers, will be written to a new file in XYZ format with the suffix "-result.xyz".

## Functions and Files

- `run-planeRANSAC`: The main function that runs the RANSAC algorithm.
- `write-lines`: Function to write lines of points to a file.
- `readXYZ`: Function to read the point cloud data file and convert string values to numbers.
- `pick-random-points`: Function to randomly pick three points from the set of points.
- `plane`: Function to calculate the plane equation given three points.
- `support`: Function to calculate the support of a plane.
- `distance-to-plane`: Function to calculate the distance of a point from a plane.
- `CSI2120/Project/Part-3 Scheme/Ransac_300215340.scm`: The Scheme source file containing the implementation of the RANSAC algorithm.

## Usage

1. Clone the repository to your local machine.
2. Load the `CSI2120/Project/Part-3 Scheme/Ransac_300215340.scm` file in your Scheme interpreter.
3. Call the `run-planeRANSAC` function with the desired parameters to run the RANSAC algorithm on your point cloud data.

## Example

```scheme
(run-planeRANSAC "Point_Cloud_1_No_Road_Reduced.xyz" 0.99 0.5 0.8)

```
