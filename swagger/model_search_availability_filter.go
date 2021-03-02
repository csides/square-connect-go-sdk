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

// A query filter to search for availabilities by.
type SearchAvailabilityFilter struct {
	StartAtRange *TimeRange `json:"start_at_range"`
	// The query expression to search for availabilities matching the specified seller location IDs. This query expression is not applicable when `booking_id` is present.
	LocationId string `json:"location_id,omitempty"`
	// The list of segment filters to apply. A query with `n` segment filters returns availabilities with `n` segments per availability. It is not applicable when `booking_id` is present.
	SegmentFilters []SegmentFilter `json:"segment_filters,omitempty"`
	// The query expression to search for availabilities for an existing booking by matching the specified `booking_id` value. This is commonly used to reschedule an appointment. If this expression is specified, the `location_id` and `segment_filters` expressions are not allowed.
	BookingId string `json:"booking_id,omitempty"`
}
