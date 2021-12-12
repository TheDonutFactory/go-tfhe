package tfhe

import (
	"time"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

type Torus = int32
type Binary = uint32

type TGSWKey = TLWEKey

func Assert(condition bool) {
	if !condition {
		panic("Assertion error")
	}
}

func UniformDist(min, max int) distuv.Uniform {
	return distuv.Uniform{
		Min: float64(min),
		Max: float64(max),
	}
}

func UniformDistF(min, max float64) distuv.Uniform {
	s1 := rand.NewSource(uint64(time.Now().UnixNano()))
	return distuv.Uniform{
		Min: min,
		Max: max,
		Src: s1,
	}
}

func NormalDist(mu, sigma float64) distuv.Normal {
	s1 := rand.NewSource(uint64(time.Now().UnixNano()))
	return distuv.Normal{
		Mu:    mu,
		Sigma: sigma,
		Src:   s1,
	}
}

func ModSwitchToTorus(mu, space int32) Torus {
	gap := ((uint64(1) << 63) / uint64(space)) * 2
	return int32(uint64((uint64(mu) * gap) >> 32))
}

// LWESample
type LWESample struct {
	A []Torus
	B *Torus
	N int
}

func NewLWESample(n int) *LWESample {
	a := make([]Torus, n+1)
	return &LWESample{A: a, B: &a[n], N: n}
}

func NewLweSampleArray(n, t int) (arr []*LWESample) {
	arr = make([]*LWESample, n)
	for i := 0; i < n; i++ {
		arr[i] = NewLWESample(t)
	}
	return
}

// LWEKey
type LWEKey struct {
	Key []Binary
	N   int
}

func NewLWEKey(n int) *LWEKey {
	return &LWEKey{Key: make([]Binary, n), N: n}
}

// LweParams
type LweParams struct {
	N        int
	alphaMin float64
	alphaMax float64
}

func NewLweParams(n int, min, max float64) *LweParams {
	return &LweParams{n, min, max}
}

/*
type IntPolynomial struct {
	N     int
	Coefs []int32
}

type TorusPolynomial struct {
	N     int
	Coefs []Torus
}

func NewTorusPolynomial(n int) *TorusPolynomial {
	return &TorusPolynomial{N: n, Coefs: make([]Torus, n)}
}

func NewTorusPolynomialArray(size, n int) (arr []TorusPolynomial) {
	arr = make([]TorusPolynomial, size)
	for i := 0; i < size; i++ {
		arr[i] = TorusPolynomial{N: n, Coefs: make([]Torus, n)}
	}
	return
}
*/

// TLWESample
type TLWESample struct {
	A [][]Torus
	K int
	N int
}

func NewTLWESample(n, k int) *TLWESample {
	//arr := NewTorusPolynomialArray(k+1, n)
	arr := make([][]Torus, k+1)
	for i := range arr {
		//arr[i] = make([]Torus, n)
		arr[i] = make([]Torus, n+1)
		//return &LWESample{A: a, B: &a[n], N: n}
	}

	return &TLWESample{
		A: arr,
		N: n,
		K: k,
	}
}

func NewTLWESampleArray(size, n, k int) (arr []*TLWESample) {
	arr = make([]*TLWESample, size)
	for i := 0; i < size; i++ {
		arr[i] = NewTLWESample(n, k)
	}
	return
}

func (sample *TLWESample) NumPolys() int {
	return sample.K + 1
}

func (sample *TLWESample) ExtractPoly(index int) []Torus {
	Assert(index <= sample.K)
	return sample.A[index]
}

func (sample *TLWESample) a(index int) []Torus {
	Assert(index <= sample.K)
	return sample.ExtractPoly(index)
}

func (sample *TLWESample) b() []Torus {
	return sample.A[sample.K]
}

// TLWEKey
type TLWEKey struct {
	A [][]Binary
	K int
	N int
}

func NewTLWEKey(n, k int) *TLWEKey {
	a := make([][]Binary, k)
	for i := range a {
		a[i] = make([]Binary, n)
	}
	return &TLWEKey{A: a, N: n, K: k}
}

func NewTLWEKeyArray(size, n, k int) (arr []*TLWEKey) {
	arr = make([]*TLWEKey, size)
	for i := 0; i < size; i++ {
		arr[i] = NewTLWEKey(n, k)
	}
	return
}

func (sample *TLWEKey) NumPolys() int {
	return sample.K
}

func (sample *TLWEKey) ExtractPoly(index int) []Binary {
	Assert(index < sample.K)
	return sample.A[index]
}

func (sample *TLWEKey) ExtractLWEKey() *LWEKey {
	return &LWEKey{Key: sample.A[sample.K-1], N: sample.N}
}

// TGSWSample
type TGSWSample struct {
	A [][]*TLWESample
	Y []*TLWESample // flattned representation of A
	N int
	K int
	L int
	W int
}

func NewTGSWSample(n, k, l, w int) *TGSWSample {
	y := make([]*TLWESample, (k+1)*l)
	a := make([][]*TLWESample, k+1)
	var c int = 0
	for i := range a {
		a[i] = make([]*TLWESample, l)
		for j := range a[i] {
			a[i][j] = NewTLWESample(n, k)
			y[c] = a[i][j]
			c++
		}
	}

	return &TGSWSample{A: a, Y: y, N: n, K: k, L: l, W: w}
}

func NewTGSWSampleArray(size, n, k, l, w int) (arr []*TGSWSample) {
	arr = make([]*TGSWSample, size)
	for i := 0; i < size; i++ {
		arr[i] = NewTGSWSample(n, k, l, w)
	}
	return
}

func (sample *TGSWSample) NumTLWESamples() int {
	return (sample.K + 1) * sample.L
}

func (sample *TGSWSample) ExtractTLWESample(index int) *TLWESample {
	Assert(index < len(sample.Y))
	return sample.Y[index]
}

// KeySwitchingKey
type KeySwitchingKey struct {
	// m, l, w
	A   [][][]*LWESample
	N   int
	L   int // decomp size
	W   int // basebit
	M   int
	T   int
	raw []*LWESample
}

func NewKeySwitchingKey(n, l, w, m int) *KeySwitchingKey {

	/*
		ksk_: NewKeySwitchingKey(param.lwe_n_,
			param.keyswitching_decomp_size_,
			param.keyswitching_decomp_bits_,
			param.tlwe_n_*param.tlwe_k_),
	*/

	t := m * l << w

	ks := make([][][]*LWESample, m)
	raw := make([]*LWESample, m*l*(0x1<<w))
	var c int = 0
	for i := 0; i < m; i++ {
		ks[i] = make([][]*LWESample, l)
		for j := 0; j < l; j++ {
			ks[i][j] = make([]*LWESample, (0x1 << w))
			for k := 0; k < (0x1 << w); k++ {
				ks[i][j][k] = NewLWESample(n)
				raw[c] = ks[i][j][k]
				c++
			}
		}
	}

	return &KeySwitchingKey{
		A: ks, raw: raw, N: n, L: l, W: w, M: m, T: t,
	}

}

/*
func (me *KeySwitchingKey) GetLWESampleIndex(degree, digit, value int) int {
	Assert(degree < me.M)
	Assert(digit < me.L)
	Assert(value < (0x1 << me.W))
	return ((degree*me.L + digit) << me.W) + value
}


func (me *KeySwitchingKey) ExtractLWESample(index int) *LWESample {
	return me.raw[index]
}
*/

func (me *KeySwitchingKey) NumLWESamples() int {
	return me.T
}

// type BootstrappingKey = []TGSWSample
// TGSWSample

type BootstrappingKey struct {
	Bk []*TGSWSample
	//ks *KeySwitchingKey
	N int
	K int
	L int
	W int
	T int
}

func NewBootstrappingKey(n, k, l, w, t int) *BootstrappingKey {
	arr := NewTGSWSampleArray(t, n, k, l, w)
	return &BootstrappingKey{Bk: arr, N: n, K: k, L: l, W: w, T: t}
}

func NewBootstrappingKeyArray(size, n, k, l, w, t int) (arr []*BootstrappingKey) {
	arr = make([]*BootstrappingKey, size)
	for i := 0; i < size; i++ {
		arr[i] = NewBootstrappingKey(n, k, l, w, t)
	}
	return
}

func (key *BootstrappingKey) ExtractTGSWSample(index int) *TGSWSample {
	return key.Bk[index]
}
