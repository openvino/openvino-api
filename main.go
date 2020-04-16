package main

import (
	"fmt"
	"log"

	"github.com/openvino/openvino-api/src/app"
	"github.com/openvino/openvino-api/src/config"
	
)

func main() {

	config := config.New();

	app := &app.App{}
	app.Initialize(config)

	log.Printf("Serving application at PORT: %d", config.Port)
	app.Run(fmt.Sprintf(":%d", config.Port))
	
}
