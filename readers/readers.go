package readers

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(read_to []byte) (int, error) {
	for {
		read_to = append(read_to, "A"...)
	}
}

// Readers exercise main
// reader.Validate(MyReader{})
