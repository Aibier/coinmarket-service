package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"

	"github.com/Aibier/coinmarket-service/controllers"
)

// HandleRequest ...
func HandleRequest(ctx context.Context, name App) (string, error) {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":   "Coin Market Migration",
			"AppVersion": "v1",
		}).Info("Starting the app...")
	controllers.Sync()
	log.Printf("At the end of my job, let's rest now! Completed time %s", time.Now().Local().String())
	return fmt.Sprintf("Resources are saved %s by!", name.Name), nil
}

// App - the struct which contains information about our app
type App struct {
	Name    string
	Version string
}

// start server
func (app *App) start() error {


	return nil
}

func main() {
	lambda.Start(HandleRequest)
}