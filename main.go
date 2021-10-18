package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
)

var (
	// TEMPLATES 는 template 글로벌 변수입니다.
	// VFS 내부의 html파일을 담기 위해 사용합니다.
	TEMPLATES    = template.New("")
	flagHTTPPort = flag.String("http", "", "webservice port number")
	flagDBURI    = flag.String("dburi", "mongodb://localhost:27017", "mongoDB uri for application")
	flagDBName   = flag.String("dbname", "baeseoyoung", "mongoDB database name")
)

func main() {
	flag.Parse()

	if *flagHTTPPort != "" {
		ip, err := serviceIP()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Service start: http://%s%s\n", ip, *flagHTTPPort)
		webserver()
	} else {
		flag.Usage()
		os.Exit(1)
	}

}
