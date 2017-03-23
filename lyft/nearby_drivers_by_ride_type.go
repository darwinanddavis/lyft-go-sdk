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

type NearbyDriversByRideType struct {

	// driver's ride type. if driver is eligable for several ride types, he will be duplicated.
	RideType string `json:"ride_type,omitempty"`

	// list of nearby drivers group by ride type sorted by eta
	Drivers []NearbyDriver `json:"drivers,omitempty"`
}