package forbyfour

var (
	sse3Supt, fmaSupt bool
	avxSupt, avx2Supt bool
	// MatMul uses vector instructions to perform efficient
	// 4x4 matrix multiplication in the form of c = a x b
	MatMul func(a, b, c []float64)
)

type sel int8

const (
	AVX1 sel = iota
	AVX2
	SSE3_FMA
	SSE2_FMA
	SSE3
	SSE2
	DEFAULT
)

func innrPrd_default(a, b, c []float64) {
	var x, y int
	for x = 0; x < 4; x++ { //rows
		for y = 0; y < 4; y++ { //cols
			c[x*4+y] = a[x*4] * b[y]+ a[x*4+1] * b[4+y]+a[x*4+2] * b[8+y]+a[x*4+3] * b[12+y]
		}
	}
}
