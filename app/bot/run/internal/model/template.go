package model

type TStart struct {
	ID       int64
	Username string
	Notice   string
	Chats    []string
	Version  string
}

type TSearchResults struct {
	Keyword string
	Results []*TSearchResult
	Took    int64
}

type TSearchResult struct {
	Seq        int
	SenderName string
	SenderLink string
	Date       string
	Content    string
	Link       string
}
