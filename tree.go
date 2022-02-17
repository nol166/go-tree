package main

import (
	"fmt"

	"github.com/thatisuday/commando"
)

func main() {

	// configure the commando
	commando.SetExecutableName("tree").
		SetDescription("A tree command").
		SetVersion("0.0.1")

	// configure the "root" command
	commando.Register(nil).AddArgument("dir", "local directory path", "./"). // default value is "./"
											AddFlag("level,l", "level of depth to travel", commando.Int, 1). // default value is 1
											AddFlag("size,s", "size of the tree", commando.Bool, nil).       // default value is false
											SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			// print a message about the options
			fmt.Printf("Printing options of the 'root' command...\n\n")

			// print the arguments
			for k, v := range args {
				fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
			}

			// print the flags
			for k, v := range flags {
				fmt.Printf("flag -> %v: %v(%T)\n", k, v.Value, v.Value)
			}
		})

	// configure the info command
	commando.Register("info").SetShortDescription("Prints information about the directory").
		SetDescription("This command displays more information about the contents of the directory like size, permission and ownership, etc.").
		AddArgument("dir", "local directory path", "./").                  // default value is "./"
		AddFlag("level,l", "level of depth to travel", commando.Int, nil). // required
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			// print a message about the options
			fmt.Printf("Printing options of the 'info' command...\n\n")

			// print the arguments
			for k, v := range args {
				fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
			}

			// print the flags
			for k, v := range flags {
				fmt.Printf("flag -> %v: %v(%T)\n", k, v.Value, v.Value)
			}
		})

	// parse command line arguments
	commando.Parse(nil)
}
