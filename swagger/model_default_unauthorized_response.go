/*
 * Pleo.io API
 *
 * Pleo.io API
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DefaultUnauthorizedResponse struct {
	// http response status code
	StatusCode float64 `json:"statusCode,omitempty"`
	// error type
	Error_ string `json:"error,omitempty"`
	// error message
	Message string `json:"message,omitempty"`
}