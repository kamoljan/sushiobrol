package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kamoljan/sushiobrol/api"
	"github.com/kamoljan/sushiobrol/conf"
)

func initStoreDirs() {
	fmt.Println("Initializing data store...")
	for _, format := range conf.Image.Format {
		for _, screen := range conf.Image.Screen {
			dir := fmt.Sprintf("%s/%s/%s/%s/%s", conf.SushiobrolStore, conf.Image.Machine, format, screen.Density, screen.Ui)
			fmt.Print("dir = ", dir)
			if err := os.MkdirAll(dir, 0755); err != nil {
				log.Fatal("Was not able to create dirs ", err)
			}
			mkSubDirs(dir)
		}
	}
	fmt.Println("...Done")
}

func mkSubDirs(path string) {
	for i := 0; i < 256; i++ {
		for x := 0; x < 256; x++ {
			err := os.MkdirAll(fmt.Sprintf("%s/%02x/%02x", path, i, x), 0755)
			if err != nil {
				log.Fatal("Was not able to create dirs ", err)
			}
		}
	}
	fmt.Println("...65536 dirs are created")
}

func logHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		log.Println(req.URL)
		h.ServeHTTP(rw, req)
	})
}

func main() {
	initStoreDirs()
	http.HandleFunc("/", api.Post)
	// http.HandleFunc("/noimageprocess/", api.PutNoImageProcess)
	http.HandleFunc("/fid/", api.Get)
	err := http.ListenAndServe(fmt.Sprintf(":%d", conf.SushiobrolPort), logHandler(http.DefaultServeMux))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
