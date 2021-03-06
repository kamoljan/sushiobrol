package api

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"github.com/chai2010/webp"
	"github.com/nfnt/resize"
	"image"
	_ "image/gif"
	"image/jpeg"
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

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write(json.Message("ERROR", "Not supported Method"))
		return
	}
	f, _, err := r.FormFile("image")
	if err != nil {
		w.Write(json.Message("ERROR", "Can't Find Image"))
		return
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	h := sha1.New()
	var result json.Result
	var ic iconf
	ic.machine = conf.Image.Machine
	if conf.InputType == "jpeg" {
		ic.image, _, err = image.Decode(io.TeeReader(buf, h))
	} else {
		ic.image, err = webp.Decode(io.TeeReader(buf, h))
	}
	ic.hash = fmt.Sprintf("%x", h.Sum(nil))
	if err != nil {
		w.Write(json.Message("ERROR", "Unable to decode your image! Type="+conf.InputType+" error:"+err.Error()))
		return
	}
	setColor(&ic)
	for _, format := range conf.Image.Format { // jpeg, webp, ...
		for _, screen := range conf.Image.Screen {
			ic.format = format
			ic.ui = screen.Ui
			ic.density = screen.Density
			ic.width = screen.Width
			if ic.fid, err = imgToFile(&ic); err != nil {
				w.Write(json.Message("ERROR", "Unable to create a file"))
				return
			}
			fid := json.Fid{fmt.Sprintf("%s_%s", screen.Density, screen.Ui), ic.fid}
			result.Image = append(result.Image, fid)
		}
	}
	w.Write(json.Message("OK", &result))
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
	if conf.InputType == "jpeg" {
		err = jpeg.Encode(out, img, nil) // write image to file
		if err != nil {
			fmt.Println("ERROR: Unable to Encode into webp")
			return "", err
		}
	} else {
		if webp.Encode(out, img, &webp.Options{conf.Lossless, conf.Quality}); err != nil {
			fmt.Println("ERROR: Unable to Encode into webp")
			return "", err
		}
	}
	return ic.fid, err
}

func setColor(ic *iconf) {
	img1x1 := resize.Resize(1, 1, ic.image, resize.NearestNeighbor)
	red, green, blue, _ := img1x1.At(0, 0).RGBA()
	ic.color = fmt.Sprintf("%X%X%X", red>>8, green>>8, blue>>8) // removing 1 byte 9A16->9A
}
