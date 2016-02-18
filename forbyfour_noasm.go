//+build !amd64 noasm appengine

func init() {
	MatMul = innrPrd_default
}

func ForceSelect(s sel) {
	MatMul = innrPrd_default
}
