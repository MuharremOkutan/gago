package gago

import "testing"

func TestMutNormalFloat64All(t *testing.T) {
	var (
		rng     = makeRandomNumberGenerator()
		genome  = []float64{1, 2, 3}
		mutated = make([]float64, len(genome))
	)
	copy(mutated, genome)
	MutNormalFloat64(mutated, 1, rng)
	for i, v := range mutated {
		if v == genome[i] {
			t.Error("Gene should have been modified but hasn't")
		}
	}
}

func TestMutNormalFloat64None(t *testing.T) {
	var (
		rng     = makeRandomNumberGenerator()
		genome  = []float64{1, 2, 3}
		mutated = make([]float64, len(genome))
	)
	copy(mutated, genome)
	MutNormalFloat64(mutated, 0, rng)
	for i, v := range mutated {
		if v != genome[i] {
			t.Error("Gene has been modified but shouldn't have")
		}
	}
}

func TestMutUniformString(t *testing.T) {
	var (
		rng     = makeRandomNumberGenerator()
		genome  = []string{"a", "b", "c"}
		corpus  = []string{"d", "e", "f"}
		mutated = make([]string, len(genome))
	)
	copy(mutated, genome)
	MutUniformString(mutated, corpus, 3, rng)
	// Check the length of the mutated genome is consistent
	if len(mutated) != len(genome) {
		t.Error("Mutated genome has the wrong length")
	}
	// Check the new genes are present in the previous genome or in the corpus
	for _, v := range mutated {
		var inGenome = false
		for _, gene := range genome {
			if gene == v {
				inGenome = true
			}
		}
		var inCorpus = false
		for _, element := range corpus {
			if element == v {
				inCorpus = true
			}
		}
		if !inGenome && !inCorpus {
			t.Error("New genome is not present in previous genome or in corpus")
		}
	}
}

func TestMutPermuteFloat64(t *testing.T) {
	var (
		rng     = makeRandomNumberGenerator()
		genome  = []float64{1, 2, 3}
		mutated = make([]float64, len(genome))
	)
	copy(mutated, genome)
	MutPermuteFloat64(mutated, 3, rng)
	// Check the length of the mutated genome is consistent
	if len(mutated) != len(genome) {
		t.Error("Mutated genome has the wrong length")
	}
	// Check the genes in the initial genome are still present
	for _, v := range genome {
		if !elementInSlice(v, uncastFloat64s(mutated)) {
			t.Error("Gene in initial genome has disappeared")
		}
	}
}

func TestMutPermuteInt(t *testing.T) {
	var (
		rng     = makeRandomNumberGenerator()
		genome  = []int{1, 2, 3}
		mutated = make([]int, len(genome))
	)
	copy(mutated, genome)
	MutPermuteInt(mutated, 3, rng)
	// Check the length of the mutated genome is consistent
	if len(mutated) != len(genome) {
		t.Error("Mutated genome has the wrong length")
	}
	// Check the genes in the initial genome are still present
	for _, v := range genome {
		if !elementInSlice(v, uncastInts(mutated)) {
			t.Error("Gene in initial genome has disappeared")
		}
	}
}

func TestMutPermuteString(t *testing.T) {
	var (
		rng     = makeRandomNumberGenerator()
		genome  = []string{"a", "b", "c"}
		mutated = make([]string, len(genome))
	)
	copy(mutated, genome)
	MutPermuteString(mutated, 3, rng)
	// Check the length of the mutated genome is consistent
	if len(mutated) != len(genome) {
		t.Error("Mutated genome has the wrong length")
	}
	// Check the genes in the initial genome are still present
	for _, v := range genome {
		if !elementInSlice(v, uncastStrings(mutated)) {
			t.Error("Gene in initial genome has disappeared")
		}
	}
}

func TestMutSpliceFloat64(t *testing.T) {
	var (
		rng     = makeRandomNumberGenerator()
		genome  = []float64{1, 2, 3}
		mutated = make([]float64, len(genome))
	)
	copy(mutated, genome)
	MutSpliceFloat64(mutated, rng)
	// Check the length of the mutated genome is consistent
	if len(mutated) != len(genome) {
		t.Error("Mutated genome has the wrong length")
	}
	// Check the genes in the initial genome are still present
	for _, v := range genome {
		if !elementInSlice(v, uncastFloat64s(mutated)) {
			t.Error("Gene in initial genome has disappeared")
		}
	}
}

func TestMutSpliceInt(t *testing.T) {
	var (
		rng     = makeRandomNumberGenerator()
		genome  = []int{1, 2, 3}
		mutated = make([]int, len(genome))
	)
	copy(mutated, genome)
	MutSpliceInt(mutated, rng)
	// Check the length of the mutated genome is consistent
	if len(mutated) != len(genome) {
		t.Error("Mutated genome has the wrong length")
	}
	// Check the genes in the initial genome are still present
	for _, v := range genome {
		if !elementInSlice(v, uncastInts(mutated)) {
			t.Error("Gene in initial genome has disappeared")
		}
	}
}

func TestMutSpliceString(t *testing.T) {
	var (
		rng     = makeRandomNumberGenerator()
		genome  = []string{"a", "b", "c"}
		mutated = make([]string, len(genome))
	)
	copy(mutated, genome)
	MutSpliceString(mutated, rng)
	// Check the length of the mutated genome is consistent
	if len(mutated) != len(genome) {
		t.Error("Mutated genome has the wrong length")
	}
	// Check the genes in the initial genome are still present
	for _, v := range genome {
		if !elementInSlice(v, uncastStrings(mutated)) {
			t.Error("Gene in initial genome has disappeared")
		}
	}
}
