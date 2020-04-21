package routing

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"test-case/soal_5/handler"
	"test-case/soal_5/middleware"
)

type Router struct {
	handler         handler.Handler
	PORT, APP_NAME 	string
	DEBUG_MODE 		bool
}

func NewRouter(handler handler.Handler, PORT, APP_NAME string, DEBUG_MODE bool) *Router{
	return &Router{handler, PORT, APP_NAME, DEBUG_MODE,
	}
}


func (r *Router) Routing(){
	router := fasthttprouter.New()
	logMiddleware := middleware.NewLoggingMiddleware(r.DEBUG_MODE)
	router.GET("/", logMiddleware(r.handler.Root))
	router.GET("/ping", logMiddleware(r.handler.Ping))
	router.GET("/time", logMiddleware(r.handler.Time))

	fmt.Println("Listening on " + r.PORT)
	fmt.Println("Ready to serve...")
	if errorIntialRunning := fasthttp.ListenAndServe(":"+r.PORT, CORS(router.Handler)); errorIntialRunning != nil {
		fmt.Println("Failed to serve ~", errorIntialRunning)
	}
}
