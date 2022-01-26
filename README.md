# Golang-Echo-Gorm-Clean-Architecture

Implementation of Clean Architecture in Golang RestAPI projects. Maintaining a large codebase with maximum readability and minimal overhead is hard.

### How to structure the code

![Untitled Diagram](https://user-images.githubusercontent.com/15135199/150567366-5ea535bf-07ac-4049-a24b-44a0cd5c21cf.png)

- <b>Handlers</b> are responsible for receiving requests from clients and sending responses back to to clients.
- <b>Services</b> are responsible for receiving information from handlers to process bussiness logic.
- <b>Repositories</b> abstract data storage behind a common interface, allowing services to persist data to and retrieve data

### Project structure

    ├── main.go                     // running
    ├── config                      
    │   ├── config_prop.go          // config property (read all config file to struct)
    │   ├── db                      // config db
    │   └── redis                   // config redis
    ├── middlewares
    ├── routes
    │   └── route.go                // routes
    ├── features
    │   ├── user                    // user module
    │   │   ├── data.go             // request and response data
    │   │   ├── service.go          // business logic
    │   │   └── handler.go    
    │   ├── role                    // role module
    │   │   ├── data.go             // request and response data
    │   │   ├── service.go          // business logic
    │   │   └── handler.go
    ├── repositories
    │   ├── models.go               // models or entities
    │   ├── user_repo.go            // repo interface
    │   ├── role_repo.go            // repo interface
    │   ├── mysql
    │   │   ├── user_repo.go        // repo implement
    │   │   └── role_repo.go        // repo implement
    ├── config.yaml
    ├── docker-compose.yaml
    ├── Makefile

### Start app function

```go

func main() {

	var server = echo.New()

	route.SetupLogger(server)
	route.SetupGlobalErrorHandler(server)
	route.SetupMiddleware(server)
	route.SetupRoute(server)

	server.Logger.Fatal(server.Start(":8080"))
}

```

### Dependencies

- github.com/labstack/echo/v4
- github.com/labstack/echo/v4/middleware
- github.com/labstack/gommon/log
- gorm.io/gorm
- gorm.io/driver/mysql
- github.com/jinzhu/copier

### Reference

- https://about.sourcegraph.com/go/gophercon-2019-how-uber-go-es/
