package liquid_ass

import (
	"GoWebApp/util"
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {
	tpl := util.GetTemplate("liquid_ass.gohtml")
	cookie, err := r.Cookie("secret")

	data := struct {
		SecretAss bool
	}{
		SecretAss: err == nil && cookie.Value == "ass",
	}

	util.CheckErrLog(tpl.Execute(w, data))
}
