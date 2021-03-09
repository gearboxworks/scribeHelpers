/*
This package provides an "all-in-one" generic response code mechanism.
It manages the following:
- error state.
- warning state.
- ok state.
- debug state - will keep track of historic function calls.
- exit codes - for program exiting.
- generic "run states" - runnning, paused, created, restarting, removing, exited, dead.
- generic response interface.
	- allows any type to be stored and returned to functions.

Examples:

*/
package ux


/*

@TODO - Provide logrus interfacing.

*/