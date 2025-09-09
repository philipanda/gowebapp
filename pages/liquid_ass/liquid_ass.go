package liquid_ass

import (
	"GoWebApp/util"
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {
	tpl := util.GetTemplate("liquid_ass.html")
	data := struct {
	}{}

	util.CheckErrLog(tpl.Execute(w, data))
}
