package bleve

func (b *Bleve) Index(id string, data interface{}) error {
	return b.index.Index(id, data)
}
