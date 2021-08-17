package cmd

import (
	k8sClient "github.com/ek/pkg/k8s"
	"github.com/ek/pkg/migrate"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use: "migrate",
	Short: "迁移k8s中的镜像",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

var migrateOption = &migrate.Option{}

func init() {
	migrateCmd.Flags().StringVarP(&migrateOption.Namespace, "namespace", "n", "", "-n default")
	migrateCmd.Flags().StringVarP(&migrateOption.Image, "image", "i", "", "-i image")
	migrateCmd.Flags().StringVarP(&migrateOption.Repository, "repository", "r", "", "-r repository")
	migrateCmd.Flags().StringVarP(&migrateOption.Exclude, "exclude", "e", "", "-e regex")
	migrateCmd.Flags().StringVarP(&migrateOption.Config, "config", "c", "", "-c config path")
	migrateCmd.Flags().StringVarP(&migrateOption.Prefix, "prefix", "p", "", "-p prefix")
	rootCmd.AddCommand(migrateCmd)
}

func run() {

	// 获取指定镜像名称
	images := k8sClient.FetchImages(migrateOption.Namespace)
	migrate.MigrateImage(images, migrateOption)

}