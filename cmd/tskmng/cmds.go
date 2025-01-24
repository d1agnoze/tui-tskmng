package tskmng

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

const versionFile = "VERSION"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Tmg version number",
	Run: func(cmd *cobra.Command, args []string) {
		raw, err := os.ReadFile("VERSION")
		if err != nil || raw == nil {
			log.Info("Version: v0.0-dev")
		} else {
			log.Infof("Version: %s", string(raw))
		}
	},
}
