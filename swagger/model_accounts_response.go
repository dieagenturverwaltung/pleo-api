/*
 * Pleo.io API
 *
 * Pleo.io API
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AccountsResponse struct {
	Accounts []AccountsResponseAccounts `json:"accounts,omitempty"`
	Metadata *ExpensesResponseMetadata  `json:"metadata"`
}
