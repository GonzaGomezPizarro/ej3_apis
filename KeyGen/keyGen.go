package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://classify-web.herokuapp.com/api/keygen?length=32?symbols=1"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	//req.Header.Add("", "") // header vacio, no se usa.

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
