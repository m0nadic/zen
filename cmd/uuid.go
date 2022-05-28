package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"zen/internal/pkg/util"
)

var (
	upper bool
	count int
)

// uuidCmd represents the uuid command
var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generates an Universal Unique Identifier",
	Run: func(cmd *cobra.Command, args []string) {
		uuids := util.GenerateUUIDs(count, upper)
		fmt.Print(strings.Join(uuids, "\n"))
	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)
	uuidCmd.Flags().BoolVarP(&upper, "upper", "u", false, "Output in upper case")
	uuidCmd.Flags().IntVarP(&count, "count", "n", 1, "number of UUIDs to generate")
}
