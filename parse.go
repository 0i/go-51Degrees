package fiftyoneDegrees

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

func ParseDeviceInfo(data string) (*DeviceInfo, error) {
	// data = strings.Replace(data, "\n", ` `, -1)
	// data = strings.Replace(data, "\t", ` `, -1)
	// data = strings.Replace(data, "\r", ` `, -1)
	// data = strings.Replace(data, `\n`, `\\n`, -1)
	// data = strings.Replace(data, `\t`, `\\t`, -1)
	// data = strings.Replace(data, `\r`, `\\r`, -1)
	// data = strings.Replace(data, `"Kindle Fire HD 7"",`, `"Kindle Fire HD 7",`, -1)
	// data = strings.Replace(data, `"Kindle Fire HD 7" (`, `"Kindle Fire HD 7 (`, -1)
	// data = strings.Replace(data, `"PIXI 4 5"",`, `"PIXI 4 5",`, -1)
	// data = strings.Replace(data, `"Pixi 4 5"",`, `"Pixi 4 5",`, -1)
	// data = strings.Replace(data, `"eSmart 7"",`, `"eSmart 7",`, -1)
	// data = strings.Replace(data, `"Galaxy Note Pro 12.2"",`, `"Galaxy Note Pro 12.2",`, -1)
	// data = strings.Replace(data, `"Miia Tab 7"",`, `"Miia Tab 7",`, -1)
	// data = strings.Replace(data, `"MyTablet 7"",`, `"MyTablet 7",`, -1)

	res := gjson.Parse(data)
	if res.Get("Id").String() == "" {
		return nil, fmt.Errorf("Can not parse json")
	}

	deviceInfo := &DeviceInfo{}

	if v := res.Get("DeviceType").String(); v != "" {
		deviceInfo.DeviceType = parseString(v)
	}

	if v := res.Get("HardwareName").String(); v != "" {
		deviceInfo.HardwareName = parseString(v)
	}

	if v := res.Get("HardwareVendor").String(); v != "" {
		deviceInfo.HardwareVendor = parseString(v)
	}

	if v := res.Get("HardwareModel").String(); v != "" {
		deviceInfo.HardwareModel = parseString(v)
	}

	if v := res.Get("BrowserName").String(); v != "" {
		deviceInfo.BrowserName = parseString(v)
	}

	if v := res.Get("BrowserVersion").String(); v != "" {
		deviceInfo.BrowserVersion = parseString(v)
	}

	if v := res.Get("BrowserVendor").String(); v != "" {
		deviceInfo.BrowserVendor = parseString(v)
	}

	if v := res.Get("PlatformName").String(); v != "" {
		deviceInfo.PlatformName = parseString(v)
	}

	if v := res.Get("PlatformVersion").String(); v != "" {
		deviceInfo.PlatformVersion = parseString(v)
	}

	if v := res.Get("PlatformVendor").String(); v != "" {
		deviceInfo.PlatformVendor = parseString(v)
	}

	if v := res.Get("ScreenPixelsWidth").String(); v != "" {
		deviceInfo.ScreenPixelsWidth = parseUint16(v)
	}

	if v := res.Get("ScreenPixelsHeight").String(); v != "" {
		deviceInfo.ScreenPixelsHeight = parseUint16(v)
	}

	if v := res.Get("CPU").String(); v != "" {
		deviceInfo.CPU = parseString(v)
	}

	if v := res.Get("CPUCores").String(); v != "" {
		deviceInfo.CPUCores = parseUint8(v)
	}

	if v := res.Get("DeviceRAM").String(); v != "" {
		deviceInfo.DeviceRAM = parseUint16(v)
	}

	if v := res.Get("IsCrawler").String(); v != "" {
		deviceInfo.IsCrawler = parseBool(v)
	}

	if v := res.Get("PriceBand").String(); v != "" {
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
