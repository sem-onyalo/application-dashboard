package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/sem-onyalo/application-dashboard/service/response"

	"github.com/sem-onyalo/application-dashboard/service"
	"github.com/sem-onyalo/application-dashboard/service/request"
)

// TODO: move to config service
const (
	webAPIVersion               = "0.1"
	defaultWebAppPort           = 8080
	webAppPortConfigKey         = "APPDASH_WEBAPP_PORT"
	webAppTemplatesDirConfigKey = "APPDASH_WEBAPP_TEMPLATES_DIR"
)

// App is a service for performing web application operations
type App struct {
	Config    service.Config
	Port      int
	templates *template.Template
}

// apiInfo represents the web api info
type apiInfo struct {
	Name    string
	Version string
}

// NewApp returns a reference to the web application service
func NewApp(config service.Config) (*App, error) {
	var app *App
	var err error
	var port = defaultWebAppPort

	appPortConfig := config.GetValue(request.GetConfigValue{Key: webAppPortConfigKey})
	if appPortConfig.Value != "" {
		port, err = strconv.Atoi(appPortConfig.Value)
		if err != nil {
			return app, err
		}
	}

	templatesDirConfig := config.GetValue(request.GetConfigValue{Key: webAppTemplatesDirConfigKey})
	if templatesDirConfig.Value == "" {
		return app, errors.New("Config not found: web app templates directory ")
	}

	// TODO: get root html file from config APPDASH_WEBAPP_ROOTAPPFILE
	templates := template.Must(template.ParseFiles(templatesDirConfig.Value + "\\root.htm"))

	app = &App{Config: config, Port: port, templates: templates}
	return app, nil
}

// Start runs the web application
func (a App) Start(request request.StartApp) response.StartApp {
	var port = a.Port
	if request.Port > 0 {
		port = request.Port
	}

	http.HandleFunc("/", a.rootHandler)
	http.HandleFunc(fmt.Sprintf("/api/%s", webAPIVersion), a.rootAPIHandler)
	srv := &http.Server{Addr: fmt.Sprintf(":%d", port)}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			// TODO: send to log service
			fmt.Printf("Web app listen and serve error: %s", err)
		}
	}()

	return response.StartApp{Server: srv}
}

// rootHandler is the http handler for the root path
func (a App) rootHandler(w http.ResponseWriter, r *http.Request) {
	err := a.templates.ExecuteTemplate(w, "root.htm", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// rootAPIHandler is the http handler for the api root path
func (a App) rootAPIHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(apiInfo{Name: "SysDash", Version: webAPIVersion})
}
