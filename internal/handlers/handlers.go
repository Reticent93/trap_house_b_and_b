package handlers

import (
	"encoding/json"
	"fmt"
	config2 "github.com/Reticent93/trap_house_b_and_b/internal/config"
	models2 "github.com/Reticent93/trap_house_b_and_b/internal/models"
	render2 "github.com/Reticent93/trap_house_b_and_b/internal/render"
	"log"
	"net/http"
)

//Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config2.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config2.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render2.RenderTemplate(w, r, "home.page.tmpl", &models2.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	//send data to template
	render2.RenderTemplate(w, r, "about.page.tmpl", &models2.TemplateData{
		StringMap: stringMap,
	})

}

//Reservation renders the make a room page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render2.RenderTemplate(w, r, "make-reservation.page.tmpl", &models2.TemplateData{})
}

//Dons renders the make a reservation page
func (m *Repository) Dons(w http.ResponseWriter, r *http.Request) {
	render2.RenderTemplate(w, r, "dons.page.tmpl", &models2.TemplateData{})
}

//Bastones renders the make a room page and displays form
func (m *Repository) Bastones(w http.ResponseWriter, r *http.Request) {
	render2.RenderTemplate(w, r, "bastones.page.tmpl", &models2.TemplateData{})
}

//Availability renders the make a availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render2.RenderTemplate(w, r, "search.avail.page.tmpl", &models2.TemplateData{})
}

//PostAvailability renders the make a availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("start date is %s, end date is %s", start, end)))

	w.Write([]byte("Posted to search availability"))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

//AvailabilityJSON handles request for availability and send JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      false,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "     ")

	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(out)
}

//Contact renders the make a availability page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render2.RenderTemplate(w, r, "contact.page.tmpl", &models2.TemplateData{})
}
