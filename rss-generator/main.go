package main

import (
	"github.com/ITK13201/rss-generator/cmd"
	_ "github.com/ITK13201/rss-generator/cmd/jobs"
)

func main() {
	cmd.Execute()
}
