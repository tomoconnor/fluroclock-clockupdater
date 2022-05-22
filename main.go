package main

import (
	"errors"
	"log"
	"net/http"
	"os"
)

func main() {
	update_url := os.Getenv("UPDATE_URL")
	if update_url == "" {
		log.Fatal("UPDATE_URL is not set")
	}

	if _, err := os.Stat("/etc/fclock-clock-enable"); errors.Is(err, os.ErrNotExist) {
		log.Fatal("fclock-clock is disabled")
	} else {
		resp, err := http.Post(update_url, "application/json", nil)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			log.Fatal("Update failed")
		}
	}

}
