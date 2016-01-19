package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.wordpress.org/secret-key/1.1/salt/")
	if err != nil {
		fmt.Println(fmt.Sprintf("Error %s", err.Error()))
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Sprintf("Couldn't read the HTTP body %v", err.Error()))
		panic(err)
	}
	salts := string(body)
	fmt.Println(salts)
}
