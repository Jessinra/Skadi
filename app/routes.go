package app

import "gitlab.com/trivery-id/skadi/internal/user/controller"

var AuthController *controller.AuthController

func initRoutes() {
	router.Use(
		corsMiddleware(),
		addUUIDToRequestCtxMiddleware(),
	)

	router.POST("/auth/login", AuthController.Login)
	router.POST("/auth/refresh", AuthController.RefreshToken)

	router.POST("/graphql", authenticatedUser(graphqlHandler))
	router.GET("/playground", playgroundHandler)

	router.GET("/", ping)
}
