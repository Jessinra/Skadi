package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"gitlab.com/trivery-id/skadi/graph/generated"
	"gitlab.com/trivery-id/skadi/graph/resolver"
)

func ping(c *gin.Context) {
	version := os.Getenv("VERSION")
	c.String(http.StatusOK, fmt.Sprintf("ping - %s", version))
}

func graphqlHandler(c *gin.Context) {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))
	h.ServeHTTP(c.Writer, c.Request)
}

func playgroundHandler(c *gin.Context) {
	h := playground.Handler("GraphQL", "/graphql")
	h.ServeHTTP(c.Writer, c.Request)
}
