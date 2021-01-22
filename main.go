package main

// imports .env variables either from $RUNPATH/.env or a file sourced from the config flag. Must be put in each package
import (
	"github.com/xplorfin/go-template/cmd"
	"os"
)

func main() {
	cmd.Start(os.Args)
}
