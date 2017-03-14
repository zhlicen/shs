package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/daviddengcn/go-colortext"
)

var showDetail bool

func handleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1><img src=\"%s\"></h1>", logoImage)
	fmt.Fprintf(w, "<p>[Mirror] %s %s</p>", r.Method, r.URL.Path)
	fmt.Fprintf(w, "<p>[Status] OK!</p>")
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		ct.ChangeColor(ct.White, true, ct.Black, false)
		fmt.Printf("BODY   : ")
		ct.ChangeColor(ct.White, true, ct.Yellow, false)
		fmt.Printf("%s\n", string(body))
	}
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	defer ct.ChangeColor(ct.White, false, ct.Black, false)
	ct.ChangeColor(ct.White, true, ct.Black, false)
	fmt.Printf("REQUEST: ")
	ct.ChangeColor(ct.White, true, ct.Green, false)
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)
	switch r.Method {
	case "GET":
		handleGet(w, r)
	case "POST":
		handlePost(w, r)
	default:
		ct.ChangeColor(ct.Red, true, ct.Black, false)
		fmt.Println("Unhandled")
	}
	fmt.Print("\n\n")
}

func main() {
	port := flag.Int("port", 80, "Port for listen.")
	path := flag.String("path", "/", "Service path")
	flag.BoolVar(&showDetail, "detail", false, "Show detail info")
	flag.Parse()

	http.HandleFunc(*path, requestHandler)
	fmt.Printf("Starting service [")
	ct.ChangeColor(ct.Blue, false, ct.Yellow, true)
	fmt.Printf(*path)
	ct.ChangeColor(ct.White, false, ct.Black, false)
	fmt.Printf("] at port [")
	ct.ChangeColor(ct.Blue, false, ct.Yellow, true)
	fmt.Printf("%d", *port)
	ct.ChangeColor(ct.White, false, ct.Black, false)
	fmt.Printf("]\n")

	err := http.ListenAndServe(":"+strconv.Itoa(*port), nil)

	if err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}
