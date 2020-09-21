package clientRest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"projectcorp/domain/model"
)

type ClientRest struct {}

func(*ClientRest) GetEmployee(url string, id string) (*model.Employee, error) {

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	params := request.URL.Query()
	params.Add("id", id)
	request.URL.RawQuery = params.Encode()

	resp, errDefaultClient := http.DefaultClient.Do(request)
	if errDefaultClient != nil {
		return nil,errDefaultClient
	}

	defer resp.Body.Close()
	bodyBytes, errioutil := ioutil.ReadAll(resp.Body)
	if errioutil != nil {
		return nil,errioutil
	}

	var employee *model.Employee
	errMarshal := json.Unmarshal(bodyBytes, &employee)
	if errMarshal != nil {
		return nil,errMarshal
	}

	return employee,nil


}