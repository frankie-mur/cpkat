package cmd

import (
	"fmt"

	"github.com/frankie-mur/cpkat/data"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCommand)
}

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "add a new link",
	Long:  `add a new link`,
	Run: func(cmd *cobra.Command, args []string) {
		name, link := args[0], args[1]
		createdLink, err := data.AddLink(name, link)

		if err != nil {
			panic(err)
		}

		fmt.Printf("Added new link: Name: %s -> Url: %s\n", createdLink.Name, createdLink.Link)
	},
}
