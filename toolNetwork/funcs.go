package toolNetwork


func (r *TypeNetwork) IsAvailable(port uint16) bool {
	var ok bool
	if state := r.IsNil(); state.IsError() {
		return ok
	}

	ok = true
	for _, p := range r.Listeners {
		if p.Port == port {
			ok = false
			break
		}
	}

	return ok
}
func (r *TypeNetwork) IsNotAvailable(port uint16) bool {
	return !r.IsAvailable(port)
}
