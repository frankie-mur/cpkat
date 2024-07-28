package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "embed"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

//go:embed schema.sql
var ddl string

var rootCmd = &cobra.Command{
	Version: "0.0.1",
	Use:     "cpkat",
	Long:    `cpkat is a command line tool for managing links`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cpkat is a command line tool for managing links")
		ctx := context.Background()

		db, _ := openDB(ctx)
		// if err != nil {
		// 	return err
		// }

		fmt.Fprintf(os.Stdout, "db: %+v\n", db)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func openDB(ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		return nil, err
	}

	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return nil, err
	}

	return db, nil
}
