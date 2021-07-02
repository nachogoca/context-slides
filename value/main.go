package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type key string

const RequestIDKey = key("request-id")

func Println(ctx context.Context, msg string) {

	id, ok := ctx.Value(RequestIDKey).(int64)
	if !ok {
		log.Println("could not find any request id")
		return
	}
	log.Printf("[%d] %s", id, msg)
}

func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, RequestIDKey, id)
		f(rw, r.WithContext(ctx))
	}
}

func main() {

	http.HandleFunc("/", Decorate(handler))
	panic(http.ListenAndServe("127.0.0.1:1323", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	Println(ctx, "started request")
	defer Println(ctx, "finished request")
	err := operation(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func operation(ctx context.Context) error {
	select {
	case <-time.After(5 * time.Second):
		Println(ctx, "Completed work")
	case <-ctx.Done():
		Println(ctx, ctx.Err().Error())
		return ctx.Err()
	}
	return nil
}

// END OMIT
