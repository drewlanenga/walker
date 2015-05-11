package walker

import "testing"

var (
	testNprocs  = 8
	testNiter   = 10000
	testNsteps  = 50
	testHistory = Vector{0.1, 0.2, 0.3, 0.2, 0.4, 0.8, 0.5, 1.0}
)

func TestParWalks(t *testing.T) {
	Walks(testNiter, testNsteps, testNprocs, testHistory)
}

func TestSeqWalks(t *testing.T) {
	walks(testNiter, testNsteps, testHistory)
}

func BenchmarkParWalks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Walks(testNiter, testNsteps, testNprocs, testHistory)
	}
}

func BenchmarkWalks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		walks(testNiter, testNsteps, testHistory)
	}
}
