package Models

type Wish struct {
	id          string
	wishListId  string
	title       string
	description string
	ClaimedBy   string
	link        string
	image       []byte
}
