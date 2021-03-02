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

// An object representing a team member's wage information.
type WageSetting struct {
	// The unique ID of the `TeamMember` whom this wage setting describes.
	TeamMemberId string `json:"team_member_id,omitempty"`
	// <b>Required</b> The ordered list of jobs that the team member is assigned to. The first job assignment is considered the team member's \"Primary Job\". <br> <b>Min Length 1    Max Length 12</b>
	JobAssignments []JobAssignment `json:"job_assignments,omitempty"`
	// Whether the team member is exempt from the overtime rules of the seller country.
	IsOvertimeExempt bool `json:"is_overtime_exempt,omitempty"`
	// Used for resolving concurrency issues; request will fail if version provided does not match server version at time of request. If not provided, Square executes a blind write, potentially overwriting data from another write. Read about [optimistic concurrency](https://developer.squareup.com/docs/working-with-apis/optimistic-concurrency) in Square APIs for more information.
	Version int32 `json:"version,omitempty"`
	// The timestamp in RFC 3339 format describing when the wage setting object was created. Ex: \"2018-10-04T04:00:00-07:00\" or \"2019-02-05T12:00:00Z\"
	CreatedAt string `json:"created_at,omitempty"`
	// The timestamp in RFC 3339 format describing when the wage setting object was last updated. Ex: \"2018-10-04T04:00:00-07:00\" or \"2019-02-05T12:00:00Z\"
	UpdatedAt string `json:"updated_at,omitempty"`
}
