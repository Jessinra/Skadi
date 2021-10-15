package app

func initRoutes() {
	router.Use(
		parseJWT(),
	)

	router.GET("/", ping)

	router.POST("/graphql", graphqlHandler)
	router.GET("/playground", playgroundHandler)
}
