package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	// START OMIT
	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/photos", nil)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

	// END OMIT
}
