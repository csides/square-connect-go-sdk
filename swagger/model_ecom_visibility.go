/*
 * Square
 *
 * Use Square APIs to manage and run business including payment, customer, product, inventory, and employee management.
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// EcomVisibility : Determines item visibility in Ecom (Online Store) and Online Checkout.
type EcomVisibility string

// List of EcomVisibility
const (
	ECOM_VISIBILITY_DO_NOT_USE_EcomVisibility EcomVisibility = "ECOM_VISIBILITY_DO_NOT_USE"
	UNINDEXED_EcomVisibility                  EcomVisibility = "UNINDEXED"
	UNAVAILABLE_EcomVisibility                EcomVisibility = "UNAVAILABLE"
	HIDDEN_EcomVisibility                     EcomVisibility = "HIDDEN"
	VISIBLE_EcomVisibility                    EcomVisibility = "VISIBLE"
)
