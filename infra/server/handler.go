package server

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gentom/elastico/app"
	"github.com/labstack/echo"
)

type defineIndexRequest struct {
	URL       string           `json:"url"`
	Username  string           `json:"username"`
	Password  string           `json:"password"`
	IndexName string           `json:"index_name"`
	Mapping   *json.RawMessage `json:"mapping"`
}

type commonResponse struct {
	StatusCode string `json:"status_code"`
}

// DefineIndex func
func DefineIndex(c echo.Context) error {
	var response commonResponse
	request := &defineIndexRequest{}
	c.Bind(request)

	ctx := context.Background()
	app := app.Start(request.URL, request.Username, request.Password)
	err := app.DefineIndex(ctx, request.IndexName, request.Mapping)
	if err != nil {
		fmt.Println(err)
		response.StatusCode = "FAILURE"
		return c.JSON(500, response)
	}
	response.StatusCode = "SUCCESS"
	return c.JSON(500, response)
}
