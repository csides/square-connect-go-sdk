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

type TerminalCheckout struct {
	// A unique ID for this `TerminalCheckout`
	Id          string `json:"id,omitempty"`
	AmountMoney *Money `json:"amount_money"`
	// An optional user-defined reference ID which can be used to associate this `TerminalCheckout` to another entity in an external system. For example, an order ID generated by a third-party shopping cart. Will also be associated with any payments used to complete the checkout.
	ReferenceId string `json:"reference_id,omitempty"`
	// An optional note to associate with the checkout, as well any payments used to complete the checkout.
	Note          string                 `json:"note,omitempty"`
	DeviceOptions *DeviceCheckoutOptions `json:"device_options"`
	// The duration as an RFC 3339 duration, after which the checkout will be automatically canceled. TerminalCheckouts that are `PENDING` will be automatically `CANCELED` and have a cancellation reason of `TIMED_OUT`.  Default: 5 minutes from creation  Maximum: 5 minutes
	DeadlineDuration string `json:"deadline_duration,omitempty"`
	// The status of the `TerminalCheckout`. Options: `PENDING`, `IN_PROGRESS`, `CANCEL_REQUESTED`, `CANCELED`, `COMPLETED`
	Status       string              `json:"status,omitempty"`
	CancelReason *ActionCancelReason `json:"cancel_reason,omitempty"`
	// A list of ids for payments created by this `TerminalCheckout`.
	PaymentIds []string `json:"payment_ids,omitempty"`
	// The time when the `TerminalCheckout` was created as an RFC 3339 timestamp.
	CreatedAt string `json:"created_at,omitempty"`
	// The time when the `TerminalCheckout` was last updated as an RFC 3339 timestamp.
	UpdatedAt string `json:"updated_at,omitempty"`
}
