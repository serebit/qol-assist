package cmd

import "github.com/spf13/cobra"

var listUsersCmd = &cobra.Command{
	Use: "list-users",
	Short: "List users on the system",
	Long: "List users on the system and their associated groups",
	Aliases: []string{"lu"},
	Run: listUsers,
}

func init() {
	RootCmd.AddCommand(listUsersCmd)
}

func listUsers(_ *cobra.Command, _ []string) {

}
