package biz

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

var _ handler = (*GoodHandler)(nil)

type GoodHandler struct {
}

func NewGoodHandler() *GoodHandler {
	return &GoodHandler{}
}
func (g *GoodHandler) RegisterRoutes(server *gin.Engine) {
	gg := server.Group("/goods")
	gg.POST("/listByKeyword", g.listByKeyword)
	gg.POST("/listByCategory", g.listByCategory)
	gg.POST("/detail", g.detail)
	gg.POST("/add_good", g.addGood)
	gg.POST("/update_good", g.updateGood)
	gg.POST("/delete_good", g.deleteGood)
}

func (g *GoodHandler) listByKeyword(ctx *gin.Context) {

}
func (g *GoodHandler) listByCategory(ctx *gin.Context) {
	type Req struct {
		Page     int    `json:"page"`
		Size     int    `json:"size"`
		Category string `json:"category"`
	}
	var req Req
	if err := ctx.ShouldBindJSON(&req); err != nil {
		zap.S().Errorf("[list] invalid params: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 查询商品列表

}

func (g *GoodHandler) detail(ctx *gin.Context) {

}

func (g *GoodHandler) addGood(ctx *gin.Context) {

}

func (g *GoodHandler) updateGood(ctx *gin.Context) {

}

func (g *GoodHandler) deleteGood(ctx *gin.Context) {

}
