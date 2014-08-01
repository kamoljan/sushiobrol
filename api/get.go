package api

import (
	"fmt"
	"github.com/kamoljan/sushiobrol/conf"
	"html"
	"log"
	"net/http"
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

func parsePath(eid string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s", conf.SushiobrolStore, eid[:4], eid[5:9], eid[10:16], eid[17:21], eid[22:24], eid[24:26], eid)
}
