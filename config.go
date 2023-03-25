/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package polybase

// The constant DefaultName contains the default client name, which is used as the value of
// the 'X-Polybase-Client' HTTP header.
const DefaultName = "default"

// The Config structure defines the configuration used to set up the go-polybase client.
type Config struct {
	// The URL field defines a url to a node or any other Polybase API gateway to which the
	// client will send requests. You can use pre-defined internal url values or specify your
	// own url values.
	//
	// Internal values:
	//
	//	- DefaultURL (Default)
	//	- TestnetURL
	URL string `json:"url"`

	// The Name field defines the client name used as the value of the 'X-Polybase-Client' HTTP
	// header in requests to the Polybase API.
	//
	// Additionally, for better analysis, the prefix "durudex/go-polybase:" is added to each name.
	// This allows for easier identification of the module or library from which requests are made.
	//
	// Optional. Default DefaultName.
	Name string `json:"name"`

	// The DefaultNamespace field defines the default namespace that will be added to the
	// collection name when creating a new instance.
	//
	// Optional.
	DefaultNamespace string `json:"defaultName"`

	// The RecoverHandler field defines the handler that will be called in case of a panic.
	//
	// Panics usually occur during development and may indicate passing an incorrect type or a lack
	// of connection to the internet or the Polybase API.
	//
	// Optional. Default DefaultRecover.
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
