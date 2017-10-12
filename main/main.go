package main

import (
	"fmt"
	"os"

	"github.com/abondar24/GoBase/basic"
	reporter "github.com/abondar24/GoBase/issue_reporter"
	"github.com/abondar24/GoBase/links"
	"github.com/abondar24/GoBase/server"
	"gopkg.in/alecthomas/kingpin.v2"

	"log"
	"strconv"

	"github.com/abondar24/GoBase/geometry"
	 "github.com/abondar24/GoBase/network"
	"github.com/abondar24/GoBase/client"
)

var (
	base       = kingpin.New("base", "Go base Demos")
	srv        = base.Command("server", "Start a new server")
	serverArgs = srv.Arg("server", "Available serverArgs.").Required().Strings()
	clt        = base.Command("client","Run client")
	clientArgs = clt.Arg("client","Available clients").Required().Strings()
	link       = base.Command("links", "Fetch from url")
	linkArgs   = link.Arg("type", "Fetch type").Required().Strings()
	issues     = base.Command("issues", "Search issues on github")
	issuesArgs = issues.Arg("type", "Output type").Required().Strings()
	geom       = base.Command("geometry", "Geometry Demo")
	geomArgs   = geom.Arg("points", "Point coordinates").Required().Strings()
	basics     = base.Command("basic", "Basic staff")
	baseArgs   = basics.Arg("bases", "Basic features").Required().Strings()
	ntw        = base.Command("network", "Network examples")
	netArgs    = ntw.Arg("network","Which example to run").Required().Strings()
)

func main() {
	fmt.Println("Ok,let's GO")

	cmd := kingpin.MustParse(base.Parse(os.Args[1:]))

	switch cmd {

	case basics.FullCommand():
		basicRun(*baseArgs)

	case geom.FullCommand():
		geomRun(*geomArgs)

	case issues.FullCommand():
		issuesRun(*issuesArgs)

	case link.FullCommand():
		linksRun(*linkArgs)

	case srv.FullCommand():
		serverRun(*serverArgs)

	case clt.FullCommand():
		clientRun(*clientArgs)

	case ntw.FullCommand():
		networkRun(*netArgs)
	}

	os.Exit(0)
}


func basicRun(args []string) {
	argMap := getArgMap(args)

	if argMap["defer"] {
		basic.DeferPanicDemo()
	}

	if argMap["dirs"] {
		basic.DirectoryTraversal(os.Args[3:])
	}

	if argMap["files"] {
		basic.DuplicatesFilesInput(os.Args[3:])
	}

	if argMap["json"] {
		basic.Jsonmarshaller()
	}

	if argMap["pipes"] {
		basic.PipelineDemo()
	}

	if argMap["spinner"] {
		basic.SpinnerDemo()
	}

	if argMap["surface"] {
		basic.SurfaceDrawer()
	}

	if argMap["toposort"] {
		basic.TopoSortDemo()
	}

	if argMap["utils-echo"] {
		basic.Echo()
	}

	if argMap["utils-sha"] {
		basic.CountSha()
	}

	if argMap["utils-anon"] {
		basic.AnonFuncDemo()
	}

	if argMap["utils-slow"] {
		basic.BigSlowOperation()
	}

}

func geomRun(args []string) {
	p1x, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		log.Fatal(err)
	}

	p1y, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		log.Fatal(err)
	}

	p2x, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		log.Fatal(err)
	}

	p2y, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		log.Fatal(err)
	}

	geometry.GeometryDemo(p1x, p1y, p2x, p2y)
}

func issuesRun(args []string) {
	argMap := getArgMap(args)

	if argMap["html"] {
		reporter.SearchIssues(os.Args[3:], reporter.HTMLMode)
	}

	if argMap["console"] {
		reporter.SearchIssues(os.Args[3:], reporter.ConsoleMode)
	}

}

func linksRun(args []string) {
	argMap := getArgMap(args)

	if argMap["linkArgs-console"] {
		_, err := links.FetchToConsole(os.Args[3:], links.FetchMode)
		if err != nil {
			log.Fatal(err)
		}
	}

	if argMap["linkArgs-file"] {

		fName, flen, err := links.FetchToFile(os.Args[3])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("File %v , len %v\n", fName, flen)
	}

	if argMap["linkArgs-concurent"] {

		links.ConcurentFetcher(os.Args[3:])

	}

	if argMap["links-from-url"] {
		lnks, err := links.FindLinksFromUrl(os.Args[3])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(lnks)
	}

	if argMap["links-from-urls"] {
		html, err := links.FetchToConsole(os.Args[3:], links.LinksMode)
		if err != nil {
			log.Fatal(err)
		}

		lnks, err := links.FindLinks(html)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(lnks)
	}

	if argMap["links-breadth"] {
		links.FindLinksBreadthFirst(links.Crawl, os.Args[3:])

	}

	if argMap["links-tags"] {
		html, err := links.FetchToConsole(os.Args[3:], links.LinksMode)
		if err != nil {
			log.Fatal(err)
		}
		links.ShowTags(html)
	}

	if argMap["links-tags-tree"] {
		html, err := links.FetchToConsole(os.Args[3:], links.LinksMode)
		if err != nil {
			log.Fatal(err)
		}
		links.ShowTagsTree(html)
	}

}

func serverRun(args []string) {
	argMap := getArgMap(args)

	if argMap["srv-server"] {
		server.WebServer()
	}

	if argMap["clock"] {
		server.ClockServer()
	}

	if argMap["echo"] {
		server.EchoServer()
	}

	if argMap["chat"] {
		server.ChatServer()
	}

	if argMap["rest"] {
		server.RestServer()
	}

	if argMap["daytime-tcp"] {
		server.DaytimeTcpServer()
	}

	if argMap["daytime-udp"] {
		server.DaytimeUdpServer()
	}

	if argMap["daytime-asn1"] {
		server.DaytimeAsn1Server()
	}

	if argMap["multi"] {
		server.MultithreadServer()
	}

	if argMap["json"] {
		server.JsonServer()
	}

	if argMap["gob"] {
		server.GobServer()
	}

	if argMap["ftp"] {
		server.FtpServer()
	}

	if argMap["utf16"] {
		server.Utf16Server()
	}
}

func clientRun(args []string) {
	argMap := getArgMap(args)

	if argMap["netcat"] {
		client.Netcat()
	}

	if argMap["tcp"]{
		client.TcpClient(os.Args[3])
	}

	if argMap["concurrent"] {
		client.NetcatChannel()
	}

	if argMap["daytime-udp"] {
		client.DaytimeUdpClient(os.Args[3])
	}

	if argMap["daytime-asn1"] {
		client.DaytimeAsn1Client(os.Args[3])
	}

	if argMap["json"] {
		client.JsonClient(os.Args[3])
	}

	if argMap["gob"] {
		client.GobClient(os.Args[3])
	}

	if argMap["ftp"] {
		client.FtpClient(os.Args[3])
	}

	if argMap["utf16"] {
		client.Utf16Client(os.Args[3])
	}

}

func networkRun(args []string)  {
	argMap := getArgMap(args)

	if argMap["get-mask"] {
		network.GetMask(os.Args[3])
	}

	if argMap["resolve-ip"]{
		network.ResolveIP(os.Args[3])
	}

	if argMap["host-lookup"]{
		network.HostLookup(os.Args[3])
	}

	if argMap["port-lookup"]{
		network.PortLookup(os.Args[3],os.Args[4])
	}

	if argMap["ping"]{
		network.Ping(os.Args[3])
	}

	if argMap["asn1"]{
		network.Asn1Marshall(os.Args[3])
	}

	if argMap["json"]{
		network.JsonMarshall()
	}


	if argMap["base64"]{
		network.Base64Encoder()
	}
}

func getArgMap(args []string) map[string]bool {
	argMap := map[string]bool{}
	for _, arg := range args {
		argMap[arg] = true
	}

	return argMap
}

