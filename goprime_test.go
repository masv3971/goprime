package goprime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	_, err := New()
	assert.NoError(t, err)
}

func TestPrime(t *testing.T) {
	tts := []struct {
		name string
		have PrimeNumbers
		arg  PrimeNumber
		want PrimeNumber
	}{
		{
			name: "First",
			have: primeData,
			arg:  0,
			want: 2,
		},
		{
			name: "Second",
			have: primeData,
			arg:  1,
			want: 3,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.have.Prime(tt.arg)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNext(t *testing.T) {
	tts := []struct {
		name string
		have PrimeNumber
		want PrimeNumber
	}{
		{
			name: "first",
			have: 2,
			want: 3,
		},
		{
			name: "second",
			have: 3,
			want: 5,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New()
			assert.NoError(t, err)

			got, err := tt.have.Next()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFirst(t *testing.T) {
	got := primeData.First()
	assert.Equal(t, PrimeNumber(2), got)
}

func TestLast(t *testing.T) {
	got := primeData.Last()
	assert.Equal(t, PrimeNumber(999983), got)
}

func TestRandom(t *testing.T) {
	got, err := primeData.Random()
	assert.NoError(t, err)
	assert.Equal(t, PrimeNumber(782581), got)
}

func TestPrimeData_checkForEven(t *testing.T) {
	for _, p := range primeData {
		if p%2 == 0 && p != 2 {
			t.Errorf("Found even %d", p)
		}
	}
}

func TestCollection(t *testing.T) {
	_, err := New()
	assert.NoError(t, err)

	type args struct {
		first, last int
	}
	tts := []struct {
		name string
		args args
		want PrimeNumbers
	}{
		{
			name: "First three primes",
			args: args{
				first: 0,
				last:  2,
			},
			want: PrimeNumbers{2, 3, 5},
		},
		{
			name: "from third to fifth prime",
			args: args{
				first: 3,
				last:  5,
			},
			want: PrimeNumbers{7, 11, 13},
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			got := primeData.Collection(tt.args.first, tt.args.last)
			assert.Equal(t, tt.want, got)

		})
	}
}

func TestCountAmountPrimes(t *testing.T) {
	t.Logf("Amount of primes: %d", len(primeData))
}
