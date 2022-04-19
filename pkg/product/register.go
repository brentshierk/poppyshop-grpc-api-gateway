package product

import (
	"github.com/brentshierk/poppyshop-grpc-api-gateway/pkg/auth"
	"github.com/brentshierk/poppyshop-grpc-api-gateway/pkg/config"
	"github.com/brentshierk/poppyshop-grpc-api-gateway/pkg/product/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
a := auth.InitAuthMiddleware(authSvc)

svc := &ServiceClient{
Client: InitServiceClient(c),
}

	productRoutes := r.Group("/product")
	productRoutes.Use(a.AuthRequired)
	productRoutes.POST("/", svc.CreateProduct)
	productRoutes.GET("/:id", svc.FindOne)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FineOne(ctx, svc.Client)
}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}