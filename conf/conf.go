package conf

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const (
	SushiobrolId    = 1
	SushiobrolStore = "/var/sushiobrol/store/"
	SushiobrolPort  = 9090
	CacheMaxAge     = 30 * 24 * 60 * 60 // 30 days
	Mime            = "image/webp"
)

var Image ImageConf

// {
// 	"machine": "0001",
// 	"format":["jpeg", "webp"],
// 	"screen":[
// 	    {"density": "xlarge", "ui": "view", "width": 640, "height": 543},
//    	{"density": "xlarge", "ui": "list", "width": 320, "height": 284},
//    	{"density": "large",  "ui": "view", "width": 480, "height": 320},
//    	{"density": "large",  "ui": "list", "width": 240, "height": 190},
//    	{"density": "medium", "ui": "view", "width": 320, "height": 256},
//    	{"density": "medium", "ui": "list", "width": 160, "height": 102}
//  ]
// }
type ImageConf struct {
	Machine string    `json:"machine"`
	Format  []string  `json:"format"`
	Screen  []Density `json:"screen"`
}

type Density struct {
	Density string `json:"density"`
	Ui      string `json:"ui"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
}

func initImageConf() {
	file, _ := os.Open("conf/image.json")
	decoder := json.NewDecoder(file)
	Image = ImageConf{}
	err := decoder.Decode(&Image)
	if err != nil {
		log.Fatal("Was not able to decode conf file", err)
	}
	fmt.Println(Image.Machine)
}

func init() {
	initImageConf()
}
