package fiftyoneDegrees

import (
	"strconv"
	"strings"

	"github.com/bitly/go-simplejson"
)

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

func ParseDeviceInfo(data []byte) (*DeviceInfo, error) {
	json, err := simplejson.NewJson(data)
	if err != nil {
		return nil, err
	}

	deviceInfo := &DeviceInfo{}
	if v, err := json.Get("DeviceType").String(); err == nil {
		deviceInfo.DeviceType = parseString(v)
	}

	if v, err := json.Get("HardwareName").String(); err == nil {
		deviceInfo.HardwareName = parseString(v)
	}

	if v, err := json.Get("HardwareVendor").String(); err == nil {
		deviceInfo.HardwareVendor = parseString(v)
	}

	if v, err := json.Get("HardwareModel").String(); err == nil {
		deviceInfo.HardwareModel = parseString(v)
	}

	if v, err := json.Get("BrowserName").String(); err == nil {
		deviceInfo.BrowserName = parseString(v)
	}

	if v, err := json.Get("BrowserVersion").String(); err == nil {
		deviceInfo.BrowserVersion = parseString(v)
	}

	if v, err := json.Get("BrowserVendor").String(); err == nil {
		deviceInfo.BrowserVendor = parseString(v)
	}

	if v, err := json.Get("PlatformName").String(); err == nil {
		deviceInfo.PlatformName = parseString(v)
	}

	if v, err := json.Get("PlatformVersion").String(); err == nil {
		deviceInfo.PlatformVersion = parseString(v)
	}

	if v, err := json.Get("PlatformVendor").String(); err == nil {
		deviceInfo.PlatformVendor = parseString(v)
	}

	if v, err := json.Get("ScreenPixelsWidth").String(); err == nil {
		deviceInfo.ScreenPixelsWidth = parseUint16(v)
	}

	if v, err := json.Get("ScreenPixelsHeight").String(); err == nil {
		deviceInfo.ScreenPixelsHeight = parseUint16(v)
	}

	if v, err := json.Get("CPU").String(); err == nil {
		deviceInfo.CPU = parseString(v)
	}

	if v, err := json.Get("CPUCores").String(); err == nil {
		deviceInfo.CPUCores = parseUint8(v)
	}

	if v, err := json.Get("DeviceRAM").String(); err == nil {
		deviceInfo.DeviceRAM = parseUint16(v)
	}

	if v, err := json.Get("IsCrawler").String(); err == nil {
		deviceInfo.IsCrawler = parseBool(v)
	}

	if v, err := json.Get("PriceBand").String(); err == nil {
		deviceInfo.PriceBand = parsePriceBand(v)
	}

	return deviceInfo, nil
}

func parseString(v string) *string {
	if v == "Unknown" {
		return nil
	}

	return &v
}

func parseBool(v string) *bool {
	if v == "Unknown" {
		return nil
	}

	ret := false
	if v == "True" {
		ret = true
	}
	return &ret
}

func parseUint16(v string) *uint16 {
	if v == "Unknown" {
		return nil
	}

	if i, err := strconv.ParseUint(v, 10, 64); err == nil {
		ret := uint16(i)
		return &ret
	}
	return nil
}

func parseUint8(v string) *uint8 {
	if v == "Unknown" {
		return nil
	}

	if i, err := strconv.ParseUint(v, 10, 64); err == nil {
		ret := uint8(i)
		return &ret
	}
	return nil
}

func parsePriceBand(v string) *ModelPriceBand {
	if v == "Unknown" {
		return nil
	}

	sp := strings.Split(v, " - ")
	if len(sp) != 2 {
		return nil
	}
	minF, _ := strconv.ParseFloat(sp[0], 64)
	maxF, _ := strconv.ParseFloat(sp[1], 64)
	if minF <= 0 || maxF <= 0 {
		return nil
	}
	min := uint(minF * 100)
	max := uint(maxF * 100)
	avg := (min + max) / 2

	return &ModelPriceBand{
		Min: min,
		Max: max,
		Avg: avg,
	}
}
