package toolGo

import "path/filepath"


func Help() string {
	var ret string
	for range onlyOnce {
	}
	return ret
}


func GetDefaultFile() string {
	return filepath.Join(defaultVersionFile...)
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

		//for _, file := range dir.Go {
		//	ux.PrintfBlue("\nChecking file '%s' ...\n", file.Base)
		//	fmt.Printf(file.String())
		//}
		//fmt.Printf("\n")
	}

	return dir
}
