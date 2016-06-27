package main
import (
	"os"
	"net/http"
	"path"
	"io"
)

//improved url_fetcher which can write to local file
func main() {

	fetchURL(os.Args[1:])

}


func fetchURL(url []string) (filename string, n int64, err error) {

	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f,err := os.Create(local)
	if err!=nil{
		return "", 0, err
	}
    n,err = io.Copy(f,resp.Body)
	if closeErr :=f.Close(); err ==nil{ //close,but prefer error from copy if any
		err = closeErr
	}
	return local,n,err
}