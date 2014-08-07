package api

import (
	"fmt"
	"github.com/kamoljan/sushiobrol/conf"
	"html"
	"log"
	"net/http"
	"strings"
)

/*
 * http://localhost:9090/fid/
 * 0001-webp-xlarge-view-da39a3ee5e6b4b0d3255bfef95601890afd80709-EBEBEB-640-871
 */
func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", conf.Mime)
	w.Header().Set("Cache-Control", fmt.Sprintf("max-age=%d", conf.CacheMaxAge))
	fid := html.EscapeString(r.URL.Path[5:]) // cut "/fid/"
	path := parsePath(fid)
	log.Println("GET: fid = " + fid)
	log.Println("GET: path = " + path)
	http.ServeFile(w, r, path)
}

func parsePath(fid string) string {
	a := strings.Split(fid, "-")
	return fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s", conf.SushiobrolStore, a[0], a[1], a[2], a[3], a[4][:2], a[4][2:4], fid)
}
