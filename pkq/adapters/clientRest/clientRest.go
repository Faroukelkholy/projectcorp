package clientRest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"projectcorp/config"
	"projectcorp/pkq/domain/model"
)

var cfg = config.Parse()

type ClientRest struct {}

func(*ClientRest) GetEmployee(idParam string) (*model.Employee,error) {
	requestURL := fmt.Sprintf(cfg.GetEmployeesURL+"%s",idParam)
	request, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	log.Println("request.URL :",request.URL.String())

	resp, errDefaultClient := http.DefaultClient.Do(request)
	if errDefaultClient != nil {
		log.Println("restClient  errDefaultClient:",errDefaultClient)
		return nil,errDefaultClient
	}

	defer resp.Body.Close()
	bodyBytes, errioutil := ioutil.ReadAll(resp.Body)
	if errioutil != nil {
		log.Println("restClient  errioutil:",errioutil)
		return nil,errioutil
	}

	var employee model.Employee
	errMarshal := json.Unmarshal(bodyBytes, &employee)
	if errMarshal != nil {
		log.Println("restClient  errMarshal:",errMarshal)
		return nil,errMarshal
	}

	return &employee,nil

}