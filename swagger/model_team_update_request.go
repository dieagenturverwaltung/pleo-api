/*
 * Pleo.io API
 *
 * Pleo.io API
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type TeamUpdateRequest struct {
	// Team name
	Name string `json:"name"`
	// Team code
	Code string `json:"code,omitempty"`
}