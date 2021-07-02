package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", handler)
	panic(http.ListenAndServe("127.0.0.1:1323", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	log.Println("started request")
	defer log.Println("finished request")
	go func() {
		err := operation(ctx)
		if err != nil {
			log.Printf("could not complete operation: %s", err.Error())
		}
	}()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func operation(ctx context.Context) error {
	select {
	case <-time.After(5 * time.Second):
		log.Println("Completed work")
	case <-ctx.Done():
		log.Println(ctx.Err())
		return ctx.Err()
	}
	return nil
}

// END OMIT
