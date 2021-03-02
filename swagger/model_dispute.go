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

// Represents a dispute a cardholder initiated with their bank.
type Dispute struct {
	// The unique ID for this `Dispute`, generated by Square.
	DisputeId   string         `json:"dispute_id,omitempty"`
	AmountMoney *Money         `json:"amount_money,omitempty"`
	Reason      *DisputeReason `json:"reason,omitempty"`
	State       *DisputeState  `json:"state,omitempty"`
	// The time when the next action is due, in RFC 3339 format.
	DueAt           string           `json:"due_at,omitempty"`
	DisputedPayment *DisputedPayment `json:"disputed_payment,omitempty"`
	// The IDs of the evidence associated with the dispute.
	EvidenceIds []string   `json:"evidence_ids,omitempty"`
	CardBrand   *CardBrand `json:"card_brand,omitempty"`
	// The timestamp when the dispute was created, in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`
	// The timestamp when the dispute was last updated, in RFC 3339 format.
	UpdatedAt string `json:"updated_at,omitempty"`
	// The ID of the dispute in the card brand system, generated by the card brand.
	BrandDisputeId string `json:"brand_dispute_id,omitempty"`
	// The timestamp when the dispute was reported, in RFC 3339 format.
	ReportedDate string `json:"reported_date,omitempty"`
	// The current version of the `Dispute`.
	Version int32 `json:"version,omitempty"`
	// The ID of the location where the dispute originated.
	LocationId string `json:"location_id,omitempty"`
}
