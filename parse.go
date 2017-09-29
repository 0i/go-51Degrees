package fiftyoneDegrees

import (
	"strconv"
	"strings"

	"github.com/bitly/go-simplejson"
)

func ParseDeviceInfo(data string) (*DeviceInfo, error) {
	// data = strings.Replace(data, "\n", ` `, -1)
	// data = strings.Replace(data, "\t", ` `, -1)
	// data = strings.Replace(data, "\r", ` `, -1)
	// data = strings.Replace(data, `"Kindle Fire HD 7"",`, `"Kindle Fire HD 7",`, -1)
	// data = strings.Replace(data, `"Kindle Fire HD 7" (`, `"Kindle Fire HD 7 (`, -1)
	// data = strings.Replace(data, `"PIXI 4 5"",`, `"PIXI 4 5",`, -1)
	// data = strings.Replace(data, `"Pixi 4 5"",`, `"Pixi 4 5",`, -1)
	// data = strings.Replace(data, `"eSmart 7"",`, `"eSmart 7",`, -1)
	// data = strings.Replace(data, `"Galaxy Note Pro 12.2"",`, `"Galaxy Note Pro 12.2",`, -1)
	// data = strings.Replace(data, `"Miia Tab 7"",`, `"Miia Tab 7",`, -1)
	// data = strings.Replace(data, `"MyTablet 7"",`, `"MyTablet 7",`, -1)

	json, err := simplejson.NewJson([]byte(data))
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
