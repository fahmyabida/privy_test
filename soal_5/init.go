package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"test-case/soal_5/handler"
	"test-case/soal_5/routing"
)

const (
	ENVIRONMENT = "ENVIRONMENT"
	DEVELOPMENT = "DEVELOPMENT"
	PRODUCTION = "PRODUCTION"
)

func RunApplication(){
	PORT, APP_NAME, DEBUG_MODE := getEnvironmentVariable()
	fmt.Println(strings.ToUpper(APP_NAME), "is running...")
	fmt.Println("debug mode =", DEBUG_MODE)
	handler := handler.NewHandler(APP_NAME)
	router := routing.NewRouter(handler, PORT, APP_NAME, DEBUG_MODE)
	router.Routing()
}

func getEnvironmentVariable() (PORT, APP_NAME string, DEBUG_MODE bool){
	ENV := os.Getenv(ENVIRONMENT)
	if ENV == DEVELOPMENT {
		PORT 			= os.Getenv("DEVELOPMENT_PORT")
		APP_NAME 		= os.Getenv("DEVELOPMENT_APP_NAME")
		DEBUG_MODE, _	= strconv.ParseBool(os.Getenv("DEVELOPMENT_DEBUG_MODE"))
	} else if ENV == PRODUCTION {
		PORT 			= os.Getenv("PRODUCTION_PORT")
		APP_NAME 		= os.Getenv("PRODUCTION_APP_NAME")
		DEBUG_MODE, _	= strconv.ParseBool(os.Getenv("PRODUCTION_DEBUG_MODE"))
	} else {
		log.Fatal("ENVIRONMENT NOT LISTED")
	}
	return PORT, APP_NAME, DEBUG_MODE
}
