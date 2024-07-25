package main

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"

	"github.com/frankie-mur/cpkat/links"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string

func run() error {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		return err
	}

	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return err
	}

	queries := links.New(db)

	newLink, err := queries.CreateLink(ctx, links.CreateLinkParams{
		Name: "test",
		Link: "https://google.com",
	})
	if err != nil {
		return err
	}
	fmt.Printf("New link: %+v\n", newLink)

	link, err := queries.GetLink(ctx, newLink.ID)
	if err != nil {
		return err
	}

	fmt.Printf("Fetched link: %+v\n", link)

	allLinks, err := queries.ListLinks(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("List of links: %v\n", allLinks)

	return nil
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
