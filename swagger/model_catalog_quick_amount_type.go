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

// CatalogQuickAmountType : Determines the type of a specific Quick Amount.
type CatalogQuickAmountType string

// List of CatalogQuickAmountType
const (
	DO_NOT_USE_CatalogQuickAmountType CatalogQuickAmountType = "QUICK_AMOUNT_TYPE_DO_NOT_USE"
	MANUAL_CatalogQuickAmountType     CatalogQuickAmountType = "QUICK_AMOUNT_TYPE_MANUAL"
	AUTO_CatalogQuickAmountType       CatalogQuickAmountType = "QUICK_AMOUNT_TYPE_AUTO"
)
