package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	req, err := http.NewRequest(http.MethodGet, "https://wipdev.netlify.app", nil)
	if err != nil {
		return fmt.Errorf("error forming the request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending the request: %v", err)
	}

	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	fmt.Println(string(resBody))

	return err
}
