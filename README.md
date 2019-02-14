

# Run in Unix-like enviroment

### Setup the enviroment
```
export GOPATH="`pwd`"
```

### When running for the first time get the gorilla mux
```
go get github.com/gorilla/mux
```

### Setup the .config.json (see config.json.example)
This program uses https://openweathermap.org/ api. So you need a api key for that

### Build and run

```
go build api
./api <path to config>
```
