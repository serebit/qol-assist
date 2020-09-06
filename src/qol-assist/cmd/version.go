package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const QolAssistVersion = "0.5.0"

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "show version",
	Long: "Print the qol-assist version and exit",
	Run: printVersion,
}

func init() {
	RootCmd.AddCommand(versionCmd)
}

func printVersion(_ *cobra.Command, _ []string) {
	fmt.Printf("qol-assist version %v\n\nCopyright © 2017-2020 Solus Project\n", QolAssistVersion)
	fmt.Println("Licensed under the Apache License, Version 2.0")
}
