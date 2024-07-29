package main

import (
	"github.com/7samll7/cmd"
	"github.com/7samll7/handle"
)

func main() {
	platform, content, language := cmd.CommandArgs()
	handle.TxtTrans(platform, content, language)
}
