package app

func initRoutes() {
	router.GET("/", ping)

	router.POST("/graphql", graphqlHandler)
	router.GET("/playground", playgroundHandler)
}
