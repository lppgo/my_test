package check

import (
	"github.com/spf13/cobra"
)

type RequestCheckOptions struct {
	configPath string
}

func NewRequestCheckOptions() *RequestCheckOptions {
	o := &RequestCheckOptions{}

	return o
}

func NewVPARequestCheck() *cobra.Command {
	o := NewRequestCheckOptions()
	cmd := &cobra.Command{
		Use:   "check [json file]",
		Short: "check request validity by json file",
		Long:  "check request by new request json file",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Run(args); err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}

// Run 具体的实现
func (c *RequestCheckOptions) Run(args []string) error {
	return nil
}
