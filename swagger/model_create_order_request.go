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

type CreateOrderRequest struct {
	Order *Order `json:"order,omitempty"`
	// A value you specify that uniquely identifies this order among orders you've created.  If you're unsure whether a particular order was created successfully, you can reattempt it with the same idempotency key without worrying about creating duplicate orders.  See [Idempotency](https://developer.squareup.com/docs/basics/api101/idempotency) for more information.
	IdempotencyKey string `json:"idempotency_key,omitempty"`
}
