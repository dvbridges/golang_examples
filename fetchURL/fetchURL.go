// Demo: How to fetch a URL and print the content

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error collecting %v: %v", resp, err)
			os.Exit(1)
		}

		/** Alternative method that copies input to Stdout

		output := bufio.NewReader(resp.Body)
		cp, err := io.Copy(os.Stdout, output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error collecting %v: %v", cp, err)
			os.Exit(1)
		}

		*/

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error collecting %v: %v", b, err)
			os.Exit(1)
		}

		fmt.Printf("%s: %s", resp.Status, b)

	}
}
