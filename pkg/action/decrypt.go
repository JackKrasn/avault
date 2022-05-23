package action

type Decrypt struct {
	Password string
}

func NewDecrypt() *Decrypt {
	return &Decrypt{}
}
