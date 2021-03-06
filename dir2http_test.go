package main

import (
  "testing"
  "net/http"
  "strconv"
  "io/ioutil"
  "strings"
)

var client *http.Client 

func init() {
  setRootPath("./test")
  go startServer(1234) 
  
  client = &http.Client{
    CheckRedirect: func(req *http.Request, via []*http.Request) error {
        return http.ErrUseLastResponse
    },
  }
}

func TestServerAcceptRequests(t *testing.T) {
  res, err := client.Get("http://localhost:1234")
  if err != nil {
    t.Fatal(err)
  }
  if res == nil {
    t.Fatal("Response is nil")
  }
  defer res.Body.Close()
}

func TestResponseOk(t *testing.T) {
  res, err := client.Get("http://localhost:1234")
  if err != nil {
    t.Fatal(err)
  }
  defer res.Body.Close()
  
  if res.StatusCode != 200 {
    t.Errorf("Expected the response status code to be %d, but got %d.", 200, res.StatusCode)
  }
}

func TestIndexReturnedForRoot(t *testing.T) {
  res, err := client.Get("http://localhost:1234")
  if err != nil {
    t.Fatal(err)
  }
  defer res.Body.Close()
  
  if res.StatusCode != 200 {
    t.Errorf("Expected the response status code to be %d, but got %d.", 200, res.StatusCode)
  }
  l, err := strconv.Atoi(res.Header.Get("Content-Length"))
  if err != nil {
    t.Fatal("Wrong 'Content-Length':", err)
  }
  if l <= 0 {
    t.Errorf("Expected the 'Content-Length' header to be greater than zero, but got '%v'.", res.Header.Get("Content-Length"))
  }
  
  b, err := ioutil.ReadAll(res.Body)
  if err != nil {
    t.Fatal(err)
  }
  body := string(b)
  if !strings.Contains(body, "<title>dir2http - index</title>") {
    t.Errorf("Expected the response body to contain '%v', but got '%v'.", "<title>dir2http - index</title>", body)
  }
}

func TestIndexReturnedForRootWithSlash(t *testing.T) {
  res, err := client.Get("http://localhost:1234/")
  if err != nil {
    t.Fatal(err)
  }
  defer res.Body.Close()
  
  if res.StatusCode != 200 {
    t.Errorf("Expected the response status code to be %d, but got %d.", 200, res.StatusCode)
  }
  l, err := strconv.Atoi(res.Header.Get("Content-Length"))
  if err != nil {
    t.Fatal("Wrong 'Content-Length':", err)
  }
  if l <= 0 {
    t.Errorf("Expected the 'Content-Length' header to be greater than zero, but got '%v'.", res.Header.Get("Content-Length"))
  }
  
  b, err := ioutil.ReadAll(res.Body)
  if err != nil {
    t.Fatal(err)
  }
  body := string(b)
  if !strings.Contains(body, "<title>dir2http - index</title>") {
    t.Errorf("Expected the response body to contain '%v', but got '%v'.", "<title>dir2http - index</title>", body)
  }
}

func TestSubDirIndexIsReturned(t *testing.T) {
  res, err := client.Get("http://localhost:1234/page/")
  if err != nil {
    t.Fatal(err)
  }
  defer res.Body.Close()
  
  if res.StatusCode != 200 {
    t.Errorf("Expected the response status code to be %d, but got %d.", 200, res.StatusCode)
  }
  l, err := strconv.Atoi(res.Header.Get("Content-Length"))
  if err != nil {
    t.Fatal("Wrong 'Content-Length':", err)
  }
  if l <= 0 {
    t.Errorf("Expected the 'Content-Length' header to be greater than zero, but got '%v'.", res.Header.Get("Content-Length"))
  }
  
  b, err := ioutil.ReadAll(res.Body)
  if err != nil {
    t.Fatal(err)
  }
  body := string(b)
  if !strings.Contains(body, "<title>dir2http - page</title>") {
    t.Errorf("Expected the response body to contain '%v', but got '%v'.", "<title>dir2http - index</title>", body)
  }
}

func TestPageIsReturned(t *testing.T) {
  res, err := client.Get("http://localhost:1234/page/next.html")
  if err != nil {
    t.Fatal(err)
  }
  defer res.Body.Close()
  
  if res.StatusCode != 200 {
    t.Errorf("Expected the response status code to be %d, but got %d.", 200, res.StatusCode)
  }
  l, err := strconv.Atoi(res.Header.Get("Content-Length"))
  if err != nil {
    t.Fatal("Wrong 'Content-Length':", err)
  }
  if l <= 0 {
    t.Errorf("Expected the 'Content-Length' header to be greater than zero, but got '%v'.", res.Header.Get("Content-Length"))
  }
  
  b, err := ioutil.ReadAll(res.Body)
  if err != nil {
    t.Fatal(err)
  }
  body := string(b)
  if !strings.Contains(body, "<title>dir2http - next</title>") {
    t.Errorf("Expected the response body to contain '%v', but got '%v'.", "<title>dir2http - index</title>", body)
  }
}

func TestNonExistingResourceReturnsNotFound(t *testing.T) {
  res, err := client.Get("http://localhost:1234/this-does-not-exist")
  if err != nil {
    t.Fatal(err)
  }
  defer res.Body.Close()
  
  if res.StatusCode != 404 {
    t.Errorf("Expected the response status code to be %d, but got %d.", 404, res.StatusCode)
  }
}

func TestNonExistingDirReturnsNotFound(t *testing.T) {
  res, err := client.Get("http://localhost:1234/this-does-not-exist/")
  if err != nil {
    t.Fatal(err)
  }
  defer res.Body.Close()
  
  if res.StatusCode != 404 {
    t.Errorf("Expected the response status code to be %d, but got %d.", 404, res.StatusCode)
  }
}