package structs

type StringMsg struct {
	p *string
}

func (msg *StringMsg) Inject(message string) {
	msg.p = &message
}

func (msg *StringMsg) Extract() []byte {
	return []byte(*msg.p)
}
