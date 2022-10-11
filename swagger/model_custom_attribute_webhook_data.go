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

type CustomAttributeWebhookData struct {
	// The type of the event data object. The value is `\"custom_attribute\"`.
	Type_ string `json:"type,omitempty"`
	// The ID of the event data object.
	Id     string                            `json:"id,omitempty"`
	Object *CustomAttributeWebhookDataObject `json:"object,omitempty"`
}