package ptr

// helper package to return pointer representation of a data type

func String(s string) *string {
	return &s
}

func Int(i int) *int {
	return &i
}

func Uint64(i uint64) *uint64 {
	return &i
}

func Float64(f float64) *float64 {
	return &f
}
