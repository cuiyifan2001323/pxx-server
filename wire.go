// @Author 齐静春
// @Date 2023-11-14 21:55:00
//
//go:build:wireinject
package main

//
//func InitWebServer() *gin.Engine {
//	wire.Build(
//		ioc.InitDB,
//		ioc.InitGinMiddlewares,
//		dao.NewUserDao,
//		repository.NewUserRepository,
//		service.NewUserService,
//		ioc.InitWebServer,
//		web.NewUserHandle,
//	)
//	return gin.Default()
//}
