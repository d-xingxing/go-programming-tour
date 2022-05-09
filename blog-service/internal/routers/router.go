package routers

import (
	_ "github.com/d-xingxing/go-programming-tour/blog-service/docs"
	"github.com/d-xingxing/go-programming-tour/blog-service/global"
	"github.com/d-xingxing/go-programming-tour/blog-service/internal/middleware"
	"github.com/d-xingxing/go-programming-tour/blog-service/internal/routers/api"
	v1 "github.com/d-xingxing/go-programming-tour/blog-service/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

// NewRouter 可以从此处添加路由
func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 新增auth相关的路由
	r.POST("/auth", api.GetAuth)

	// 上传文件服务
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	// 提供静态资源访问
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	article := v1.NewArticle()
	tag := v1.NewTag()
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		// 标签管理
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags/:id", tag.Get)
		apiv1.GET("/tags", tag.List)

		// 文章管理
		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}

	return r
}
