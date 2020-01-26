# dir2http: Directory to a HTTP Server

Makes a local directory accessible via HTTP.

## Usage
```
dir2http 1234 ./test

curl http://localhost:1234 
curl http://localhost:1234/page/
curl http://localhost:1234/page/next.html
curl http://localhost:1234/image.jpg --output image.jpg
```
