package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var version bool

var RootCmd = &cobra.Command{
	Use:   "nodeInfo",
	Short: "nodeInfo",
	Long:  `nodeInfo`,
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&version, "version", "v", false, "the nodeInfo-server version")
}
