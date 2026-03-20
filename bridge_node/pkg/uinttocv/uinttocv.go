package uinttocv

import "encoding/hex"

// UintToCV converts a Go uint64 into a Clarity Value (CV) encoded uint128.
//
// Clarity ABI format for uint:
// - 1 byte  : type tag (0x01 for uint)
// - 16 bytes: unsigned integer (uint128, big-endian)
//
// Total size = 17 bytes
func UintToCV(val uint64) string {
	// Allocate 17 bytes:
	// [0]      → type tag
	// [1..16]  → uint128 value (big-endian)
	b := make([]byte, 17)

	// Set Clarity type tag for uint
	b[0] = 0x01

	// Encode uint64 into the LOWER 8 bytes of uint128 (big-endian)
	// Layout after this loop:
	// b[9..16] = actual value
	// b[1..8]  = zero padding (since uint64 < uint128)
	for i := 0; i < 8; i++ {
		b[16-i] = byte(val >> uint(8*i))
	}

	// NOTE:
	// We intentionally leave the upper 8 bytes (b[1..8]) as zero.
	// This is required because Clarity expects full uint128 width.

	// Return hex-encoded CV with 0x prefix (Stacks RPC format)
	return "0x" + hex.EncodeToString(b)
}
