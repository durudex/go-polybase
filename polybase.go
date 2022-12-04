package polybase

const URL string = "https://testnet.polybase.xyz/v0"

type Polybase interface {
	Collection(name string) Collection
}

type Record[T any] struct {
	Block Block `json:"block"`
	Data  T     `json:"data"`
}

type Block struct {
	Hash string `json:"hash"`
}

type Cursor struct {
	After  string `json:"after"`
	Before string `json:"before"`
}

func New() Polybase { return &polybase{client: NewClient()} }

type polybase struct{ client Client }

func (p *polybase) Collection(name string) Collection {
	return NewCollection(name, p.client)
}
