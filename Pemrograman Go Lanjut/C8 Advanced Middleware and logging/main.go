// package main
// import (
// 	"fmt"
// 	"github.com/labstack/echo"
// 	"net/http"
// 	"github.com/labstack/echo/middleware"

// )

// func main() {
// 	e := echo.New()

// 	//middleware
// 	e.Use(middlewareOne)
// 	e.Use(middlewareTwo)
// 	//middleware 3rd party
// 	e.Use(echo.WrapMiddleware(middlewareSomething))

// 	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
//         Format: "method=${method}, uri=${uri}, status=${status}\n",
//     }))

// 	e.GET("/index", func(c echo.Context) (err error){
// 		fmt.Println("Thereeee!")

// 		return c.JSON(http.StatusOK, true)
// 	})

// 	e.Logger.Fatal(e.Start(":9000"))
// }

// func middlewareOne(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error{
// 		fmt.Println("from middleware one")
// 		return next(c)
// 	}
// }

// func middlewareTwo(next echo.HandlerFunc) echo.HandlerFunc{
// 	return func(c echo.Context) error{
// 		fmt.Println("from middleware two")

// 		return next(c)
// 	}
// }

// func middlewareSomething(next http.Handler) http.Handler{
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
// 		fmt.Println("from middleware something")
// 		next.ServeHTTP(w, r)
// 	})
// }

package main

import (
	"fmt"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {
	e := echo.New()

	e.Use(middlewareLogging)
	e.HTTPErrorHandler = errorHandler

	e.GET("/index", func(c echo.Context) error {
		return c.JSON(http.StatusOK, true)
	})

	lock := make(chan error)
	go func(lock chan error) { lock <- e.Start(":9000") }(lock)

	time.Sleep(1 * time.Millisecond)
	makeLogEntry(nil).Warning("application started without ssl/tls enabled")

	err := <-lock
	if err != nil {
		makeLogEntry(nil).Panic("failed to start application")
	}
}

func makeLogEntry(c echo.Context) *log.Entry {
	if c == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return log.WithFields(log.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": c.Request().Method,
		"uri":    c.Request().URL.String(),
		"ip":     c.Request().RemoteAddr,
	})
}

func middlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		makeLogEntry(c).Info("incoming request")
		return next(c)
	}
}

func errorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if ok {
		report.Message = fmt.Sprintf("http error %d - %v", report.Code, report.Message)
	} else {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	makeLogEntry(c).Error(report.Message)
	c.HTML(report.Code, report.Message.(string))
}
