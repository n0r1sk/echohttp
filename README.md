# echohttp
A small go http server that responses local and http information. For more information see the Dockerfile

# build instructions for Alpine Linux container
The binary was build with the following settings to create a runable Alpine Linux binary.
```GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o echohttp```
