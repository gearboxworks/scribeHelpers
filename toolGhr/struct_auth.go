package toolGhr

import (
	"github.com/gearboxworks/scribeHelpers/toolRuntime"
	"github.com/gearboxworks/scribeHelpers/ux"
	"os"
)


type TypeAuth struct {
	Token    string		// `goptions:"-s, --security-token, description='Github token ($GITHUB_TOKEN if set). required if repo is private.'"`
	//User     string `goptions:"-u, --user, description='Github repo user or organisation (required if $GITHUB_USER not set)'"`
	AuthUser string		// `goptions:"-a, --Auth-user, description='Username for authenticating to the API (falls back to $GITHUB_AUTH_USER or $GITHUB_USER)'"`

	runtime  *toolRuntime.TypeRuntime
	state    *ux.State
}
func (auth *TypeAuth) IsNil() *ux.State {
	return ux.IfNilReturnError(auth)
}

func NewAuth(runtime *toolRuntime.TypeRuntime) *TypeAuth {
	var auth TypeAuth
	runtime = runtime.EnsureNotNil()

	for range onlyOnce {
		auth = TypeAuth{
			Token:    os.Getenv("GITHUB_TOKEN"),
			//User:     os.Getenv("GITHUB_USER"),
			AuthUser: os.Getenv("GITHUB_AUTH_USER"),

			runtime: runtime,
			state:   ux.NewState(runtime.CmdName, runtime.Debug),
		}
		if auth.AuthUser == "" {
			auth.AuthUser = os.Getenv("GITHUB_USER")
		}
	}
	auth.state.SetPackage("")
	auth.state.SetFunctionCaller()
	return &auth
}

func (auth *TypeAuth) isValid() *ux.State {
	if state := ux.IfNilReturnError(auth); state.IsError() {
		return state
	}

	for range onlyOnce {
		auth.state = auth.state.EnsureNotNil()

		if auth.Token == "" {
			auth.state.SetError("$GITHUB_TOKEN is empty")
			break
		}

		//if Auth.User == "" {
		//	Auth.state.SetError("$GITHUB_USER is empty")
		//	break
		//}

		if auth.AuthUser == "" {
			auth.state.SetError("$GITHUB_AUTH_USER is empty")
			break
		}
	}

	return auth.state
}

func (auth *TypeAuth) Set(a TypeAuth) *ux.State {
	if state := auth.IsNil(); state.IsError() {
		return state
	}
	auth.state.SetFunction()

	for range onlyOnce {
		auth.state = a.isValid()
		if auth.state.IsNotOk() {
			break
		}

		//options := Options{}
		//goptions.ParseAndFail(&options)
		//
		//if options.Version {
		//	fmt.Printf("github-Release v%s\n", github.VERSION)
		//	return nil
		//}
		//
		//if len(options.Verbs) == 0 {
		//	goptions.PrintHelp()
		//	return nil
		//}
		//
		//VERBOSITY = len(options.Verbosity)
		//github.VERBOSITY = VERBOSITY
		//
		//if cmd, found := commands[options.Verbs]; found {
		//	err := cmd(options)
		//	if err != nil {
		//		if !options.Quiet {
		//			fmt.Fprintln(os.Stderr, "error:", err)
		//		}
		//		return err
		//	}
		//}
	}

	return auth.state
}
