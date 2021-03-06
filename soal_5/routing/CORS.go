package routing

import "github.com/valyala/fasthttp"

var (
	corsAllowCredentials 	= "true"
	corsAllowHeaders     	= "*"
	corsAllowExposeHeader	= "*"
	corsAllowMethods     	= "*"
	corsAllowOrigin      	= "*"
)

func CORS(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)
		ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
		ctx.Response.Header.Set("Access-Control-Expose-Headers", corsAllowExposeHeader)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", corsAllowMethods)
		ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)
		next(ctx)
	}
}
