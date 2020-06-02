package loader

/*
This directory contains a common set of generic helpers. 3rd party helper packages can be added to the JTC package
using the


All helper packages can be used in one of two ways:
1. Direct
	- That is; native GoLang.
	- Methods are hanging off a structure with a "Type" prefix - EG TypeGit, TypeSystem.
		- This is to logically separate low level types from high level types.

2. Templates, (aka Helpers)
	- All, (or most), low level functions/methods are available as templates.
	- For template sanity, (avoiding function name clashes), it is better to present all functions as methods.
	- Methods are "typed" directly from the low-level types - EG TypeGit => HelperGit, TypeSystem => HelperSystem.
		- This is to avoid unnecessary type duplication.
	- Helper types have corresponding "Reflect()" functions to ease translation back and forth between low level types.
	- Any functions with a "Helper" prefix can be accessed within templates, all other functions are ignored.
		- This is to allow for functions to be "package visible", yet "template invisible".
		- EG - HelperNewGit - is available as "{{ $git := NewGit }}"
		- EG - HelperNewSystem - is available as "{{ $sys := NewSystem }}"
	- You can still access Template functions as GoLang native functions.
		- They just have more reflection sanity checking.
			- Because "templates" are considered "user scripts" and humans are humans.

*/
