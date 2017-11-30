package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	url := args[1]
	interactions, err := strconv.Atoi(args[2])

	if err != nil {
		fmt.Printf("please use a valid number of interactions")
	} else {
		responses := make(chan string)

		for i := 0; i < interactions; i++ {
			t := i
			go func() {
				resp, err := http.Get(url)
				if err == nil {
					bodyBytes, err2 := ioutil.ReadAll(resp.Body)
					if err2 == nil {
						bodyString := string(bodyBytes)
						responses <- bodyString
					} else {

						fmt.Printf("Request number %d failed\n", t)
					}
				} else {

					fmt.Printf("Request number %d failed\n", t)
				}
			}()
		}
		areEqual := true
		previous := ""
		for i := 0; i < interactions; i++ {
			current := <-responses
			fmt.Printf("%s \n", current)
			if previous != "" {
				areEqual = previous == current
			}
			previous = current
		}

		if areEqual {

			fmt.Printf("All equal ")
		} else {
			fmt.Printf("not equal")
		}
	}
}
