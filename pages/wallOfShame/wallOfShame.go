package wallOfShame

import (
	"os/exec"
	"time"
	"strings"
	"regexp"
	"net/http"
	"GoWebApp/util"
)

type WallOfShameRecord struct {
	User string
	Method string
	Ip string
	Time string
}


func WallOfShameUpdater() {
	path := util.GetDataPath("wallOfShame.txt")
	for {
	  exec.Command("bash", "-c", "lastb -w > " + path).Run()
	  time.Sleep(time.Second * 10);
	}
}

func GetWallOfShame() []WallOfShameRecord {
	wosBytes := util.GetData("wallOfShame.txt")
	wos := strings.Split(string(wosBytes), "\n")

	reg := regexp.MustCompile(`(?P<user>[^ ]*) [ ]*(?P<method>[^ ]*) [ ]*(?P<ip>[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}) [ ]*(?P<time>.*)`)
	var records []WallOfShameRecord

	for _, row := range wos {
	  matches := reg.FindAllStringSubmatch(row, -1)
	  for _, match := range matches {    
		var record WallOfShameRecord
  
		i := reg.SubexpIndex("user")
		if i > 0 {
		  record.User = match[i]
		}
		i = reg.SubexpIndex("method")
		if i > 0 {
		  record.Method = match[i]
		}
		i = reg.SubexpIndex("ip")
		if i > 0 {
		  record.Ip = match[i]
		}
		i = reg.SubexpIndex("time")
		if i > 0 {
		  record.Time = match[i]
		}
  
		records = append(records, record)
	  }
	}
	return records
}

func Page(w http.ResponseWriter, r *http.Request) {
	tpl := util.GetTemplate("wallOfShame.html")
	records := GetWallOfShame()

	data := struct {
	  WallOfShame []WallOfShameRecord
	}{
	  WallOfShame: records,
	}
  
	err := tpl.Execute(w, data)
	util.CheckErrPanic(err)
}
  