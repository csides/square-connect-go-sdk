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

// Defines the fields included in the response body for requests to __ListCustomerSegments__.  One of `errors` or `segments` is present in a given response (never both).
type ListCustomerSegmentsResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// The list of customer segments belonging to the associated Square account.
	Segments []CustomerSegment `json:"segments,omitempty"`
	// A pagination cursor to be used in subsequent calls to __ListCustomerSegments__ to retrieve the next set of query results. Only present only if the request succeeded and additional results are available.  See the [Pagination guide](https://developer.squareup.com/docs/working-with-apis/pagination) for more information.
	Cursor string `json:"cursor,omitempty"`
}
