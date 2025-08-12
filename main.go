package main

import (
	"GoWebApp/pages/index"
	"GoWebApp/pages/wallOfShame"
	"GoWebApp/util"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	ex, _ := os.Executable()
	util.RootDirectory = filepath.Dir(ex)

	util.Log("Server Starts")

	fs := http.FileServer(http.Dir(util.GetStaticFilesPath()))
	http.Handle("/styles/", fs)
	http.Handle("/data/", fs)
	http.Handle("/images/", fs)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		util.Log("Request for / from " + r.RemoteAddr)
		index.Page(w, r)
	})

	http.HandleFunc("/wallOfShame", func(w http.ResponseWriter, r *http.Request) {
		util.Log("Request for /wallOfShame from " + r.RemoteAddr)
		wallOfShame.Page(w, r)
	})

	util.Log("Endpoints configured")

	go wallOfShame.WallOfShameUpdater()
	util.Log("Goroutines launched")

	util.Log("All set!")
	util.LogFatal(http.ListenAndServe(":80", nil))
	//logFatal(http.ListenAndServeTLS(":9990", "/etc/letsencrypt/live/philipanda.top/fullchain.pem", "/etc/letsencrypt/live/philipanda.top/privkey.pem", nil))
}
