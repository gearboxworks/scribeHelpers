package toolCobraHelp

import (
	"fmt"
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/spf13/cobra"
	"strings"
)


func (tc *TypeCommands) AddCommands(name string, parent *cobra.Command, child ...*cobra.Command) {
	for range onlyOnce {
		if tc.Commands == nil {
			tc.Commands = make(Cmds)
		}
		if tc.Commands[name] == nil {
			tc.Commands[name] = []*cobra.Command{}
		}

		parent.AddCommand(child...)

		for _, c := range child {
			tc.Commands[name] = append(tc.Commands[name], c)
		}
	}
}


func (tc *TypeCommands) HelpCommand(section string) string {
	var ret string
	for range onlyOnce {
		if _, ok := tc.Commands[section]; !ok {
			break
		}

		for _, c := range tc.Commands[section] {
			ret += ux.SprintfGreen("\t\t%s", c.Name())
			ret += ux.SprintfBlue(" - %s\n", c.Short)
		}

		tc.State.SetOk()
	}
	return ret
}


func (tc *TypeCommands) HelpCommands() string {
	var ret string
	for range onlyOnce {
		for n, _ := range tc.Commands {
			ret += ux.SprintfCyan("\t%s\n", n)
			ret += tc.HelpCommand(n)
			//ret += ux.SprintfCyan("\n")
		}

		//{{- range .Commands }}
		//{{- if (or .IsAvailableCommand (eq .Name "help")) }}
		//{{ rpad (SprintfGreen .Name) .NamePadding}}	- {{ .Short }}{{ end }}
		//{{- end }}
	}
	return ret
}


func (tc *TypeCommands) ParseHelpFlags(cmd *cobra.Command) bool {
	var ok bool
	for range onlyOnce {
		fl := cmd.Flags()

		//// Show HelpVariables.
		//ok, _ = fl.GetBool(FlagHelpVariables)
		//if ok {
		//	tc.ShowHelpVariables()
		//	break
		//}
		//
		//// Show HelpFunctions.
		//ok, _ = fl.GetBool(FlagHelpFunctions)
		//if ok {
		//	tc.ShowHelpFunctions()
		//	break
		//}

		// Show HelpExamples.
		ok, _ = fl.GetBool(FlagHelpExamples)
		if ok {
			tc.ShowHelpExamples()
			break
		}

		// Show all help.
		ok, _ = fl.GetBool(FlagHelpAll)
		if ok {
			tc.ShowHelpAll()
			break
		}
	}
	return ok
}


func (tc *TypeCommands) _GetUsage(c *cobra.Command) string {
	var str string

	if c.Parent() == nil {
		str += ux.SprintfCyan("%s [flags] ", c.Name())
	} else {
		str += ux.SprintfCyan("%s [flags] ", c.Parent().Name())
		str += ux.SprintfGreen("%s ", c.Use)
	}

	if c.HasAvailableSubCommands() {
		str += ux.SprintfGreen("[command] ")
		str += ux.SprintfCyan("<args> ")
	}

	return str
}


func (tc *TypeCommands) _GetVersion(c *cobra.Command) string {
	var str string

	if c.Parent() == nil {
		str = ux.SprintfBlue("%s ", tc.runtime.CmdName)
		str += ux.SprintfCyan("v%s", tc.runtime.CmdVersion)
	}

	return str
}


func (tc *TypeCommands) SetHelp(c *cobra.Command) {
	var tmplHelp string
	var tmplUsage string

	//fmt.Printf("%s", rootCmd.UsageTemplate())
	//fmt.Printf("%s", rootCmd.HelpTemplate())

	cobra.AddTemplateFunc("GetUsage", tc._GetUsage)
	cobra.AddTemplateFunc("GetVersion", tc._GetVersion)

	cobra.AddTemplateFunc("HelpCommands", tc.HelpCommands)

	cobra.AddTemplateFunc("SprintfBlue", ux.SprintfBlue)
	cobra.AddTemplateFunc("SprintfCyan", ux.SprintfCyan)
	cobra.AddTemplateFunc("SprintfGreen", ux.SprintfGreen)
	cobra.AddTemplateFunc("SprintfMagenta", ux.SprintfMagenta)
	cobra.AddTemplateFunc("SprintfRed", ux.SprintfRed)
	cobra.AddTemplateFunc("SprintfWhite", ux.SprintfWhite)
	cobra.AddTemplateFunc("SprintfYellow", ux.SprintfYellow)

	// 	{{ with .Parent }}{{ SprintfCyan .Name }}{{ end }} {{ SprintfGreen .Name }} {{ if .HasAvailableSubCommands }}{{ SprintfGreen "[command]" }}{{ end }}

	// {{ HelpCommands }}

	//{{- range .Commands }}
	// {{- if (or .IsAvailableCommand (eq .Name "help")) }}
	// {{ rpad (SprintfGreen .Name) .NamePadding}}	- {{ .Short }}{{ end }}
	//{{- end }}

	tmplUsage += `
{{ SprintfBlue "Usage: " }}
	{{ GetUsage . }}

{{- if gt (len .Aliases) 0 }}
{{ SprintfBlue "\nAliases:" }} {{ .NameAndAliases }}
{{- end }}

{{- if .HasExample }}
{{ SprintfBlue "\nExamples:" }}
	{{ .Example }}
{{- end }}

{{- if .HasAvailableSubCommands }}
{{ SprintfBlue "\nWhere " }}{{ SprintfGreen "[command]" }}{{ SprintfBlue " is one of:" }}
{{- range .Commands }}
{{- if (or .IsAvailableCommand (eq .Name "help")) }}
	{{ rpad (SprintfGreen .Name) .NamePadding}}     	- {{ .Short }}{{ end }}
{{- end }}
{{- end }}


{{- if .HasAvailableLocalFlags }}
{{ SprintfBlue "\nFlags:" }}
{{ .LocalFlags.FlagUsages | trimTrailingWhitespaces }}
{{- end }}

{{- if .HasAvailableInheritedFlags }}
{{ SprintfBlue "\nGlobal Flags:" }}
{{ .InheritedFlags.FlagUsages | trimTrailingWhitespaces }}
{{- end }}

{{- if .HasHelpSubCommands }}
{{- SprintfBlue "\nAdditional help topics:" }}
{{- range .Commands }}
{{- if .IsAdditionalHelpTopicCommand }}
	{{ rpad (SprintfGreen .CommandPath) .CommandPathPadding }} {{ .Short }}
{{- end }}
{{- end }}
{{- end }}

{{- if .HasAvailableSubCommands }}
{{ SprintfBlue "\nUse" }} {{ SprintfCyan .CommandPath }} {{ SprintfCyan "help" }} {{ SprintfGreen "[command]" }} {{ SprintfBlue "for more information about a command." }}
{{- end }}
`

	tmplHelp = `{{ GetVersion . }}

{{ SprintfBlue "Commmand:" }} {{ SprintfCyan .Use }}

{{ SprintfBlue "Description:" }} 
	{{ with (or .Long .Short) }}
{{- . | trimTrailingWhitespaces }}
{{- end }}

{{- if or .Runnable .HasSubCommands }}
{{ .UsageString }}
{{- end }}
`

	//c.SetHelpCommand(c)
	//c.SetHelpFunc(PrintHelp)
	c.SetHelpTemplate(tmplHelp)
	c.SetUsageTemplate(tmplUsage)
}


func (tc *TypeCommands) ShowHelpAll() {
	for range onlyOnce {
		tc.ShowHelpExamples()
	}

	tc.State.SetOk()
}


func (tc *TypeCommands) ShowHelpVariables() {
	for range onlyOnce {
		ux.PrintfBlue("Keys accessible within your template file:\n")
	}

	tc.State.SetOk()
}


func (tc *TypeCommands) ShowHelpExamples() {
	for range onlyOnce {
		var examples Examples

		//examples = append(examples, Example {
		//	Command: "run",
		//	Args:    []string{"MyScript.sh.tmpl", "config.json"},
		//	Info:    "Same again using 'run'. This will execute the MyScript.sh output file afterwards.",
		//})


		ux.PrintflnBlue("Examples:")
		for _, v := range examples {
			fmt.Printf("# %s\n\t%s %s\n\n",
				ux.SprintfBlue(v.Info),
				ux.SprintfCyan("%s %s", tc.runtime.Cmd, v.Command),
				ux.SprintfWhite(strings.Join(v.Args, " ")),
			)
		}
	}

	tc.State.SetOk()
}


type Example struct {
	Command string
	Args []string
	Info string
}
type Examples []Example


//func (at *TypeCommands) AddExample(example Example) {
//	for range onlyOnce {
//		at.HelpExamples = append(at.HelpExamples, example)
//	}
//}