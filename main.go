package main

import (
	"fmt"
	"net/http"
  "GoWebApp/pages/index"
  "GoWebApp/pages/wallOfShame"
  "GoWebApp/util"
)

func main() {
  util.Log("Serwer startuje")

  fs := http.FileServer(http.Dir(util.AppRootDirectory + "static/"))

  http.Handle("/styles/", fs)
  http.Handle("/data/", fs)

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Request for / from " + r.RemoteAddr)
    index.Page(w, r)
  })

  http.HandleFunc("/wallOfShame", func(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Request for /wallOfShame from " + r.RemoteAddr)
    wallOfShame.Page(w, r)
  })

  util.Log("Endpointy ustawione")
  
  go wallOfShame.WallOfShameUpdater()
  util.Log("Gorutyny odpalone")

  util.Log("Startujemy!")
  util.LogFatal(http.ListenAndServe(":9990", nil));
  //logFatal(http.ListenAndServeTLS(":9990", "/etc/letsencrypt/live/philipanda.top/fullchain.pem", "/etc/letsencrypt/live/philipanda.top/privkey.pem", nil))
}

