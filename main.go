package main

import (
	"errors"
	"flag"
	"log"

	"github.com/18F/cf-db-connect/connector"

	"code.cloudfoundry.org/cli/plugin"
)

const SUBCOMMAND = "connect-to-db"

// DBConnectPlugin is the struct implementing the interface defined by the core CLI. It can
// be found at  "code.cloudfoundry.org/cli/plugin/plugin.go"
type DBConnectPlugin struct{}

func (c *DBConnectPlugin) parseOptions(args []string) (options connector.Options, err error) {
	metadata := c.GetMetadata()
	command := metadata.Commands[0]
	flags := flag.NewFlagSet(command.Name, flag.ExitOnError)

	err = flags.Parse(args[1:])
	if err != nil {
		return
	}

	nonFlagArgs := flags.Args()
	if len(nonFlagArgs) != 2 {
		err = errors.New("Wrong number of arguments")
		return
	}

	options = connector.Options{
		AppName:             nonFlagArgs[0],
		ServiceInstanceName: nonFlagArgs[1],
	}
	return
}

// Run is the entry point when the core CLI is invoking a command defined
// by the plugin. The first parameter, plugin.CliConnection, is a struct that can
// be used to invoke cli commands. The second paramter, args, is a slice of
// strings. args[0] will be the name of the command, and will be followed by
// any additional arguments a cli user typed in.
func (c *DBConnectPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	// check to ensure it's the right subcommand, not others like CLI-MESSAGE-UNINSTALL
	if args[0] != SUBCOMMAND {
		return
	}

	opts, err := c.parseOptions(args)
	if err != nil {
		log.Fatalln(err)
	}

	err = connector.Connect(cliConnection, opts)
	if err != nil {
		log.Fatalln(err)
	}
}

func (c *DBConnectPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "DBConnect",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 15,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     SUBCOMMAND,
				HelpText: "Open a shell that's connected to a database service instance",
				UsageDetails: plugin.Usage{
					Usage: "\n   cf " + SUBCOMMAND + " <app_name> <service_instance_name>",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(DBConnectPlugin))
}
