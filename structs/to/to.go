package to

func Slice[A, B any](elems []A, f func(A) B) []B {
	result := make([]B, 0, len(elems))
	if elems == nil {
		return result
	}
	for _, elem := range elems {
		result = append(result, f(elem))
	}
	return result
}
