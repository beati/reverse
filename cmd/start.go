package cmd

import (
	"github.com/beati/reverse/reverse"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:          "start",
	Short:        "Start starts the reverse proxy.",
	Long:         "Start starts the reverse proxy.",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		config := &reverse.Config{}
		err := viper.UnmarshalKey("Reverse", &config)
		if err != nil {
			return err
		}

		return reverse.Start(config)
	},
}

func init() {
	RootCmd.AddCommand(startCmd)
}
