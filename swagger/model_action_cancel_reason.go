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

type ActionCancelReason string

// List of ActionCancelReason
const (
	TERMINAL_CHECKOUT_CANCEL_REASON_DO_NOT_USE_ActionCancelReason ActionCancelReason = "TERMINAL_CHECKOUT_CANCEL_REASON_DO_NOT_USE"
	BUYER_CANCELED_ActionCancelReason                             ActionCancelReason = "BUYER_CANCELED"
	SELLER_CANCELED_ActionCancelReason                            ActionCancelReason = "SELLER_CANCELED"
	TIMED_OUT_ActionCancelReason                                  ActionCancelReason = "TIMED_OUT"
)
