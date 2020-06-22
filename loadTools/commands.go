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


func (at *TypeScribeArgs) LoadCommands(cmd *cobra.Command, subCmd bool) *ux.State {
	for range onlyOnce {
		if cmd == nil {
			break
		}

		var rootCmd = &cobra.Command{
			Use:   CmdRoot,
			Short: ux.SprintfMagenta(CmdRoot) + ux.SprintfBlue(" - The ultimate scripting toolkit."),
			Long: ux.SprintfBlue(`The ultimate scripting toolkit.
Feed me a JSON and GoLang template file, I'll do the rest.

This utility allows generation of any config file or script based off GoLang templates.
Helper functions provide generic access to O/S, Git, file copying and SSH.
See help for further information:
`) +
				ux.SprintfWhite("  Functions: %s --FlagHelpFunctions\n", CmdRoot) +
				ux.SprintfWhite("  Variables: %s --FlagHelpVariables\n", CmdRoot) +
				ux.SprintfWhite("  Examples:  %s --FlagHelpExamples\n", CmdRoot),
			Run: at.CmdRoot,
		}

		var toolsCmd = &cobra.Command{
			Use:   CmdTools,
			Short: ux.SprintfMagenta(CmdRoot) + ux.SprintfBlue(" - Show all built-in template helpers."),
			Long:  ux.SprintfMagenta(CmdRoot) + ux.SprintfBlue(" - Show all built-in template helpers."),
			Run:   at.CmdTools,
		}

		var convertCmd = &cobra.Command{
			Use:   CmdConvert,
			Short: ux.SprintfMagenta(CmdRoot) + ux.SprintfBlue(" - Convert a template file to the resulting output file."),
			Long: ux.SprintfMagenta(CmdRoot) + ux.SprintfBlue(" - Convert a template file to the resulting output file."),
			Run: at.CmdConvert,
		}

		var loadCmd = &cobra.Command{
			Use:   CmdLoad,
			Short: ux.SprintfMagenta(CmdRoot) + ux.SprintfBlue(" - Load and execute a template file."),
			Long: ux.SprintfMagenta(CmdRoot) + ux.SprintfBlue(" - Load and execute a template file."),
			Run: at.CmdLoad,
			DisableFlagParsing: false,
		}

		var runCmd = &cobra.Command{
			Use:   CmdRun,
			Short: ux.SprintfMagenta(CmdRoot) + ux.SprintfBlue(" - Execute resulting output file as a BASH script."),
			Long: ux.SprintfMagenta(CmdRoot) + ux.SprintfBlue(`Execute resulting output file as a BASH script.
You can also use this command as the start to '#!' scripts.
For example: #!/usr/bin/env scribe --json gearbox.json run
`),
			Run: at.CmdRun,
		}


		if subCmd {
			at.cmd = rootCmd
			cmd.AddCommand(rootCmd)
		} else {
			at.cmd = cmd
		}


		//if subCmd {
		//	at.cmd.AddCommand(rootCmd)
		//	at.SetHelp(rootCmd)
		//
		//	rootCmd.AddCommand(convertCmd)
		//	rootCmd.AddCommand(toolsCmd)
		//	rootCmd.AddCommand(loadCmd)
		//	rootCmd.AddCommand(runCmd)
		//} else {
			at.cmd.AddCommand(convertCmd)
			at.cmd.AddCommand(toolsCmd)
			at.cmd.AddCommand(loadCmd)
			at.cmd.AddCommand(runCmd)
			at.SetHelp(at.cmd)
		//}

		at.cmd.Flags().StringVarP(&at.Scribe.File, FlagScribeFile, "s", DefaultScribeFile, ux.SprintfBlue("Alternative SCRIBE file."))
		at.cmd.Flags().StringVarP(&at.Json.File, FlagJsonFile, "j", DefaultJsonFile, ux.SprintfBlue("Alternative JSON file."))
		at.cmd.Flags().StringVarP(&at.Template.File, FlagTemplateFile, "t", DefaultTemplateFile, ux.SprintfBlue("Alternative template file."))
		at.cmd.Flags().StringVarP(&at.Output.File, FlagOutputFile, "o", DefaultOutFile, ux.SprintfBlue("Output file."))
		at.cmd.Flags().StringVarP(&at.WorkingPath.File, FlagWorkingPath, "p", DefaultWorkingPath, ux.SprintfBlue("Set working path."))

		at.cmd.Flags().BoolVarP(&at.Chdir, FlagChdir, "c", false, ux.SprintfBlue("Change to directory containing %s", DefaultJsonFile))
		at.cmd.Flags().BoolVarP(&at.RemoveTemplate, FlagRemoveTemplate, "", false, ux.SprintfBlue("Remove template file afterwards."))
		at.cmd.Flags().BoolVarP(&at.ForceOverwrite, FlagForce, "f", false, ux.SprintfBlue("Force overwrite of output files."))
		at.cmd.Flags().BoolVarP(&at.RemoveOutput, FlagRemoveOutput, "", false, ux.SprintfBlue("Remove output file afterwards."))
		at.cmd.Flags().BoolVarP(&at.QuietProgress, FlagQuiet, "q", false, ux.SprintfBlue("Silence progress in shell scripts."))
		at.cmd.Flags().BoolVarP(&at.Verbose, FlagVerbose, "", false, ux.SprintfBlue("Verbose output."))

		at.cmd.Flags().BoolVarP(&at.Debug, FlagDebug ,"d", false, ux.SprintfBlue("DEBUG mode."))

		at.cmd.Flags().BoolVarP(&at.HelpAll, FlagHelpAll, "", false, ux.SprintfBlue("Show all help."))
		at.cmd.Flags().BoolVarP(&at.HelpVariables, FlagHelpVariables, "", false, ux.SprintfBlue("Help on template variables."))
		at.cmd.Flags().BoolVarP(&at.HelpFunctions, FlagHelpFunctions, "", false, ux.SprintfBlue("Help on template functions."))
		at.cmd.Flags().BoolVarP(&at.HelpExamples, FlagHelpExamples, "", false, ux.SprintfBlue("Help on template examples."))
	}
	return at.State
}


func (at *TypeScribeArgs) CmdRoot(cmd *cobra.Command, args []string) {
	for range onlyOnce {
		if at.ParseScribeFlags(cmd) {
			break
		}

		at.ProcessArgs(cmd.Use, args)
		if at.State.IsNotOk() {
			_ = cmd.Help()
			break
		}

		// Show help if no commands specified.
		if len(args) == 0 {
			_ = cmd.Help()
			at.State.SetOk()
			break
		}
	}
}


func (at *TypeScribeArgs) CmdTools(cmd *cobra.Command, args []string) {
	for range onlyOnce {
		at.State = at.ProcessArgs(cmd.Use, args)
		// Ignore errors as there's no args.

		at.PrintTools()
		at.State.Clear()
	}
}


func (at *TypeScribeArgs) CmdConvert(cmd *cobra.Command, args []string) {
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


func (at *TypeScribeArgs) CmdLoad(cmd *cobra.Command, args []string) {
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


func (at *TypeScribeArgs) CmdRun(cmd *cobra.Command, args []string) {
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


func (at *TypeScribeArgs) FlagHide(flag string) {
	for range onlyOnce {
		//err := at.cmd.Flags().MarkHidden(flag)
		//if err != nil {
		//	at.State.SetError(err)
		//	break
		//}

		f := at.cmd.Flag(flag)
		if f == nil {
			at.State.SetError("Unknown flag '%s'.", flag)
			break
		}

		f.Hidden = true
	}
}


func (at *TypeScribeArgs) FlagSetDefault(flag string, def string) {
	for range onlyOnce {
		f := at.cmd.Flag(flag)
		if f == nil {
			at.State.SetError("Unknown flag '%s'.", flag)
			break
		}

		f.DefValue = def
	}
}
