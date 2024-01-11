package filters

func CascateValidator(filter []TextFilter, text string) error {
	for _, f := range filter {
		result := f(text)

		if result != nil {
			return result
		}
	}

	return nil
}
