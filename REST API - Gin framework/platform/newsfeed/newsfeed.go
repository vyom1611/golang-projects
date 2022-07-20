package newsfeed

type Getter interface {
	//Getting a slice of all items
	GetAll() []Item
}

type Adder interface {
	//Taking in one item to be added to Repo
	Add(item Item)
}

//JSON schema
type Item struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

type Repo struct {
	Items []Item
}

func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}
func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

func (r *Repo) GetAll() []Item {
	return r.Items
}
