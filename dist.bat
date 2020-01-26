rm -R .out
mkdir .out

set GOOS=linux
set GOARCH=amd64

go build -o .out/linux/dir2http

set GOOS=darwin
set GOARCH=amd64

go build -o .out/darwin/dir2http

set GOOS=windows
set GOARCH=amd64

go build -o .out/windows/dir2http.exe

zip -j .out/dir2http-linux-amd64.zip .out/linux/dir2http
zip -j .out/dir2http-darwin-amd64.zip .out/darwin/dir2http 
zip -j .out/dir2http-win-amd64.zip .out/windows/dir2http.exe