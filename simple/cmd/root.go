package cmd

import (
        "fmt"
        "github.com/gofunct/spawnd/cookiecutter-grpcgo/simple/config"
        "os"

        "github.com/spf13/cobra"
)

var (
	cookie *config.CookieCutter
)

func init() {
	cookie = config.NewCookieCutterCookieCutter("SIMPLE")
}


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
        Use:   "simple",
        Short: "A brief description of your application",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
        if err := rootCmd.Execute(); err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
}

func init() {
        cobra.OnInitialize()

        // Cobra also supports local flags, which will only run
        // when this action is called directly.
        // rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}