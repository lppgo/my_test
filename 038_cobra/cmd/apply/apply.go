package apply

import (
	"github.com/spf13/cobra"
)

type RequestApplyOptions struct {
	configPath string
}

func NewRequestApplyOptions() *RequestApplyOptions {
	o := &RequestApplyOptions{}

	return o
}

func NewVPARequestApply() *cobra.Command {
	o := NewRequestApplyOptions()

	cmd := &cobra.Command{
		Use:   "apply [json file]",
		Short: "apply request by json file",
		Long:  "apply request by new request json file",
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
func (c *RequestApplyOptions) Run(args []string) error {
	return nil
}
