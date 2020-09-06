package cmd

import "github.com/spf13/cobra"

var triggerCmd = &cobra.Command{
	Use: "trigger",
	Short: "Schedule migration on next boot",
	Long: "Schedules available migrations to be run after the system is rebooted",
	Aliases: []string{"t"},
	Run: trigger,
}

func init() {
	RootCmd.AddCommand(listUsersCmd)
}

func trigger(_ *cobra.Command, _ []string) {

}
