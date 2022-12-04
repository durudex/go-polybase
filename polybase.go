package polybase

const TestnetURL string = "https://testnet.polybase.xyz/v0"

type Polybase interface {
	Collection(name string) Collection
}

type Config struct {
	URL              string
	DefaultNamespace string
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

func New(cfg Config) Polybase {
	return &polybase{client: NewClient(cfg.URL), cfg: cfg}
}

type polybase struct {
	client Client
	cfg    Config
}

func (p *polybase) Collection(name string) Collection {
	if p.cfg.DefaultNamespace != "" {
		name = p.cfg.DefaultNamespace + "/" + name
	}

	return NewCollection(name, p.client)
}
