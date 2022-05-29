package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"zen/internal/pkg/util"
)

var (
	dataFile string
)

// fakeCmd represents the fake command
var fakeCmd = &cobra.Command{
	Use:   "fake",
	Short: "Expand template with fake data generation",
	Long:  `Generates fake data by template expansion`,
	Run: func(cmd *cobra.Command, args []string) {

		data, err := util.ReadData(dataFile)

		util.IfErrorExit(err, 1)

		fmt.Println(data)
		for _, arg := range args {
			err := util.ExpandTemplate(arg, os.Stdout, data)
			util.IfErrorExit(err, 1)
		}
	},
}

func init() {
	rootCmd.AddCommand(fakeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fakeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	fakeCmd.Flags().StringVarP(&dataFile, "data", "d", "", "read data from json")
}
