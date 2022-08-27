package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	data string
	root = &cobra.Command{
		Use: "tsc",
		Run: func(cmd *cobra.Command, args []string) {
			if data == "" {
				fmt.Printf("hello tsc~")
			}
		},
	}
)

func init() {
	cobra.OnInitialize(onData)
	root.PersistentFlags().StringVarP(&data, "data", "d", "", "json文件")
	root.AddCommand(redirect)
}

func Run() {
	if err := root.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
