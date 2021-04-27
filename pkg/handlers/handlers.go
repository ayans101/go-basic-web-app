package handlers

import (
	"fmt"
	"net/http"

	"github.com/ayans101/go-basic-web-app/pkg/config"
	"github.com/ayans101/go-basic-web-app/pkg/models"
	"github.com/ayans101/go-basic-web-app/pkg/render"
)

//	Repo the repository used by the handlers
var Repo *Repository

//	Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//	NewRepo creates a repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//	NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//	Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	fmt.Println(m.App.Session)
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	fmt.Println(m.App.Session)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

//	About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//	perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	//	send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
