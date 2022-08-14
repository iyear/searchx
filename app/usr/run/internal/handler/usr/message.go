package usr

import (
	"github.com/gotd/td/tg"
	"strings"
)

func messageText(msg *tg.Message) string {
	// TODO(iyear): index service messages?

	media, ok := msg.GetMedia()
	if ok {
		return strings.Join([]string{msg.Message, mediaText(media)}, " ")
	}

	return msg.Message
}

func mediaText(media tg.MessageMediaClass) string {
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
	return ""
}

func mediaDocText(document *tg.MessageMediaDocument) string {
	doc, ok := document.Document.(*tg.Document)
	if !ok {
		return ""
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
			return ""
		}
		attrs = append(attrs, text)
	}

	return strings.Join(attrs, " ")
}

func mediaWebPageText(webPage *tg.MessageMediaWebPage) string {
	switch wp := webPage.Webpage.(type) {
	case *tg.WebPage:
		// TODO(iyear): index further content
		return strings.Join([]string{wp.SiteName, wp.Title, wp.Description}, " ")
	}
	return ""
}

func mediaVenueText(venue *tg.MessageMediaVenue) string {
	return strings.Join([]string{venue.Title, venue.Address}, " ")
}

func mediaInvoiceText(invoice *tg.MessageMediaInvoice) string {
	return strings.Join([]string{invoice.Title, invoice.Description}, " ")
}
