package variablelengthquantity

import "fmt"

func EncodeVarint(input []uint32) []byte {
	var out []byte

	for _, n := range input {
		if n == 0 {
			out = append(out, 0)
			continue
		}

		var parts []byte
		for n > 0 {
			part := byte(n & 0x7F)
			parts = append(parts, part)
			n >>= 7
		}

		for i := 1; i < len(parts); i++ {
			parts[i] |= 0x80
		}

		for i := len(parts) - 1; i >= 0; i-- {
			out = append(out, parts[i])
		}
	}
	return out
}

func DecodeVarint(input []byte) ([]uint32, error) {
	var out []uint32
	var result uint32 = 0
	var shift uint = 0

	for i, b := range input {
		result = (result << 7) | uint32(b & 0x7F)

		if b&0x80 == 0 {
			out = append(out, result)
			result = 0
			shift = 0
		} else {
			shift += 7
			if shift > 28 {
				return nil, fmt.Errorf("VLQ overflow at byte index %d", i)
			}
		}
	}

	if shift != 0 {
		return nil, fmt.Errorf("incomplete VLQ sequence")
	}

	return out, nil
}
