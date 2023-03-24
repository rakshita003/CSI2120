/**
 * @author Rakshita Mathur
 * Student ID: 300215340
 * Course: CSI 2120 
 * Class: MyIterator
 */ 


 import java.util.ArrayList;
 import java.util.Iterator;
 import java.util.NoSuchElementException;
 
 class MyIterator implements Iterator<Point3D>{
    
     private int index = 0;
     private ArrayList<Point3D> cloud = new ArrayList<Point3D>();
 
     public MyIterator(ArrayList<Point3D> points){
         cloud = points;
         
     }
 
     @Override
     public Point3D next(){
         if(hasNext()){
             return cloud.get(index++);
         }
         else{
             throw new NoSuchElementException("There no points left in the cloud");
         }
     }
     @Override
     public boolean hasNext(){
         return index < cloud.size();
     }
     @Override
     public void remove(){
         if( index <= 0){
             throw new IllegalStateException("You can't delete point before first next() method call");
         }
         cloud.remove(--index);
     }
 }
 