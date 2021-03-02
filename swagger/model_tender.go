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

// Represents a tender (i.e., a method of payment) used in a Square transaction.
type Tender struct {
	// The tender's unique ID.
	Id string `json:"id,omitempty"`
	// The ID of the transaction's associated location.
	LocationId string `json:"location_id,omitempty"`
	// The ID of the tender's associated transaction.
	TransactionId string `json:"transaction_id,omitempty"`
	// The timestamp for when the tender was created, in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`
	// An optional note associated with the tender at the time of payment.
	Note               string `json:"note,omitempty"`
	AmountMoney        *Money `json:"amount_money,omitempty"`
	TipMoney           *Money `json:"tip_money,omitempty"`
	ProcessingFeeMoney *Money `json:"processing_fee_money,omitempty"`
	// If the tender is associated with a customer or represents a customer's card on file, this is the ID of the associated customer.
	CustomerId  string             `json:"customer_id,omitempty"`
	Type_       *TenderType        `json:"type"`
	CardDetails *TenderCardDetails `json:"card_details,omitempty"`
	CashDetails *TenderCashDetails `json:"cash_details,omitempty"`
	// Additional recipients (other than the merchant) receiving a portion of this tender. For example, fees assessed on the purchase by a third party integration.
	AdditionalRecipients []AdditionalRecipient `json:"additional_recipients,omitempty"`
	// The ID of the [Payment](#type-payment) that corresponds to this tender. This value is only present for payments created with the v2 Payments API.
	PaymentId string `json:"payment_id,omitempty"`
}
