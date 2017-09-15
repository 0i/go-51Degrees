package fiftyoneDegrees

type DeviceInfo struct {
	DeviceType         *string         `json:"device_type,omitempty"`
	HardwareName       *string         `json:"hardware_name,omitempty"`
	HardwareVendor     *string         `json:"hardware_vendor,omitempty"`
	HardwareModel      *string         `json:"hardware_model,omitempty"`
	BrowserName        *string         `json:"browser_name,omitempty"`
	BrowserVersion     *string         `json:"browser_version,omitempty"`
	BrowserVendor      *string         `json:"browser_vendor,omitempty"`
	PlatformName       *string         `json:"platform_name,omitempty"`
	PlatformVersion    *string         `json:"platform_version,omitempty"`
	PlatformVendor     *string         `json:"platform_vendor,omitempty"`
	ScreenPixelsWidth  *uint16         `json:"screen_pixels_width,omitempty"`
	ScreenPixelsHeight *uint16         `json:"screen_pixels_height,omitempty"`
	CPU                *string         `json:"cpu,omitempty"`
	CPUCores           *uint8          `json:"cpu_cores,omitempty"`
	DeviceRAM          *uint16         `json:"device_ram,omitempty"`
	IsCrawler          *bool           `json:"is_crawler,omitempty"`
	PriceBand          *ModelPriceBand `json:"price_band,omitempty"`
}

type ModelPriceBand struct {
	Min uint `json:"min,omitempty"`
	Max uint `json:"max,omitempty"`
	Avg uint `json:"avg,omitempty"`
}
