package idg

// Parse will parse a string id
func Parse(in string) (id ID, err error) {
	err = id.parse([]byte(in))
	return
}
