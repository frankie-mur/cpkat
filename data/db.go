package data

import (
	"context"
	"database/sql"

	_ "embed"

	"github.com/frankie-mur/cpkat/links"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string

var db *sql.DB

func OpenDB(ctx context.Context) error {
	var err error
	db, err = sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		return err
	}

	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return err
	}

	return db.Ping()
}

func GetAllLinks() ([]links.Link, error) {
	if db == nil {
		panic("db is nil")
	}
	q := links.New(db)

	allLinks, err := q.ListLinks(context.Background())
	if err != nil {
		return nil, err
	}

	return allLinks, nil
}

func AddLink(name string, link string) (links.Link, error) {
	if db == nil {
		panic("db is nil")
	}
	//TODO: probably should not call a New() everytime
	q := links.New(db)

	newLink, err := q.CreateLink(context.Background(), links.CreateLinkParams{
		Name: name,
		Link: link,
	})

	if err != nil {
		return links.Link{}, err
	}

	return newLink, nil
}

// func run() error {

// 	// queries := links.New(db)

// 	// newLink, err := queries.CreateLink(ctx, links.CreateLinkParams{
// 	// 	Name: "test",
// 	// 	Link: "https://google.com",
// 	// })
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// fmt.Printf("New link: %+v\n", newLink)

// 	// link, err := queries.GetLink(ctx, newLink.ID)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// fmt.Printf("Fetched link: %+v\n", link)

// 	// allLinks, err := queries.ListLinks(ctx)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// fmt.Printf("List of links: %v\n", allLinks)

// 	// return nil
// }
