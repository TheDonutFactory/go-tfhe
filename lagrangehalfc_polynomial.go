package tfhe

//"github.com/takatoh/fft"

type LagrangeHalfCPolynomial struct {
	coefsC []complex128
}

func NewLagrangeHalfCPolynomial(n int32) *LagrangeHalfCPolynomial {
	//Assert(n == 1024)
	return &LagrangeHalfCPolynomial{coefsC: make([]complex128, n/2)}
}

func LagrangeHalfCPolynomialClear(p *LagrangeHalfCPolynomial) {
	p.coefsC = make([]complex128, len(p.coefsC))
}

func LagrangeHalfCPolynomialSetTorusConstant(result *LagrangeHalfCPolynomial, mu Torus32) {
	muc := complex(T32tod(mu), 0.)
	for j := 0; j < len(result.coefsC); j++ {
		result.coefsC[j] = muc
	}
}

func LagrangeHalfCPolynomialAddTorusConstant(result *LagrangeHalfCPolynomial, mu Torus32) {
	muc := complex(T32tod(mu), 0.)
	for j := 0; j < len(result.coefsC); j++ {
		if j < 10 {
			//fmt.Printf("Before: (%f, %f)\n", real(result.coefsC[j]), imag(result.coefsC[j]))
			//fmt.Printf("Add Mu: (%f, %f)\n", real(muc), imag(muc))
			result.coefsC[j] += muc
			//fmt.Printf("Result: (%f, %f)\n", real(result.coefsC[j]), imag(result.coefsC[j]))
			//fmt.Println()
		} else {
			result.coefsC[j] += muc
		}
	}
}

/*
EXPORT void LagrangeHalfCPolynomialSetXaiMinusOne(LagrangeHalfCPolynomial* result, const int32_t ai) {
    LagrangeHalfCPolynomial_IMPL* result1 = (LagrangeHalfCPolynomial_IMPL*) result;
    const int32_t Ns2 = result1->proc->Ns2;
    const int32_t _2N = result1->proc->_2N;
    const cplx* omegaxminus1 = result1->proc->omegaxminus1;
    for (int32_t i=0; i<Ns2; i++)
	result1->coefsC[i]=omegaxminus1[((2*i+1)*ai)%_2N];
}
*/

/** termwise multiplication in Lagrange space */
func LagrangeHalfCPolynomialMul(result *LagrangeHalfCPolynomial, a *LagrangeHalfCPolynomial, b *LagrangeHalfCPolynomial) {
	for j := 0; j < len(result.coefsC); j++ {
		result.coefsC[j] = a.coefsC[j] * b.coefsC[j]
	}
}

/** termwise multiplication and addTo in Lagrange space */
func LagrangeHalfCPolynomialAddMul(accum *LagrangeHalfCPolynomial, a *LagrangeHalfCPolynomial, b *LagrangeHalfCPolynomial) {
	for j := 0; j < len(accum.coefsC); j++ {
		accum.coefsC[j] += a.coefsC[j] * b.coefsC[j]
	}
}

/** termwise multiplication and addTo in Lagrange space */
func LagrangeHalfCPolynomialSubMul(accum *LagrangeHalfCPolynomial, a *LagrangeHalfCPolynomial, b *LagrangeHalfCPolynomial) {
	for j := 0; j < len(accum.coefsC); j++ {
		accum.coefsC[j] += a.coefsC[j] * b.coefsC[j]
	}
}

func LagrangeHalfCPolynomialAddTo(accum *LagrangeHalfCPolynomial, a *LagrangeHalfCPolynomial) {
	for j := 0; j < len(accum.coefsC); j++ {
		accum.coefsC[j] += a.coefsC[j]
	}
}