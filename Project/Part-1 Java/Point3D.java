
/**
 * @author Rakshita Mathur
 * Student ID: 300215340
 * Course: CSI 2120 
 * Class: Point3D 
 */ 
 
public class Point3D{
    /**
     * x coordinate of the point
     */
    public double x;
    /**
     * y coordinate of the point
     */
    public double y;
    /**
     * z coordinate of the point
     */
    public double z;
    /**
     * Constructor that accepts three doubles
     * @param x coordinate of the point
     * @param y coordinate of the point
     * @param z coordinate of the point
     */
    public Point3D(double x, double y, double z){
        this.x = x;
        this.y = y;
        this.z = z;

    }
    /**
     * getter method for x coordinate
     * @return  x coordinate
     */
    public double getX() {
        return x;
    }
    /**
     * getter method for y coordinate
     * @return y coordinate
     */
    public double getY() {
        return y;
    }
    /**
     * getter method for z coordinate
     * @return z coordinate
     */
    public double getZ() {
        return z;
    }
    /**
     * converting the point to string
     * @returns the string representation of the point
     */
     public String toString() {
         return  (this.getX() + " " + this.getY() + " " + this.getZ());
     }
}