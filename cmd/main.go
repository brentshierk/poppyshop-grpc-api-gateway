package main

import (
	"github.com/brentshierk/poppyshop-grpc-api-gateway/pkg/auth"
	"github.com/brentshierk/poppyshop-grpc-api-gateway/pkg/config"
	"github.com/brentshierk/poppyshop-grpc-api-gateway/pkg/order"
	"github.com/brentshierk/poppyshop-grpc-api-gateway/pkg/product"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}