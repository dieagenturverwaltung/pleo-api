/*
 * Pleo.io API
 *
 * Pleo.io API
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AccountCreationRequest struct {
	// Account number
	AccountNumber string `json:"accountNumber,omitempty"`
	// Description of this account
	Name string `json:"name,omitempty"`
	// unique Tax Code ID
	TaxCodeId string `json:"taxCodeId,omitempty"`
	// unique Category ID of the account
	AccountCategoryId string `json:"accountCategoryId,omitempty"`
	// Is the account hidden/archived for the user
	Hidden bool `json:"hidden,omitempty"`
}
