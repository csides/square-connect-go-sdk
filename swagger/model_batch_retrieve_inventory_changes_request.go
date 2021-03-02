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

type BatchRetrieveInventoryChangesRequest struct {
	// The filter to return results by `CatalogObject` ID. The filter is only applicable when set. The default value is null.
	CatalogObjectIds []string `json:"catalog_object_ids,omitempty"`
	// The filter to return results by `Location` ID.  The filter is only applicable when set. The default value is null.
	LocationIds []string `json:"location_ids,omitempty"`
	// The filter to return results by `InventoryChangeType` values other than `TRANSFER`. The default value is `[PHYSICAL_COUNT, ADJUSTMENT]`.
	Types []InventoryChangeType `json:"types,omitempty"`
	// The filter to return results by `InventoryState` values.  Use the  `states` filter as its replacement.
	Statuses []InventoryState `json:"statuses,omitempty"`
	// The filter to return `ADJUSTMENT` query results by `InventoryState`. This filter is only applied when set. The default value is null.
	States []InventoryState `json:"states,omitempty"`
	// The filter to return results with their `calculated_at` value   after the given time as specified in an RFC 3339 timestamp.  The default value is the UNIX epoch of (`1970-01-01T00:00:00Z`).
	UpdatedAfter string `json:"updated_after,omitempty"`
	// The filter to return results with their `created_at` or `calculated_at` value   strictly before the given time as specified in an RFC 3339 timestamp.  The default value is the UNIX epoch of (`1970-01-01T00:00:00Z`).
	UpdatedBefore string `json:"updated_before,omitempty"`
	// A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for the original query.  See the [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination) guide for more information.
	Cursor string `json:"cursor,omitempty"`
}
