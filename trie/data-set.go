package trie

/*
#cgo CFLAGS: -I . -Wimplicit-function-declaration
#cgo darwin LDFLAGS: -lm
#cgo linux LDFLAGS: -lm -lrt
#include "51Degrees.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"strings"
	"unsafe"

	fd "github.com/0i/go-51Degrees"
)

type DataSet struct {
	cDataSet *C.fiftyoneDegreesDataSet
}

func NewDataSet(fileName string, properties []string) (*DataSet, error) {
	dataSet := &DataSet{
		cDataSet: new(C.fiftyoneDegreesDataSet),
	}

	var cFileName *C.char = C.CString(fileName)
	defer C.free(unsafe.Pointer(cFileName))

	var cProperties *C.char = C.CString(strings.Join(properties, ","))
	defer C.free(unsafe.Pointer(cProperties))

	status := C.fiftyoneDegreesInitWithPropertyString(
		cFileName,
		dataSet.cDataSet,
		cProperties,
	)

	if status == 0 {
		return dataSet, nil
	}

	switch status {
	case C.DATA_SET_INIT_STATUS_INSUFFICIENT_MEMORY:
		return nil, fmt.Errorf("Insufficient memory to load [%s].", fileName)
	case C.DATA_SET_INIT_STATUS_CORRUPT_DATA:
		return nil, fmt.Errorf("Device data file [%s] is corrupted.", fileName)
	case C.DATA_SET_INIT_STATUS_INCORRECT_VERSION:
		return nil, fmt.Errorf("Device data file [%s] is not correct version.", fileName)
	case C.DATA_SET_INIT_STATUS_FILE_NOT_FOUND:
		return nil, fmt.Errorf("Device data file [%s] not found.", fileName)
	}

	return nil, fmt.Errorf("Unknown error code: %d", status)
}

func (d *DataSet) Detection(userAgent string) (*fd.DeviceInfo, error) {
	var cUserAgent *C.char = C.CString(userAgent)
	defer C.free(unsafe.Pointer(cUserAgent))

	deviceOffsets := C.fiftyoneDegreesGetDeviceOffset(d.cDataSet, cUserAgent)

	resultLength := 50000
	buff := make([]byte, resultLength)
	length := C.fiftyoneDegreesProcessDeviceJSON(
		d.cDataSet,
		deviceOffsets,
		(*C.char)(unsafe.Pointer(&buff[0])),
		C.int(resultLength),
	)

	return fd.ParseDeviceInfo(buff[:length])
}
