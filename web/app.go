package web

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/sem-onyalo/application-dashboard/service"
	"github.com/sem-onyalo/application-dashboard/service/request"
	"github.com/sem-onyalo/application-dashboard/service/response"
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
	Association service.Association
	Config      service.Config
	Endpoint    service.Endpoint
	Port        int
	templates   *template.Template
}

// apiInfo represents the web api info
type apiInfo struct {
	Name    string
	Version string
}

// NewApp returns a reference to the web application service
func NewApp(association service.Association, config service.Config, endpoint service.Endpoint) (*App, error) {
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

	app = &App{
		Association: association,
		Config:      config,
		Endpoint:    endpoint,
		Port:        port,
		templates:   templates,
	}
	return app, nil
}

// Start runs the web application
func (a App) Start(request request.StartApp) response.StartApp {
	var port = a.Port
	if request.Port > 0 {
		port = request.Port
	}

	a.setRoutes()
	srv := &http.Server{Addr: fmt.Sprintf(":%d", port)}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			// TODO: send to log service
			fmt.Printf("Web app listen and serve error: %s", err)
		}
	}()

	return response.StartApp{Server: srv}
}

// setRoutes sets the HTTP API routes
func (a App) setRoutes() {
	http.HandleFunc("/", a.rootHandler)
	http.HandleFunc("/assets/", a.assetHandler)
	http.HandleFunc(fmt.Sprintf("/api/v%s", webAPIVersion), a.rootAPIHandler)
	http.HandleFunc(fmt.Sprintf("/api/v%s/associations", webAPIVersion), a.associationsHandler)
	http.HandleFunc(fmt.Sprintf("/api/v%s/endpoints", webAPIVersion), a.endpointsHandler)
	http.HandleFunc(fmt.Sprintf("/api/v%s/endpoints/incidences", webAPIVersion), a.endpointsIncidencesHandler)
	http.HandleFunc(fmt.Sprintf("/api/v%s/endpoints/tests", webAPIVersion), a.endpointsTestsHandler)
	http.HandleFunc(fmt.Sprintf("/api/v%s/endpoint-tests", webAPIVersion), a.endpointTestsHandler)
}

// apiResponseHandler handles the response for successful requests
func (a App) apiResponseHandler(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

// rootHandler is the http handler for the root path
func (a App) rootHandler(w http.ResponseWriter, r *http.Request) {
	err := a.templates.ExecuteTemplate(w, "root.htm", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// assetHandler handles static resource files
func (a App) assetHandler(w http.ResponseWriter, r *http.Request) {
	var contentType string
	if strings.HasSuffix(r.URL.Path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(r.URL.Path, ".js") {
		contentType = "text/javascript"
	} else if strings.HasSuffix(r.URL.Path, ".svg") {
		contentType = "image/svg+xml"
	} else {
		contentType = "text/plain"
	}

	templatesDirConfig := a.Config.GetValue(request.GetConfigValue{Key: webAppTemplatesDirConfigKey})
	file, err := os.Open(templatesDirConfig.Value + r.URL.Path)
	if err != nil {
		// TODO: send to log service
		fmt.Printf("Error opening url path %s: %s\n", r.URL.Path, err)
		w.WriteHeader(404)
	} else {
		defer file.Close()
		w.Header().Set("Content-Type", contentType)
		reader := bufio.NewReader(file)
		reader.WriteTo(w)
	}
}

// rootAPIHandler is the http handler for the api root path
func (a App) rootAPIHandler(w http.ResponseWriter, r *http.Request) {
	a.apiResponseHandler(w, apiInfo{Name: "SysDash", Version: webAPIVersion})
}

// assocaitionsHandler handles association requests
func (a App) associationsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		a.associationsPostHandler(w, r)
	default:
		http.Error(w, "Unsupported HTTP method for path associations/", http.StatusBadRequest)
	}
}

// associationsPostHandler handles association POST requests
func (a App) associationsPostHandler(w http.ResponseWriter, r *http.Request) {
	var serviceRequest request.CreateAssociation
	json.NewDecoder(r.Body).Decode(&serviceRequest)
	serviceResponse, err := a.Association.Create(serviceRequest)
	if err != nil {
		// TODO: also send to log service with err
		http.Error(w, "Create association failed", http.StatusInternalServerError)
	}
	a.apiResponseHandler(w, serviceResponse)
}

// endpointsHandler handles endpoint requests
func (a App) endpointsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.endpointsGetHandler(w, r)
	default:
		http.Error(w, "Unsupported HTTP method for path endpoints/", http.StatusBadRequest)
	}
}

// endpointsIncidencesHandler handler endpoint incidences requests
func (a App) endpointsIncidencesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		a.endpointsIncidencesPostHandler(w, r)
	default:
		http.Error(w, "Unsupported HTTP method for path endpoints/incidences", http.StatusBadRequest)
	}
}

// endpointsIncidencesPostHandler handler endpoint incidences POST requests
func (a App) endpointsIncidencesPostHandler(w http.ResponseWriter, r *http.Request) {
	var serviceRequest request.CreateEndpointIncident
	json.NewDecoder(r.Body).Decode(&serviceRequest)
	serviceResponse, err := a.Endpoint.CreateIncident(serviceRequest)
	if err != nil {
		// TODO: also send to log service with err
		http.Error(w, "Create endpoint incident failed", http.StatusInternalServerError)
	}
	a.apiResponseHandler(w, serviceResponse)
}

// endpointsTestsHandler handler endpoint test requests
func (a App) endpointsTestsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		testAllEndpoints, err := a.Endpoint.TestAll()
		if err != nil {
			// TODO: also send to log service with err
			http.Error(w, "Test all endpoints request failed", http.StatusInternalServerError)
		}
		a.apiResponseHandler(w, testAllEndpoints.EndpointTests)
	default:
		http.Error(w, "Unsupported HTTP method for path endpoints/tests/", http.StatusBadRequest)
	}
}

// endpointsGetHandler handles GET endpoint requests
func (a App) endpointsGetHandler(w http.ResponseWriter, r *http.Request) {
	getAllEndpoints, err := a.Endpoint.GetAll()
	if err != nil {
		// TODO: also send to log service with err
		http.Error(w, "Get all endpoints failed", http.StatusInternalServerError)
	}

	a.apiResponseHandler(w, getAllEndpoints.Endpoints)
}

// endpointTestsHandler handles endpoint-test requests
func (a App) endpointTestsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getEndpointTests, err := a.Endpoint.GetTests()
		if err != nil {
			// TODO: also send to log service with err
			http.Error(w, "Get endpoint tests request failed", http.StatusInternalServerError)
		}
		a.apiResponseHandler(w, getEndpointTests)
	default:
		http.Error(w, "Unsupported HTTP method for path endpoint-tests/", http.StatusBadRequest)
	}
}
