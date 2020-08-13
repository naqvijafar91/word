package word

import (
	"io/ioutil"
)

type FileInjestor interface {
	Injest(absPathToFile string) ([]byte, error)
}

type fileInjestor struct {
}

func (fi *fileInjestor) Injest(absPathToFile string) ([]byte, error) {
	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile(absPathToFile)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func NewFileInjestor() FileInjestor {
	return &fileInjestor{}
}
