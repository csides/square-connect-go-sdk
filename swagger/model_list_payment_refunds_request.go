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

// Retrieves a list of refunds for the account making the request.  The maximum results per page is 100.
type ListPaymentRefundsRequest struct {
	// The timestamp for the beginning of the requested reporting period, in RFC 3339 format.  Default: The current time minus one year.
	BeginTime string `json:"begin_time,omitempty"`
	// The timestamp for the end of the requested reporting period, in RFC 3339 format.  Default: The current time.
	EndTime string `json:"end_time,omitempty"`
	// The order in which results are listed: - `ASC` - Oldest to newest. - `DESC` - Newest to oldest (default).
	SortOrder string `json:"sort_order,omitempty"`
	// A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for the original query.  For more information, see [Pagination](https://developer.squareup.com/docs/basics/api101/pagination).
	Cursor string `json:"cursor,omitempty"`
	// Limit results to the location supplied. By default, results are returned for all locations associated with the seller.
	LocationId string `json:"location_id,omitempty"`
	// If provided, only refunds with the given status are returned. For a list of refund status values, see [PaymentRefund](#type-paymentrefund).  Default: If omitted, refunds are returned regardless of their status.
	Status string `json:"status,omitempty"`
	// If provided, only refunds with the given source type are returned. - `CARD` - List refunds only for payments where `CARD` was specified as the payment source.  Default: If omitted, refunds are returned regardless of the source type.
	SourceType string `json:"source_type,omitempty"`
	// The maximum number of results to be returned in a single page.  It is possible to receive fewer results than the specified limit on a given page.  If the supplied value is greater than 100, no more than 100 results are returned.  Default: 100
	Limit int32 `json:"limit,omitempty"`
}
