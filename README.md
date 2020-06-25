# scribeTools
This repository contains packages that are used within the scribe framework.


## Package format
Methods and functions conform to a standard layout.

That is: functions/method returns and method receiver structures follow a similar structure.

The intent is to make everything consistent and therefore easy to understand.


### ux.State package
The entire codebase has this structure embedded throughout. Complex structures will always contain this structure and may or may not reference child structures with this included as well.

The ux package contains all the critical and important UX interfacing methods/functions and can return errors, warnings, ok and debug states.
As well as handling variable return types via a simple interface method, keeping track of state changes and providing console colouring based on states.

It's intent is to be simple and functional.


### Method/function names

Names starting with:
- `Get` - Will return a structure/variable that you expect to see. Example: `loadGit.GetRepo` will return a Git repository structure.
- `Is` - Will return a boolean... Always.
- `New` - Always instantiates and initializes a structure. This may, or may not call further `New` methods.
- `Tool` - Is always a function name and never a method. Any functions with this prefix will have the `Tool` prefix removed, and the result will be available to the Scribe template framework.

### TODO

Add file mime type detection to toolPath
- https://github.com/h2non/filetype

Add in VFSGEN support
- https://github.com/shurcooL/vfsgen

Add better type handling within ux.State
- https://go101.org/article/type-system-overview.html

