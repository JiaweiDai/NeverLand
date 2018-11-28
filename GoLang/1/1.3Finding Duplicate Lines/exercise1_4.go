package main 

import(
	"os"
	"bufio"
	"fmt"
)

func countsLine(f *os.File , counts map[string]int){
	input := bufio.NewScanner(f)
	for input.Scan(){
		counts[input.Text()]++
	}
}

func printFileName(fileName string , counts map[string]int) {
	for _ , n := range counts{
		if n > 1{
			fmt.Println(fileName)
			break
		}	
	}
}

func main() {
	counts := make(map[string]int)

	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Println("usage: go run excerice1_4.go file name")
	} else {
		for _, fileName := range files{
			f , err := os.Open(fileName)
			if err != nil{
				fmt.Fprintf(os.Stderr, "excerice1_4 %v\n" , err)
				continue
			}

			countsLine(f , counts)
			f.Close()

			printFileName(fileName , counts)
		}
	}

	for line , n := range counts{
		if n > 1{
			fmt.Printf("%s\t%d\n", line , n)
		}
	}

}