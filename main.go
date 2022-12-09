package main

import (
	"fmt"
	"io/ioutil"

	http "github.com/mostafa/mal4md/net/http"
)

func main() {
	client := &http.Client{}

	if resp, err := client.Get("http://httpbin.test.k6.io/uuid"); err != nil {
		fmt.Println(err)
	} else {
		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(body))
		}
		resp.Body.Close()
	}

	if req, err := http.NewRequest("GET", "http://httpbin.test.k6.io/headers", nil); err != nil {
		fmt.Println(err)
	} else {
		req.Header.Add("If-None-Match", `W/"wyzzy"`)
		if resp, err := client.Do(req); err != nil {
			fmt.Println(err)
		} else {
			if body, err := ioutil.ReadAll(resp.Body); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(string(body))
			}
			resp.Body.Close()
		}
	}
}
