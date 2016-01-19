package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func getAppName() string {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 1 {
		panic(fmt.Errorf("App Name Not provided"))
	}
	arg := argsWithoutProg[0]
	return arg
}

func main() {
	appName := getAppName()
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
	reg := regexp.MustCompile(`'(.+)',\s+'(.+)'`)
	matches := reg.FindAllStringSubmatch(salts, -1)
	command := fmt.Sprintf("dokku config:set %s", appName)
	for _, matches := range matches {
		command = command + fmt.Sprintf(" %s='%s'", matches[1], matches[2])
	}
	fmt.Println(command)
}
