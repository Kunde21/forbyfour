package forbyfour

import "testing"

var a, b, res []float64

func init() {
	a, b = make([]float64, 16), make([]float64, 16)

	for i, j := 0, 0; i < 16; i, j = i+1, j+1 {
		a[i], b[i] = float64(j), float64(j)
	}
	res = []float64{56, 62, 68, 74, 152, 174, 196, 218, 248, 286, 324, 362, 344, 398, 452, 506}
}

func TestBroadcast(t *testing.T) {
	c := make([]float64, 16)
	testBroadcast(a, b, c)
	for i, v := range res {
		if c[i] != v {
			t.Logf("Value fault at %v.  Expected %v got %v", i, v, c[i])
			t.Fail()
		}
	}
}

func TestInnrPrd_avx1(t *testing.T) {
	c := make([]float64, 16)
	innrPrd_avx(a, b, c)
	for i, v := range res {
		if c[i] != v {
			t.Logf("Value fault at %v.  Expected %v got %v", i, v, c[i])
			t.Fail()
		}
	}
}

func TestInnrPrd_avx2(t *testing.T) {
	c := make([]float64, 16)

	innrPrd_avx2(a, b, c)
	for i, v := range res {
		if c[i] != v {
			t.Logf("Value fault at %v.  Expected %v got %v", i, v, c[i])
			t.Fail()
		}
	}
}

func TestInnrPrd_sse3_fma(t *testing.T) {
	c := make([]float64, 16)
	innrPrd_sse3_fma(a, b, c)
	for i, v := range res {
		if c[i] != v {
			t.Logf("Value fault at %v.  Expected %v got %v", i, v, c[i])
			t.Fail()
		}
	}
}

func TestInnrPrd_sse2_fma(t *testing.T) {
	c := make([]float64, 16)
	innrPrd_sse2_fma(a, b, c)
	for i, v := range res {
		if c[i] != v {
			t.Logf("Value fault at %v.  Expected %v got %v", i, v, c[i])
			t.Fail()
		}
	}
}

func TestInnrPrd_sse3(t *testing.T) {
	c := make([]float64, 16)
	innrPrd_sse3(a, b, c)
	for i, v := range res {
		if c[i] != v {
			t.Logf("Value fault at %v.  Expected %v got %v", i, v, c[i])
			t.Fail()
		}
	}
}

func TestInnrPrd_sse2(t *testing.T) {
	c := make([]float64, 16)
	innrPrd_sse2(a, b, c)
	for i, v := range res {
		if c[i] != v {
			t.Logf("Value fault at %v.  Expected %v got %v", i, v, c[i])
			t.Fail()
		}
	}
}

func TestNaive(t *testing.T) {
	c := make([]float64, 16)

	innrPrd_default(a, b, c)
	for i, v := range res {
		if c[i] != v {
			t.Logf("Value fault at %v.  Expected %v got %v", i, v, c[i])
			t.Fail()
		}
	}
}

func BenchmarkInnrPrd_sse3_fma(t *testing.B) {
	if !(sse3Supt || fmaSupt) {
		t.Skip()
	}
	c := make([]float64, 16)

	t.ResetTimer()
	t.ReportAllocs()
	for i := 0; i < t.N; i++ {
		innrPrd_sse3_fma(a, b, c)
	}
}

func BenchmarkInnrPrd_sse2_fma(t *testing.B) {
	if !fmaSupt {
		t.Skip()
	}
	c := make([]float64, 16)

	t.ResetTimer()
	t.ReportAllocs()
	for i := 0; i < t.N; i++ {
		innrPrd_sse2_fma(a, b, c)
	}
}

func BenchmarkInnrPrd_avx1(t *testing.B) {
	if !avxSupt {
		t.Skip()
	}
	c := make([]float64, 16)

	t.ResetTimer()
	t.ReportAllocs()
	for i := 0; i < t.N; i++ {
		innrPrd_avx(a, b, c)
	}
}

func BenchmarkInnrPrd_avx2(t *testing.B) {
	if !avxSupt {
		t.Skip()
	}
	c := make([]float64, 16)

	t.ResetTimer()
	t.ReportAllocs()
	for i := 0; i < t.N; i++ {
		innrPrd_avx2(a, b, c)
	}
}

func BenchmarkInnrPrd_broadcast(t *testing.B) {
	if !avx2Supt {
		t.Skip()
	}
	c := make([]float64, 16)

	t.ResetTimer()
	t.ReportAllocs()
	for i := 0; i < t.N; i++ {
		testBroadcast(a, b, c)
	}
}

func BenchmarkInnrPrd_broadcast2(t *testing.B) {
	if !avx2Supt {
		t.Skip()
	}
	c := make([]float64, 16)

	t.ResetTimer()
	t.ReportAllocs()
	for i := 0; i < t.N; i++ {
		testBroadcast2(a, b, c)
	}
}

func BenchmarkNaive(t *testing.B) {
	c := make([]float64, 16)

	t.ResetTimer()
	t.ReportAllocs()
	for i := 0; i < t.N; i++ {
		innrPrd_default(a, b, c)
	}
}

func BenchmarkInnrPrd_sse3(t *testing.B) {
	if !sse3Supt {
		t.Skip()
	}
	c := make([]float64, 16)

	t.ResetTimer()
	t.ReportAllocs()
	for i := 0; i < t.N; i++ {
		innrPrd_sse3(a, b, c)
	}
}

func BenchmarkInnrPrd_sse2(t *testing.B) {
	c := make([]float64, 16)

	t.ResetTimer()
	t.ReportAllocs()
	for i := 0; i < t.N; i++ {
		innrPrd_sse2(a, b, c)
	}
}
