package entity

type UpFile struct {
	Filepath string `json:"filepath"`
	Filename string `json:"filename"`
	Metadata string `json:"metadata"`
}
