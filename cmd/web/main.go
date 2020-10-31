package main

import (
	"os"
	"flag"
	
	"log"
	"net/http"
	
)


func main() {
	addr:=flag.String("addr","8000","HTTP network address")
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer:=http.FileServer(http.Dir("./ui/static/"))

	infoLog:=log.New(os.Stdout,"INFO/t",log.Ldate|log.Ltime)
	errorLog:=log.New(os.Stderr,"ERROR/t",log.Ldate|log.Ltime|log.Lshortfile)

	mux.Handle("/static/",http.StripPrefix("/static",fileServer))
	infoLog.Printf("Starting server on %s",*addr)
	err := http.ListenAndServe(":8000", mux)
	errorLog.Fatal(err)
}
