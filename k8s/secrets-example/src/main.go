package main

import (
	"fmt" //to print messages to stdout
	"log" //logging :)
	"os"

	//our web server that will host the mock
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

var configuration []byte
var secret []byte

func Response(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello")
}

func Status(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "ok")
}
func getEnvs() {
	username := os.Getenv("SECRET_USERNAME")
	password := os.Getenv("SECRET_PASSWORD")
	log.Printf("Got username %s and password %s", username, password)
}
func main() {

	fmt.Println("starting secrets example...")
	getEnvs()
	router := fasthttprouter.New()
	router.GET("/", Response)
	router.GET("/status", Status)

	log.Fatal(fasthttp.ListenAndServe(":5001", router.Handler))
}
