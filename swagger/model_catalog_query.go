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

// A query composed of one or more different types of filters to narrow the scope of targeted objects when calling the `SearchCatalogObjects` endpoint.  Although a query can have multiple filters, only certain query types can be combined per call to [SearchCatalogObjects](#endpoint-Catalog-SearchCatalogObjects). Any combination of the following types may be used together: - [exact_query](#type-CatalogExactQuery) - [prefix_query](#type-CatalogPrefixQuery) - [range_query](#type-CatalogRangeQuery) - [sorted_attribute_query](#type-CatalogSortedAttribute) - [text_query](#type-CatalogTextQuery) All other query types cannot be combined with any others.  When a query filter is based on an attribute, the attribute must be searchable. Searchable attributes are listed as follows, along their parent types that can be searched for with applicable query filters.  * Searchable attribute and objects queryable by searchable attributes ** - `name`:  `CatalogItem`, `CatalogItemVariation`, `CatalogCategory`, `CatalogTax`, `CatalogDiscount`, `CatalogModifier`, 'CatalogModifierList`, `CatalogItemOption`, `CatalogItemOptionValue` - `description`: `CatalogItem`, `CatalogItemOptionValue` - `abbreviation`: `CatalogItem` - `upc`: `CatalogItemVariation` - `sku`: `CatalogItemVariation` - `caption`: `CatalogImage` - `display_name`: `CatalogItemOption`  For example, to search for [CatalogItem](#type-CatalogItem) objects by searchable attributes, you can use the `\"name\"`, `\"description\"`, or `\"abbreviation\"` attribute in an applicable query filter.
type CatalogQuery struct {
	SortedAttributeQuery                   *CatalogQuerySortedAttribute                   `json:"sorted_attribute_query,omitempty"`
	ExactQuery                             *CatalogQueryExact                             `json:"exact_query,omitempty"`
	SetQuery                               *CatalogQuerySet                               `json:"set_query,omitempty"`
	PrefixQuery                            *CatalogQueryPrefix                            `json:"prefix_query,omitempty"`
	RangeQuery                             *CatalogQueryRange                             `json:"range_query,omitempty"`
	TextQuery                              *CatalogQueryText                              `json:"text_query,omitempty"`
	ItemsForTaxQuery                       *CatalogQueryItemsForTax                       `json:"items_for_tax_query,omitempty"`
	ItemsForModifierListQuery              *CatalogQueryItemsForModifierList              `json:"items_for_modifier_list_query,omitempty"`
	ItemsForItemOptionsQuery               *CatalogQueryItemsForItemOptions               `json:"items_for_item_options_query,omitempty"`
	ItemVariationsForItemOptionValuesQuery *CatalogQueryItemVariationsForItemOptionValues `json:"item_variations_for_item_option_values_query,omitempty"`
}
