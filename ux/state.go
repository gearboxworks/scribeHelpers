package ux

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)


type StateGetter interface {
	Print()
	IsError() bool
	IsWarning() bool
	IsOk() bool
	SetError(format string, args ...interface{})
	SetWarning(format string, args ...interface{})
	SetOk(format string, args ...interface{})
	ClearError()
	ClearAll()
	IsRunning() bool
	IsPaused() bool
	IsCreated() bool
	IsRestarting() bool
	IsRemoving() bool
	IsExited() bool
	IsDead() bool
	SetString(s string)
}

type State struct {
	prefix      string
	prefixArray []string
	_Executable string
	_Package    string
	_Function   string

	_Fatal      error
	_Error      error
	_Warning    error
	_Ok         error
	_Debug      error
	ExitCode    int
	debug       RuntimeDebug
	verbose     bool

	RunState    string

	Output      string
	_Separator  string
	OutputArray []string
	response    TypeResponse
	//Response    interface{}
}


const DefaultSeparator = "\n"


func NewState(name string, debugMode bool) *State {
	me := State{}
	me.Clear()
	me.debug.Enabled = debugMode
	me._Executable = name	// @TODO - Add this to debugging.
	//if debugMode {
	//	PrintflnWarning("%s - DEBUG MODE ENABLED", name)
	//}

	return &me
}

func (state *State) EnsureNotNil() *State {
	for range onlyOnce {
		if state == nil {
			state = NewState("", false)
		}
		state.SetOk()
	}
	return state
}


func EnsureStateNotNil(p *State) *State {
	for range onlyOnce {
		if p == nil {
			p = NewState("", false)
		}
		p.Clear()
	}
	return p
}


func IsInterfaceNil(ref interface{}) bool {
	var ok bool
	for range onlyOnce {
		if ref == nil {
			ok = true
			break
		}

		if reflect.ValueOf(ref).Kind() != reflect.Ptr {
			break
		}

		rv := reflect.ValueOf(ref)
		r := rv.IsNil()
		if r {
			ok = true
			break
		}
		if rv.Interface() == nil {
			ok = true
			break
		}

		e := rv.Elem()
		//PrintflnCyan("PTR==%s", e.Kind().String())
		if e.Kind().String() == "ptr" {
			//PrintflnCyan("POINTER TO POINTER")
			ok = IsInterfaceNil(e.Addr().Elem().Interface())
			break
		}

		// @TODO - Something fishy going on here.
		// When we have been given a pointer, we need to be able to discover if it's nil.
		// DEBUG
		//l := reflect.ValueOf(ref).Kind()
		//l := e.Kind()
		//fmt.Printf("HERE -> %s %s\n",
		//	rv.Kind().String(),
		//	rv.Elem().Kind().String(),
		//	)
		//
		//switch v := ref.(type) {
		//	case nil:
		//		fmt.Printf("YES! %s %v %v %v %v %v\n", l, v, rv.IsValid(), rv.IsNil(), rv.IsZero(), e.Len())
		//	case interface{}:
		//		fmt.Printf("%s %v %v %v %v %v\n", l, v, rv.IsValid(), rv.IsNil(), rv.IsZero(), e.Len())
		//		test := v.(interface{})
		//		if test == nil {
		//			fmt.Printf("YES! %s %v %v %v %v %v\n", l, v, rv.IsValid(), rv.IsNil(), rv.IsZero(), e.Len())
		//		}
		//	default:
		//		fmt.Printf("%s %v %v %v %v %v\n", l, v, rv.IsValid(), rv.IsNil(), rv.IsZero(), e.Len())
		//}
		//fmt.Printf("")
	}
	return ok
}


func IfNilReturnError(ref interface{}, name ...string) *State {
	if IsInterfaceNil(ref) {
		s := NewState("", true)
		s._Fatal = errors.New("SW ERROR")
		s.ExitCode = 255
		PanicDump()
		return s
	}

	state := SearchStructureForUxState(ref, name...)
	if state == nil {
		// If the structure doesn't contain *ux.State
		state = NewState(strings.Join(name, " "), false)
	}
	state.SetOk()

	return state
}


// Search a given structure for the State object and return it's pointer.
func SearchStructureForUxState(ref interface{}, name ...string) *State {
	var state *State

	for range onlyOnce {
		v := reflect.ValueOf(ref)
		var e reflect.Value

		// We're doing these checks to ensure ease of future expansion.
		if v.Kind() == reflect.Ptr {
			e = v.Elem()
			if e.Kind().String() == "ptr" {
				//PrintflnCyan("POINTER TO POINTER")	@TODO - DEBUG
				state = SearchStructureForUxState(e.Addr().Elem().Interface(), name...)
				break
			}
		} else if v.Kind() == reflect.Struct {
			// We can't handle a non-pointer, otherwise we get this...
			// reflect.flag.mustBeAssignable using unaddressable value
			//e = v
			Panic(PanicErrorNotGivenAPointer, v.String())
		} else {
			break
		}

		typeOfT := e.Type()
		for i := 0; i < e.NumField(); i++ {
			if typeOfT.Field(i).Name == "State" {
				state = e.Field(i).Interface().(*State)
				if state == nil {
					// Make sure *ux.State isn't nil.
					state = NewState(strings.Join(name, " "), false)
					e.Field(i).Set(reflect.ValueOf(state))
				}
				break
			}
		}
	}

	return state
}


func (state *State) Clear() {
	if state == nil {
		StatePanic(state)
		return
	}
	state._Debug = nil
	state._Fatal = nil
	state._Error = nil
	state._Warning = nil
	state._Ok = errors.New("")
	state.ExitCode = 0

	state.Output = ""
	state._Separator = DefaultSeparator
	state.OutputArray = []string{}
	state.response = *NewResponse()
}


func (state *State) GetPrefix() string {
	return state.prefix
}
func (state *State) GetPackage() string {
	return state._Package
}
func (state *State) GetFunction() string {
	return state._Function
}
func (state *State) SetPackage(s string) {
	if s == "" {
		// Discover package name.
		//pc, file, no, ok := runtime.Caller(1)
		pc, _, _, ok := runtime.Caller(1)
		if ok {
			//s = file + ":" + string(no)
			details := runtime.FuncForPC(pc)
			s = filepath.Base(details.Name())
			sa := strings.Split(s, ".")
			if len(sa) > 0 {
				s = sa[0]
			}
		}
	}

	state._Package = s
	if state._Function == "" {
		state.prefix = state._Package
	} else {
		state.prefix = state._Package + "." + state._Function + "()"
		state.prefixArray = append(state.prefixArray, state.prefix)
	}
}
func (state *State) SetFunction(s ...string) {
	st := strings.Join(s, " ")
	if st == "" {
		// Discover function name.
		//pc, file, no, ok := runtime.Caller(1)
		pc, _, _, ok := runtime.Caller(1)
		if ok {
			//s = file + ":" + string(no)
			details := runtime.FuncForPC(pc)
			foo := details.Name()
			st = filepath.Base(foo)
			sa := strings.Split(st, ".")
			switch {
				case len(sa) > 2:
					st = sa[2]
				case len(sa) > 1:
					st = sa[1]
				case len(sa) > 0:
					st = sa[0]
			}
		}
	}

	state._Function = st
	if state._Package == "" {
		state.prefix = state._Function + "()"
	} else {
		state.prefix = state._Package + "." + state._Function + "()"
	}

	state.prefixArray = append(state.prefixArray, state.prefix)
}
func (state *State) SetFunctionCaller() {
	var s string
	// Discover function name.
	pc, _, _, ok := runtime.Caller(2)
	if ok {
		//s = file + ":" + string(no)
		details := runtime.FuncForPC(pc)
		s = filepath.Base(details.Name())
		sa := strings.Split(s, ".")
		if len(sa) > 0 {
			s = sa[1]
		}
	}

	state.SetFunction(s)
}


func (state *State) GetState() *bool {
	var b bool
	return &b
}
func (state *State) SetState(p *State) {
	if state == nil {
		swerr := NewState("", true)
		swerr._Fatal = errors.New("SW ERROR")
		swerr.ExitCode = 255
		return
	}
	state._Error =      p._Error
	state._Warning =    p._Warning
	state._Ok =         p._Ok
	state._Debug =      p._Debug
	state.ExitCode =    p.ExitCode
	state.Output =      p.Output
	state.OutputArray = p.OutputArray
	state.SetResponse(p.response.data)
	state.RunState =    p.RunState
}


func (state *State) Sprint() string {
	var ret string

	e := ""
	if state.ExitCode != 0 {
		e = fmt.Sprintf("Exit(%d) - ", state.ExitCode)
	}

	pa := ""
	if len(state.prefixArray) > 0 {
		pa = fmt.Sprintf("[%s] - ", state.prefixArray[0])
	}

	ou := ""
	if state.Output != "" {
		ou = "\n" + SprintfOk("%s ", state.Output)
	}

	switch {
	case state._Error != nil:
		ret = SprintfError("ERROR: %s%s%s%s", pa, e, state._Error, ou)
		//ret = SprintfError("ERROR: %s%s%s", e, state._Error, ou)

	case state._Warning != nil:
		ret = SprintfWarning("WARNING: %s%s%s%s", pa, e, state._Warning, ou)
		//ret = SprintfWarning("WARNING: %s%s%s", e, state._Warning, ou)

	case state._Ok != nil:
		if state._Ok.Error() != "" {
			ret = SprintfOk("%s%s", state._Ok, ou)
		}

	case state.debug.Enabled:
		if state._Debug != nil {
			ret = SprintfDebug("%s%s", state._Debug, ou)
		}
	}

	return ret
}
func (state *State) SprintSimple() string {
	var ret string

	e := ""
	if state.ExitCode != 0 {
		e = fmt.Sprintf("Exit(%d) - ", state.ExitCode)
	}

	//pa := ""
	//if len(state.prefixArray) > 0 {
	//	pa = fmt.Sprintf("[%s] - ", state.prefixArray[0])
	//}

	ou := ""
	if state.Output != "" {
		ou = "\n" + SprintfOk("%s ", state.Output)
	}

	switch {
		case state._Error != nil:
			//ret = SprintfError("ERROR: %s%s%s%s", pa, e, state._Error, ou)
			ret = SprintfError("ERROR: %s%s%s", e, state._Error, ou)

		case state._Warning != nil:
			//ret = SprintfWarning("WARNING: %s%s%s%s", pa, e, state._Warning, ou)
			ret = SprintfWarning("WARNING: %s%s%s", e, state._Warning, ou)

		case state._Ok != nil:
			if state._Ok.Error() != "" {
				ret = SprintfOk("%s%s", state._Ok, ou)
			}

		case state.debug.Enabled:
			if state._Debug != nil {
				ret = SprintfDebug("%s%s", state._Debug, ou)
			}
	}

	return ret
}
func (state *State) DebugPrint() {
	if !state.debug.Enabled {
		return
	}
	_, _ = fmt.Fprintf(os.Stderr, state.Sprint() + "\n")
}
func (state *State) SprintError() string {
	var ret string

	for range onlyOnce {
		if state == nil {
			break
		}

		if state._Ok != nil {
			// If we have an OK response.
			break
		}

		ret = state.Sprint()
	}

	return ret
}
func (state *State) SprintResponse() string {
	return state.SprintSimple()
}

var _lastResponse = ""	// Avoid doubling up on same response.
func (state *State) PrintResponse() {
	thisResponse := state.SprintSimple()
	if thisResponse == _lastResponse {
		return
	}
	_lastResponse = thisResponse
	_, _ = fmt.Fprintf(os.Stdout, thisResponse + "\n")
}


func (state *State) IsError() bool {
	var ok bool

	if state == nil {
		//fmt.Printf("DUH\n")
		ok = true
	} else if state._Error != nil {
		ok = true
	}

	return ok
}

func (state *State) IsWarning() bool {
	var ok bool

	if state._Warning != nil {
		ok = true
	}

	return ok
}

func (state *State) IsOk() bool {
	var ok bool

	if state._Ok != nil {
		ok = true
	}

	return ok
}
func (state *State) IsNotOk() bool {
	ok := true

	for range onlyOnce {
		if state._Warning != nil {
			break
		}
		if state._Error != nil {
			break
		}
		ok = false
	}

	return ok
}

func (state *State) OkHas(s string) bool {
	var ok bool

	if strings.Contains(state._Ok.Error(), s) {
		ok = true
	}

	return ok
}

func (state *State) WarningHas(s string) bool {
	var ok bool

	if strings.Contains(state._Warning.Error(), s) {
		ok = true
	}

	return ok
}

func (state *State) ErrorHas(s string) bool {
	var ok bool

	if strings.Contains(state._Error.Error(), s) {
		ok = true
	}

	return ok
}

func (state *State) DebugHas(s string) bool {
	var ok bool

	if strings.Contains(state._Debug.Error(), s) {
		ok = true
	}

	return ok
}


func (state *State) SetExitCode(e int) {
	if state == nil {
		return
	}
	state.ExitCode = e
}
func (state *State) GetExitCode() int {
	return state.ExitCode
}


func (state *State) SetError(error ...interface{}) {
	for range onlyOnce {
		if state == nil {
			StatePanic(state)
			break
		}
		state.debug.fetchRuntimeDebug(2)

		state._Ok = nil
		state._Warning = nil

		if len(error) == 0 {
			state._Error = errors.New("ERROR")
			break
		}

		if error[0] == nil {
			state._Error = nil
			break
		}

		debugPrefix := ""
		if state.debug.Enabled {
			debugPrefix = SprintfCyan("%s:%d [%s] - ", state.debug.File, state.debug.Line, state.debug.Function)
		}
		state._Error = errors.New(debugPrefix + _Sprintf(error...))
		if state.debug.Enabled {
			state.PrintResponse()
		}
	}
}
func (state *State) GetError() error {
	return state._Error
}


func (state *State) SetWarning(warning ...interface{}) {
	for range onlyOnce {
		if state == nil {
			StatePanic(state)
			break
		}
		state.debug.fetchRuntimeDebug(2)

		state._Ok = nil
		state._Error = nil

		if len(warning) == 0 {
			state._Warning = errors.New("WARNING")
			break
		}

		if warning[0] == nil {
			state._Warning = nil
			break
		}

		debugPrefix := ""
		if state.debug.Enabled {
			debugPrefix = SprintfCyan("%s:%d [%s] - ", state.debug.File, state.debug.Line, state.debug.Function)
		}
		state._Warning = errors.New(debugPrefix + _Sprintf(warning...))
		if state.debug.Enabled {
			state.PrintResponse()
		}
	}
}
func (state *State) GetWarning() error {
	return state._Warning
}


func (state *State) SetOk(msg ...interface{}) {
	for range onlyOnce {
		if state == nil {
			StatePanic(state)
			break
		}
		state.debug.fetchRuntimeDebug(2)

		state._Error = nil
		state._Warning = nil
		state.ExitCode = 0

		if len(msg) == 0 {
			state._Ok = errors.New("")
			break
		}

		if msg[0] == nil {
			state._Ok = errors.New("")
			break
		}

		debugPrefix := ""
		if state.debug.Enabled {
			debugPrefix = SprintfCyan("%s:%d [%s] - ", state.debug.File, state.debug.Line, state.debug.Function)
		}
		state._Ok = errors.New(debugPrefix + _Sprintf(msg...))
		if state.debug.Enabled {
			state.PrintResponse()
		}
	}
}
func (state *State) GetOk() error {
	return state._Ok
}


func (state *State) ClearError() {
	state._Error = nil
}


func (state *State) IsRunning() bool {
	var ok bool
	if state.RunState == StateRunning {
		ok = true
	}
	return ok
}

func (state *State) IsPaused() bool {
	var ok bool
	if state.RunState == StatePaused {
		ok = true
	}
	return ok
}

func (state *State) IsCreated() bool {
	var ok bool
	if state.RunState == StateCreated {
		ok = true
	}
	return ok
}

func (state *State) IsRestarting() bool {
	var ok bool
	if state.RunState == StateRestarting {
		ok = true
	}
	return ok
}

func (state *State) IsRemoving() bool {
	var ok bool
	if state.RunState == StateRemoving {
		ok = true
	}
	return ok
}

func (state *State) IsExited() bool {
	var ok bool
	if state.RunState == StateExited {
		ok = true
	}
	return ok
}

func (state *State) IsDead() bool {
	var ok bool
	if state.RunState == StateDead {
		ok = true
	}
	return ok
}


func (state *State) ExitOnNotOk() string {
	if state.IsNotOk() {
		_, _ = fmt.Fprintf(os.Stderr, state.Sprint() + "\n")
		os.Exit(state.ExitCode)
	}
	return ""
}


func (state *State) ExitOnError() string {
	if state.IsWarning() {
		_, _ = fmt.Fprintf(os.Stderr, state.Sprint() + "\n")
	}
	if state.IsError() {
		_, _ = fmt.Fprintf(os.Stderr, state.Sprint() + "\n")
		os.Exit(state.ExitCode)
	}
	return ""
}


func (state *State) ExitOnWarning() string {
	if state.IsWarning() {
		_, _ = fmt.Fprintf(os.Stderr, state.Sprint() + "\n")
		os.Exit(state.ExitCode)
	}
	return ""
}


func (state *State) Exit(e int) string {
	state.ExitCode = e
	_, _ = fmt.Fprintf(os.Stdout, state.Sprint())
	os.Exit(state.ExitCode)
	return ""
}


func Exit(e int64, msg ...interface{}) string {
	ret := _Sprintf(msg...)
	if e == 0 {
		_, _ = fmt.Fprintf(os.Stdout, SprintfOk(ret))
	} else {
		_, _ = fmt.Fprintf(os.Stderr, SprintfError(ret))
	}
	os.Exit(int(e))
	return ""	// Will never get here.
}


func _Sprintf(msg ...interface{}) string {
	var ret string

	for range onlyOnce {
		if len(msg) == 0 {
			break
		}

		value := reflect.ValueOf(msg[0])
		switch value.Kind() {
			case reflect.String:
				if len(msg) == 1 {
					ret = fmt.Sprintf(msg[0].(string))
				} else {
					ret = fmt.Sprintf(msg[0].(string), msg[1:]...)
				}

			default:
				if len(msg) == 1 {
					ret = fmt.Sprintf("%v", msg)
				} else {
					var es string
					for _, e := range msg {
						es += fmt.Sprintf("%v ", e)
					}
					es = strings.TrimSuffix(es, " ")
					ret = es
				}
		}

		//ret = fmt.Sprintf(msg[0].(string), msg[1:]...)
	}

	return ret
}

func (state *State) DebugEnable() {
	for range onlyOnce {
		if state == nil {
			break
		}
		state.debug.Enabled = true
	}
}
func (state *State) DebugDisable() {
	for range onlyOnce {
		if state == nil {
			break
		}
		state.debug.Enabled = false
	}
}
func (state *State) DebugSet(d bool) {
	for range onlyOnce {
		if state == nil {
			break
		}
		state.debug.Enabled = d
	}
}

func (state *State) SetDebug(msg ...interface{}) {
	for range onlyOnce {
		if state == nil {
			break
		}
		state.debug.fetchRuntimeDebug(2)

		if len(msg) == 0 {
			state._Debug = errors.New("DEBUG")
			break
		}

		if msg[0] == nil {
			state._Debug = errors.New("DEBUG")
			break
		}

		debugPrefix := ""
		if state.debug.Enabled {
			debugPrefix = SprintfCyan("%s:%d [%s] - ", state.debug.File, state.debug.Line, state.debug.Function)
		}
		state._Debug = errors.New(debugPrefix + _Sprintf(msg...))
		if state.debug.Enabled {
			state.PrintResponse()
		}
	}
}
func (state *State) GetDebug() error {
	return state._Debug
}

func (state *State) SetVerboseMode(m bool) {
	for range onlyOnce {
		if state == nil {
			break
		}
		state.verbose = m
	}
}
func (state *State) EnableVerboseMode() {
	for range onlyOnce {
		if state == nil {
			break
		}
		state.verbose = true
	}
}
func (state *State) DisableVerboseMode() {
	for range onlyOnce {
		if state == nil {
			break
		}
		state.verbose = false
	}
}
func (state *State) IsVerboseMode() bool {
	if state == nil {
		return false
	}
	return state.verbose
}

func (state *State) IsResponseOfType(t string) bool {
	return state.response.IsOfType(t)
}


func (state *State) IsResponseNotOfType(t string) bool {
	return state.response.IsNotOfType(t)
}
