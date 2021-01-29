package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

type Data []struct {
	BRL string `json:""`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		res, err := http.Get("https://api.exchangeratesapi.io/latest?base=BRL&symbols=BRL,USD")

		if err != nil {
			log.Fatal(err)
		}

		body, readErr := ioutil.ReadAll(res.Body)

		if readErr != nil {
			log.Fatal(readErr)
		}

		fmt.Println(string(body))
	})

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
