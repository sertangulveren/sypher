package main

import (
	"flag"
	"github.com/sertangulveren/sypher/internal/commander"
	"github.com/sertangulveren/sypher/internal/shared"
	"os"
)



func main() {
	flag.Usage = shared.Help
	flag.Parse()

	if len(os.Args) < 2 {
		// no command
		shared.Help()
	}

	switch os.Args[1] {
	case "gen":
		commander.Generate()
	case "print":
		commander.Print()
	case "edit":
		commander.Edit()
	case "gitignore":
		commander.GenerateGitIgnore()
	default:
		shared.Help()
	}
}
