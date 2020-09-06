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
	fmt.Printf("qol-assist version %v\n\nCopyright © 2017-2020 Solus Project\n\n", QolAssistVersion)
	fmt.Println("qol-assist is free software; you can redistribute it and/or modify")
	fmt.Println("it under the terms of the GNU General Public License as published by")
	fmt.Println("the Free Software Foundation; either version 2 of the License, or")
	fmt.Println("(at your option) any later version.")
}
