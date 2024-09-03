package cmd

import "github.com/spf13/cobra"

var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "",
	GroupID: "",
	Version: "",
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}

func init() {

}
