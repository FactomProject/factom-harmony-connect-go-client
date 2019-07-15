/*
 * Harmony Connect
 *
 * An easy to use API that helps you access the Factom blockchain.
 *
 * API version: 1.0.19
 * Contact: harmony-support@factom.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package connectclient

// A link to the directory block.
type DBlockLink struct {
	// The Key Merkle Root for this directory block.
	Keymr string `json:"keymr"`
	// The Factom block height of this block.
	Height int32 `json:"height"`
	// An API link to retrieve all available information about this directory block.
	Href string `json:"href"`
}