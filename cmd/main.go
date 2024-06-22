package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/antoniofmoraes/multithreading-challenge/internal"
)

func main() {
	cepinfoViacepChannel := make(chan *internal.CepInfoResponse)
	cepinfoBrasilapiChannel := make(chan *internal.CepInfoResponse)

	go func() {
		cepinfo, err := internal.GetCepInfoFromViacep("81900140")
		if err != nil {
			log.Print(err)
			return
		}
		cepinfoViacepChannel <- cepinfo
	}()

	go func() {
		cepinfo, err := internal.GetCepInfoFromBrasilapi("81900140")
		if err != nil {
			log.Print(err)
			return
		}
		cepinfoBrasilapiChannel <- cepinfo
	}()

	var cepinfo *internal.CepInfoResponse
	select {
	case cepinfo = <-cepinfoViacepChannel:
		fmt.Println("Cep information got from Viacep API")
	case cepinfo = <-cepinfoBrasilapiChannel:
		fmt.Println("Cep information got from BrasilAPI")
	case <-time.After(time.Second):
		log.Fatalf("Timeout")
	}

	cepinfoJson, err := json.Marshal(cepinfo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(cepinfoJson))
}
