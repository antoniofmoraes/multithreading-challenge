package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/antoniofmoraes/multithreading-challenge/internal"
)

func main() {
	cep := getCepArg()

	cepinfoViacepChannel := make(chan *internal.CepInfoResponse)
	cepinfoBrasilapiChannel := make(chan *internal.CepInfoResponse)

	go func() {
		cepinfo, err := internal.GetCepInfoFromViacep(cep)
		if err != nil {
			log.Print(err)
			return
		}
		cepinfoViacepChannel <- cepinfo
	}()

	go func() {
		cepinfo, err := internal.GetCepInfoFromBrasilapi(cep)
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

func getCepArg() string {
	if len(os.Args) < 2 {
		log.Fatalf("Cep argument needed. Example: go run cmd/main.go 82590-300")
	}

	return internal.RemoveNonDigits(os.Args[1])
}
