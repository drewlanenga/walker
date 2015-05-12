package walker

import (
	"testing"

	"github.com/bmizerany/assert"
)

var (
	testNprocs  = 8
	testNiter   = 10000
	testNsteps  = 50
	testHistory = Vector{0.1, 0.2, 0.3, 0.2, 0.4, 0.8, 0.5, 1.0}
	testDest    = 5.4
)

func TestParWalks(t *testing.T) {
	plower, _ := Walks(testNiter, testNsteps, testNprocs, testDest, testHistory)
	assert.T(t, plower < 0.1)
}

func TestSeqWalks(t *testing.T) {
	walks(testNiter, testNsteps, testDest, testHistory)
}

func BenchmarkParWalks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Walks(testNiter, testNsteps, testNprocs, testDest, testHistory)
	}
}

func BenchmarkWalks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		walks(testNiter, testNsteps, testDest, testHistory)
	}
}
