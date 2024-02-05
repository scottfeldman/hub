package hub

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/merliot/device"
)

func (h *Hub) apiCreate(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	model := r.URL.Query().Get("model")
	name := r.URL.Query().Get("name")

	thinger, err := h.server.CreateThing(id, model, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	child := thinger.(device.Devicer)
	child.CopyWifiAuth(h.WifiAuth)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Child id '%s' created", id)
}

func (h *Hub) apiDelete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := h.server.DeleteThing(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Child id '%s' deleted", id)
}

func (h *Hub) apiSave(w http.ResponseWriter, r *http.Request) {
	if err := h.saveChildren(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Children saved")
}

func (h *Hub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch strings.TrimPrefix(r.URL.Path, "/") {
	case "state":
		device.ShowState(h.templates, w, h)
	case "create":
		h.apiCreate(w, r)
	case "delete":
		h.apiDelete(w, r)
	case "save":
		h.apiSave(w, r)
	case "models":
		device.RenderTemplate(h.templates, w, "models.tmpl", h)
	default:
		h.API(h.templates, w, r)
	}
}