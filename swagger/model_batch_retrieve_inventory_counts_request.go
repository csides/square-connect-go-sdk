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

type BatchRetrieveInventoryCountsRequest struct {
	// The filter to return results by `CatalogObject` ID. The filter is applicable only when set.  The default is null.
	CatalogObjectIds []string `json:"catalog_object_ids,omitempty"`
	// The filter to return results by `Location` ID.  This filter is applicable only when set. The default is null.
	LocationIds []string `json:"location_ids,omitempty"`
	// The filter to return results with their `calculated_at` value  after the given time as specified in an RFC 3339 timestamp.  The default value is the UNIX epoch of (`1970-01-01T00:00:00Z`).
	UpdatedAfter string `json:"updated_after,omitempty"`
	// A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for the original query.  See the [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination) guide for more information.
	Cursor string `json:"cursor,omitempty"`
	// The filter to return results by `InventoryState`. The filter is only applicable when set. Ignored are untracked states of `NONE`, `SOLD`, and `UNLINKED_RETURN`. The default is null.
	States []InventoryState `json:"states,omitempty"`
}
