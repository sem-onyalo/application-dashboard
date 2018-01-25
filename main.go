package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sem-onyalo/application-dashboard/service/request"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sem-onyalo/application-dashboard/core/interactor"
	"github.com/sem-onyalo/application-dashboard/web"
)

func main() {

	configService := interactor.NewConfig()

	appService, err := web.NewApp(configService)
	if err != nil {
		// TODO: send to log service
		fmt.Printf("Create web app failed: %v", err)
	}

	app := appService.Start(request.StartApp{})
	// TODO: send to log service
	fmt.Printf("app server listening on: %v", app.Server.Addr)

	var input string
	fmt.Print("> ")
	fmt.Scanln(&input)
	for input != "quit" {
		fmt.Printf("Command not supported: %s\n", input)
		fmt.Print("> ")
		fmt.Scanln(&input)
	}

	if err = app.Server.Shutdown(nil); err != nil {
		fmt.Printf("App shutdown error: %s\n", err)
	}
	return

	// --------------------------------------------------

	databaseService, err := interactor.NewDatabase(configService)
	if err != nil {
		panic(err)
	}

	endpointService := interactor.NewEndpoint(databaseService)

	response, err := endpointService.GetAll()

	if err != nil {
		fmt.Printf("Get all endpoints failed: %v", err)
		return
	}

	for _, ep := range response.Endpoints {
		timerStart := time.Now()
		// TODO: move this to an http service
		rsp, err := http.Get(ep.URL)
		timerEnd := time.Now()
		timerElapsed := timerEnd.Sub(timerStart)

		var result string
		if err != nil {
			result = "NULL"
		} else {
			result = rsp.Status
		}

		fmt.Printf("%s %s %fs\n", ep.Name, result, timerElapsed.Seconds())
	}
}
