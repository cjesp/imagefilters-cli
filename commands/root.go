package commands

import (
	"github.com/spf13/cobra"
)

// var Workers int
// var Time bool
var (
	input, output string
	sigmaS, sigmaI float64
	Workers int
	Time bool
)

var RootCmd = &cobra.Command{
	Use: "imageTool",
	Short: "A small tool for manipulating images.",
	Long: `Image tool allows one to manipulate images
with different algorithms, currently only blurring is 
supported.`,
	Run: func(cmd *cobra.Command, args []string)  {
		
	},
}

func init()  {
	RootCmd.PersistentFlags().IntVar(&Workers, "workers", 1, 
	"number of workers (goroutines/threads)")
	RootCmd.PersistentFlags().BoolVar(&Time, "time", false, 
	"output execution time")
}

