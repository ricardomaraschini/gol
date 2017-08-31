package filereader

import "io/ioutil"

// Decoder is an interface to an entity able to decode the file content
type Decoder interface {
	Decode([]byte) error
}

// New returns a new FileReader
func New(opts ...Setter) FileReader {
	fp := FileReader{}
	for _, o := range opts {
		o(&fp)
	}
	return fp
}

// FileReader reads a file and uses Dec to decode it
type FileReader struct {
	Dec      Decoder
	FilePath string
}

// Parse reads the file content and uses Dec to decode it
func (f *FileReader) Parse() ([]byte, error) {
	b, err := ioutil.ReadFile(f.FilePath)
	if err != nil {
		return make([]byte, 0), err
	}

	if f.Dec == nil {
		return b, nil
	}

	return b, f.Dec.Decode(b)
}
