package main

import (
	"gopkg.in/gographics/imagick.v2/imagick"
	"imgdemo"
)

/*
gitCommit an gitTag, are injected at compile time, via the flag:
go build -ldflags "-X main.GitCommit=${GIT_COMMIT} -X main.GitTag=${GIT_TAG}" -a  .
they can be viewed in the version handler

*/

func main() {

	imagick.Initialize()
	defer imagick.Terminate()

	myapp := imgdemo.NewServer()

	myapp.Run()
}
