package main

import (
	"github.com/subiz/go-elasticsearch/internal/cmd/generate/commands"
	_ "github.com/subiz/go-elasticsearch/internal/cmd/generate/commands/gensource"
	_ "github.com/subiz/go-elasticsearch/internal/cmd/generate/commands/genstruct"
	_ "github.com/subiz/go-elasticsearch/internal/cmd/generate/commands/gentests"
)

func main() {
	commands.Execute()
}
