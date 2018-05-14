package main

import (
	"fmt"
)

type MetadataPregenStatus uint8
const (
	// MetadataPregenPending indicates that pregen has not been completed
	MetadataPregenPending MetadataPregenStatus = iota
	// MetadataPregenComplete indicates that pregen has completed successfully
	MetadataPregenComplete MetadataPregenStatus = 2
	// MetadataPregenFailure indicates that pregen failed
	MetadataPregenFailure MetadataPregenStatus = 4
	// MetadataPregenInProgress indicates that pregen is currently in progress
	MetadataPregenInProgress MetadataPregenStatus = 6
)

const (
	RecMetadataPregenBitMask   = 0x6
)


func main() {
	attr := uint32(MetadataPregenComplete)
	fmt.Println(attr)
	fmt.Println(attr&RecMetadataPregenBitMask)
}
