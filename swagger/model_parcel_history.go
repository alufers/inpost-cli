/*
 * Inpost Mobile
 *
 * Extracted from retrofit
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type ParcelHistory struct {
	Date time.Time `json:"date,omitempty"`
	Status string `json:"status,omitempty"`
}
