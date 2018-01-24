package main

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sem-onyalo/application-dashboard/core/interactor"
)

func main() {
	configService := interactor.NewConfig()

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
		rsp, err := http.Get(ep.URL)
		timerEnd := time.Now()
		timerElapsed := timerEnd.Sub(timerStart)

		var result string
		if err != nil {
			result = "FAIL"
		} else {
			result = "PASS"
		}

		fmt.Printf("%s %s %s %fs\n", ep.Name, result, rsp.Status, timerElapsed.Seconds())
	}
}
