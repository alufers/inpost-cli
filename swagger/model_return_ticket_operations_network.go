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

type ReturnTicketOperationsNetwork struct {
	AutoArchivableSince time.Time `json:"autoArchivableSince,omitempty"`
	Delete bool `json:"delete,omitempty"`
	Highlight bool `json:"highlight,omitempty"`
	ManualArchive bool `json:"manualArchive,omitempty"`
	RefreshUntil time.Time `json:"refreshUntil,omitempty"`
	Send bool `json:"send,omitempty"`
}
