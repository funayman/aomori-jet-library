package controller

import "github.com/funayman/aomori-library/router"

func init() {
	router.RouteStatic("/js/", "www/js")
	router.RouteStatic("/css/", "www/css")
	router.RouteStatic("/assets/", "www/assets")
	// router.Instance().Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("www/js/"))))
	// router.Instance().PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("www/js/"))))
}
