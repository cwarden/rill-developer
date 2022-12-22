package docs

import (
	"fmt"

	"github.com/rilldata/rill/cli/pkg/browser"
	"github.com/spf13/cobra"
)

var docsURL = "https://docs.rilldata.com"

// docsCmd represents the docs command.
func DocsCmd() *cobra.Command {
	docsCmd := &cobra.Command{
		Use:   "docs",
		Short: "Open docs.rilldata.com",
		Run: func(cmd *cobra.Command, args []string) {
			err := browser.Open(docsURL)
			if err != nil {
				fmt.Printf("Could not open browser. Copy this URL into your browser: %s\n", docsURL)
			}
		},
	}
	return docsCmd
}
