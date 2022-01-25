package cmd

import (
	"fmt"
	"github.com/ek/pkg/k8s"
	"github.com/spf13/cobra"
)

var imagesCmd = &cobra.Command{
	Use: "images",
	Short: "获取k8s中指定命名空间中的容器镜像",
	Run: func(cmd *cobra.Command, args []string) {
		k8sClient := k8s.NewClient()
		images := k8sClient.FetchImages(namespace)
		for _, image := range images {
			fmt.Println(image)
		}
	},
}
var namespace = ""

func init() {
	imagesCmd.Flags().StringVarP(&namespace, "namespace", "n", namespace, "-n namespace")
	rootCmd.AddCommand(imagesCmd)
}
