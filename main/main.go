package main

import (
	"fmt"
	"github.com/goBase/basic"
	reporter "github.com/goBase/issue_reporter"
	"github.com/goBase/links"
	"github.com/goBase/server"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"

	"github.com/goBase/geometry"
	"log"
	"strconv"
)

var (
	base     = kingpin.New("base", "gobaseDemos")
	web      = base.Command("web", "Start a new server")
	servers  = web.Arg("server", "Available servers.").Required().Strings()
	link     = base.Command("links", "Fetch from url")
	fetch    = link.Arg("type", "Fetch type").Required().Strings()
	issues   = base.Command("issues", "Search issues on github")
	iss      = issues.Arg("type", "Output type").Required().Strings()
	geom     = base.Command("geometry", "Geometry Demo")
	points   = geom.Arg("points", "Point coordinates").Required().Strings()
	basics   = base.Command("basic", "Basic staff")
	baseArgs = basics.Arg("bases", "Basic features").Required().Strings()
)

func main() {
	fmt.Println("Ok,let's GO")

	cmd := kingpin.MustParse(base.Parse(os.Args[1:]))

	switch cmd {

	case basics.FullCommand():
		basicRun(*baseArgs)

	case geom.FullCommand():
		geomRun(*points)

	case issues.FullCommand():
		issuesRun(*iss)

	case link.FullCommand():
		linksRun(*fetch)

	case web.FullCommand():

		serverRun(*servers)
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

	if argMap["fractal"] {
		basic.DrawFractal()
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

	if argMap["fetch-console"] {
		_, err := links.FetchToConsole(os.Args[3:], links.FetchMode)
		if err != nil {
			log.Fatal(err)
		}
	}

	if argMap["fetch-file"] {

		fName, flen, err := links.FetchToFile(os.Args[3])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("File %v , len %v\n", fName, flen)
	}

	if argMap["fetch-concurent"] {

		links.ConcurentFetcher(os.Args[3:])

	}

	if argMap["links-from-url"] {
		links, err := links.FindLinksFromUrl(os.Args[3])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(links)
	}

	if argMap["links-from-urls"] {
		html, err := links.FetchToConsole(os.Args[3:], links.LinksMode)
		if err != nil {
			log.Fatal(err)
		}

		links, err := links.FindLinks(html)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(links)
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

	if argMap["web-server"] {
		server.RunWebServer()
	}

	if argMap["clock"] {
		server.RunClockServer()
	}

	if argMap["echo"] {
		server.RunEchoServer()
	}

	if argMap["chat"] {
		server.RunChatServer()
	}

	if argMap["client"] {
		server.Netcat()
	}

	if argMap["client-concurent"] {
		server.NetcatChannel()
	}

}

func getArgMap(args []string) map[string]bool {
	argMap := map[string]bool{}
	for _, arg := range args {
		argMap[arg] = true
	}

	return argMap
}
