echo "Hello World!"

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o linuxMain main.go

SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -o macMain main.go

