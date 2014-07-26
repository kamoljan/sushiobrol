package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kamoljan/sushiobrol/api"
	"github.com/kamoljan/sushiobrol/conf"
)

func initStore(path string) {
	log.Println("Initializing data store...")
	for i := 0; i < 256; i++ {
		for x := 0; x < 256; x++ {
			err := os.MkdirAll(fmt.Sprintf("%s/%02x/%02x", path, i, x), 0755)
			if err != nil {
				log.Fatal("Was not able to create dirs ", err)
			}
		}
	}
	log.Println("...Done") // total 65536 directories
}

func logHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		log.Println(req.URL)
		h.ServeHTTP(rw, req)
	})
}

func main() {
	initStore(conf.SushiobrolStore)
	http.HandleFunc("/", api.Put)
	// http.HandleFunc("/noimageprocess/", api.PutNoImageProcess)
	// http.HandleFunc("/egg/", api.Get)
	err := http.ListenAndServe(fmt.Sprintf(":%d", conf.SushiobrolPort), logHandler(http.DefaultServeMux))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
