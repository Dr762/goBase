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
- [Client](#client)
- [Network](#network)
- [Security](#security)


Overview
--------

Bunch of demos for GoLang . All is ran from console

Basic
-----
Some basic golanf staff

Run: ./main basic <demo>
 
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

run: ./main issues <mode> <repo> <issue> 

mode:
    - HTML 
    - Console

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
run: ./main links <action> list<urls>

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
run: ./main server web-server. In browser localhost:8000/lissajous

**Methods**
```yaml
Lissajous - create a gif
```

Server
------
Runs different servers
run: ./main server <server-type>

If you want to use chat server :  ./main server chat client-type

server-type:
   - web-server
   - clock
   - echo
   - chat
   - rest
   - daytime-tcp
   - daytime-udp
   - daytime-asn1
   - multi
   - json
   - gob
   - ftp
   - utf16
   - tls
   - file

**Methods**
```yaml
RunWebServer - run web server.
   Modes:
        - check server is operational: localhost:8000
        - run enties counter: localhost:8000/counter
        - draw lissajous: localhost:8000/lissajous  
        - draw fractal: localhost:8000/fractal

ClockServer - run clock server.
    Run: localhost:8080

EchoServer - run echo server.
    Run: localhost:5000

ChatServer - run chat server.
    Run: localhost:8090
    
RestServer - run rest server.
    Run: localhost:8080
    REST methods: 
      - /getPerson/{person_id}, GET
      - /getPersons, GET
      - /insertPerson, POST
      - /getJob/{job_id}, GET
      - /getJobForPerson/{person_id}, GET
      - /insertJob, POST

DaytimeTcpServer - run daytime tcp server.
    Run: localhost:1200
    
DaytimeUdpServer - run daytime udp server.
    Run: localhost:1300    
    
DaytimeAsn1Server - run daytime server sending data in asn1 format.
    Run: localhost:1400    
    

MultithreadServer -run multithread server.
    Run: localhost:1201    

JsonServer - run json server.
    Run: localhost:1500   
    
GobServer - run gob server.
    Run: localhost:1600  
    
FtpServer - run ftp server.
    Run: localhsot:1700   
    
Utf16Server - run server which sends message in utf16
    Run: localhost:1800    
    
TlsServer - run server with x509 certificate(certificate requrired)
    Run: localhost:1900      
    
FileServer - run fileserver
    Run: localhost:2000                      
```

Client
------
Runs different client
run: ./main client <client-type>

client-type: 
   - tcp <host>:<port>
   - netcat
   - concurrent 
   - daytime-udp <host>:<port>
   - daytime-asn1 <host>:<port>
   - json <host>:<port>
   - gob <host>:<port>
   - ftp <host>
   - utf16 <host>:<port>
   - tls <host>:<port>
   - http <host>

**Methods**
```yaml
     
TcpClient - run tcp client

Netcat - run netcat client

NetcatChannel - run concurrent tcp client based on go routines and channels.  

DaytimeUdpClient - run daytime udp client

DaytimeAsn1Client - run daytime asn1 client

JsonClient - run json client

FtpClient - run ftp client

Utf16Client - run client which accepts message in utf16

TlsClient - run client to connect to tls server

HttpClient - run a http client
          
```

Network
---------
Basic network operations
run: ./main network <operation> <param><param>


operations:
    - get-mask <ipAddr>
    - resolve-ip <hostname>
    - host-lookup <hostname>
    - port-lookup <protocol> <service>
    - ping <hostname>
    - asn1 <value>
    - json
    - base64
    - http-header <url>
    - http-get <url>
    
     
   
**Methods**
```yaml
GetMask - get mask of ip address
ResolveIP - get ip address of hostname
HostLookup - get ip addresses of host
PortLookup - get port of service
Ping - send icmp request to host
Asn1Marschall - marshalls and unmarshalls value to/from byte array
JsonMarshall - marshalls and unmarshalls value to/from json
Base64Encoder - encodes and decodes a byte array which consists of 8 digigts
ReadHttpHeader - reades headers of http response
ReadHttpGet - make a get request to url
```

Security
---------
Basic security examples related to network
run: ./main security <example> <param><param>


examples:
    - md5 <string>
    - blowfish <string>
    - gen-rsa
    - load-rsa
    - x509-gen
    - x509-load
     
        
**Methods**
```yaml
Md5Hash - calculates and outputs hashvalue of incoming string
BlowFish - uses blofish to encode/decode string
GenRsaKey - generates a rsa key and saves to the file
LoadRsaKey - loads a rsa kay from file(you need to have private.key and public.key files)
GenX509 - generates a new x509 certificate based on private.key
LoadX509 - loads and parses x509 certificate
```



