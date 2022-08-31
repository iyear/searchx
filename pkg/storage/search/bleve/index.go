package bleve

import (
	"context"
	"github.com/iyear/searchx/pkg/storage/search"
)

func (b *Bleve) Index(ctx context.Context, items []*search.Item) error {
	batch := b.index.NewBatch()

	for _, item := range items {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := batch.Index(item.ID, item.Data); err != nil {
				return err
			}
		}
	}

	return b.index.Batch(batch)
}
