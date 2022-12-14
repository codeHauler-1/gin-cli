package cmd

import (
	"fmt"
	"git.xiaoyanggroup.cn/xyyjyframework/gin-cli/app/gen"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	genCmd.AddCommand(daoCmd)
	rootCmd.AddCommand(genCmd)
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "auto generate dao/service/controller",
}

var daoCmd = &cobra.Command{
	Use:   "dao",
	Short: "auto generate dao",
	Run: func(cmd *cobra.Command, args []string) {
		vp := viper.New()
		vp.SetConfigType("yaml")
		vp.AddConfigPath(".")
		if err := vp.ReadInConfig(); err != nil {
			fmt.Println(fmt.Sprintf("Error:%s", err.Error()))
			return
		}

		dns := vp.GetString("mysql.default.link")
		if dns == "" {
			fmt.Println("Error:not found mysql link in config.yaml")
			return
		}

		if err := gen.GenerateDao(dns); err != nil {
			fmt.Println(fmt.Sprintf("Error:%s", err.Error()))
		}
		fmt.Println("done!")
	},
}
