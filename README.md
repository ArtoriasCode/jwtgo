# jwtgo

A Go (Golang) backend clean architecture project with API Gateway, Microservices, MongoDB and JWT Authentication.

The project was created for educational purposes and is not ideal. It has its shortcomings, which are gradually being corrected.

## Project architecture
The architecture of a web application consists of these layers:
- Reverse Proxy
- API Gateway
- Microservice
- Database

The architecture of the microservice application consists of these layers:
- Server
- Service
- Repository

![Image](https://raw.githubusercontent.com/ArtoriasCode/jwtgo/refs/heads/main/assets/architecture.png)

The API is accessed via Reverse Proxy, in our case it is Nginx. It handles all incoming requests and prevents access to microservices and API gateway directly.

## Major packages used in project
- **[Gin](https://pkg.go.dev/github.com/gin-gonic/gin)**: Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin. 
- **[gRPC](https://pkg.go.dev/google.golang.org/grpc)**: The Go implementation of gRPC: A high performance, open source, general RPC framework that puts mobile and HTTP/2 first. For more information see the Go gRPC docs, or jump directly into the quick start. 
- **[protobuf](https://pkg.go.dev/google.golang.org/protobuf)**: Go support for Google's protocol buffers.
- **[MongoDB](https://pkg.go.dev/go.mongodb.org/mongo-driver)**: The Official Golang driver for MongoDB.
- **[JWT](https://pkg.go.dev/github.com/golang-jwt/jwt/v5)**: Go implementation of JSON Web Tokens (JWT).
- **[Cleanenv](https://pkg.go.dev/github.com/ilyakaznacheev/cleanenv)**: Clean and minimalistic environment configuration reader for Golang.
- **[Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)**: Package bcrypt implements Provos and Mazières's bcrypt adaptive hashing algorithm.
- **[Logrus](https://pkg.go.dev/github.com/sirupsen/logrus)**: Structured, pluggable logging for Go.
- **[Validator](https://pkg.go.dev/github.com/go-playground/validator/v10)**: Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving.

## Request flow without JWT authentication middleware
![Image](https://raw.githubusercontent.com/ArtoriasCode/jwtgo/refs/heads/main/assets/without-jwt.png)

## Request flow with JWT authentication middleware
![Image](https://raw.githubusercontent.com/ArtoriasCode/jwtgo/refs/heads/main/assets/with-jwt.png)

## How to run the project?
First, download it and navigate to the root directory:
```bash
# Move to your workspace
cd your-workspace

# Clone the project into your workspace
git clone https://github.com/ArtoriasCode/jwtgo.git

# Move to the project root directory
cd jwtgo
```

### Run with Docker
- Create a `.env` file, similar to `.env.sample`.
- Install the [Docker](https://www.docker.com/get-started/), [Protoc](https://grpc.io/docs/protoc-installation/), [Taskfile](https://taskfile.dev/installation/) if it is not installed on your computer.
- Fill in the `.env` file with your data.
- Run the application build with the following command:

  ```bash
  task build
  ```
- Access API using http://localhost.

## Examples of API requests and responses
### SignUp endpoint
- Request:
  ```
  curl --location 'http://localhost/api/v1/auth/signup' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "email": "test@gmail.com",
    "password": "securepassword"
  }'
  ```
  
- Response:
  ```
  HTTP/1.1 200 OK
  Content-Type: application/json
  ```
  ```json
  {
    "message": "User successfully registered"
  }
  ```

### SignIn endpoint
- Request:
  ```
  curl --location 'http://localhost/api/v1/auth/signin' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "email": "test@gmail.com",
    "password": "securepassword"
  }'
  ```

- Response:
  ```
  HTTP/1.1 200 OK
  Content-Type: application/json
  Set-Cookie: access_token=access_token; Path=/; HttpOnly; SameSite=Strict
  Set-Cookie: refresh_token=refresh_token; Path=/; HttpOnly; SameSite=Strict
  ```
  ```json
  {
    "message": "User successfully logged in"
  }
  ```

### SignOut endpoint
- Request:
  ```
  curl --location 'http://localhost/api/v1/auth/signout' \
  --header 'Content-Type: application/json' \
  -b 'access_token=access_token; refresh_token=refresh_token'
  ```

- Response:
  ```
  HTTP/1.1 200 OK
  Content-Type: application/json
  Set-Cookie: access_token=access_token; Path=/; HttpOnly; SameSite=Strict
  Set-Cookie: refresh_token=refresh_token; Path=/; HttpOnly; SameSite=Strict
  ```
  ```json
  {
    "message": "User successfully logged out"
  }
  ```

### Refresh endpoint
- Request:
  ```
  curl --location 'http://localhost/api/v1/auth/refresh' \
  --header 'Content-Type: application/json' \
  -b 'access_token=access_token; refresh_token=refresh_token'
  ```

- Response:
  ```
  HTTP/1.1 200 OK
  Content-Type: application/json
  Set-Cookie: access_token=access_token; Path=/; HttpOnly; SameSite=Strict
  Set-Cookie: refresh_token=refresh_token; Path=/; HttpOnly; SameSite=Strict
  ```
  ```json
  {
    "message": "Tokens successfully updated"
  }
  ```

## Complete project folder structure
```
├── build
│   └── package
│       ├── api.Dockerfile
│       ├── auth.Dockerfile
│       └── user.Dockerfile
├── cmd
│   ├── api
│   │   └── main.go
│   ├── auth
│   │   └── main.go
│   └── user
│       └── main.go
├── configs
│   └── nginx.conf
├── deployments
│   └── docker-compose.yaml
├── internal
│   ├── app
│   │   ├── api
│   │   │   ├── config
│   │   │   │   └── config.go
│   │   │   ├── controller
│   │   │   │   └── http
│   │   │   │       ├── dto
│   │   │   │       │   └── user.go
│   │   │   │       ├── mapper
│   │   │   │       │   └── user.go
│   │   │   │       ├── middleware
│   │   │   │       │   ├── security.go
│   │   │   │       │   └── validation.go
│   │   │   │       └── v1
│   │   │   │           └── auth.go
│   │   │   └── main.go
│   │   ├── auth
│   │   │   ├── config
│   │   │   │   └── config.go
│   │   │   ├── interface
│   │   │   │   └── service
│   │   │   │       └── auth.go
│   │   │   ├── server
│   │   │   │   └── grpc
│   │   │   │       ├── dto
│   │   │   │       │   └── user.go
│   │   │   │       ├── mapper
│   │   │   │       │   └── user.go
│   │   │   │       └── v1
│   │   │   │           └── auth.go
│   │   │   ├── service
│   │   │   │   └── auth.go
│   │   │   └── main.go
│   │   └── user
│   │       ├── adapter
│   │       │   └── mongodb
│   │       │       ├── entity
│   │       │       │   └── user.go
│   │       │       ├── mapper
│   │       │       │   └── user.go
│   │       │       └── repository
│   │       │           └── user.go
│   │       ├── config
│   │       │   └── config.go
│   │       ├── entity
│   │       │   └── user.go
│   │       ├── interface
│   │       │   ├── repository
│   │       │   │   └── user.go
│   │       │   └── service
│   │       │       └── user.go
│   │       ├── server
│   │       │   └── grpc
│   │       │       ├── dto
│   │       │       │   └── user.go
│   │       │       ├── mapper
│   │       │       │   └── user.go
│   │       │       └── v1
│   │       │           └── user.go
│   │       ├── service
│   │       │   └── user.go
│   │       └── main.go
│   └── pkg
│       ├── error
│       │   ├── auth.go
│       │   ├── jwt.go
│       │   ├── repository.go
│       │   └── server.go
│       ├── interface
│       │   └── service
│       │       ├── jwt.go
│       │       └── password.go
│       ├── proto
│       │   ├── auth
│       │   │   ├── auth.pb.go
│       │   │   └── auth_grpc.pb.go
│       │   └── user
│       │       ├── user.pb.go
│       │       └── user_grpc.pb.go
│       ├── request
│       │   ├── schema
│       │   │   └── response.go
│       │   └── response.go
│       └── service
│           ├── schema
│           │   └── jwt.go
│           ├── jwt.go
│           └── password.go
├── pkg
│   ├── client
│   │   └── mongodb.go
│   ├── logging
│   │   └── logger.go
│   └── proto
│       ├── auth
│       │   └── auth.proto
│       └── user
│           └── user.proto
├── .env
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
├── README.md
└── taskfile.yaml
```