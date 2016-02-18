// +build !noasm !appengine

package forbyfour

func init() {
	initasm()
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

func initasm()

func innrPrd_avx(a, b, c []float64)

func innrPrd_avx2(a, b, c []float64)

func innrPrd_sse3_fma(a, b, c []float64)

func innrPrd_sse2_fma(a, b, c []float64)

func innrPrd_sse3(a, b, c []float64)

func innrPrd_sse2(a, b, c []float64)
