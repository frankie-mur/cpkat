package cmd

import (
	"fmt"

	"github.com/frankie-mur/cpkat/data"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCommand)
}

var listCommand = &cobra.Command{
	Use:  "list",
	Long: `show list of all the links`,
	Run: func(cmd *cobra.Command, args []string) {
		links, err := data.GetAllLinks()
		if err != nil {
			panic(err)
		}

		for _, link := range links {
			fmt.Println(link.Name, link.Link)
		}
	},
}
