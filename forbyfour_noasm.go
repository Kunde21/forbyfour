//+build !amd64 noasm appengine

func init() {
	MatMul = innrPrd_default
}

func ForceSelect(s int8) {
	MatMul = innrPrd_default
}
