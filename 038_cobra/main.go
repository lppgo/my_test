package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"conbra-demo/cmd/apply"
	"conbra-demo/cmd/check"
	"conbra-demo/config"
)

func main() {
	var cmdCheck = check.NewVPARequestCheck()
	var cmdApply = apply.NewVPARequestApply()

	var rootCmd = &cobra.Command{Use: "example"}

	flags := rootCmd.PersistentFlags()
	addFlags(flags)

	rootCmd.AddCommand(cmdApply, cmdCheck)
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func addFlags(flags *pflag.FlagSet) {
	flags.StringVar(&config.Cfg.KubeConfig, "config", "none", "config file[/.xx.yaml]")
	flags.StringVar(&config.Cfg.Mode, "mode", "cpu", "mode[cpu  or all]")
}

/*


Usage:
  example [command]

Available Commands:
  apply       apply request by json file
  check       check request validity by json file
  completion  generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
      --config string   config file[/.xx.yaml] (default "none")
  -h, --help            help for example
      --mode string     mode[cpu  or all] (default "cpu")

Use "example [command] --help" for more information about a command.

*/
