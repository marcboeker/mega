# Go + Graphql + Ent + Echo Boilerplate

It took me a while to figure out how I can use [GraphQL](https://github.com/99designs/gqlgen) with the [ent ORM](https://entgo.io/) and
serve the GraphQL endpoint via the [Echo framework](https://echo.labstack.com/).

I also wanted [proper configuration management](https://github.com/spf13/viper) and a clean structure for my
application logic. It also should be easily testable, because I only write tests if it's convenient.

That's why I came up with **Mega** which could stand for something like
"**M**erge **e**nt with **G**raphQL **a**wesomeness" or "**M**y **e**arliest **G**o **a**dventures".

## What's inside the box

- Go (of course)
- GraphQL via [gqlgen](https://github.com/99designs/gqlgen)
- [ent](https://entgo.io/) as an ORM
- [Echo web framework](https://echo.labstack.com/)
- [Wire](https://github.com/google/wire) for dependency injection
- [Viper](https://github.com/spf13/viper) for configuration management
- [Testify](https://github.com/stretchr/testify) for making testing easy
- [Make](https://www.gnu.org/software/make/manual/make.html) for automation

The challenge was to glue everything as modular as possible together to make replacing single components easy. I'm also not a fan of reflection, so the used libraries like ent, gqlgen, and Wire generate the necessary code instead. It adds little overhead, as you always have to run `make ent|graphql|wire` after each modification. But it adds a lot of type safety and makes developing so much easier.

## Give me some structure

One of the advantages of Go is its flexibility when it comes to structuring your application. It could also be challenging and requires continuous refactoring as your application grows.

I've worked on several web applications using Go in the past and tried a lot of different approaches. From throwing everything in one package (hey my PHP friends from the past) to making anything as modular as possible (waving over to the Java guys).

I've tried [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) and also domain-driven design. Both have pros and cons, but they felt too overkill for most of my Go projects. So I've ended up with the idea of encapsulating my application logic into small services that have a well-defined interface, are easily testable, and can be used with a variety of API frontends like REST, GraphQL, or gRPC.

### Services to the rescue

All my application logic is contained in different [services](https://github.com/marcboeker/tree/services). Each service provides an [interface](https://github.com/marcboeker/tree/services/user/service.go) to describe its API. Each service should but must not contain a [test suite](https://github.com/marcboeker/tree/services/user/service.go).

If you have a closer look at a service, you could see the following components:

```
type Service interface {
	Get(ctx context.Context, id uuid.UUID) (*ent.User, error)
	Create(ctx context.Context, user model.AddUserInput) (*ent.User, error)
}
```

The `Service` interface defines all the offered methods. It also helps to decouple your code if you reference via interfaces instead of specific types.

```
type User struct {
	client *db.Client
}
```

The `User` struct bundles all service methods together and contains the needed dependencies for the service. If you need to send emails, you could, for example, inject a mailer dependency that provides functionality to send mails.

```
func New(client *db.Client) Service {
	return &User{client: client}
}
```

The `New` method simply acts as an initializer.

```
func (s *User) Get(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	return s.client.User.Get(ctx, id)
}
```

And last but not least, the `Get` method on the `User` struct is the actual application logic.

## Getting started

I've decided to use `make` for automating all the tasks. That's why all commands can be run via `make`.

- `make ent` - Generates the required code from your defined models.
- `make graphql` - Generates the GraphQL resolvers and models from the schema.
- `make wire` - Generates the dependency injection code.
- `make all` - Executes the `ent`, `graphql`, and `wire` task all together.
- `make run` - Starts the web server.
- `make test` - Executes all tests.

After each modification of the GraphQL schema, the models or the dependency graph you need to run either `make ent|graphql|wire` or `make all`.

### Defining your GraphQl schema

The schema is located under the `/graph/schema` directory. It contains two different types of files. The "main" schema in `schema.graphqls` contains common types like a `Timestamp`.
And the service-specific schemas like `user.graphqls` contain the query and mutation definitions for the respective service.
Splitting your schema into multiple files also makes it much easier to implement the resolvers later, as they only contain your service-specific queries and mutations and not those of the whole application.

### Configuration

As I'm using Viper as configuration management, it is easy to have specific configurations for different environments. Just copy the `config.yaml` in the repository's root and adapt the setting for a different environment.

The default `make run` command takes the default config file named `config.yml` in the repo's root.

To use your own config file, simply supply the `-config mynewconfig.yml` argument to the `go run` command.

`go run cmd/main.go -config dev.yml`

### First steps

If you execute `make run`, it will start the server on http://127.0.0.1:8080. The default GraphQL query endpoint is under [127.0.0.1:8080/query](http://127.0.0.1:8080/query). There is also a GraphQL playground which is available via [127.0.0.1:8080/playground](http://127.0.0.1:8080/playground).

To insert your first user, simply run the following mutation:

```
mutation {
  createUser(user: {name: "Foo", age: 10}) {
    id,
    createdAt
  }
}
```

It will return the following response:

```
{
  "data": {
    "createUser": {
      "id": "da612c7d-494c-4164-9afc-d353c3e923cf",
      "createdAt": "2022-01-24T00:09:11+01:00"
    }
  }
}
```

You can see a new user with the ID `da612c7d-494c-4164-9afc-d353c3e923cf` is created.

To query for that user, you can use the following query:

```
query {
  user(id: "39b066c7-4b34-4a41-9d21-0b55b16ff0eb") {
    id
    createdAt
  }
}
```

And it returns the same response as the mutation above:

```
{
  "data": {
    "user": {
      "id": "39b066c7-4b34-4a41-9d21-0b55b16ff0eb",
      "createdAt": "2022-01-24T00:17:02+01:00"
    }
  }
}
```
