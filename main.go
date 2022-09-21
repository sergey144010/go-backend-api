package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const serverPort = 80

func main() {

	var dataPath string
	envPath, ok := os.LookupEnv("DATA")
	if !ok {
		dataPath = "/data"
	} else {
		dataPath = envPath
	}

	var serverOne string = fmt.Sprintf(":%v", serverPort)
	fmt.Printf("Starting server on port %v\n", serverPort)
	log.Fatal(http.ListenAndServe(serverOne, handler(dataPath)))
}

func handler(dataPath string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		files, err := ioutil.ReadDir(dataPath)
		if err != nil {
			panic(err)
		}

		for _, file := range files {
			absFilePath := dataPath + "/" + file.Name()
			nameParts := nameParts(file.Name())
			method := method(nameParts)
			nameParts = cleanName(nameParts)
			route := route(nameParts)

			if req.URL.Path == route && req.Method == strings.ToUpper(method) {
				content, err := ioutil.ReadFile(absFilePath)
				if err != nil {
					panic(err)
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(content)

				break
			}
		}
	})
}

func nameParts(name string) []string {
	return strings.Split(name, ".")
}

func method(nameParts []string) string {
	return nameParts[len(nameParts)-2]
}

func cleanName(nameParts []string) []string {
	return append(nameParts[:len(nameParts)-2])
}

func route(nameParts []string) string {
	return "/" + strings.Join(nameParts, "/")
}
