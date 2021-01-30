package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/urfave/negroni"
)

type Data struct {
	Rates *Rates `json:"rates"`
}

type Rates struct {
	BRL string `json:"brl"`
	USD string `json:"usd"`
	EUR string `json:"eur"`
}

func main() {
	mux := http.NewServeMux()

	param := ""

	param = os.Args[1]

	if len(os.Args) > 1 {
		param = os.Args[1]
	} else {
		param = "BRL"
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		res, err := http.Get("https://api.exchangeratesapi.io/latest?base=" + param + "&symbols=BRL,USD,EUR")

		if err != nil {
			log.Fatal(err)
		}

		body, readErr := ioutil.ReadAll(res.Body)

		if readErr != nil {
			log.Fatal(readErr)
		}

		fmt.Fprintf(w, string(body))
	})

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
