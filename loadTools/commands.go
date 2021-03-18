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
	"fmt"
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)


func (at *TypeScribeArgs) LoadCommands(cmd *cobra.Command, subCmd bool) *ux.State {
	for range onlyOnce {
		if cmd == nil {
			break
		}

		var rootCmd = &cobra.Command {
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
			DisableFlagParsing: false,
		}
		at.SelfCmd = rootCmd


		var toolsCmd = &cobra.Command {
			Use:   CmdTools,
			Short: ux.SprintfMagenta(CmdRoot) + ux.SprintfBlue(" - Show all built-in template helpers."),
			Long:  ux.SprintfMagenta(CmdRoot) + ux.SprintfBlue(" - Show all built-in template helpers."),
			Run:   at.CmdTools,
			DisableFlagParsing: false,
		}

		var convertCmd = &cobra.Command {
			Use:   CmdConvert,
			Short: ux.SprintfMagenta(CmdRoot) + ux.SprintfBlue(" - Convert a template file to the resulting output file."),
			Long: ux.SprintfMagenta(CmdRoot) + ux.SprintfBlue(" - Convert a template file to the resulting output file."),
			Run: at.CmdConvert,
			DisableFlagParsing: false,
		}

		var loadCmd = &cobra.Command {
			Use:   CmdLoad,
			Short: ux.SprintfMagenta(CmdRoot) + ux.SprintfBlue(" - Load and execute a template file."),
			Long: ux.SprintfMagenta(CmdRoot) + ux.SprintfBlue(" - Load and execute a template file."),
			Run: at.CmdLoad,
			DisableFlagParsing: false,
		}

		var runCmd = &cobra.Command {
			Use:   CmdRun,
			Short: ux.SprintfMagenta(CmdRoot) + ux.SprintfBlue(" - Execute resulting output file as a BASH script."),
			Long: ux.SprintfMagenta(CmdRoot) + ux.SprintfBlue(`Execute resulting output file as a BASH script.
You can also use this command as the start to '#!' scripts.
For example: #!/usr/bin/env scribe --json gearbox.json run
`),
			Run: at.CmdRun,
			DisableFlagParsing: false,
		}


		var flags *pflag.FlagSet
		if subCmd {
			at.cmd = rootCmd
			cmd.AddCommand(rootCmd)
			flags = at.cmd.Flags()
			rootCmd.DisableFlagParsing = true
			toolsCmd.DisableFlagParsing = true
			convertCmd.DisableFlagParsing = true
			loadCmd.DisableFlagParsing = true
			runCmd.DisableFlagParsing = true
		} else {
			at.cmd = cmd
			flags = at.cmd.PersistentFlags()
		}
		flags = at.cmd.PersistentFlags()

		at.cmd.AddCommand(convertCmd)
		at.cmd.AddCommand(toolsCmd)
		at.cmd.AddCommand(loadCmd)
		at.cmd.AddCommand(runCmd)
		at.SetHelp(at.cmd)

		if at.Scribe.DefaultFile == "" {
			at.Scribe.DefaultFile = DefaultScribeFile
		}
		if at.Json.DefaultFile == "" {
			at.Json.DefaultFile = DefaultJsonFile
		}
		if at.Template.DefaultFile == "" {
			at.Template.DefaultFile = DefaultTemplateFile
		}
		if at.Output.DefaultFile == "" {
			at.Output.DefaultFile = DefaultOutFile
		}
		if at.WorkingPath.DefaultFile == "" {
			at.WorkingPath.DefaultFile = DefaultWorkingPath
		}

		flags.StringVarP(&at.Scribe.Value, FlagScribeFile, "s", at.Scribe.DefaultFile, ux.SprintfBlue("Alternative SCRIBE file."))
		flags.StringVarP(&at.Json.Value, FlagJsonFile, "j", at.Json.DefaultFile, ux.SprintfBlue("Alternative JSON file."))
		flags.StringVarP(&at.Template.Value, FlagTemplateFile, "t", at.Template.DefaultFile, ux.SprintfBlue("Alternative template file."))
		flags.StringVarP(&at.Output.Value, FlagOutputFile, "o", at.Output.DefaultFile, ux.SprintfBlue("Output file."))
		flags.StringVarP(&at.WorkingPath.Value, FlagWorkingPath, "p", at.WorkingPath.DefaultFile, ux.SprintfBlue("Set working path."))

		flags.BoolVarP(&at.Chdir, FlagChdir, "c", false, ux.SprintfBlue("Change to directory containing %s", DefaultJsonFile))
		flags.BoolVarP(&at.RemoveTemplate, FlagRemoveTemplate, "", false, ux.SprintfBlue("Remove template file afterwards."))
		flags.BoolVarP(&at.ForceOverwrite, FlagForce, "f", false, ux.SprintfBlue("Force overwrite of output files."))
		flags.BoolVarP(&at.RemoveOutput, FlagRemoveOutput, "", false, ux.SprintfBlue("Remove output file afterwards."))
		flags.BoolVarP(&at.QuietProgress, FlagQuiet, "q", false, ux.SprintfBlue("Silence progress in shell scripts."))
		flags.BoolVarP(&at.Verbose, FlagVerbose, "", false, ux.SprintfBlue("Verbose output."))

		flags.BoolVarP(&at.Debug, FlagDebug ,"d", false, ux.SprintfBlue("DEBUG mode."))

		flags.BoolVarP(&at.HelpAll, FlagHelpAll, "", false, ux.SprintfBlue("Show all help."))
		flags.BoolVarP(&at.HelpVariables, FlagHelpVariables, "", false, ux.SprintfBlue("Help on template variables."))
		flags.BoolVarP(&at.HelpFunctions, FlagHelpFunctions, "", false, ux.SprintfBlue("Help on template functions."))
		flags.BoolVarP(&at.HelpExamples, FlagHelpExamples, "", false, ux.SprintfBlue("Help on template examples."))

		cobra.EnableCommandSorting = false
	}
	return at.State
}


func (at *TypeScribeArgs) AddConfigOption(persistent bool, hidden bool) *ux.State {
	for range onlyOnce {
		if at.cmd == nil {
			at.State.SetError("Need to call LoadCommands first.")
			break
		}

		var fs *pflag.FlagSet
		if persistent {
			fs = at.cmd.PersistentFlags()
		} else {
			fs = at.cmd.Flags()
		}

		at.DiscoverConfigDir()
		if at.State.IsNotOk() {
			// No error - just ignore.
			at.State.SetOk()
			break
		}

		//fileDir := filepath.Join(at.Runtime.User.HomeDir, ".gearbox")
		//stat, err := os.Stat(fileDir)
		//if os.IsNotExist(err) {
		//	at.State.SetError("path does not exist - %s", err)
		//	break
		//}
		//if !stat.IsDir() {
		//	at.State.SetError("config file directory '%s' is not a directory", fileDir)
		//	break
		//}
		//
		//fileName := fmt.Sprintf("%s-config.json", prefix)
		//filePath := filepath.Join(fileDir, fileName)

		fs.StringVar(&at.ConfigFile, FlagConfigFile, at.ConfigPath, ux.SprintfBlue("Alternative command option config file."))

		_ = viper.BindPFlag(FlagScribeFile, at.cmd.Flags().Lookup(FlagScribeFile))
		_ = viper.BindPFlag(FlagJsonFile, at.cmd.Flags().Lookup(FlagJsonFile))
		_ = viper.BindPFlag(FlagTemplateFile, at.cmd.Flags().Lookup(FlagTemplateFile))
		_ = viper.BindPFlag(FlagOutputFile, at.cmd.Flags().Lookup(FlagOutputFile))
		_ = viper.BindPFlag(FlagWorkingPath, at.cmd.Flags().Lookup(FlagWorkingPath))

		_ = viper.BindPFlag(FlagChdir, at.cmd.Flags().Lookup(FlagChdir))
		_ = viper.BindPFlag(FlagRemoveTemplate, at.cmd.Flags().Lookup(FlagRemoveTemplate))
		_ = viper.BindPFlag(FlagForce, at.cmd.Flags().Lookup(FlagForce))
		_ = viper.BindPFlag(FlagRemoveOutput, at.cmd.Flags().Lookup(FlagRemoveOutput))
		_ = viper.BindPFlag(FlagQuiet, at.cmd.Flags().Lookup(FlagQuiet))
		_ = viper.BindPFlag(FlagVerbose, at.cmd.Flags().Lookup(FlagVerbose))

		_ = viper.BindPFlag(FlagDebug, at.cmd.Flags().Lookup(FlagDebug))

		_ = viper.BindPFlag(FlagHelpAll, at.cmd.Flags().Lookup(FlagHelpAll))
		_ = viper.BindPFlag(FlagHelpVariables, at.cmd.Flags().Lookup(FlagHelpVariables))
		_ = viper.BindPFlag(FlagHelpFunctions, at.cmd.Flags().Lookup(FlagHelpFunctions))
		_ = viper.BindPFlag(FlagHelpExamples, at.cmd.Flags().Lookup(FlagHelpExamples))

		cobra.OnInitialize(at.initConfig)

		if hidden {
			err := at.cmd.PersistentFlags().MarkHidden(FlagConfigFile)
			if err != nil {
				at.State.SetError(err)
			}
		}
	}

	return at.State
}


func (at *TypeScribeArgs) DiscoverConfigDir() *ux.State {
	if state := ux.IfNilReturnError(at); state.IsError() {
		return state
	}

	for range onlyOnce {
		var files []string
		files = append(files, filepath.Join(at.Runtime.CmdDir, at.Runtime.CmdFile + "-config.json"))
		files = append(files, filepath.Join(at.Runtime.CmdDir, "scribe-config.json"))
		files = append(files, filepath.Join(at.Runtime.CmdDir, "scribe.json"))

		files = append(files, filepath.Join(at.Runtime.User.HomeDir, ".gearbox", at.Runtime.CmdFile + "-config.json"))
		files = append(files, filepath.Join(at.Runtime.User.HomeDir, ".gearbox", "scribe-config.json"))
		files = append(files, filepath.Join(at.Runtime.User.HomeDir, ".gearbox", "scribe.json"))

		files = append(files, fmt.Sprintf("%c%s", filepath.Separator, filepath.Join("usr", "local", "gearbox", "etc", at.Runtime.CmdFile + "-config.json")))
		files = append(files, fmt.Sprintf("%c%s", filepath.Separator, filepath.Join("usr", "local", "gearbox", "etc", "scribe-config.json")))
		files = append(files, fmt.Sprintf("%c%s", filepath.Separator, filepath.Join("usr", "local", "gearbox", "etc", "scribe.json")))

		files = append(files, fmt.Sprintf("%c%s", filepath.Separator, filepath.Join("opt", "gearbox", "etc", at.Runtime.CmdFile + "-config.json")))
		files = append(files, fmt.Sprintf("%c%s", filepath.Separator, filepath.Join("opt", "gearbox", "etc", "scribe-config.json")))
		files = append(files, fmt.Sprintf("%c%s", filepath.Separator, filepath.Join("opt", "gearbox", "etc", "scribe.json")))

		at.ConfigDir = ""
		at.ConfigFile = ""
		for _, check := range files {
			at.isFileExisting(check)
			if at.State.IsNotOk() {
				continue
			}

			at.ConfigPath = check
			at.ConfigDir = filepath.Dir(check)
			at.ConfigFile = filepath.Base(check)
			at.State.SetOk()
			break
		}
	}

	//if at.State.IsNotOk() {
	//	at.State.SetWarning("No config file found")
	//}

	return at.State
}


func (at *TypeScribeArgs) isDirExisting(path string) *ux.State {
	if state := ux.IfNilReturnError(at); state.IsError() {
		return state
	}

	for range onlyOnce {
		stat, err := os.Stat(path)
		if os.IsNotExist(err) {
			at.State.SetWarning("path does not exist - %s", err)
			break
		}

		if !stat.IsDir() {
			at.State.SetWarning("config file directory '%s' is not a directory", path)
			break
		}
	}

	return at.State
}


func (at *TypeScribeArgs) isFileExisting(path string) *ux.State {
	if state := ux.IfNilReturnError(at); state.IsError() {
		return state
	}

	for range onlyOnce {
		at.isDirExisting(filepath.Dir(path))
		if at.State.IsNotOk() {
			break
		}

		stat, err := os.Stat(path)
		if os.IsNotExist(err) {
			at.State.SetWarning("path does not exist - %s", err)
			break
		}

		if stat.IsDir() {
			at.State.SetWarning("config file '%s' is a directory", path)
			break
		}
	}

	return at.State
}


// initConfig reads in config file and ENV variables if set.
func (at *TypeScribeArgs) initConfig() {
	if state := ux.IfNilReturnError(at); state.IsError() {
		return
	}
	var err error

	for range onlyOnce {
		if at.ConfigPath != "" {
			viper.SetConfigFile(at.ConfigPath)
			// Use config file from the flag.
		} else {
			// Search config in home directory with name "launch" (without extension).
			dir := filepath.Join(at.Runtime.User.HomeDir, ".gearbox")

			viper.AddConfigPath(dir)
			viper.SetConfigName("scribe")
		}

		//viper.AutomaticEnv() // read in environment variables that match

		// If a config file is found, read it in.
		err = viper.ReadInConfig()
		if err != nil {
			at.State.SetError(err)
			ux.PrintflnWarning("config file, (%s), error - %s", at.ConfigPath, err)
			break
		}

		//ux.Printf("using config file '%s'\n", viper.ConfigFileUsed())
		//_ = viper.WriteConfig()
	}
}


func (at *TypeScribeArgs) ReadConfig() *ux.State {
	if state := ux.IfNilReturnError(at); state.IsError() {
		return state
	}

	for range onlyOnce {
		// If a config file is found, read it in.
		err := viper.ReadInConfig()
		if err != nil {
			at.State.SetError(err)
			break
		}

		at.State.SetOk()
		//ux.Printf("using config file '%s'\n", viper.ConfigFileUsed())
		//_ = viper.WriteConfig()
	}

	return at.State
}


func (at *TypeScribeArgs) WriteConfig() *ux.State {
	if state := ux.IfNilReturnError(at); state.IsError() {
		return state
	}

	for range onlyOnce {
		err := viper.WriteConfig()
		if err != nil {
			at.State.SetError(err)
			break
		}

		at.State.SetOk()
	}

	return at.State
}


func (at *TypeScribeArgs) CmdRoot(cmd *cobra.Command, args []string) {
	if state := ux.IfNilReturnError(at); state.IsError() {
		return
	}

	for range onlyOnce {
		var ok bool
		ok, at.State = at.CheckSubCommand(cmd, args)
		if ok {
			break
		}

		if at.ParseScribeFlags(cmd) {
			break
		}

		at.ProcessArgs(cmd.Use, args)
		if at.State.IsNotOk() {
			_ = cmd.Help()
			break
		}

		at.Load()
		if at.State.IsNotOk() {
			at.State.PrintResponse()
			break
		}

		// Show help if no commands specified.
		if len(args) == 0 {
			_ = cmd.Help()
			at.State.SetOk()
			break
		}
	}

	return
}


func (at *TypeScribeArgs) CmdTools(cmd *cobra.Command, args []string) {
	if state := ux.IfNilReturnError(at); state.IsError() {
		return
	}

	for range onlyOnce {
		var ok bool
		ok, at.State = at.CheckSubCommand(cmd, args)
		if ok {
			break
		}

		at.State = at.ProcessArgs(cmd.Use, args)
		// Ignore errors as there's no args.

		at.Load()
		if at.State.IsNotOk() {
			at.State.PrintResponse()
			break
		}

		at.PrintTools()
		at.State.Clear()
	}

	return
}


func (at *TypeScribeArgs) CmdConvert(cmd *cobra.Command, args []string) {
	if state := ux.IfNilReturnError(at); state.IsError() {
		return
	}

	for range onlyOnce {
		var ok bool
		ok, at.State = at.CheckSubCommand(cmd, args)
		if ok {
			break
		}

		at.RemoveTemplate = true
		at.Output.Arg = SelectConvert

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

	return
}


func (at *TypeScribeArgs) CmdLoad(cmd *cobra.Command, args []string) {
	if state := ux.IfNilReturnError(at); state.IsError() {
		return
	}

	for range onlyOnce {
		var ok bool
		ok, at.State = at.CheckSubCommand(cmd, args)
		if ok {
			break
		}

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
		if at.State.IsNotOk() {
			break
		}
	}

	return
}


func (at *TypeScribeArgs) CmdRun(cmd *cobra.Command, args []string) {
	if state := ux.IfNilReturnError(at); state.IsError() {
		return
	}

	for range onlyOnce {
		var ok bool
		ok, at.State = at.CheckSubCommand(cmd, args)
		if ok {
			break
		}

		at.ExecShell = true
		at.Output.Arg = SelectConvert
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

	return
}


func (at *TypeScribeArgs) GetCmd() *cobra.Command {
	var ret *cobra.Command
	if state := at.IsNil(); state.IsError() {
		return ret
	}
	return at.SelfCmd
}


func (at *TypeScribeArgs) CmdHelp() *ux.State {
	if state := at.IsNil(); state.IsError() {
		return state
	}

	err := at.SelfCmd.Help()
	if err != nil {
		at.State.SetError(err)
	}
	return at.State
}


func (at *TypeScribeArgs) FlagGet(flag string) *pflag.Flag {
	var ret *pflag.Flag
	for range onlyOnce {
		ret = at.cmd.Flag(flag)
		if ret == nil {
			at.State.SetError("Unknown flag '%s'.", flag)
			break
		}
	}
	return ret
}


func (at *TypeScribeArgs) FlagHide(flag string) *ux.State {
	if state := ux.IfNilReturnError(at); state.IsError() {
		return state
	}

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

	return at.State
}


func (at *TypeScribeArgs) FlagSetDefault(flag string, def string) *ux.State {
	if state := ux.IfNilReturnError(at); state.IsError() {
		return state
	}

	for range onlyOnce {
		f := at.cmd.Flag(flag)
		if f == nil {
			at.State.SetError("Unknown flag '%s'.", flag)
			break
		}

		f.DefValue = def
	}

	return at.State
}
