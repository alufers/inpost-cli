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

type ReturnTicketEventLogNetwork struct {
	Date time.Time `json:"date,omitempty"`
	Name string `json:"name,omitempty"`
	Type_ string `json:"type,omitempty"`
}
