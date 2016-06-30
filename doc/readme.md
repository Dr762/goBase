Go lang demos
======================

Table of Contents
-----------------

- [Overview](#overview)

Modules:
- [Basic](#basic)
- [Geometry](#geometry)
- [IssueReporter](#issue_reporter)
- [Links](#links)
- [Lissajous](#lissajous)
- [Server](#server)


Overview
--------

Bunch of demos from Kernigan's GoLang book. All is ran from console

Basic
-----

Defer panic: 
    #shows deffered action and throws panic
    run: ./main basic defer

    
Directory traversal: 
    #shows volume and files number in specified dirs
    run: ./main basic dirs list<directories>
    
Files: 
    #count and show duplicate lines in files
    run: ./main basic files list<files>
    
JsonMarshaller: 
    #marshalls struct to json
    run: ./main basic json

MandelBrot: 
   #draws a fractal picture. pipe it to a jpeg.file 
   run: ./main basic fractal 

Pipeline: 
   #shows go routine interaction via pipes 
   run: ./main basic pipes 

Spinner:
   #counts fibonacci numbers with delay
   run: ./main basic spinner 

Surface:
   #draws and svg 
   run: ./main basic surface 

Toposort:
   #sorts a map
   run: ./main basic toposort
   
Utils:
   #small util and demo functions
   Echo:
       #shows cmd arguments
       run: ./main basic utils-echo
   
   SHA:
       #counts sha sum of specified string
       run: ./main basic utils-sha
   
   AnonFunc:
       #show a call of a anon func
       run: ./main basic utils-anon
   
   BigSlow:
       #turns on a slow operation and traces it
       run: ./main basic utils-slow

Geometry
-----
#type struct and methods demo

run: ./main geometry x1 y1 x2 y2   

**Types**

Point:
   X: float64
   Y: float64

ColoredPoint:
   Color: color.RGBA

**Methods**
PointDistance - distance between two points
ScaleBy - scale point coordinates by factor
Add - add one point to another
Sub - substract one point from another
TranslateBy - move point with offset
Distance - perimeter distance
GeometryDemo - run all methods with specified coordinates
        