package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"projectcorp/config"
	"projectcorp/pkq/domain/model"
)

var cfg = config.Parse()

type Client struct{}

func (*Client) GetEmployee(id string) (e *model.Employee, err error) {
	URL := fmt.Sprintf(cfg.GetEmployeesURL+"%s", id)
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}

	resp, errhc := http.DefaultClient.Do(req)
	if errhc != nil {
		return nil, errhc
	}
	defer resp.Body.Close()

	bytes, errio := ioutil.ReadAll(resp.Body)
	if errio != nil {
		return nil, errio
	}

	err = json.Unmarshal(bytes, &e)
	fmt.Println("inside hc e:", e)
	if err != nil {
		return nil, err
	}

	return e, nil
}
