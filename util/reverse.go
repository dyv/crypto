package util

// Reverse copies src into dst in byte-reversed order and returns dst,
// such that src[0] goes into dst[len-1] and vice versa.
// dst and src may be the same slice but otherwise must not overlap.
//
// XXX would be nice to have this function in the 'bytes' standard package
func Reverse(dst,src []byte) []byte {
	l := len(dst)
	if len(src) != l {
		panic("Reverse requires equal-length slices")
	}
	for i, j := 0, l-1; i < (l+1)/2; {
		dst[i], dst[j] = src[j], src[i]
		i++
		j--
	}
	return dst
}

