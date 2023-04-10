# Prolog RANSAC Algorithm

This Prolog implementation provides predicates for the Random Sample Consensus (RANSAC) algorithm. RANSAC is a robust algorithm used for fitting models to data points with noise or outliers. This implementation includes four predicates:

1. `random3points/2`: This predicate randomly selects three points from a list of points and returns them as a triplet of points in the form `[[x1, y1, z1], [x2, y2, z2], [x3, y3, z3]]`.

2. `plane/2`: This predicate computes the coefficients of a plane equation given three points, represented as a list of points in the form `[[x1, y1, z1], [x2, y2, z2], [x3, y3, z3]]`. The plane equation is represented as a list of coefficients `[a, b, c, d]` from the equation `ax + by + cz = d`.

3. `support/4`: This predicate finds the support points of a plane given a list of points and an epsilon value. It returns the number of support points `N` that are within the specified distance `Eps` from the plane defined by the coefficients `Plane`.

4. `ransac_number_of_iterations/3`: This predicate calculates the number of iterations required by RANSAC with given confidence and percentage parameters `Confidence` and `Percentage`.

# Author
Written by Rakshita Mathur.

