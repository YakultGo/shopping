package biz

import "github.com/gin-gonic/gin"

var _ handler = (*GoodHandler)(nil)

type GoodHandler struct {
}

func NewGoodHandler() *GoodHandler {
	return &GoodHandler{}
}
func (g *GoodHandler) RegisterRoutes(server *gin.Engine) {
	gg := server.Group("/goods")
	gg.POST("/list", g.list)
	gg.POST("/detail", g.detail)
	gg.POST("/add_good", g.addGood)
	gg.POST("/update_good", g.updateGood)
	gg.POST("/delete_good", g.deleteGood)
}

func (g *GoodHandler) list(ctx *gin.Context) {

}

func (g *GoodHandler) detail(ctx *gin.Context) {

}

func (g *GoodHandler) addGood(ctx *gin.Context) {

}

func (g *GoodHandler) updateGood(ctx *gin.Context) {

}

func (g *GoodHandler) deleteGood(ctx *gin.Context) {

}
