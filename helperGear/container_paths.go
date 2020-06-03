package helperGear

import "github.com/newclarity/scribeHelpers/ux"

type VolumeMounts map[string]string

func (m *VolumeMounts) Add(local string, remote string) bool {
	var ok bool

	for range OnlyOnce {
		if state := ux.IfNilReturnError(m); state.IsError() {
			break
		}

		if local == DefaultPathNone {
			ok = true
			break
		}

		if remote == DefaultPathNone {
			ok = true
			break
		}

		(*m)[local] = remote
	}

	return ok
}


type SshfsMounts map[string]string

func (m *SshfsMounts) Add(local string, remote string) bool {
	var ok bool

	for range OnlyOnce {
		if state := ux.IfNilReturnError(m); state.IsError() {
			break
		}

		if local == DefaultPathNone {
			ok = true
			break
		}

		if remote == DefaultPathNone {
			ok = true
			break
		}

		(*m)[local] = remote
	}

	return ok
}
