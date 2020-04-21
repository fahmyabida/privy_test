package middleware

import (
	"github.com/valyala/fasthttp"
	"log"
)

type RequestHandler func(ctx *fasthttp.RequestCtx)

type Middleware func(h RequestHandler) fasthttp.RequestHandler

func NewLoggingMiddleware(debugMode bool) Middleware {
	return func(h RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			if debugMode {
				log.Println("URL="+string(ctx.Request.RequestURI())," | ", "request-body="+string(ctx.Request.Body()))
				h(ctx)
				log.Println("URL="+string(ctx.Request.RequestURI())," | ", "response-body="+string(ctx.Response.Body()))
			} else {
				h(ctx)
			}
		}
	}
}
