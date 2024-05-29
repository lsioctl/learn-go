package graph

import "graphql-poc/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	// map to store the characters
	CharacterStore map[string]model.Character
}
