package colors

import "strconv"
пше
type Hex string
type RGB struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

func (h Hex) ToRGB() (RGB, error) {
	return Hex2RGB(h)
}

func Hex2RGB(hex Hex) (RGB, error) {
	var rgb RGB
	values, err := strconv.ParseUint(string(hex), 16, 32)
	if err != nil {
		return RGB{}, err
	}

	rgb = RGB{
		Red:   uint8(values >> 16),
		Green: uint8((values >> 8) & 0xff),
		Blue:  uint8(values & 0xff),
	}

	return rgb, nil
}
