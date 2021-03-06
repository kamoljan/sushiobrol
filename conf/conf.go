package conf

import (
	"encoding/json"
	"log"
	"os"
)

const (
	InputType       = "jpeg"
	SushiobrolId    = 1
	SushiobrolStore = "var/sushiobrol/store"
	SushiobrolPort  = 9090
	CacheMaxAge     = 30 * 24 * 60 * 60 // 30 days
	Mime            = "image/jpeg"
	Lossless        = true
	Quality         = 80
)

var Image ImageConf

type ImageConf struct {
	Machine string    `json:"machine"`
	Format  []string  `json:"format"`
	Hash    string    `json:"hash"`
	Color   string    `json:"color"`
	Screen  []Density `json:"screen"`
}

type Density struct {
	Density string `json:"density"`
	Ui      string `json:"ui"`
	Width   uint   `json:"width"`
	Height  uint   `json:"height"`
}

func initImageConf() {
	file, _ := os.Open("conf/image.json")
	decoder := json.NewDecoder(file)
	Image = ImageConf{}
	err := decoder.Decode(&Image)
	if err != nil {
		log.Fatal("Was not able to decode conf file", err)
	}
}

func init() {
	initImageConf()
}
