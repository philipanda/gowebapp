package main

import (
	"GoWebApp/pages/guestbook"
	"GoWebApp/pages/index"
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

	http.HandleFunc("/guestbook", func(w http.ResponseWriter, r *http.Request) {
		util.Log("Request for /guestbook from " + r.RemoteAddr)
		guestbook.Page(w, r)
	})

	util.Log("Endpoints configured")
	util.Log("All set!")
	util.LogFatal(http.ListenAndServe(":80", nil))
	//logFatal(http.ListenAndServeTLS(":9990", "/etc/letsencrypt/live/philipanda.top/fullchain.pem", "/etc/letsencrypt/live/philipanda.top/privkey.pem", nil))
}
