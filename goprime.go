package goprime

import (
	"errors"
	"math/rand"
)

type (
	// PrimeNumber type
	PrimeNumber uint64

	// PrimeNumbers type
	PrimeNumbers []PrimeNumber

	// Next type
	Next map[PrimeNumber]PrimeNumber
)

var (
	primes   PrimeNumbers
	balanced PrimeNumbers
	next     Next

	// ErrOutOfRange out of range error
	ErrOutOfRange = errors.New("Error: Out of range")
)

// Client type
type Client struct {
	Primes PrimeNumbers
	Next   Next
}

// New creates a new instance of goprime
func New() (*Client, error) {
	c := &Client{
		Primes: primeData,
		Next:   Next{},
	}

	for index := range c.Primes {
		if index == len(c.Primes)-1 {
			break
		}
		currentPrime, err := c.Primes.Prime(PrimeNumber(index))
		if err != nil {
			return nil, err
		}
		nextPrime, err := c.Primes.Prime(PrimeNumber(index) + 1)
		if err != nil {
			return nil, err
		}

		c.Next[currentPrime] = nextPrime
	}

	primes = c.Primes
	next = c.Next

	return c, nil
}

// Prime return prime i
func (p PrimeNumbers) Prime(i PrimeNumber) (PrimeNumber, error) {
	if len(p) < int(i) {
		return 0, ErrOutOfRange
	}
	return p[i], nil
}

// Next return the next prime
func (p PrimeNumber) Next() (PrimeNumber, error) {
	next, ok := next[p]
	if !ok {
		return 0, ErrOutOfRange
	}
	return next, nil
}

// First return the first prime
func (p PrimeNumbers) First() PrimeNumber {
	return p[0]
}

// Last return the last prime (in this collection, lol..)
func (p PrimeNumbers) Last() PrimeNumber {
	return p[len(p)-1]
}

// Collection return a collection from s to e
func (p PrimeNumbers) Collection(s, e int) PrimeNumbers {
	return p[s : e+1]
}

// Random return a random prime from the collection
func (p PrimeNumbers) Random() (PrimeNumber, error) {
	index := rand.Intn(len(p))
	return p.Prime(PrimeNumber(index))
}
