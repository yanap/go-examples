package main

import (
	"net/url"
)

func main() {
	timezone := "Asia/Tokyo"
	es := url.PathEscape(timezone)
	println(es)
	var loc = url.PathEscape("Asia/Tokyo")
	println(loc)
}
