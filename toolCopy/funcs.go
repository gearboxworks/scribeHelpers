package toolCopy


func (c *TypeOsCopy) SetSourcePath(path ...string) bool {
	var ok bool

	for range onlyOnce {
		ok = c.Source.SetPath(path...)
		if !ok {
			break
		}

		if c.Destination.IsValid() {
			c.Valid = true
		}

		ok = true
	}

	return ok
}
func (c *TypeOsCopy) GetSourcePath() string {
	return c.Source.GetPath()
}


func (c *TypeOsCopy) SetDestinationPath(path ...string) bool {
	var ok bool

	for range onlyOnce {
		ok = c.Destination.SetPath(path...)
		if !ok {
			break
		}

		if c.Source.IsValid() {
			c.Valid = true
		}

		ok = true
	}

	return ok
}
func (c *TypeOsCopy) GetDestinationPath() string {
	return c.Destination.GetPath()
}


func (c *TypeOsCopy) SetExcludePaths(paths ...string) bool {
	return c.Exclude.SetPaths(paths...)
}
func (c *TypeOsCopy) AddExcludePaths(paths ...string) bool {
	return c.Exclude.AddPaths(paths...)
}
func (c *TypeOsCopy) GetExcludePaths() *PathArray {
	return c.Exclude.GetPaths()
}


func (c *TypeOsCopy) SetIncludePaths(paths ...string) bool {
	return c.Include.SetPaths(paths...)
}
func (c *TypeOsCopy) AddIncludePaths(paths ...string) bool {
	return c.Include.AddPaths(paths...)
}
func (c *TypeOsCopy) GetIncludePaths() *PathArray {
	return c.Include.GetPaths()
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
