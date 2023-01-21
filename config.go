/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package polybase

// Config structure stores the configuration for interacting with Polybase database.
type Config struct {
	// URL field stores the URL address to which the Polybase client will be connected.
	URL string

	// Name field stores the name of the client, which will be transmitted via the
	// X-Polybase-Client HTTP header.
	//
	// Optional. Default value "default".
	Name string

	// DefaultNamespace field stores the namespace used in the Polybase.Collection method as
	// a name prefix.
	//
	// Optional.
	DefaultNamespace string
}

func (c *Config) configure() {
	if c.Name == "" {
		c.Name = "default"
	}
}
