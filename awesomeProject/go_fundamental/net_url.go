package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func main() {
	input := "http://www.h-aws-dtv-use1-009-tds.cds1.yospace.com/esamserver/dispatcher/mds/[${repoll}/]receive?dId=dfwhls&cid=${streamId}&mode=0&submode=1&mt=${tag}&prefetch=180&prebid=ov_tds_config_live_v6&showmode=true&format=2"

	u, err := url.Parse(input)
	if err != nil {
		log.Fatal(err)
	}
	hostname := strings.TrimPrefix(u.Hostname(), "www.")
	//hostname := u.Hostname()
	fmt.Println("hostname = ", hostname)
	fmt.Println("u.RawQuery = ", u.RawQuery)

	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println("dId = ", m["dId"][0])
	fmt.Println("cid = ", m["cid"][0])
	fmt.Println("prefetch = ", m["prefetch"][0])
}
