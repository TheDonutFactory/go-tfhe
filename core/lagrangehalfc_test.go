package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/thedonutfactory/go-tfhe/fft"
	. "github.com/thedonutfactory/go-tfhe/types"
)

func printTorusPolynomial(t *TorusPolynomial) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d, ", t.Coefs[i])
	}
	fmt.Printf("\n")
}

func printLagrange(p *LagrangeHalfCPolynomial) {
	for i := 0; i < 10; i++ {
		fmt.Printf("(%f, %f i), ", real(p.Coefs[i]), imag(p.Coefs[i]))
	}
	fmt.Printf("\n")
}

//TEST(LagrangeHalfcTest, fftIsBijective) {
func TestFftIsBijective(t *testing.T) {
	assert := assert.New(t)
	NBTRIALS := 1
	toler := 1e-9
	var N int32 = 1024
	for trials := 0; trials < NBTRIALS; trials++ {
		//a := NewTorusPolynomial(N)
		acopy := NewTorusPolynomial(N)
		b := NewTorusPolynomial(N)
		afft := NewLagrangeHalfCPolynomial(N)
		a := NewTorusPolynomial(N)
		torusPolynomialUniform(a)
		fmt.Printf("torusPolynomialUniform(a):")
		printTorusPolynomial(a)

		TorusPolynomialCopy(acopy, a)
		fmt.Printf("torusPolynomialCopy(acopy, a):")
		printTorusPolynomial(acopy)

		TorusPolynomialIfft(afft, a.Coefs)
		fmt.Printf("torusPolynomialIfft(afft, a):")
		printLagrange(afft)

		b.Coefs = TorusPolynomialFft(afft)
		fmt.Printf("torusPolynomialFft(b, afft):")
		printTorusPolynomial(b)

		fmt.Printf("A: \n")
		printTorusPolynomial(a)

		fmt.Printf("B: \n")
		printTorusPolynomial(b)

		assert.EqualValues(torusPolynomialNormInftyDist(a, acopy), 0)
		assert.LessOrEqual(torusPolynomialNormInftyDist(a, b), toler)
	}
}

//TEST(LagrangeHalfcTest, LagrangeHalfCPolynomialClear) {
func TestLagrangeHalfCPolynomialClear(t *testing.T) {
	assert := assert.New(t)
	NBTRIALS := 10
	var N int32 = 1024
	for trials := 0; trials < NBTRIALS; trials++ {
		a := NewTorusPolynomial(N)
		zero := NewTorusPolynomial(N)
		afft := NewLagrangeHalfCPolynomial(N)
		afft.Clear()
		torusPolynomialUniform(a)
		torusPolynomialClear(zero)
		a.Coefs = TorusPolynomialFft(afft)
		assert.EqualValues(torusPolynomialNormInftyDist(zero, a), 0)
	}
}

/** sets to this Torus32 constant */
//EXPORT void LagrangeHalfCPolynomialSetTorusConstant(LagrangeHalfCPolynomial* result, const Torus32 mu);
func TestLagrangeHalfCPolynomialSetTorusConstant(t *testing.T) {
	assert := assert.New(t)
	NBTRIALS := 10
	var N int32 = 1024
	for trials := 0; trials < NBTRIALS; trials++ {
		mu := UniformTorus32Dist()
		a := NewTorusPolynomial(N)
		cste := NewTorusPolynomial(N)
		afft := NewLagrangeHalfCPolynomial(N)
		torusPolynomialUniform(a)

		//tested function
		afft.SetTorusConstant(mu)
		a.Coefs = TorusPolynomialFft(afft)

		//expected result
		torusPolynomialClear(cste)
		cste.Coefs[0] = mu

		assert.EqualValues(torusPolynomialNormInftyDist(cste, a), 0)
	}
}

//EXPORT void LagrangeHalfCPolynomialAddTorusConstant(LagrangeHalfCPolynomial* result, const Torus32 cst);
func TestLagrangeHalfCPolynomialAddTorusConstant(t *testing.T) {
	assert := assert.New(t)
	NBTRIALS := 1
	var N int32 = 1024
	toler := 1e-9
	for trials := 0; trials < NBTRIALS; trials++ {
		var mu int32 = UniformTorus32Dist()
		a := NewTorusPolynomial(N)
		aPlusCste := NewTorusPolynomial(N)
		b := NewTorusPolynomial(N)
		afft := NewLagrangeHalfCPolynomial(N)

		//torusPolynomialUniform(a)
		TorusPolynomialIfft(afft, a.Coefs)
		afft.AddTorusConstant(mu)
		b.Coefs = TorusPolynomialFft(afft)

		TorusPolynomialCopy(aPlusCste, a)
		aPlusCste.Coefs[0] += mu

		assert.LessOrEqual(torusPolynomialNormInftyDistSkipFirst(aPlusCste, b), toler)
	}
}

func TestTorusPolynomialSmallMultFFT(t *testing.T) {
	assert := assert.New(t)
	NBTRIALS := 1
	var N int32 = 4
	toler := 1e-9
	for trials := 0; trials < NBTRIALS; trials++ {
		a := NewIntPolynomial(N)
		b := NewTorusPolynomial(N)
		aB := NewTorusPolynomial(N)
		aBref := NewTorusPolynomial(N)
		a.Coefs = []int32{9, -10, 7, 6}
		b.Coefs = []int32{-5, 4, 0, -2}
		torusPolynomialMultKaratsuba(aBref, a, b)
		torusPolynomialMultFFT(aB, a, b)
		assert.LessOrEqual(torusPolynomialNormInftyDist(aB, aBref), toler)
	}
}

func TestTorusPolynomialMultFFT(t *testing.T) {
	assert := assert.New(t)
	NBTRIALS := 1
	var N int32 = 1024
	toler := 1e-9
	for trials := 0; trials < NBTRIALS; trials++ {
		a := NewIntPolynomial(N)
		b := NewTorusPolynomial(N)
		aB := NewTorusPolynomial(N)
		aBref := NewTorusPolynomial(N)
		for i := int32(0); i < N; i++ {
			a.Coefs[i] = UniformTorus32Dist()%1000 - 500
		}
		torusPolynomialUniform(b)
		//a.Coefs = []int32{9, -10, 7, 6}
		//b.Coefs = []int32{-5, 4, 0, -2}
		torusPolynomialMultKaratsuba(aBref, a, b)
		torusPolynomialMultFFT(aB, a, b)
		assert.LessOrEqual(torusPolynomialNormInftyDist(aB, aBref), toler)
	}
}

func TestTorusPolynomialAddMulRFFT(t *testing.T) {
	assert := assert.New(t)
	NBTRIALS := 1
	var N int32 = 1024
	toler := 1e-9
	for trials := 0; trials < NBTRIALS; trials++ {
		a := NewIntPolynomial(N)
		b := NewTorusPolynomial(N)
		aB := NewTorusPolynomial(N)
		aBref := NewTorusPolynomial(N)
		for i := int32(0); i < N; i++ {
			a.Coefs[i] = UniformTorus32Dist()%1000 - 500
		}
		torusPolynomialUniform(b)
		torusPolynomialUniform(aB)
		TorusPolynomialCopy(aBref, aB)
		torusPolynomialAddMulRKaratsuba(aBref, a, b)
		torusPolynomialAddMulRFFT(aB, a, b)
		assert.LessOrEqual(torusPolynomialNormInftyDist(aB, aBref), toler)
	}
}

func TestTorusPolynomialSubMulRFFT(t *testing.T) {
	assert := assert.New(t)
	NBTRIALS := 1
	var N int32 = 1024
	toler := 1e-9
	for trials := 0; trials < NBTRIALS; trials++ {
		a := NewIntPolynomial(N)
		b := NewTorusPolynomial(N)
		aB := NewTorusPolynomial(N)
		aBref := NewTorusPolynomial(N)
		for i := int32(0); i < N; i++ {
			a.Coefs[i] = UniformTorus32Dist()%1000 - 500
		}
		torusPolynomialUniform(b)
		torusPolynomialUniform(aB)
		TorusPolynomialCopy(aBref, aB)
		torusPolynomialSubMulRKaratsuba(aBref, a, b)
		torusPolynomialSubMulRFFT(aB, a, b)
		assert.LessOrEqual(torusPolynomialNormInftyDist(aB, aBref), toler)
	}
}

//TEST(LagrangeHalfcTest, LagrangeHalfCPolynomialAddTo) {
func TestLagrangeHalfCPolynomialAddTo(t *testing.T) {
	assert := assert.New(t)
	NBTRIALS := 1
	var N int32 = 1024
	toler := 1e-9
	for trials := 0; trials < NBTRIALS; trials++ {
		a := NewTorusPolynomial(N)
		b := NewTorusPolynomial(N)
		aPlusB := NewTorusPolynomial(N)
		aPlusBbis := NewTorusPolynomial(N)
		afft := NewLagrangeHalfCPolynomial(N)
		bfft := NewLagrangeHalfCPolynomial(N)
		torusPolynomialUniform(a)
		TorusPolynomialIfft(afft, a.Coefs)
		torusPolynomialUniform(b)
		TorusPolynomialIfft(bfft, b.Coefs)
		afft.AddTo(bfft)
		aPlusBbis.Coefs = TorusPolynomialFft(afft)
		TorusPolynomialAdd(aPlusB, b, a)
		assert.LessOrEqual(torusPolynomialNormInftyDist(aPlusBbis, aPlusB), toler)
	}
}