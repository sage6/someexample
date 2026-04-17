package main

import (
	"fmt"
	"os"
	"net/http"
	"io"
	"io/ioutil"
	"path/pathname"
)

func main() {
	
	start := time.Now()
	ch := make(chan string)

	
