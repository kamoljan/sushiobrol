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

	// image_ext "github.com/chai2010/gopkg/image"
	_ "github.com/chai2010/gopkg/image/png"

	// "github.com/dgryski/go-webp"

	"github.com/kamoljan/sushiobrol/conf"
	"github.com/kamoljan/sushiobrol/json"
)

/*
   {
       "image": {
       	"machine": "0001",
           "hash": "8787bec619ff019fd17fe02599a384d580bf6779",
           "color": "ACA0AC",
           "width": 840,
           "height": 756,
	       "density":[
          	    {
          	    	"name": "xlarge",
          	    	"value": [
          	    	    {"ui": "view", "value": {"width": 640, "height": 543}},
          	    	    {"ui": "list", "value": {"width": 320, "height": 284}}
          	    	]
          	    },
          	    {
          	        "name": "large",
          	    	"value": [
          	    	    {"ui": "view", "value": {"width": 480, "height": 320}},
          	    	    {"ui": "list", "value": {"width": 240, "height": 190}}
          	    	]
          	    },
          	    {
          	    	"name": "medium",
          	    	"value": [
          	    	    {"ui": "view", "value": {"width": 320, "height": 256}},
          	    	    {"ui": "list", "value": {"width": 160, "height": 102}}
          	    	]
          	    }
           ]
       }
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
	        origin/04/0d/0001-04...8b-ACA0AC-640-543
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
