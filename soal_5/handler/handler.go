package handler

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

type Handler struct {
	APP_NAME string
}

func NewHandler(APPNAME string) Handler{
	return Handler{APPNAME}
}

func(h *Handler) Root(ctx *fasthttp.RequestCtx){
	fmt.Fprintln(ctx, "Welcome to the "+h.APP_NAME)
}

func(h *Handler) Ping(ctx *fasthttp.RequestCtx){
	fmt.Fprintln(ctx, "pong")
}

func(h *Handler) Time(ctx *fasthttp.RequestCtx){
	fmt.Fprintln(ctx, "waktu")
}

