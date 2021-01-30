package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/urfave/negroni"
)

func main() {
	// Mux start
	mux := http.NewServeMux()

	// Argument parameter handling
	param := ""
	if len(os.Args) > 1 {
		param = os.Args[1]
	} else {
		param = "BRL"
	}

	// Default route declaration
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		// Exchange Rates API link that will be consumed
		res, err := http.Get("https://api.exchangeratesapi.io/latest?base=" + param + "&symbols=BRL,USD,EUR")

		// Error Handling
		if err != nil {
			log.Fatal(err)
		}

		// JSON data handling
		body, readErr := ioutil.ReadAll(res.Body)

		// JSON error handling
		if readErr != nil {
			log.Fatal(readErr)
		}

		fmt.Fprintf(w, string(body))
	})

	// Negroni default declarations
	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(mux)

	// API port that will be used
	http.ListenAndServe(":3000", n)
}
