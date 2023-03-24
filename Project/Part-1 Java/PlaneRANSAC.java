/**
 * @author Rakshita Mathur
 * Student ID: 300215340
 * Course: CSI 2120 
 * Class: PlaneRANSAC
 */ 

import java.io.*;

public class PlaneRANSAC {
    
    /**
     * The point cloud object pc
     */
    private PointCloud pc;

    /**
     * The distance threshold eps
     */
    private double eps;

    /**
     * Constructor that accepts a point cloud object
     * @param pc the point cloud object
     */
    public PlaneRANSAC(PointCloud pc) {
        this.pc = pc;
    }

    /**
     * Setter method for eps
     * @param eps the distance threshold
     */
    public void setEps(double eps) {
        this.eps = eps;
    }

    /**
     * Getter method for eps
     * @return eps the distance threshold
     */
    public double getEps() {
        return eps;
    }

    /**
     * Method to calculate the number of iterations
     * @param confidence the confidence level
     * @param percentageOfPointsOnPlane the percentage of points on the plane
     * @return the number of iterations
     */
    public int getNumberOfIterations(double confidence, double percentageOfPointsOnPlane) {
        double p = percentageOfPointsOnPlane;
        double q = 1 - Math.pow(p,3);
        double e = 1 - confidence;
        return (int) Math.ceil(Math.log(e) / Math.log(q));
    }

    /**
     * Method to run the RANSAC algorithm
     * @param numberOfIterations the number of iterations
     * @param filename the name of the file to save the points on the plane
     * saves the points after the processing of algorithm in the file
     */
    public void run(int numberOfIterations, String filename) throws IOException {
        Point3D p1 = pc.getPoint();
        Point3D p2 = pc.getPoint();
        Point3D p3 = pc.getPoint();
        Plane3D plane = new Plane3D(p1, p2, p3);
        int maxInliers = 0;
        for (int i = 0; i < numberOfIterations; i++) {
            p1 = pc.getPoint();
            p2 = pc.getPoint();
            p3 = pc.getPoint();
            Plane3D newPlane = new Plane3D(p1, p2, p3);
            int inliers = 0;
            for (Point3D pt : pc) {
                if (newPlane.getDistance(pt) < eps) {
                    inliers++;
                }
            }
            if (inliers > maxInliers) {
                maxInliers = inliers;
                plane = newPlane;
            }
        }
        PointCloud planePoints = new PointCloud();
        for (Point3D pt : pc) {
            if (plane.getDistance(pt) < eps) {
                planePoints.addPoint(pt);
            }
        }
        planePoints.save(filename);
    }

    /**
     * Main method to test the PlaneRANSAC class
     * @param args[0] the name of the file containing the point cloud
     * @param args[1] the distance threshold
     * @param args[2] the confidence level
     * @param args[3] the percentage of points on the plane
     * @throws IOException
     * @output A xyz file containing the points on the plane
     */
    public static void main(String[] args) throws IOException {
        
        // getting the file name from the command line and senind it to the PointCloud class to read the points and store them in the arraylist
        PointCloud pc = new PointCloud(args[0]);

        // creating the object of the PlaneRANSAC class and passing the point cloud object to it
        PlaneRANSAC ransac = new PlaneRANSAC(pc);

        // setting the distance threshold 
        ransac.setEps(Double.parseDouble(args[1]));

        // getting the number of iterations
        int iterations = ransac.getNumberOfIterations(Double.parseDouble(args[2]), Double.parseDouble(args[3]));

        // setting the file as the name of the datafile which we got from the command line with the number of iterations and the extension .xyz
        String filename = args[0].split("\\.")[0] + "_plane_with_iterations" + iterations + ".xyz";

        // running the RANSAC algorithm with the number of iterations and the file name
        ransac.run(iterations, filename);
    }

}