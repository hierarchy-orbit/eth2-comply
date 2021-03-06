/*
 * Eth2 Beacon Node API
 *
 * API specification for the beacon node, which enables users to query and participate in Ethereum 2.0 phase 0 beacon chain.
 *
 * API version: Dev - Eth2Spec v0.12.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package eth2spec
// InlineResponse400 struct for InlineResponse400
type InlineResponse400 struct {
	// Either specific error code in case of invalid request or http status code
	Code float32 `json:"code,omitempty"`
	// Message describing error
	Message string `json:"message,omitempty"`
	// Optional stacktraces, sent when node is in debug mode
	Stacktraces []string `json:"stacktraces,omitempty"`
}
