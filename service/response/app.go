package response

import "net/http"

// StartApp represents the response to a start web app request
type StartApp struct {
	Server *http.Server
}
