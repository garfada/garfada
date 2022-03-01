package meal

import "github.com/gin-gonic/gin"

const (
	Route = "/meal"
)

type GinRoutes struct{}

func (r *GinRoutes) Create(context *gin.Context) {
	context.JSON(200, ModelResponse{})
}
func (r *GinRoutes) Read(context *gin.Context) {
	context.JSON(200, ModelResponse{})
}
func (r *GinRoutes) Update(context *gin.Context) {
	context.JSON(200, ModelResponse{})
}
func (r *GinRoutes) Delete(context *gin.Context) {
	context.JSON(200, ModelResponse{})
}
