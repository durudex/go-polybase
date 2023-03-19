/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package polybase

const DefaultName = "default"

// Config structure stores the configuration for interacting with Polybase database.
type Config struct {
	// The URL field defines a url to a node or any other Polybase API gateway to which the
	// client will send requests. You can use pre-defined internal url values or specify your
	// own url values.
	URL string `json:"url"`

	// Name field stores the name of the client, which will be transmitted via the
	// X-Polybase-Client HTTP header.
	//
	// Optional. Default value "default".
	Name string `json:"name"`

	// DefaultNamespace field stores the namespace used in the Polybase.Collection method as
	// a name prefix.
	//
	// Optional.
	DefaultNamespace string `json:"defaultName"`

	RecoverHandler RecoverHandler `json:"-"`
}

func (c *Config) configure() {
	if c.URL == "" {
		c.URL = DefaultURL
	}
	if c.Name == "" {
		c.Name = DefaultName
	}
	if c.RecoverHandler == nil {
		c.RecoverHandler = DefaultRecover
	}
}
