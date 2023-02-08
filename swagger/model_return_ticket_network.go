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

type ReturnTicketNetwork struct {
	AcceptedDate time.Time `json:"acceptedDate,omitempty"`
	Address string `json:"address,omitempty"`
	CreatedDate time.Time `json:"createdDate,omitempty"`
	DeliveredDate time.Time `json:"deliveredDate,omitempty"`
	Description string `json:"description,omitempty"`
	EventLog []ReturnTicketEventLogNetwork `json:"eventLog,omitempty"`
	ExpiryDate time.Time `json:"expiryDate,omitempty"`
	FormType string `json:"formType,omitempty"`
	Operations *ReturnTicketOperationsNetwork `json:"operations,omitempty"`
	OrderNumber string `json:"orderNumber,omitempty"`
	OrganizationName string `json:"organizationName,omitempty"`
	ParcelSize string `json:"parcelSize,omitempty"`
	QrCode string `json:"qrCode,omitempty"`
	ReturnCode string `json:"returnCode,omitempty"`
	ReturnReason string `json:"returnReason,omitempty"`
	Rma string `json:"rma,omitempty"`
	SentDate time.Time `json:"sentDate,omitempty"`
	ShipmentNumber string `json:"shipmentNumber,omitempty"`
	Status string `json:"status,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}
