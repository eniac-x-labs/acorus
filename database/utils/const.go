package utils

var (
	// The postgres parameter counter for a given query is stored via a  uint16,
	// resulting in a parameter limit of 65535. In order to avoid reaching this limit
	// we'll utilize a batch size of 3k for inserts, well below as long as the the number
	// of columns < 20.
	BatchInsertSize int = 3_000
)
