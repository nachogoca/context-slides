package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	e.GET("/", handler)
	e.Logger.Fatal(e.Start(":1323"))
}

// START OMIT
func handler(c echo.Context) error {
	err := operation(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "Ok")

}

func operation(ctx context.Context) error {
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("Completed work")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		return ctx.Err()
	}
	return nil
}

// END OMIT
