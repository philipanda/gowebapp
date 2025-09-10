package secret

import "net/http"

func Page(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	code := r.FormValue("secret")
	http.SetCookie(w, &http.Cookie{
		Name:     "secret",
		Value:    code,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
