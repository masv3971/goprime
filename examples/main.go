package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/masv3971/goprime"
)

func main() {
	c, err := goprime.New()
	if err != nil {
		panic(err)
	}

	for i, p := range c.Primes {
		fmt.Println("prime", p)

		if i == len(c.Primes)-1 {
			break
		}

		next, err := p.Next()
		if err != nil {
			panic(err)
		}

		fmt.Println("next prime", next)
	}

	fmt.Println("collection: ", c.Primes.Collection(3, 5))

	// keep in mind that PrimeNumber have to be a valid prime, Next() won't find non prime numbers.
	next, err := goprime.PrimeNumber(2).Next()
	if err != nil {
		panic(err)
	}
	fmt.Println("next prime", next)

	fmt.Println("first prime in collection", c.Primes.First())
	fmt.Println("last prime in collection", c.Primes.Last())

	// Always seed random!
	rand.Seed(time.Now().Unix())
	r, err := c.Primes.Random()
	if err != nil {
		panic(err)
	}
	fmt.Println("random prime", r)

	p10, err := c.Primes.Prime(10)
	if err != nil {
		panic(err)
	}
	fmt.Println("the 10th prime number", p10)
}
