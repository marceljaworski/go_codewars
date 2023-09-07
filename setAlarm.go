func SetAlarm(employed, vacation bool) bool {
	switch {
	case employed == true && vacation == true:
		return false
	case employed == true && vacation == false:
		return true
	case employed == false && vacation == true:
		return false
	case employed == false && vacation == false:
		return false
	}
	return false
	// return employed && !vacation
}