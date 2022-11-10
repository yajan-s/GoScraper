package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func checkStatus(ch chan string) {
	resp, err := http.Get("https://api.nike.com/launch/launch_views/v2/5c196c26-cc06-39d1-b79b-cb086e93282a")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	i := strings.Index(sb, "launchState")
	ch <- strings.Split(strings.Split(sb[i:], " ")[2], "\"")[1]
}

func main() {
	ch := make(chan string)
	for i := 0; i < 1000; i++ {

		go checkStatus(ch)
		val := <-ch

		fmt.Println("Status: " + val)

	}
	close(ch)
}
