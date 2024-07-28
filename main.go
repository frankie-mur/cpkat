package main

import (
	"github.com/frankie-mur/cpkat/cmd"
)

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

func main() {
	cmd.Execute()
}
