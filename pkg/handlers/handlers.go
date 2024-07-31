package handlers

import (
	"fmt"
	"github.com/malpania/beerproj/pkg/config"
	"github.com/malpania/beerproj/pkg/models"
	"github.com/malpania/beerproj/pkg/render"
	"net/http"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

// NewRepository creates a repository variable
func NewRepository(app *config.AppConfig) *Repository {
	return &Repository{App: app}
}

// NewHandlers sets the repository for the handlers.
func NewHandlers(repo *Repository) {
	Repo = repo
}

func (m *Repository) Home(writer http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Value Again"
	remoteIP := request.RemoteAddr
	session := m.App.Session
	session.Put(request.Context(), "remote_ip", remoteIP)
	session.Put(request.Context(), "abc", stringMap["test"])
	render.RenderTemplate(writer, "home.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
func (m *Repository) About(writer http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello About Again"
	abc := m.App.Session.GetString(request.Context(), "abc")
	fmt.Println("abc", abc)
	remoteIP := m.App.Session.GetString(request.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	stringMap["abc"] = abc
	render.RenderTemplate(writer, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
