package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/utf6/go-blog/api"
	v1 "github.com/utf6/go-blog/api/v1"
	_ "github.com/utf6/go-blog/docs"
	"github.com/utf6/go-blog/middleware"
	"github.com/utf6/go-blog/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)
	r.POST("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新文章
		apiv1.PUT("/articles/:id", v1.EditTag)
		//文章详情
		apiv1.GET("/articles/:id", v1.GetArticle)
		//删除文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)

	}
	return r
}