package main

import (
	"github.com/7samll7/cmd"
	"github.com/7samll7/handle"
)

func main() {
	platform, content, language := cmd.CommandArgs()
	if platform != "" && content != "" && language != "" {
		handle.TxtTrans(platform, content, language)
	}
}
