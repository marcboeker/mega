package resolver

import (
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/marcboeker/mega/graph/generated"
	"github.com/marcboeker/mega/service"
)

type Resolver struct {
	*service.Mapper
	*log.Logger
}

func NewSchema(services *service.Mapper, logger *log.Logger) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{services, logger},
	})
}
