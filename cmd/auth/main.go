package main

import (
	"log"
	"net/http"
	"os"

	"github.com/codex-veritas/oauth2/pkg/oauth2"
)

func main() {
	l := log.New(os.Stdout, "MAIN  ", log.LstdFlags|log.Lshortfile)

	oauthServer := &oauth2.Server{}

	mux := http.NewServeMux()
	mux.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			l.Println("fail to parse form:", err)
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		param := oauth2.AuthorizeParam{}
		param.ResponseType = r.Form.Get("response_type")
		if param.ResponseType == "" {
			http.Error(w, "missing parameter: response_type", http.StatusBadRequest)
			return
		}
		param.ClientID = r.Form.Get("client_id")
		if param.ClientID == "" {
			http.Error(w, "missing parameter: client_id", http.StatusBadRequest)
			return
		}
		param.RedirectURI = r.Form.Get("redirect_uri")
		if param.RedirectURI == "" {
			http.Error(w, "missing parameter: redirect_uri", http.StatusBadRequest)
			return
		}
		param.Scope = r.Form.Get("scope")
		if param.Scope == "" {
			http.Error(w, "missing parameter: scope", http.StatusBadRequest)
			return
		}
		param.State = r.Form.Get("state")
		if param.State == "" {
			http.Error(w, "missing parameter: state", http.StatusBadRequest)
			return
		}

		uri, err := oauthServer.Authorize(param)

		if _, ok := err.(oauth2.ErrNotImplemented); ok {
			http.Error(w, err.Error(), http.StatusNotImplemented)
			return
		}
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("location", uri)
		w.WriteHeader(http.StatusFound)
	})

	mux.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		// Token are only use for server app
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	l.Println("listening on :8080")
	l.Fatal(srv.ListenAndServe())
}
