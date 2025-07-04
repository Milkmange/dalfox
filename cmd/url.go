package cmd

import (
	"github.com/hahwul/dalfox/v2/internal/printing"
	"github.com/hahwul/dalfox/v2/pkg/scanning"
	"github.com/spf13/cobra"
)

// urlCmd represents the url command for scanning a single target URL
var urlCmd = &cobra.Command{
	Use:   "url [target] [flags]",
	Short: "Use single target mode",
	Run:   runURLCmd,
}

// runURLCmd handles execution of the URL command to scan a single target
func runURLCmd(cmd *cobra.Command, args []string) {
	printing.Banner(options)
	if len(args) == 0 {
		printUrlErrorAndUsage()
		return
	}

	printing.Summary(options, args[0])
	printing.DalLog("SYSTEM", "Using single target mode", options)
	if options.Format == "json" {
		printing.DalLog("PRINT", "[", options)
	}
	_, _ = scanning.Scan(args[0], options, "Single")
	if options.Format == "json" {
		printing.DalLog("PRINT", "{}]", options)
	}
}

// printUrlErrorAndUsage displays error messages and usage examples for the URL command
func printUrlErrorAndUsage() {
	printing.DalLog("ERROR", "Input target url", options)
	printing.DalLog("ERROR", "e.g dalfox url https://google.com/?q=1", options)
}

// init registers the URL command and applies custom help formatting
func init() {
	rootCmd.AddCommand(urlCmd)

	// Apply custom help format to this subcommand
	ApplySubCommandCustomHelp(urlCmd)
}
