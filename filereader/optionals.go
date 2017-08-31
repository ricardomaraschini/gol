package filereader

// Setter is a function that sets a property on a FileReader
type Setter func(*FileReader)

// WithDecoder sets the decoder for the FileReader
func WithDecoder(d Decoder) Setter {
	return func(f *FileReader) {
		f.Dec = d
	}
}

// WithFilePath sets the file path for the FileReader
func WithFilePath(fp string) Setter {
	return func(f *FileReader) {
		f.FilePath = fp
	}
}
