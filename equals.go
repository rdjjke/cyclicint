package cyclicint

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func Equals[T Integer](a, b, eps T) bool {
	left, right := a-eps, a+eps
	if left < right { // no integer overflow
		return left <= b && b <= right
	} else { // integer overflow
		return b <= left || right <= b
	}
}

