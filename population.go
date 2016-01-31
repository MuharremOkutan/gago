package genalg

import (
	"math/rand"
	"time"
)

// A Population contains demes.
type Population struct {
	Demes         []Deme
	Nbdemes       int
	Nbindividuals int
	Ff            func([]float64) float64
	Best          Individual
	Genes         int
	Boundary      float64
	Couples       int
	Offsprings    int
	Rate          float64
	Std           float64
	Contestants   int
}

// GA is default configuration
var GA = Population{
	// Number of demes
	Nbdemes: 1,
	// Number of individuals in each deme
	Nbindividuals: 30,
	// Initial random boundaries
	Boundary: 100.0,
	// Number of couples that generate offsprings
	Couples: 10,
	// Number of offspings generated by each couple
	Offsprings: 2,
	// Mutation rate
	Rate: 0.1,
	// Mutation normal distribution standard deviation
	Std: 1,
	// Number of contestants for each tournament selection round
	Contestants: 3,
}

// Initialize each deme in the population and assign an initial fitness to each
// individual in each deme.
func (pop *Population) Initialize(ff func([]float64) float64, variables int) {
	// Fitness function
	pop.Ff = ff
	// Number of genes in each individual
	pop.Genes = variables
	// Create the demes
	pop.Demes = make([]Deme, 1)
	// Best individual (dummy instantiation)
	pop.Best = Individual{make([]float64, pop.Genes), 0.0}
	// Set a new random seed
	rand.Seed(time.Now().UTC().UnixNano())
	for i := range pop.Demes {
		// Create the deme
		deme := Deme{pop.Nbindividuals, make([]Individual, pop.Nbindividuals)}
		// Initialize the deme
		deme.initialize(pop.Genes, pop.Boundary)
		// Add it to the population
		pop.Demes[i] = deme
		// Initial evaluation
		pop.Demes[i].evaluate(pop.Ff)
		// Random best individual to start with
		pop.Best = pop.Demes[0].individuals[0]
	}
}

// FindBest stores the best individual over all demes.
func (pop *Population) FindBest() {
	leaders := make(Individuals, len(pop.Demes))
	for i, deme := range pop.Demes {
		leaders[i] = deme.individuals[0]
	}
	leaders.sort()
	if leaders[0].Fitness < pop.Best.Fitness {
		pop.Best = leaders[0]
	}
}

// Enhance each deme in the population.
func (pop *Population) Enhance() {
	for i := range pop.Demes {
		pop.Demes[i].crossover(pop.Couples, pop.Offsprings)
		pop.Demes[i].mutate(pop.Rate, pop.Std)
		pop.Demes[i].evaluate(pop.Ff)
		pop.Demes[i].tournament(pop.Contestants)
		pop.Demes[i].sort()
		pop.FindBest()
	}
}
