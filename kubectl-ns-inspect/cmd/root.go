package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kubectl-ns-inspect",
	Short: "Inspect if a namespace has no resources",
	Long:  `This plugin inspects a given Kubernetes namespace to check if it contains any resources.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("kubectl-ns-inspect is a Kubernetes plugin to check namespaces.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
