package toolGo


func Help(path ...string) *TypeGo {
	dir := New(nil)

	for range onlyOnce {
		if len(path) == 0 {
			break
		}

		if dir.State.IsNotOk() {
			break
		}

		dir.State = dir.Find(path...)
		if dir.State.IsNotOk() {
			break
		}

		dir.State = dir.Parse()
		if dir.State.IsNotOk() {
			break
		}

		//for _, file := range dir.Files {
		//	ux.PrintfBlue("\nChecking file '%s' ...\n", file.Path)
		//	fmt.Printf(file.String())
		//}
		//fmt.Printf("\n")
	}

	return dir
}


func ParseGoFiles(path ...string) *TypeGo {
	dir := New(nil)

	for range onlyOnce {
		if len(path) == 0 {
			break
		}

		if dir.State.IsNotOk() {
			break
		}

		dir.State = dir.Find(path...)
		if dir.State.IsNotOk() {
			break
		}

		dir.State = dir.Parse()
		if dir.State.IsNotOk() {
			break
		}

		//for _, file := range dir.Files {
		//	ux.PrintfBlue("\nChecking file '%s' ...\n", file.Path)
		//	fmt.Printf(file.String())
		//}
		//fmt.Printf("\n")
	}

	return dir
}
