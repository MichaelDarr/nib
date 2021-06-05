package cmd

import (
	"github.com/spf13/cobra"

	"github.com/MichaelDarr/nib/internal"
)

// rootCmd is the base command which all others are added to
var rootCmd = &cobra.Command{
	Use:     "nib",
	Short:   "Generate SQL migrations from GORM models",
	Version: internal.Version,
}

// Execute is the entrypoint to the CLI
func Execute() {
	rootCmd.Execute()
}
