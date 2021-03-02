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

// OrderServiceChargeCalculationPhase : Represents a phase in the process of calculating order totals. Service charges are applied __after__ the indicated phase.  [Read more about how order totals are calculated.](https://developer.squareup.com/docs/orders-api/how-it-works#how-totals-are-calculated)
type OrderServiceChargeCalculationPhase string

// List of OrderServiceChargeCalculationPhase
const (
	SERVICE_CHARGE_CALCULATION_PHASE_DO_NOT_USE_OrderServiceChargeCalculationPhase OrderServiceChargeCalculationPhase = "SERVICE_CHARGE_CALCULATION_PHASE_DO_NOT_USE"
	SUBTOTAL_PHASE_OrderServiceChargeCalculationPhase                              OrderServiceChargeCalculationPhase = "SUBTOTAL_PHASE"
	TOTAL_PHASE_OrderServiceChargeCalculationPhase                                 OrderServiceChargeCalculationPhase = "TOTAL_PHASE"
)
