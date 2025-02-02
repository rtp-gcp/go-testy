package bitlib

func BitAnd(value1 uint8, value2 uint8) uint8 {
	return value1 & value2
}

func BitClear(value1 uint8, value2 uint8) uint8 {
	return value1 &^ value2
}
