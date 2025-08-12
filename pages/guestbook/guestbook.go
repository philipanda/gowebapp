package guestbook

import (
	"GoWebApp/util"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"
)

type GuestbookRecord struct {
	Ip        string `json:"ip"`
	Time      string `json:"time"`
	Uri       string `json:"uri"`
	Useragent string `json:"useragent"`
}

func guestbook_path() string {
	return util.GetStaticFilesPath() + "/guestbook.csv"
}

func GetGuestbook() []GuestbookRecord {
	bytes, err := os.ReadFile(guestbook_path())
	util.CheckErrLog(err)
	var guests []GuestbookRecord
	json.Unmarshal(bytes, &guests)
	util.CheckErrLog(err)
	return guests
}

func RegisterGuest(entry GuestbookRecord) {
	guests := GetGuestbook()
	if strings.Contains(entry.Ip, "192.168") {
		util.Log("Guestbook: Access from local network. Not registering.")
		return
	}
	guests = append(guests, entry)
	bytes, err := json.Marshal(guests)
	if util.CheckErrLog(err) {
		return
	}
	err = os.WriteFile(guestbook_path(), bytes, os.FileMode(os.O_RDWR))
	util.CheckErrLog(err)
}
func Page(w http.ResponseWriter, r *http.Request) {
	tpl := util.GetTemplate("guestbook.html")
	records := GetGuestbook()

	data := struct {
		Guestbook []GuestbookRecord
	}{
		Guestbook: records,
	}
	util.Log(r.Header)
	RegisterGuest(GuestbookRecord{Ip: r.Header.Get("X-Forwarded-For"), Time: time.Now().Format(time.ANSIC), Uri: r.RequestURI, Useragent: r.Header.Get("User-Agent")})
	err := tpl.Execute(w, data)
	util.CheckErrPanic(err)
}
