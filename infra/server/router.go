package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func APIRouter(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		output := fmt.Sprintf("%#v", c.Request().Header)
		os.Stdout.Write([]byte(output + "\n"))
		ip := c.RealIP()
		fmt.Println("IPAddr: ", ip)
		return c.String(http.StatusOK, "ʕ◔ϖ◔ʔ"+"\n"+ip)
	})
	e.POST("/index/mapping/define", DefineIndex)
}
