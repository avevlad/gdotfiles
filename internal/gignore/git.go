package gignore

type Git struct {
	url string
}

func NewGit(url string) *Git {
	return &Git{url}
}

func (s Git) CheckLocalExist() {
	// return []JobSpec{j}, nil
}

func (s Git) Clone() {

}

func (s Git) Update() {
}
