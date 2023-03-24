/**
 * @author Rakshita Mathur
 * Student ID: 300215340
 * Course: CSI 2120 
 * Class: Plane3D 
 */ 

public class Plane3D {
    private double a;
    private double b;
    private double c;
    private double d;

    // Constructor for a plane with three points

    public Plane3D(Point3D p1, Point3D p2, Point3D p3) {
        a = (p2.getY() - p1.getY()) * (p3.getZ() - p1.getZ()) - (p2.getZ() - p1.getZ()) * (p3.getY() - p2.getY());
        b = (p2.getZ() - p1.getZ()) * (p3.getX() - p1.getX()) - (p2.getX() - p1.getX()) * (p3.getZ() - p2.getZ());
        c = (p2.getX() - p1.getX()) * (p3.getY() - p1.getY()) - (p2.getY() - p1.getY()) * (p3.getX() - p2.getX());
        d = -a * p1.getX() - b * p1.getY() - c * p1.getZ();
    }
    
    // Constructor for a plane with a, b, c, d

    public Plane3D(double a, double b, double c, double d) {
        this.a = a;
        this.b = b;
        this.c = c;
        this.d = d;
    }

    // Constructor for a plane with a normal vector and a point

    public double getDistance(Point3D pt) {
        return Math.abs(a * pt.getX() + b * pt.getY() + c * pt.getZ() + d) / Math.sqrt(a * a + b * b + c * c);
    }
     
 
}
