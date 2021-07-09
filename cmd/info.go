package cmd

import (
	"log"
	"skylark/internal/pkg/mtime"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "show info",
	Long:  "show info",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "get current time",
	Long:  "get current time",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := mtime.GetNowTime()
		log.Printf("result: %s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

var versionCmd = &cobra.Command{
	Use:   "ver",
	Short: "get verison",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("soft version is %s", "0.0.1")
	},
}

func init() {
	infoCmd.AddCommand(nowTimeCmd)
	infoCmd.AddCommand(versionCmd)
}
