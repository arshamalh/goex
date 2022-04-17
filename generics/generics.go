package generics

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

// Generics exercise main
// Index works on a slice of ints
// si := []int{10, 20, 15, -10}
// fmt.Println(generics.Index(si, 15))
// Index also works on a slice of strings
// ss := []string{"foo", "bar", "baz"}
// fmt.Println(generics.Index(ss, "hello"))
