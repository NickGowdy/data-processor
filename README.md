<br />
<div align="center">

  <h1 align="center">Data Processor</h1>

  <p align="center">
    <h2 align="center">A Golang program for concurrently registering devices.</h2>
  </p>
</div>

## About 

This program uses concurrency to register devices.

### Built With

* Golang
* Docker

## Getting Started

To run this code, we first need a `.env` file in the root of the project. Once this is done, add these vars:

```
BASE_URL=some_url
TIMEOUT=30000
CODE_REGISTRATION_LIMIT=100
```

### Prerequisites

- You will need Golang to run this program which can be downloaded at: [https://go.dev/](https://go.dev/)
- This program can also be run using docker, this can be downloaded at: [https://www.docker.com/](https://www.docker.com/)

## Usage

Then to run locally, use: `go run main.go`.

Alternatively, this code can also be run via docker. To build the docker image use: 
```
docker build -t data-processor . --build-arg BASE_URL=${BASE_URL} --build-arg TIMEOUT=${TIMEOUT} --build-arg CODE_REGISTRATION_LIMIT=${CODE_REGISTRATION_LIMIT}.
```

To run the docker image: 
```
docker run data-processor
```

To run the tests use: 
```
go test ./...
``` 
To check code coverage: 
```
go test -coverprofile=coverage.out ./... ; go tool cover -html=coverage.out
```

And to run benchmark tests for CPU and memory consumption:
```
go test -bench=. -benchmem
```

Finally to check for race conditions, use: 
```
go test -race ./...
```

## Checklist

- [x] Implement solution
- [x] Write unit tests
- [x] Write readme
- [ ] More unit tests around unhappy paths for higher code coverage

## Contact

Email - nickgowdy87@gmail.com

Website <a href="http://www.nickgowdy.com/" target="_blank">http://www.nickgowdy.com/</a>

Github <a href="https://github.com/nickgowdy" target="_blank">https://github.com/nickgowdy</a>




