# Programming challenge - Weather Retrieval endpoint

We would like to install a weather monitor in our basement office. Due to security concerns, we should only use our own software. Therefore, we need to build a simple weather retrieval service that allows our colleagues to check the current weather conditions outside of our offices.

## What should I do?

Please create a very simple **production ready** application that exposes an API endpoint for retrieving the current temperature in a city specified by the user. The temperature should be returned in the following three formats: Celsius, Fahrenheit, Kelvin. You should retrieve the temperature from a public weather service, e.g. OpenWeather API.

## What can I use?

The programming language to be used for this challenge has already been agreed upon during the interview process. You are free to use any libraries, packages, or modules commonly available for the chosen language.

## Userful links

- OpenWeather API - https://openweathermap.org/api

## Execution instructions

You can start the app by running the following commands:

```bash
cd cmd/my-api
make run
```

You can execute the tests by first instaling [mockery](https://vektra.github.io/mockery/latest/installation/#go-install) and then running:

```bash
mockery
```

From the root of the project. This will generate the mocks needed for `Test_AddWeatherRoutes`. You can then run:

```bash
go test ./...
```

Alternatively the mocks are included in the project already so you could use the ones already there.

### Example queries

You might need to install `jq` on your local system.

```bash
curl 'http://localhost:1234/weather/city/pinar%20del%20rio' | jq
```

```bash
curl http://localhost:1234/weather/city/berlin | jq
```
