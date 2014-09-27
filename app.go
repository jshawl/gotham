package main

import (
  "io/ioutil"
  "fmt"
)

func main() {
  file, err := ioutil.ReadFile("markdown.md")
  if err != nil {
    panic( err )
  }
  fmt.Println( file )
}