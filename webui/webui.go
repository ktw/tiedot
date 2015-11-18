package webui

import (
	_"html/template"
	"net/http"
	"tiedot/tdlog"
)

var WebUI string

func RegisterWebUI() {
	if WebUI == "" || WebUI == "none" || WebUI == "no" || WebUI == "false" {
		tdlog.Noticef("Web control panel is disabled on your request")
		return
	}
	http.HandleFunc("/"+WebUI, handleWebUI)
	tdlog.Noticef("Web control panel is accessible at /%s", WebUI)
}

func handleWebUI(w http.ResponseWriter, r *http.Request) {
//	w.Write("Hey!")
}