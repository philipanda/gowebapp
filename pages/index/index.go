package index

import (
	"GoWebApp/util"
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {
	tpl := util.GetTemplate("index.gohtml")
	data := struct {
	}{}

	util.CheckErrLog(tpl.Execute(w, data))
}
