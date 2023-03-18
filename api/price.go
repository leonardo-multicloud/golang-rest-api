package api

import (
	"io/ioutil"
	"log"
	"net/http"
)

type HTTPClientInterface interface {
	Do(req *http.Request) (*http.Response, error)
}

type Economy struct {
	HTTPClient HTTPClientInterface
}

func (e *Economy) GetPrice(w http.ResponseWriter, r *http.Request) {

	apiURL := "https://economia.awesomeapi.com.br/last/USD-BRL,EUR-BRL,BTC-BRL"

	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	//res, err := http.DefaultClient.Do(req) Removido com a injeção de dependencia.
	res, err := e.HTTPClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
	log.Println("GET: route: /price")
}
