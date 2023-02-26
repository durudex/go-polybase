/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package polybase

// TestnetURL constant stores the URL of the latest version of the Polybase testnet.
const TestnetURL string = "https://testnet.polybase.xyz/v0"

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
