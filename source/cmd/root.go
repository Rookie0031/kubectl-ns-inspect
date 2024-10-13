package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var dryRun bool

// MARK: root command definition
var rootCmd = &cobra.Command{
	Use:   "ns-inspect",
	Short: "Inspect if a namespace has no resources",
	Long:  `This plugin inspects a given Kubernetes namespace to check if it contains any resources.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Kubernetes client settings + inspect namespace
		kubeconfig := os.Getenv("KUBECONFIG")
		if kubeconfig == "" {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				log.Fatalf("Error getting user home directory: %v", err)
			}
			kubeconfig = filepath.Join(homeDir, ".kube", "config")
		}

		// Kubernetes client configuration
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			log.Fatalf("Error building kubeconfig: %s", err.Error())
		}
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			log.Fatalf("Error creating Kubernetes client: %s", err.Error())
		}

		// Get namespace list
		namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			log.Fatalf("Error getting namespaces: %s", err.Error())
		}

		// Inspect each namesapce
		for _, ns := range namespaces.Items {
			fmt.Printf("Checking namespace: %s\n", ns.Name)

			empty, resourceCounts, err := inspectNamespace(clientset, ns.Name)
			if err != nil {
				fmt.Printf("Error checking namespace %s: %s\n", ns.Name, err.Error())
				continue
			}

			if empty {
				fmt.Printf("Namespace %s is empty of primary resources.\n", ns.Name)

				fmt.Printf("Resources remaining in namespace %s:\n", ns.Name)
				for resourceType, count := range resourceCounts {
					fmt.Printf("- %s: %d\n", resourceType, count)
				}

				// Check delete
				if !dryRun {
					fmt.Printf("\nDo you want to delete the namespace '%s'? (yes/no): ", ns.Name)
					var input string
					fmt.Scanln(&input)

					if strings.ToLower(input) == "yes" {
						fmt.Printf("Deleting namespace: %s\n", ns.Name)
						err := clientset.CoreV1().Namespaces().Delete(context.TODO(), ns.Name, metav1.DeleteOptions{})
						if err != nil {
							fmt.Printf("Error deleting namespace %s: %s\n", ns.Name, err.Error())
						} else {
							fmt.Printf("Successfully deleted namespace: %s\n", ns.Name)
						}
					} else {
						fmt.Printf("Skipped deletion of namespace: %s\n", ns.Name)
					}
				}
			}
		}

		if dryRun {
			fmt.Println("\nDry-run mode: no namespaces were deleted.")
		}
	},
}

// MARK: Execute
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// MARK: init
func init() {
	// --dry-run flag for test
	rootCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "List namespaces without deleting them")
}
