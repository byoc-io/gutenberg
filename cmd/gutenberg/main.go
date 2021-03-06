package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	app string = "server"
)

func commandRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:  app,
		Long: "A fast, flexible and powerful publishing platform",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(2)
		},
	}

	rootCmd.AddCommand(commandServe())
	rootCmd.AddCommand(commandVersion())

	return rootCmd
}

func main() {
	if err := commandRoot().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}
}
