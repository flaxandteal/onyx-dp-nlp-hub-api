# dp-nlp-hub
A simple proxy hub

# dp-nlp-hub
## Description

A Go application microservice to consolidate enhanced search functionality. 
In particular, this will call out to the dedicated thin wrappers for ML models where required.

### Available scripts

- `make help` - Displays a help menu with available `make` scripts
- `make all` - Runs audit test and build commands
- `make run_docker_container` - Runs container name: from image name: nlp_hub
- `make build_docker` - Builds ./Dockerfile image name: nlp_hub
- `make build` - Build bin file in folder build
- `make test` - Runs all tests with -cover -race flags
- `make convey` - Runs only convey tests
- `make debug` - Runs application locally with debug mode on
- `make update` - Go gets all of the dependencies and downloads them

### Configuration

| Environment variable         | Default   | Description
| ---------------------------- | --------- | -----------
| BIND_ADDR                    | :3002     | The host and port to bind to
| GRACEFUL_SHUTDOWN_TIMEOUT    | 5s        | The graceful shutdown timeout in seconds (`time.Duration` format)
| HEALTHCHECK_INTERVAL         | 30s       | Time between self-healthchecks (`time.Duration` format)
| HEALTHCHECK_CRITICAL_TIMEOUT | 90s       | Time to wait until an unhealthy dependent propagates its state to make this app unhealthy (`time.Duration` format)
|	BERLIN_BASE           | http://localhost:3001/berlin/search |The url where the berlin api is available
|	SCRUBBER_BASE               | http://localhost:3002/scrubber/search | The url where the scrubber apiapi  is available
|	CATEGORY_BASE           | http://localhost:80/categories |The url where the scrubber api is available

## Quick setup

### Docker

```shell
make build_docker
make run_docker_container
```

### Locally

```shell
make update
make debug
```

## Dependencies

- `github.com/ONSdigital/dp-component-test v0.9.0`
- `github.com/ONSdigital/dp-healthcheck v1.5.0`
- `github.com/ONSdigital/dp-net/v2 v2.8.0`
- `github.com/ONSdigital/log.go/v2 v2.3.0`
- `github.com/cucumber/godog v0.12.6`
- `github.com/dghubble/sling v1.4.1`
- `github.com/gorilla/mux v1.8.0`
- `github.com/kelseyhightower/envconfig v1.4.0`
- `github.com/pkg/errors v0.9.1`
- `github.com/smartystreets/goconvey v1.7.2`
- `github.com/stretchr/testify v1.8.1`
- `go version go1.19.5 linux/amd64 `

## Usage

Running the project either locally or in docker will expose port 3002.

```shell
curl 'http://localhost:5000/health' 
```
This will return results of the form:

```json
{
    "status": "OK",
    "version": {
        "build_time": "2020-09-26T11:30:18Z",
        "git_commit": "6584b786caac36b6214ffe04bf62f058d4021538",
        "language": "go",
        "language_version": "go1.20",
        "version": "v0.1.0"
    },
    "uptime": 1372,
    "start_time": "2023-02-23T13:10:04.744563463Z",
    "checks": []
}
```

```shell
curl 'http://localhost:5000/search?q=dentists'
```
This will return results of the form:

```json
{
    "Scrubber": {
        "query": "dentists",
        "results": {
            "areas": null,
            "industries": null
        },
        "time": ""
    },
    "Category": null,
    "Berlin": {
        "query": {
            "codes": null,
            "exact_matches": null,
            "normalized": "",
            "not_exact_matches": null,
            "raw": "",
            "stop_words": null
        },
        "results": null,
        "time": ""
    }
}
```

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

### License

Copyright © 2023, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.

