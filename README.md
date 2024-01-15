# Latitude CLI

A Command Line Interface to interact with our API. It uses [go-swagger](https://github.com/go-swagger/go-swagger/) to generate the CLI based on the swagger definitions from our API.

## Dependencies

- [Golang](https://go.dev/doc/install)
- [go-swagger](https://github.com/go-swagger/go-swagger/)
- [api-spec-converter](https://github.com/LucyBot-Inc/api-spec-converter)

## Generating the CLI

go-swagger is not compatible with OpenAPI 3.0, so before trying to generate the CLI, we need to generate a schema with swagger 2.0. To do so, you can run this command:

```sh
api-spec-converter --from=openapi_3 --to=swagger_2 --syntax=json --order=alpha PATH_TO_API_REPO/swagger/v3/swagger.json > swagger.json
```

- Then, on the root directory of this project, run this command to generate the CLI:

```sh
swagger generate cli --target=. --spec=swagger.json --cli-app-name lsh
```

> Note: If you run into errors when running this command, check the troubleshooting section.

- After generating the CLI, build the executable with:

```sh
# Build the executable
go build -o lsh cmd/lsh/main.go

# Copy to your GOPATH so you can run this from any directory
cp lsh $(go env GOPATH)
```

## Usage

The first step is to setup your credentials:

```sh
lsh login $YOUR_LATITUDE_API_KEY
```

This will store your credentials locally in a config file, so you can run commands without having to pass the credentials every time.

Everything is set up! :tada: You should be able to run CLI commands (check `lsh help` for more info).

To check if the CLI was successfully generated, you can make a request to an API endpoint, for example: `lsh roles get-roles`. This should return a list of available roles that can be assigned to users.

## Troubleshooting

TODO
