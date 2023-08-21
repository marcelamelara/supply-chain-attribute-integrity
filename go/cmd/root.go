package cmd

import (
	"os"
	
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scai-gen",
	Short: "A CLI tool for generating SCAI metadata",
}

var outFile string

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&outFile,
		"out-file",
		"o",
		"",
		"Filename to write out the JSON-encoded object",
	)
	rootCmd.MarkPersistentFlagRequired("out-file")
	
	rootCmd.AddCommand(rdCmd)
	rootCmd.AddCommand(assertCmd)
	rootCmd.AddCommand(reportCmd)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}