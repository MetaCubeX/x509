package x509

import (
	"math/bits"
	"sort"
)

func slicesBinarySearchFunc[S ~[]E, E, T any](x S, target T, cmp func(E, T) int) (int, bool) {
	n := len(x)
	// Define cmp(x[-1], target) < 0 and cmp(x[n], target) >= 0 .
	// Invariant: cmp(x[i - 1], target) < 0, cmp(x[j], target) >= 0.
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if cmp(x[h], target) < 0 {
			i = h + 1 // preserves cmp(x[i - 1], target) < 0
		} else {
			j = h // preserves cmp(x[j], target) >= 0
		}
	}
	// i == j, cmp(x[i-1], target) < 0, and cmp(x[j], target) (= cmp(x[i], target)) >= 0  =>  answer is i.
	return i, i < n && cmp(x[i], target) == 0
}

func slicesSortFunc[S ~[]E, E any](x S, cmp func(a, b E) int) {
	sort.Slice(x, func(i, j int) bool { return cmp(x[i], x[j]) < 0 })
}

func slicesReverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func slicesRepeat[S ~[]E, E any](x S, count int) S {
	if count < 0 {
		panic("cannot be negative")
	}

	const maxInt = ^uint(0) >> 1
	hi, lo := bits.Mul(uint(len(x)), uint(count))
	if hi > 0 || lo > maxInt {
		panic("the result of (len(x) * count) overflows")
	}

	newslice := make(S, int(lo)) // lo = len(x) * count
	n := copy(newslice, x)
	for n < len(newslice) {
		n += copy(newslice[n:], newslice[:n])
	}
	return newslice
}
func slicesEqual[S ~[]E, E comparable](s1, s2 S) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func slicesEqualFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, s2 S2, eq func(E1, E2) bool) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v1 := range s1 {
		v2 := s2[i]
		if !eq(v1, v2) {
			return false
		}
	}
	return true
}

func slicesIndexFunc[S ~[]E, E any](s S, f func(E) bool) int {
	for i := range s {
		if f(s[i]) {
			return i
		}
	}
	return -1
}

func slicesDeleteFunc[S ~[]E, E any](s S, del func(E) bool) S {
	i := slicesIndexFunc(s, del)
	if i == -1 {
		return s
	}
	// Don't start copying elements until we find one to delete.
	for j := i + 1; j < len(s); j++ {
		if v := s[j]; !del(v) {
			s[i] = v
			i++
		}
	}
	clearSlice(s[i:]) // zero/nil out the obsolete elements, for GC
	return s[:i]
}

func clearSlice[S ~[]E, E any](s S) {
	var zero E
	for i := range s {
		s[i] = zero
	}
}

func slicesClone[S ~[]E, E any](s S) S {
	// Preserve nil in case it matters.
	if s == nil {
		return nil
	}
	return append(S([]E{}), s...)
}
