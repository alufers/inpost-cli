/*
 * Inpost Mobile
 *
 * Extracted from retrofit
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ValidationRequest struct {
	GeoPoint *GeoPoint `json:"geoPoint,omitempty"`
	Parcel *ParcelCompartment `json:"parcel,omitempty"`
}
