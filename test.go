package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

var data = `
search:
  &f foo:
    type: node
    constraints:
      - leisure=fitness_station
  bar:
    type: node
    constraints:
      - leisure=fitness_station
    distance_constraint:
        from: *f
        distance: 500
`

type DistanceConstraint struct {
	From     yaml.Node `yaml:"from"`
	Distance int       `yaml:"distance"`
}

type Y struct {
	Type        string             `yaml:"type"`
	Constraints []string           `yaml:"constraints,omitempty"`
	Distance    DistanceConstraint `yaml:"distance_constraint,omitempty"`
}

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type T struct {
	Search map[string]Y `yaml:"search"`
}

func main() {
	t := T{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))
}
