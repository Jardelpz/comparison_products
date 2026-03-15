package utils

func EmptyToNil(v any) any {
	switch val := v.(type) {
	case string:
		if val == "" {
			return nil
		}
	case int:
		if val == 0 {
			return nil
		}
	case float64:
		if val == 0 {
			return nil
		}
	default:
		return v
	}
	return v
}
