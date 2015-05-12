package walker

import (
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var (
	rnd   = rand.New(rand.NewSource(time.Now().UnixNano()))
	rndmu = &sync.Mutex{}
)

type Vector []float64

func (v Vector) Diff() Vector {
	y := make(Vector, len(v)-1)
	for i := 0; i < len(y); i++ {
		y[i] = v[i+1] - v[i]
	}
	return y
}

func Walks(niter, nsteps, ncpu int, dest float64, history Vector) (float64, float64) {
	runtime.GOMAXPROCS(ncpu)
	destinations := make(Vector, niter)

	steps := history.Diff()

	c := make(chan int, ncpu)
	for i := 0; i < niter; i++ {
		go destinations.Walk(i, nsteps, steps, c)
	}

	// drain the channel
	for i := 0; i < ncpu; i++ {
		<-c // wait for one task to complete
	}

	return compare(destinations, dest)
}

func (v Vector) Walk(i, nsteps int, steps Vector, c chan int) {
	dest := walk(nsteps, steps)

	v[i] = dest
	c <- 1
}

func walk(nsteps int, steps Vector) float64 {
	var dest float64
	for i := 0; i < nsteps; i++ {
		rndmu.Lock()
		which := rnd.Intn(len(steps))
		rndmu.Unlock()

		dest += steps[which]
	}
	return dest
}

func walks(niter, nsteps int, dest float64, history Vector) (float64, float64) {
	destinations := make(Vector, niter)
	for i := 0; i < niter; i++ {
		destinations[i] = walk(nsteps, history)
	}
	return compare(destinations, dest)
}

func compare(destinations Vector, dest float64) (float64, float64) {
	n := float64(len(destinations))
	var nlow, nhigh int

	for _, d := range destinations {
		if d > dest {
			nlow++
		} else if d < dest {
			nhigh++
		}
	}

	return float64(nlow) / n, float64(nhigh) / n
}
