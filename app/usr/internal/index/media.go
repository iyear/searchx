package index

import "github.com/gotd/td/tg"

func MediaText(media tg.MessageMediaClass) []string {
	switch m := media.(type) {
	case *tg.MessageMediaDocument:
		return mediaDocText(m)
	case *tg.MessageMediaWebPage:
		return mediaWebPageText(m)
	case *tg.MessageMediaVenue:
		return mediaVenueText(m)
	// case *tg.MessageMediaGame: // TODO(iyear): index media game?
	case *tg.MessageMediaInvoice:
		return mediaInvoiceText(m)
	}
	return nil
}

func mediaDocText(document *tg.MessageMediaDocument) []string {
	doc, ok := document.Document.(*tg.Document)
	if !ok {
		return nil
	}

	attrs := make([]string, 0)
	text := ""

	for _, attr := range doc.Attributes {
		switch a := attr.(type) {
		case *tg.DocumentAttributeAudio:
			text = a.Title
		case *tg.DocumentAttributeFilename:
			text = a.FileName
		case *tg.DocumentAttributeSticker:
			// sticker will have filename with "sticker.webp" and usr don't index sticker
			return nil
		}
		attrs = append(attrs, text)
	}

	return attrs
}

func mediaWebPageText(webPage *tg.MessageMediaWebPage) []string {
	switch wp := webPage.Webpage.(type) {
	case *tg.WebPage:
		// TODO(iyear): index further content
		return []string{wp.SiteName, wp.Title, wp.Description}
	}
	return nil
}

func mediaVenueText(venue *tg.MessageMediaVenue) []string {
	return []string{venue.Title, venue.Address}
}

func mediaInvoiceText(invoice *tg.MessageMediaInvoice) []string {
	return []string{invoice.Title, invoice.Description}
}
