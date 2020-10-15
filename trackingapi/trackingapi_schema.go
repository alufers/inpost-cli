package trackingapi

import "time"

type TrackingAPISchema struct {
	TrackingNumber   string            `json:"tracking_number"`
	Service          string            `json:"service"`
	Type             string            `json:"type"`
	Status           string            `json:"status"`
	CustomAttributes CustomAttributes  `json:"custom_attributes"`
	TrackingDetails  []TrackingDetails `json:"tracking_details"`
	ExpectedFlow     []interface{}     `json:"expected_flow"`
	CreatedAt        time.Time         `json:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at"`
}
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
type Address struct {
	Line1 string `json:"line1"`
	Line2 string `json:"line2"`
}
type TargetMachineDetail struct {
	Name                string   `json:"name"`
	OpeningHours        string   `json:"opening_hours"`
	LocationDescription string   `json:"location_description"`
	Location            Location `json:"location"`
	Address             Address  `json:"address"`
	Type                []string `json:"type"`
	Location247         bool     `json:"location247"`
}
type CustomAttributes struct {
	Size                string              `json:"size"`
	TargetMachineID     string              `json:"target_machine_id"`
	TargetMachineDetail TargetMachineDetail `json:"target_machine_detail"`
	EndOfWeekCollection bool                `json:"end_of_week_collection"`
}
type TrackingDetails struct {
	Status       string      `json:"status"`
	OriginStatus string      `json:"origin_status"`
	Agency       interface{} `json:"agency"`
	Datetime     time.Time   `json:"datetime"`
}