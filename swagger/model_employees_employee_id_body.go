/*
 * Pleo.io API
 *
 * Pleo.io API
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EmployeesEmployeeIdBody struct {
	// Employee first name
	FirstName string `json:"firstName,omitempty"`
	// Employee last name
	LastName string `json:"lastName,omitempty"`
	// Email
	Email string `json:"email,omitempty"`
	// Job title
	JobTitle string `json:"jobTitle,omitempty"`
	// Phone number
	Phone string `json:"phone,omitempty"`
	// External Id
	Code string `json:"code,omitempty"`
}
