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

/*
{
	"machine": "0001",
	"density":["xlarge", "large", "medium"],
	"ui": ["list", "view"]
}
*/
type ImageConf struct {
	Machine string
	Density []string
	UI      []string
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
