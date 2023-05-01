package request

import (
	"encoding/json"
	"net/http"

	"github.com/beego/beego/v2/core/logs"
)

func Get(urlp string, headers map[string]string, target interface{}) error {

	req, err := http.NewRequest("GET", urlp, nil)
	if err != nil {
		logs.Error("Error reading request. ", err)
	}

	// Se intenta acceder a cabecera, si no existe, se realiza peticion normal.

	defer func() {
		//Catch
		if r := recover(); r != nil {

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				logs.Error("Error reading response. ", err)
			}

			defer resp.Body.Close()
			json.NewDecoder(resp.Body).Decode(target)
		}
	}()

	//try
	setHeaders(&req.Header, headers)
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		logs.Error("Error reading response. ", err)
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func setHeaders(req *http.Header, headers map[string]string) {
	for k, v := range headers {
		req.Set(k, v)
	}
}
