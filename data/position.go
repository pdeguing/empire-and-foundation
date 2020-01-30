package data

// GetPositionCode converts the individual components of a position to a single number.
func GetPositionCode(region, system, orbit, suborbit int) int {
	return suborbit + orbit << 4 + system << 8 + region << 16
}

// ParsePositionCode extracts the individual components from the position code.
func ParsePositionCode(c int) (region, system, orbit, suborbit int) {
	region = c >> 16
	system = (c >> 8) & 0xFF
	orbit = (c >> 4) & 0xF
	suborbit = c & 0xF
	return
}
