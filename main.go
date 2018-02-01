package main

import (
	"fmt"

	"github.com/sem-onyalo/application-dashboard/core/interactor"
	"github.com/sem-onyalo/application-dashboard/service/request"
	"github.com/sem-onyalo/application-dashboard/web"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	configService := interactor.NewConfig()

	databaseService, err := interactor.NewDatabase(configService)
	if err != nil {
		fmt.Printf("Create database service failed: %s\n", err)
		return
	}

	associationService := interactor.NewAssociation(databaseService)
	resolutionService := interactor.NewResolution(databaseService)
	incidentService := interactor.NewIncident(databaseService, resolutionService)
	endpointService := interactor.NewEndpoint(databaseService, incidentService)

	appService, err := web.NewApp(associationService, configService, endpointService, incidentService)
	if err != nil {
		// TODO: send to log service
		fmt.Printf("Create web app service failed: %s\n", err)
		return
	}

	app := appService.Start(request.StartApp{})
	// TODO: send to log service
	fmt.Printf("App server listening on: %s\n", app.Server.Addr)

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
}
