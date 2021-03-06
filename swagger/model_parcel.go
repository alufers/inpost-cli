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

type Parcel struct {
	CashOnDelivery *CashOnDelivery `json:"cashOnDelivery,omitempty"`
	EndOfWeekCollection bool `json:"endOfWeekCollection,omitempty"`
	ExpiryDate time.Time `json:"expiryDate,omitempty"`
	IsMobileCollectPossible bool `json:"isMobileCollectPossible,omitempty"`
	IsObserved bool `json:"isObserved,omitempty"`
	MultiCompartment *MultiCompartment `json:"multiCompartment,omitempty"`
	OpenCode string `json:"openCode,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	PickupDate time.Time `json:"pickupDate,omitempty"`
	PickupPoint *DeliveryPointData `json:"pickupPoint,omitempty"`
	QrCode string `json:"qrCode,omitempty"`
	ReturnedToSenderDate time.Time `json:"returnedToSenderDate,omitempty"`
	SenderName string `json:"senderName,omitempty"`
	ShipmentNumber string `json:"shipmentNumber,omitempty"`
	ShipmentType string `json:"shipmentType,omitempty"`
	Status string `json:"status,omitempty"`
	StatusHistory []ParcelHistory `json:"statusHistory,omitempty"`
	StoredDate time.Time `json:"storedDate,omitempty"`
}
