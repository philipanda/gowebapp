package index

import (
	"net/http"
	"GoWebApp/util"
)


func Page(w http.ResponseWriter, r *http.Request) {
	tpl := util.GetTemplate("index.html")
	data := struct {
	  }{
	  }
	
	util.CheckErrLog(tpl.Execute(w, data))
}