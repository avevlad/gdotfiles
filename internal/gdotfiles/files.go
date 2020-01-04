package gdotfiles

type Files struct {
}

func NewFiles() *Files {
	return &Files{}
}

func (s Files) List() []string {
	return []string{}
}
