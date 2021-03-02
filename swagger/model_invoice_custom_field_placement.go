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

// InvoiceCustomFieldPlacement : Indicates where to render a custom field on the Square-hosted invoice page and in emailed or PDF  copies of the invoice.
type InvoiceCustomFieldPlacement string

// List of InvoiceCustomFieldPlacement
const (
	UNKNOWN_PLACEMENT_DO_NOT_USE_InvoiceCustomFieldPlacement InvoiceCustomFieldPlacement = "UNKNOWN_PLACEMENT_DO_NOT_USE"
	ABOVE_LINE_ITEMS_InvoiceCustomFieldPlacement             InvoiceCustomFieldPlacement = "ABOVE_LINE_ITEMS"
	BELOW_LINE_ITEMS_InvoiceCustomFieldPlacement             InvoiceCustomFieldPlacement = "BELOW_LINE_ITEMS"
)
