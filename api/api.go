package api

/*
#cgo LDFLAGS: -lwebp -L/opt/local/lib/

#include <stdlib.h>
#include <webp/encode.h>
*/
import "C"

import (
	"bytes"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"

	// _ "code.google.com/p/go.image/webp"
	// _ "code.google.com/p/vp8-go/webp"
	// "github.com/chai2010/gopkg/image/webp"
	"github.com/dgryski/go-webp"

	//	"github.com/kamoljan/sushiobrol/conf"
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

	var output []byte
	switch t := img.(type) {
	case *image.NRGBA:
		output = webp.WebPEncodeLosslessRGBA(t.Pix, t.Rect.Dx(), t.Rect.Dy(), t.Stride)
	case *image.RGBA:
		output = webp.WebPEncodeLosslessRGBA(t.Pix, t.Rect.Dx(), t.Rect.Dy(), t.Stride)
	default:
		log.Println("unknown type: ", reflect.TypeOf(img))
		w.Write(json.Message("ERROR", "Unknown type"))
		return
	}

	ioutil.WriteFile("test/gopher.webp", output, 0666)
	webp.FreeWebP(output)
}

// func convertToJPEG(w io.Writer, r io.Reader) error {
// 	img, _, err := image.Decode(r)
// 	if err != nil {
// 		w.Write(json.Message("ERROR", "Unable to decode your image"))
// 		return err
// 	}
// 	return jpeg.Encode(w, img)
// }
