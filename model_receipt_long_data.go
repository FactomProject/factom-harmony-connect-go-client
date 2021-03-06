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

// Contains the receipt information for the desired entry.
type ReceiptLongData struct {
	// The timestamp for this entry. Sent in [ISO 8601 Format](https://en.wikipedia.org/wiki/ISO_8601). For example: `YYYY-MM-DDThh:mm:ssZ`
	CreatedAt string `json:"created_at"`
	// The raw data that makes up the entry.
	EntrySerialized string `json:"entry_serialized"`
	// The unique identitfier of the entry.
	EntryHash string `json:"entry_hash"`
	// The branch of the merkle tree that represents this entry. Presented as an array of Merkle nodes.
	MerkleBranch []MerkleNode `json:"merkle_branch"`
	Eblock EBlockLink `json:"eblock"`
	Dblock DBlockLink `json:"dblock"`
}
