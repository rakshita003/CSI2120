/**
 * @author Rakshita Mathur
 * Student ID: 300215340
 * Course: CSI 2120 
 * Class: PointCloud
 */ 

import java.util.*;
import java.io.*;

public class PointCloud implements Iterable<Point3D> {
    /**
     * Arraylist of points of type Point3D
     */
    ArrayList<Point3D> points =  new ArrayList<Point3D>();
    /**
     * Default Constructor that creates an empty point cloud using the 3d points 
     */
    public PointCloud() {
        points = new ArrayList<Point3D>();
    }
    /**
     * Constructor that creates a point cloud that reads the points from a file
     * @param filename name of the file
     * @throws FileNotFoundException
     */
    public PointCloud(String filename) throws FileNotFoundException{

        File file = new File(filename);
        
        Scanner r = new Scanner(file);
        r.nextLine();
        while (r.hasNextLine()) {
            String[] parts = r.nextLine().split("\\s+");
            Point3D pt = new Point3D(
            Double.parseDouble(parts[0]), 
            Double.parseDouble(parts[1]), 
            Double.parseDouble(parts[2])
            );
            points.add(pt);
        }
        r.close();
    }
    /**
     * adds a point to the point cloud
     * @param pt point to be added
     */
    public void addPoint(Point3D pt) {
        points.add(pt);
    }
    /**
     * returns the number of points in the point cloud
     * @return number of points
     */
    public Point3D getPoint() {
        Random rand = new Random();
        return points.get(rand.nextInt(points.size()));
    }
    /**
     * returns the number of points in the point cloud
     * @return number of points
     */
    public void save(String filename) throws IOException {
        PrintWriter out = new PrintWriter(new File(filename));
        out.println("\tx\t\t\ty\t\t\tz");
        for (Point3D pt : points) {
        
            out.println(pt);
        }
        out.close();
    }
    /**
     * returns the number of points in the point cloud
     * @return number of points
     */
    public Iterator<Point3D> iterator() {
        Iterator<Point3D> it = new MyIterator(points);
        return it;
    }

}


   