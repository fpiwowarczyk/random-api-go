# Random API

Simple rest api that is using random.org and count standard deviation of returend values. 

## How to run
There are two ways 

### Run with go
You can run program by using that commans in project home directory 

```
go run .
```

### Run with docker
You can build dockerfile by building file with project
```
docker build ./random-api-go --tag random-api
```

then you can run container with 
```
docker run -d -p 8080:8080 random-api
```

## How to use

To call api you can use tools like postman or curl or use your browser by typing 

`localhost:8080/random/mean?requests=1&length=5`

where:
requests - concurrent requests to random.org api 

length - number of numbers to get in every request

## CI and Tests

Code is tested with about 70% of coverage to run tests use:

```
go test -v ./...

# OR for coverage report 

go test -coverprofile=coverage.out ./...
```

also there is running CI/CD runned with github actions where tests are checked always after pushing to main branch. 