# Golang-Echo-Gorm-Clean-Architecture

Implementation of Clean Architecture in Golang RestAPI projects.


![Untitled Diagram](https://user-images.githubusercontent.com/15135199/150567366-5ea535bf-07ac-4049-a24b-44a0cd5c21cf.png)

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
    │   ├── models.go               // models
    │   ├── user_repo.go            // repo interface
    │   ├── role_repo.go            // repo interface
    │   ├── mysql
    │   │   ├── user_repo.go        // repo implement
    │   │   └── role_repo.go        // repo implement
    ├── config.yaml
    ├── docker-compose.yaml
    ├── Makefile


