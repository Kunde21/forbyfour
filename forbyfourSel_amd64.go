// +build !noasm !appengine

package forbyfour

func ForceSelect(s int8) {
	switch {
	case s == AVX1 && avxSupt:
		MatMul = innrPrd_avx
	case s == AVX2 && avxSupt:
		MatMul = innrPrd_avx2
	case s == SSE3_FMA && sse3Supt && fmaSupt:
		MatMul = innrPrd_sse3_fma
	case s == SSE2_FMA && fmaSupt:
		MatMul = innrPrd_sse2_fma
	case s == SSE3 && sse3Supt:
		MatMul = innrPrd_sse3
	case s == SSE2:
		MatMul = innrPrd_sse2
	case s == DEFAULT:
		MatMul = innrPrd_default
	default:
		switch {
		case avx2Supt || avxSupt:
			MatMul = innrPrd_avx2
		case sse3Supt && fmaSupt:
			MatMul = innrPrd_sse3_fma
		case fmaSupt:
			MatMul = innrPrd_sse2_fma
		case sse3Supt:
			MatMul = innrPrd_sse3
		default:
			MatMul = innrPrd_sse2
		}
	}
}
