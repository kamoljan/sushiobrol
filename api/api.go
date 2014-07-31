package api

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"github.com/chai2010/webp"
	"github.com/nfnt/resize"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/kamoljan/sushiobrol/conf"
	"github.com/kamoljan/sushiobrol/json"
)

type iconf struct {
	image                                                image.Image
	machine, format, density, ui, hash, color, fid, path string
	width, height                                        uint
}

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

	var result json.Result
	var iconf iconf
	iconf.machine = conf.Image.Machine
	iconf.image = img
	iconf.hash = fmt.Sprintf("%x", sha1.Sum(buf.Bytes()))
	setColor(&iconf)
	for _, format := range conf.Image.Format { // jpeg, webp, ...
		for _, screen := range conf.Image.Screen {
			iconf.format = format
			iconf.ui = screen.Ui
			iconf.density = screen.Density
			iconf.width = screen.Width
			iconf.fid, err = imgToFile(&iconf)
			if err != nil {
				w.Write(json.Message("ERROR", "Unable to create a file"))
				return
			}
			fid := json.Fid{fmt.Sprintf("%s_%s", screen.Density, screen.Ui), iconf.fid}
			result.Image = append(result.Image, fid)
		}
	}
	w.Write(json.Message("OK", &result))
}

func setColor(ic *iconf) {
	img1x1 := resize.Resize(1, 1, ic.image, resize.NearestNeighbor)
	red, green, blue, _ := img1x1.At(0, 0).RGBA()
	ic.color = fmt.Sprintf("%X%X%X", red>>8, green>>8, blue>>8) // removing 1 byte 9A16->9A
}

func imgToFile(ic *iconf) (string, error) {
	img := resize.Resize(ic.width, 0, ic.image, resize.NearestNeighbor)
	ic.height = uint(img.Bounds().Size().Y)
	ic.fid = fmt.Sprintf("%s-%s-%s-%s-%s-%s-%d-%d", ic.machine, ic.format, ic.density, ic.ui, ic.hash, ic.color, ic.width, ic.height)
	dir := fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s", conf.SushiobrolStore, ic.machine, ic.format, ic.density, ic.ui, ic.hash[0:2], ic.hash[2:4])
	ic.path = fmt.Sprintf("%s/%s", dir, ic.fid)
	out, err := os.Create(ic.path)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer out.Close()
	if webp.Encode(out, img, &webp.Options{conf.Lossless, conf.Quality}); err != nil {
		fmt.Println("ERROR: Unable to Encode into webp")
		return "", err
	}
	return ic.fid, err
}
