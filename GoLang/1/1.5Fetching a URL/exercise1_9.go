package main 

import (
	"os"
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
)

func main() {
	for _, url := range os.Args[1:]{
		if ok := strings.HasPrefix(url , "http://") ; !ok {
			url = "http://" + url
		}

		resp , err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr , "fetch : %v\n" , err)
			os.Exit(1)
		}

		fmt.Println("#############")
		code := resp.Status
		code1 := resp.StatusCode

		fmt.Println("code : ", code)
		fmt.Println("code1 : ", code1)
		fmt.Println("#############")
		
		b , err := ioutil.ReadAll(resp.Body)

		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr , "reading %s : %v\n" , url , err)
			os.Exit(1)
		}

		fmt.Printf("%s\n", b)

	}
}