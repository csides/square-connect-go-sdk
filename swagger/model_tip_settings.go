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

type TipSettings struct {
	// Indicates whether tipping is enabled for this checkout. Defaults to false.
	AllowTipping bool `json:"allow_tipping,omitempty"`
	// Indicates whether tip options should be presented on their own screen before presenting the signature screen during card payment. Defaults to false.
	SeparateTipScreen bool `json:"separate_tip_screen,omitempty"`
	// Indicates whether custom tip amounts are allowed during the checkout flow. Defaults to false.
	CustomTipField bool `json:"custom_tip_field,omitempty"`
	// A list of tip percentages that should be presented during the checkout flow. Specified as up to 3 non-negative integers from 0 to 100 (inclusive). Defaults to [15, 20, 25]
	TipPercentages []int32 `json:"tip_percentages,omitempty"`
	// Enables the \"Smart Tip Amounts\" behavior. Exact tipping options depend on the region the Square seller is active in.  In the United States and Canada, tipping options will be presented in whole dollar amounts for payments under 10 USD/CAD respectively.  If set to true, the tip_percentages settings is ignored. Defaults to false.  To learn more about smart tipping, see [Accept Tips with the Square App](https://squareup.com/help/us/en/article/5069-accept-tips-with-the-square-app)
	SmartTipping bool `json:"smart_tipping,omitempty"`
}
