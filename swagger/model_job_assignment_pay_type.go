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

// JobAssignmentPayType : Enumerates the possible pay types that a job can be assigned.
type JobAssignmentPayType string

// List of JobAssignmentPayType
const (
	STATUS_DO_NOT_USE_JobAssignmentPayType JobAssignmentPayType = "STATUS_DO_NOT_USE"
	NONE_JobAssignmentPayType              JobAssignmentPayType = "NONE"
	HOURLY_JobAssignmentPayType            JobAssignmentPayType = "HOURLY"
	SALARY_JobAssignmentPayType            JobAssignmentPayType = "SALARY"
)
