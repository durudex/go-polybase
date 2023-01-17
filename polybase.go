/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package polybase

// TestnetURL constant stores the URL of the latest version of the Polybase testnet.
const TestnetURL string = "https://testnet.polybase.xyz/v0"

// Polybase interface stores methods for interacting with the Polybase database.
type Polybase interface {
	// Collection method returns an implementation for interacting with specified Polybase
	// collection. It takes one argument, which is the name of the collection, it can be a
	// full name with namespace, but if you don't use the DefaultNamespace value in the
	// configuration.
	Collection(name string) Collection
}

// Block structure stores data about a block from the blockchain.
type Block struct {
	// The hash field stores the hash of a block from the blockchain.
	Hash string `json:"hash"`
}

// Cursor structure stores data used for pagination.
type Cursor struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

// New function returns a new Polybase client.
func New(cfg Config) Polybase {
	cfg.configure()

	return &polybase{client: NewClient(cfg), cfg: cfg}
}

// polybase structure implements all methods of the Polybase interface.
type polybase struct {
	client Client
	cfg    Config
}

// Collection method returns an implementation for interacting with specified Polybase collection.
func (p *polybase) Collection(name string) Collection {
	if p.cfg.DefaultNamespace != "" {
		name = p.cfg.DefaultNamespace + "/" + name
	}

	return newCollection(name, p.client)
}
