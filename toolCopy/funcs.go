package toolCopy


func (c *TypeOsCopy) SetSourcePath(path ...string) bool {
	var ok bool

	for range onlyOnce {
		ok = c.Paths.Source.SetPath(path...)
		if !ok {
			break
		}

		if c.Paths.Destination.IsValid() {
			c.Valid = true
		}

		ok = true
	}

	return ok
}
func (c *TypeOsCopy) GetSourcePath() string {
	return c.Paths.Source.GetPath()
}


func (c *TypeOsCopy) SetDestinationPath(path ...string) bool {
	var ok bool

	for range onlyOnce {
		ok = c.Paths.Destination.SetPath(path...)
		if !ok {
			break
		}

		if c.Paths.Source.IsValid() {
			c.Valid = true
		}

		ok = true
	}

	return ok
}
func (c *TypeOsCopy) GetDestinationPath() string {
	return c.Paths.Destination.GetPath()
}


func (c *TypeOsCopy) SetExcludePaths(paths ...string) bool {
	return c.Paths.Exclude.SetPaths(paths...)
}
func (c *TypeOsCopy) AddExcludePaths(paths ...string) bool {
	return c.Paths.Exclude.AddPaths(paths...)
}
func (c *TypeOsCopy) GetExcludePaths() *PathArray {
	return c.Paths.Exclude.GetPaths()
}


func (c *TypeOsCopy) SetIncludePaths(paths ...string) bool {
	return c.Paths.Include.SetPaths(paths...)
}
func (c *TypeOsCopy) AddIncludePaths(paths ...string) bool {
	return c.Paths.Include.AddPaths(paths...)
}
func (c *TypeOsCopy) GetIncludePaths() *PathArray {
	return c.Paths.Include.GetPaths()
}


func (c *TypeOsCopy) SetMethodRsync() bool {
	return c.Method.SelectMethod(ConstMethodRsync)
}
func (c *TypeOsCopy) SetMethodTar() bool {
	return c.Method.SelectMethod(ConstMethodTar)
}
func (c *TypeOsCopy) SetMethodCpio() bool {
	return c.Method.SelectMethod(ConstMethodCpio)
}
func (c *TypeOsCopy) SetMethodSftp() bool {
	return c.Method.SelectMethod(ConstMethodSftp)
}
func (c *TypeOsCopy) SetMethodCp() bool {
	return c.Method.SelectMethod(ConstMethodCp)
}


func (c *TypeOsCopy) GetMethodOptions() interface{} {
	return c.Method.GetOptions()
}
func (c *TypeOsCopy) GetMethodName() string {
	return c.Method.GetName()
}
func (c *TypeOsCopy) GetMethodPath() string {
	return c.Method.GetPath()
}
func (c *TypeOsCopy) GetMethodAllowRemote() bool {
	return c.Method.GetAllowRemote()
}
func (c *TypeOsCopy) GetMethodAvailable() bool {
	return c.Method.GetAvailable()
}


func (c *TypeOsCopy) SetOverwrite() {
	c.Paths.Destination.SetOverwriteable()
}
func (c *TypeOsCopy) CanOverwrite() bool {
	return c.Paths.Destination.CanOverwrite()
}
