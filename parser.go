package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	logFile := "/var/log/grafana/grafana.log"
	logFileInput, err := ioutil.ReadFile(logFile)

	searchIt(string(logFileInput))

	if err != nil {
		fmt.Println(logFile, "Cannot be accessed")
	}
}

func searchIt(file string) {
	//logErrors := "lvl=eror"
	r, _ := regexp.Compile("lvl=eror\\s+msg=(\"[\\w*\\s*]*\\.?\")")
	results := r.FindAllString(file, -1)
	//newResults := r.FindStringSubmatch(results)
	for _, res := range results {
		newResults := r.FindStringSubmatch(res)
		fmt.Println(newResults[1])
	}
}
