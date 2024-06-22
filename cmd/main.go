package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/antoniofmoraes/multithreading-challenge/internal"
)

func main() {
	cepinfo, err := internal.GetCepInfoFromBrasilapi("81900140")
	if err != nil {
		log.Fatal(err)
	}

	cepinfoJson, err := json.Marshal(cepinfo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(cepinfoJson))
}
