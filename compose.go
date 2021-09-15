package check

func Nest(main error, nested error) error {
	if nested == nil {
		return main
	}
	if main == nil {
		return nested
	}
	return &wrapped{main, nested}
}
