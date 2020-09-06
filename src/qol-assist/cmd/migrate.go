package cmd

import "github.com/spf13/cobra"

var migrateCmd = &cobra.Command{
	Use: "migrate",
	Short: "Perform migration functions",
	Long: "Applies migrations that are available on the system",
	Aliases: []string{"m"},
	Run: migrate,
}

func init() {
	RootCmd.AddCommand(listUsersCmd)
}

func migrate(_ *cobra.Command, _ []string) {

}
