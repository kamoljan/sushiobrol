package api

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"os"

	// "github.com/chai2010/webp"

	"github.com/kamoljan/sushiobrol/conf"
	"github.com/kamoljan/sushiobrol/json"
)

/*
{
	"machine": "0001",
	"format":["jpeg", "webp"],
	"screen":[
	    {"density": "xlarge", "ui": "view", "width": 640, "height": 543},
   	    {"density": "xlarge", "ui": "list", "width": 320, "height": 284},
   	    {"density": "large",  "ui": "view", "width": 480, "height": 320},
   	    {"density": "large",  "ui": "list", "width": 240, "height": 190},
   	    {"density": "medium", "ui": "view", "width": 320, "height": 256},
   	    {"density": "medium", "ui": "list", "width": 160, "height": 102}
    ]
}
*/
func Put(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.Write(json.Message("ERROR", "Not supported Method"))
		return
	}

	reader, err := r.MultipartReader()
	if err != nil {
		w.Write(json.Message("ERROR", "Client should support multipart/form-data"))
		return
	}

	buf := bytes.NewBufferString("")
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if part.FileName() == "" { // if empy skip this iteration
			continue
		}
		_, err = io.Copy(buf, part)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	defer r.Body.Close()

	img, _, err := image.Decode(buf)
	if err != nil {
		w.Write(json.Message("ERROR", "Unable to decode your image"))
		return
	}
	fmt.Sprintf("test/%s", buf.Bytes())

	// "machine": "0001",
	// "format":["jpeg", "webp"],
	// "screen":[
	//     {"density": "xlarge", "ui": "view", "width": 640, "height": 543},
	//     {"density": "xlarge", "ui": "list", "width": 320, "height": 284},
	//     {"density": "large",  "ui": "view", "width": 480, "height": 320},
	//     {"density": "large",  "ui": "list", "width": 240, "height": 190},
	//     {"density": "medium", "ui": "view", "width": 320, "height": 256},
	//     {"density": "medium", "ui": "list", "width": 160, "height": 102}
	// ]

	for _, format := range conf.Image.Format { // jpeg, webp, ...
		fmt.Println("conf.Image.Format =", format)
		for _, screen := range conf.Image.Screen {
			fmt.Println("Screen =", screen)
			fmt.Println("Screen.Density =", screen.Density)
			fmt.Println("Screen.Ui =", screen.Ui)
			fmt.Println("Screen.Width =", screen.Width)
			fmt.Println("Screen.Height =", screen.Height)
		}
	}

	// PATH
	path := fmt.Sprintf("test/%x", sha1.Sum(buf.Bytes()))

	fmt.Printf("conf.Image.Machine = %s", conf.Image.Machine)

	// CREATE DESTINATION FILE
	out, err := os.Create(path)
	if err != nil {
		w.Write(json.Message("ERROR", "Unable to create a file"))
		return
	}
	defer out.Close()

	err = jpeg.Encode(out, img, nil) // write image to file
	if err != nil {
		log.Println("Unable to save your image to file")
		w.Write(json.Message("ERROR", "Unable to encode into jpeg"))
		return
	}
}

/*
	0001/
	    webp/
	        origin/view/04/0d/0001-04...8b-ACA0AC-640-543
    	    origin/list/04/0d/0001-04...8b-ACA0AC-320-271
	        xlarge/view/04/0d/0001-04...8b-ACA0AC-640-543
	        xlarge/list/04/0d/0001-04...8b-ACA0AC-320-284
	        large/view/04/0d/0001-04...8b-ACA0AC-480-320
	        large/list/04/0d/0001-04...8b-ACA0AC-240-190
	        medium/view/04/0d/0001-04...8b-ACA0AC-320-256
	        medium/list/04/0d/0001-04...8b-ACA0AC-160-102
        jpeg/
*/
func genPath(file string) string {
	path := fmt.Sprintf(conf.SushiobrolStore+"%s/%s/%s", file[5:7], file[7:9], file)
	log.Println(path)
	return path
}

func genFile(eid string, color string, width, height int) string {
	file := fmt.Sprintf("%04x_%s_%s_%d_%d", conf.SushiobrolId, eid, color, width, height)
	log.Println(file)
	return file
}
