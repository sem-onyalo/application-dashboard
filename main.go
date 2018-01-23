package main

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sem-onyalo/application-dashboard/core/interactor"
)

func main() {
	databaseInteractor, err := interactor.NewDatabase()
	if err != nil {
		panic(err)
	}

	endpointInteractor := interactor.NewEndpoint(databaseInteractor)

	response := endpointInteractor.GetAll()

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
