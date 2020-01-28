package main

import (
  "net/http"
  "os"
  "log"
  "fmt"
  "strconv"
  "path/filepath"
)

const INDEX_FILE = "index.html"

var root string

func main() {
  args := os.Args[1:]
  if len(args) < 1 {
    printUsage()
    os.Exit(1)
  }
  port, err := strconv.Atoi(args[0])
  if err != nil {
    fmt.Println("Wrong port!")
    printUsage()
    os.Exit(1)
  }
  if len(args) > 1 {
    root = args[1]
  } else {
    root = "."
  }
  root = filepath.Clean(root)
  
  info, err := os.Stat(root)
  if os.IsNotExist(err) || !info.IsDir() {
    fmt.Println(root, "is not a directory!")
    os.Exit(1)
  }
  
  startServer(port)
}

func printUsage() {
  fmt.Println("Usage: <port> [path/to/root:.]")
}

func setRootPath(p string) {
  root = p
}

func startServer(port int) {
  log.Println("Starting dir2http server")
  address := ":" + strconv.Itoa(port)
  handler := http.FileServer(http.Dir(root))
  log.Fatal(http.ListenAndServe(address, handler))
}