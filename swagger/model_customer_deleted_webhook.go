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

// Published when a [customer](#type-Customer) is deleted.  For more information, see [Use Customer Webhooks](https://developer.squareup.com/docs/customers-api/use-the-api/customer-webhooks).  The `customer` object in the event notification does not include the following fields: `cards`, `group_ids`, `segment_ids`, and `groups` (deprecated).
type CustomerDeletedWebhook struct {
	// The ID of the seller associated with the event.
	MerchantId string `json:"merchant_id,omitempty"`
	// The type of event. For this object, the value is `customer.deleted`.
	Type_ string `json:"type,omitempty"`
	// The unique ID of the event, which is used for [idempotency support](https://developer.squareup.com/docs/webhooks-api/what-it-does#idempotency-support).
	EventId string `json:"event_id,omitempty"`
	// The timestamp of when the event was created, in RFC 3339 format.
	CreatedAt string                      `json:"created_at,omitempty"`
	Data      *CustomerDeletedWebhookData `json:"data,omitempty"`
}
