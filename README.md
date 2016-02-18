# forbyfour
Simple go library implementing matrix multiplication on 4x4 matrices using vectorized assembly

# Use
Exported function `MatMul(a,b,c []float64)` multiplies two 4x4 matricies c = a x b.  All 3 slices must be initialized with 16 elements.  No input validation is done within the exported function.  This is in the name of performance and benchmarking on the target machine is suggested. 

On startup, the library will load the implementation with the highest vector instruction set available.  Specific implementations can be selected using `ForceSelect()`.  This will not load an unsupported instruction set.