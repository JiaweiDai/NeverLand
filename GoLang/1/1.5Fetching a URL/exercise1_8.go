package main 

import(
	"os"
	"strings"
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	for _, url := range os.Args[1:]{
		if ok := strings.HasPrefix(url , "http://") ; !ok{
			url = "http://" + url
		}

		resp , err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr , "fetch : %v" , err)
			os.Exit(0)
		}

		b , err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr , "reading %s : %v" , url , err)
		}

		fmt.Printf("%s" , b)

	}

}