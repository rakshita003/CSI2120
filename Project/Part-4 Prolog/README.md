Prolog RANSAC Algorithm
This Prolog implementation provides predicates for the Random Sample Consensus (RANSAC) algorithm. RANSAC is a robust algorithm used for fitting models to data points with noise or outliers. This implementation includes four predicates:

random3points/2: This predicate randomly selects three points from a list of points and returns them as a triplet of points in the form [[x1, y1, z1], [x2, y2, z2], [x3, y3, z3]].

plane/2: This predicate computes the coefficients of a plane equation given three points, represented as a list of points in the form [[x1, y1, z1], [x2, y2, z2], [x3, y3, z3]]. The plane equation is represented as a list of coefficients [a, b, c, d] from the equation ax + by + cz = d.

support/4: This predicate finds the support points of a plane given a list of points and an epsilon value. It returns the number of support points N that are within the specified distance Eps from the plane defined by the coefficients Plane.

ransac_number_of_iterations/3: This predicate calculates the number of iterations required by RANSAC with given confidence and percentage parameters Confidence and Percentage.

Usage
To use these predicates in your Prolog program, simply include the ransac.pl file in your project and import the predicates as needed. For example:

prolog
Copy code
:- consult('ransac.pl'). % Include the ransac.pl file

% Use the predicates in your program
...
Make sure to provide the required input arguments as specified in the documentation of each predicate for correct usage.

Example
Here's an example of how you can use the RANSAC algorithm predicates in Prolog:

prolog
Copy code
:- consult('ransac.pl').

% Example data
points([[1, 2, 3], [4, 5, 6], [7, 8, 9], [10, 11, 12], [13, 14, 15]]).

% Randomly select 3 points
random3points(Points, Point3),

% Compute plane equation from 3 points
plane(Point3, Plane),

% Find support points of plane within a certain distance
Eps = 0.1,
support(Plane, Points, Eps, N),

% Calculate number of iterations for RANSAC with given confidence and percentage
Confidence = 0.99,
Percentage = 0.3,
ransac_number_of_iterations(Confidence, Percentage, Iterations).
License
This Prolog RANSAC Algorithm implementation is released under the MIT License.

Author
Written by Rakshita Mathur.

References
RANSAC Algorithm - Wikipedia
Prolog - SWI-Prolog Documentation
