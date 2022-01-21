# Golang-Echo-Gorm-Clean-Architecture

Implementation of Clean Architecture in Golang RestAPI projects.

<br>

![Untitled Diagram](https://user-images.githubusercontent.com/15135199/150567366-5ea535bf-07ac-4049-a24b-44a0cd5c21cf.png)



├── api
│   ├── handler
│   │   ├── handler.go
│   │   ├── helpers.go
│   │   └── model.go
│   └── middleware
│       └── middleware.go
├── cloudformation
│   └── apigateway
│       ├── apig.yaml
│       └── create_apigateway.sh
