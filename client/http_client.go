package client

import (
	"net/url"
	"log"
	"net/http"
	"fmt"
	"os"
	"strings"
)

func HttpClient(host string) {

	hostUrl, err := url.Parse(host)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}

	request, err := http.NewRequest("GET", hostUrl.String(), nil)

	//utf-8 only
	request.Header.Add("Accept-Charset", "UTF-8;q=1, ISO-8859-1;q=0")
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(request)
	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	chSet := getCharset(response)
	fmt.Printf("Got charset %s\n", chSet)
	if chSet != "UTF-8" {
		fmt.Println("Can't handle", chSet)
		os.Exit(4)
	}

	var buf [512]byte
	reader := response.Body
	fmt.Println("Got body")
	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			os.Exit(0)
		}
		fmt.Print(string(buf[0:n]))
	}
	os.Exit(0)

}

func getCharset(response *http.Response) string {
	contentType := response.Header.Get("Content-Type")
	if contentType == "" {
		return "UTF-8"
	}

	idx := strings.Index(contentType,"charSet:")
	if idx ==-1{
		return "UTF-8"
	}

	return strings.Trim(contentType[idx:],"")
}
