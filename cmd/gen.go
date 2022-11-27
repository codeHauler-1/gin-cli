package cmd

import (
	"fmt"
	"github.com/codeHauler-1/gin-cli/app/init/project"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(genCmd)
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "auto generate dao",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error:not fount projectName")
			fmt.Println("Example:gin-cli initProject {projectName}")
			return
		}

		projectName := args[0]
		if err := project.NewProject(projectName); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(fmt.Sprintf("project %s init success!", projectName))
		fmt.Println("you can exec:")
		fmt.Println(fmt.Sprintf("     cd %s && go mod tidy && go run main.go", projectName))
		fmt.Println("done!")
	},
}