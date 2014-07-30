package api

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"image"
	// "image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/chai2010/webp"
	"github.com/nfnt/resize"

	"github.com/kamoljan/sushiobrol/conf"
	"github.com/kamoljan/sushiobrol/json"
)

type IConf struct {
	Image   image.Image
	Machine string
	Format  string
	Density string
	Ui      string
	Hash    string
	Color   string
	Width   uint
	Height  uint
	Fid     string
	Path    string
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
	var iConf IConf
	iConf.Machine = conf.Image.Machine
	iConf.Image = img
	iConf.Hash = fmt.Sprintf("%x", sha1.Sum(buf.Bytes()))
	setColor(&iConf)
	for _, format := range conf.Image.Format { // jpeg, webp, ...
		for _, screen := range conf.Image.Screen {
			iConf.Format = format
			iConf.Ui = screen.Ui
			iConf.Density = screen.Density
			iConf.Width = screen.Width
			if imgToFile(&iConf); err != nil {
				w.Write(json.Message("ERROR", "Unable to create a file"))
				return
			}
		}
	}
}

func setColor(ic *IConf) {
	img1x1 := resize.Resize(1, 1, ic.Image, resize.NearestNeighbor)
	red, green, blue, _ := img1x1.At(0, 0).RGBA()
	ic.Color = fmt.Sprintf("%X%X%X", red>>8, green>>8, blue>>8) // removing 1 byte 9A16->9A
}

func imgToFile(ic *IConf) error {
	img := resize.Resize(ic.Width, 0, ic.Image, resize.NearestNeighbor)
	ic.Height = uint(img.Bounds().Size().Y)
	ic.Fid = fmt.Sprintf("%s-%s-%s-%s-%s-%s-%d-%d", ic.Machine, ic.Format, ic.Density, ic.Ui, ic.Hash, ic.Color, ic.Width, ic.Height)
	dir := fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s", conf.SushiobrolStore, ic.Machine, ic.Format, ic.Density, ic.Ui, ic.Hash[0:2], ic.Hash[2:4])
	ic.Path = fmt.Sprintf("%s/%s", dir, ic.Fid)
	out, err := os.Create(ic.Path)
	defer out.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	if webp.Encode(out, img, &webp.Options{conf.Lossless, conf.Quality}); err != nil {
		fmt.Println("ERROR: Unable to Encode into webp")
		return err
	}
	logIConf(ic)
	return err
}

func logIConf(ic *IConf) {
	fmt.Println("Path =", ic.Path)
}
