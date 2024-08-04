package main

import (
	"context"
	_ "embed"

	"github.com/frankie-mur/cpkat/cmd"
	"github.com/frankie-mur/cpkat/data"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := data.OpenDB(context.Background())
	if err != nil {
		panic(err)
	}
	cmd.Execute()
}
