package api

import (
	"github.com/d-xingxing/go-programming-tour/blog-service/global"
	"github.com/d-xingxing/go-programming-tour/blog-service/internal/service"
	"github.com/d-xingxing/go-programming-tour/blog-service/pkg/app"
	"github.com/d-xingxing/go-programming-tour/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	//time.Sleep(1 * time.Second)
	//panic("GetAuth panic test")
	//_, err = ctxhttp.Get(c.Request.Context(), http.DefaultClient, "https://www.google.com/")
	//if err != nil {
	//	log.Fatalf("ctxhttp.Get ewrr: %v", err)
	//}

	svc := service.New(c.Request.Context()) // 这里曾经写错为service.New(c)
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenError)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}
