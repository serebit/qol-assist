package cmd

import (
	"github.com/spf13/cobra"
)

var CLIDebug bool

var RootCmd = &cobra.Command{
	Use:   "qol-assist",
	Short: "qol-assist is QoL assistance to help Solus roll!",
}
