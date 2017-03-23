/* 
 * Lyft API
 *
 * Drive your app to success with Lyft's API
 *
 * OpenAPI spec version: 1.0.0
 * Contact: api-support@lyft.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package lyft

type CancellationRequest struct {

	// Token affirming the user accepts the cancellation fee. Required if a cancellation fee is in effect.
	CancelConfirmationToken string `json:"cancel_confirmation_token,omitempty"`
}