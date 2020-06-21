/*
Commands available:
	- scribe <args> run [files]
			- Sets

	- scribe <args> load [files]
			- .

	- scribe <args> convert [files]
			- .

	- scribe help [option]
			- functions	- Show available helper functions.
			- variables	- Show available variables.
			- examples	- Show examples.
			- all		- Show all help.

Where <args> is:
  -c, --chdir             Change to directory containing scribe.json
  -d, --debug             DEBUG mode.
  -f, --force             Force overwrite of output files.
  -h, --help              help for scribe
      --help-all          Show all help.
      --help-examples     Help on template examples.
      --help-functions    Help on template functions.
      --help-variables    Help on template variables.
      --json <file>       Alternative JSON file. (default "scribe.json")
      --out <file>        Output file. (default "/dev/stdout")
  -q, --quiet             Silence progress in shell scripts.
  -o, --rm-out            Remove output file afterwards.
  -r, --rm-tmpl           Remove template file afterwards.
      --template <file>   Alternative template file. (default "scribe.tmpl")
  -v, --version           Display version of scribe

Where [file] is:
	- [filename.json|filename.tmpl]	- Specify either JSON or template file and the other will be auto-discovered.
									- Can use --json or --template to help auto-discovery.
	- <filename.json filename.tmpl>	- Specify both JSON and template file.
	- <ENV_JSON ENV_TEMPLATE>		-

*/
package loadTools

import (
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/spf13/cobra"
)


func (at *TypeScribeArgs) LoadCommands(cmd *cobra.Command) *ux.State {
	for range onlyOnce {
		var toolsCmd = &cobra.Command{
			Use:   CmdTools,
			Short: ux.SprintfMagenta("scribe") + ux.SprintfBlue(" - Show all built-in template helpers."),
			Long:  ux.SprintfMagenta("scribe") + ux.SprintfBlue(" - Show all built-in template helpers."),
			Run:   at.cmdTools,
		}

		var convertCmd = &cobra.Command{
			Use:   CmdConvert,
			Short: ux.SprintfMagenta("scribe") + ux.SprintfBlue(" - Convert a template file to the resulting output file."),
			Long: ux.SprintfMagenta("scribe") + ux.SprintfBlue(" - Convert a template file to the resulting output file."),
			Run: at.cmdConvert,
		}

		var loadCmd = &cobra.Command{
			Use:   CmdLoad,
			Short: ux.SprintfMagenta("scribe") + ux.SprintfBlue(" - Load and execute a template file."),
			Long: ux.SprintfMagenta("scribe") + ux.SprintfBlue(" - Load and execute a template file."),
			Run: at.cmdLoad,
			DisableFlagParsing: false,
		}

		var runCmd = &cobra.Command{
			Use:   CmdRun,
			Short: ux.SprintfMagenta("scribe") + ux.SprintfBlue(" - Execute resulting output file as a BASH script."),
			Long: ux.SprintfMagenta("scribe") + ux.SprintfBlue(`Execute resulting output file as a BASH script.
You can also use this command as the start to '#!' scripts.
For example: #!/usr/bin/env scribe --json gearbox.json run
`),
			Run: at.cmdRun,
		}

		cmd.AddCommand(convertCmd)
		cmd.AddCommand(toolsCmd)
		cmd.AddCommand(loadCmd)
		cmd.AddCommand(runCmd)
	}
	return at.State
}


func (at *TypeScribeArgs) cmdTools(cmd *cobra.Command, args []string) {
	for range onlyOnce {
		at.State = at.ProcessArgs(cmd.Use, args)
		// Ignore errors as there's no args.

		at.PrintTools()
		at.State.Clear()
	}
}


func (at *TypeScribeArgs) cmdConvert(cmd *cobra.Command, args []string) {
	for range onlyOnce {
		at.RemoveTemplate = true
		at.Output.File = SelectConvert

		at.State = at.ProcessArgs(cmd.Use, args)
		if at.State.IsNotOk() {
			break
		}

		at.State = at.Load()
		if at.State.IsNotOk() {
			break
		}

		at.PrintflnNotify("Converting file '%s' => '%s'", at.Template.GetPath(), at.Output.GetPath())
		at.State = at.Run()
	}
}


func (at *TypeScribeArgs) cmdLoad(cmd *cobra.Command, args []string) {
	for range onlyOnce {
		at.State = at.ProcessArgs(cmd.Use, args)
		if at.State.IsNotOk() {
			break
		}

		at.State = at.Load()
		if at.State.IsNotOk() {
			break
		}

		at.PrintflnNotify("Loading template '%s' and saving result to '%s'", at.Template.GetPath(), at.Output.GetPath())
		at.State = at.Run()
	}
}


func (at *TypeScribeArgs) cmdRun(cmd *cobra.Command, args []string) {
	for range onlyOnce {
		at.ExecShell = true
		at.Output.File = SelectConvert
		at.StripHashBang = true

		/*
			Allow this to be used as a UNIX script.
			The following should be placed on the first line.
			#!/usr/bin/env scribe load
		*/

		at.State = at.ProcessArgs(cmd.Use, args)
		if at.State.IsNotOk() {
			break
		}

		at.State = at.Load()
		if at.State.IsNotOk() {
			break
		}

		at.PrintflnNotify("Loading template '%s' and saving result to '%s'", at.Template.GetPath(), at.Output.GetPath())
		at.State = at.Run()
	}
}
