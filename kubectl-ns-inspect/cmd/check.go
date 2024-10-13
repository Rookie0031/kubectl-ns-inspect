package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var namespace string

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if a namespace has no resources",
	Long:  `This command checks a given namespace to see if it contains any resources.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 이곳에서 kubectl 명령어를 실행하여 네임스페이스의 리소스를 검사합니다.
		if namespace == "" {
			fmt.Println("Please provide a namespace.")
			return
		}

		// `kubectl get all` 명령어로 네임스페이스를 검사하는 예시입니다.
		out, err := exec.Command("kubectl", "get", "all", "-n", namespace).Output()
		if err != nil {
			fmt.Printf("Error checking namespace: %v\n", err)
			return
		}

		fmt.Printf("Resources in namespace %s:\n%s\n", namespace, string(out))
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// 플래그 추가 (예: --namespace)
	checkCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "Namespace to inspect")
}
