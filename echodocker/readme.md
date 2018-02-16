$ env GOOS=linux GOARCH=amd64 go build --tags netgo --ldflags '-extldflags "-lm -lstdc++ -static"'

$ docker build -t echoserver:1.0 .

$ docker run -d echoserver