/*
 * Pleo.io API
 *
 * Pleo.io API
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AccountsResponseAccounts struct {
	// The unique UUID identifier of the account
	Id string `json:"id"`
	// Account number
	AccountNumber string `json:"accountNumber,omitempty"`
	// unique Tax Code ID
	TaxCodeId string `json:"taxCodeId,omitempty"`
	// Description of this account
	Name string `json:"name"`
	// unique Category ID of the account
	AccountCategoryId string `json:"accountCategoryId"`
	// Is the account hidden/archived for the user
	Hidden bool `json:"hidden"`
}
