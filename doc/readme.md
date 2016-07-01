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
Some basic golanf staff

Run: ./main basic demo
 
demo:
   - defer
   - dirs list<directories>
   - json
   - pipes
   - spinner 
   - surface
   - toposort
   - utils-echo
   - utils-sha
   - utils-anon
   - utils-slow
For mandelbrot see below.   
   
**Methods**
```yaml
Defer panic - shows deffered action and throws panic
Directory traversal - shows volume and files number in specified dirs
Files - count and show duplicate lines in files
JsonMarshaller - marshalls struct to json
MandelBrot - draws a fractal picture.run: ./main web web-server. In browser localhost:8000/fractal
Pipeline - shows go routine interaction via pipes 
Spinner - counts fibonacci numbers with delay 
Surface - draws and svg 
Toposort - sorts a map   
Utils: 
   #small util and demo functions
   Echo - shows cmd arguments
   SHA - counts sha sum of specified string
   AnonFunc - show a call of a anon func
   BigSlow - turns on a slow operation and traces it
     
```

Geometry
-----
Type struct and methods demo

run: ./main geometry x1 y1 x2 y2   

**Types**
```yaml
Point:
   X: float64
   Y: float64

ColoredPoint:
   Color: color.RGBA
```

**Methods**
```yaml
PointDistance - distance between two points
ScaleBy - scale point coordinates by factor
Add - add one point to another
Sub - substract one point from another
TranslateBy - move point with offset
Distance - perimeter distance
GeometryDemo - run all methods with specified coordinates
```

IssueReporter
-------------
Shows list of issues in htmll or list mode.

run: ./main issues mode repo isssue 

mode:
    -HTML 
    -Console

**Types**
```yaml
IssuesSearchResult: 
    TotalCount: int 
    Items:      []Issue

Issue:
    Number:     int
    HTMLURL:    string 
    Title:      string
    State:      string
    User:       User
    CreatedAt:  time.Time
    Body:       string
    
User:
    Login:      string
	HTMLURL:    string

```

**Methods**
```yaml
SearchIssues - search from source in selected mode
```

Links
---------
Fetch data from url and parse it to get links
run: ./main links action list<urls>

action:
 - fetch-console
 - fetch-file
 - fetch-concurent
 - links-from-url
 - links-from-urls
 - links-breadth
 - links-tags
 - links-tags-tree

**Methods**
```yaml
FetchToConsole - fetch list of urls and show it in console
FetchToFile - fetch a single url to file
ConcurentFetcher - fetch list of urls and show fetch time
FindLinks - show list of links from the list of urls
FindLinksFromUrl - show list of links from a single url
FindLinksBreadthFirst - show links in console from a single url
ShowTags - outline stack of tags from the list of urls
ShowTagsTree - outline tree of tags from the list of urls
```


Lissajous
---------
Creates a lissajous figures gif
run: ./main web web-server. In browser localhost:8000/lissajous

**Methods**
```yaml
Lissajous - create a gif
```

Server
------
Runs different servers
run: ./main web server-type

If you want to use chat server :  ./main web chat client-type

server-type:
   - web-server
   - clock
   - echo
   - chat

client-type: 
   - client
   - concurent client

**Methods**
```yaml
RunWebServer - run web server.
   Modes:
        - check server is operational: localhost:8000
        - run enties counter: localhost:8000/counter
        - draw lissajous: localhost:8000/lissajous  
        - draw fractal: localhost:8000/fractal

RunClockServer - run clock server.
    Run: localhost:8080

RunEchoServer - run echo server.
    Run: localhost:5000

RunChatServer - run chat server.
    Run: localhost:8090

Netcat - run tcp client

NetcatChannel - run concurrent tcp client based on go routines and channels.            
```



