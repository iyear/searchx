package source

import (
	"errors"
	"github.com/bcicen/jstream"
	"github.com/fatih/color"
	"github.com/iyear/searchx/pkg/keygen"
	"github.com/iyear/searchx/pkg/models"
	"github.com/iyear/searchx/pkg/storage"
	"github.com/iyear/searchx/pkg/storage/search"
	"github.com/mitchellh/mapstructure"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	keyID       = "id"
	keyType     = "type"
	supergroup  = "supergroup"
	channel     = "channel"
	typeMessage = "message"
)

type message struct {
	ID   int         `mapstructure:"id"`
	Type string      `mapstructure:"type"`
	Time string      `mapstructure:"date_unixtime"`
	From string      `mapstructure:"from_id"`
	Text interface{} `mapstructure:"text"`
}

func Start(src, searchDriver string, searchOptions map[string]string) error {
	if searchDriver == "" {
		return errors.New("search driver can not be empty")
	}

	start := time.Now()

	options := make(map[string]interface{})
	if err := mapstructure.WeakDecode(searchOptions, &options); err != nil {
		return err
	}

	_search, err := search.New(searchDriver, options)
	if err != nil {
		return err
	}

	chatType, chatID, err := getChatInfo(src)
	if err != nil {
		return err
	}

	color.Blue("Type: %s, ID: %d\n", chatType, chatID)

	if err = index(src, chatID, _search); err != nil {
		return err
	}
	color.Blue("Index Succ... Time: %v", time.Since(start))

	return nil

}

func index(src string, chatID int64, search storage.Search) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}

	defer func(f *os.File) {
		if err = f.Close(); err != nil {
			log.Fatalln(err)
		}
	}(f)

	d := jstream.NewDecoder(f, 2)

	batchSize := 50
	items := make([]*storage.SearchItem, 0, batchSize)

	for mv := range d.Stream() {
		msg := message{}

		if mv.ValueType != jstream.Object {
			continue
		}

		if err = mapstructure.WeakDecode(mv.Value, &msg); err != nil {
			return err
		}

		if msg.ID < 0 || msg.Type != typeMessage {
			continue
		}

		text := ""

		switch r := msg.Text.(type) {
		case string:
			text = r
		case []interface{}:
			for _, tt := range r {
				switch t := tt.(type) {
				case string:
					text += t
				case map[string]interface{}:
					text += " " + t["text"].(string) + " "
				}
			}
		}

		if !strings.HasPrefix(msg.From, "user") {
			continue
		}

		if text != "" {
			items = append(items, &storage.SearchItem{
				ID: keygen.SearchMsgID(chatID, msg.ID),
				Data: &models.SearchMsg{
					ID:     strconv.Itoa(msg.ID),
					Chat:   strconv.FormatInt(chatID, 10),
					Text:   text,
					Sender: strings.TrimPrefix(msg.From, "user"),
					Date:   msg.Time,
				},
			})
		}

		if len(items) == batchSize {
			if err = search.Index(items); err != nil {
				return err
			}
			items = make([]*storage.SearchItem, 0, batchSize)
		}
	}

	if len(items) > 0 {
		if err = search.Index(items); err != nil {
			return err
		}
	}

	return nil

}

func getChatInfo(src string) (string, int64, error) {
	f, err := os.Open(src)
	if err != nil {
		return "", 0, err
	}
	defer func(f *os.File) {
		if err = f.Close(); err != nil {
			log.Fatalln(err)
		}
	}(f)

	d := jstream.NewDecoder(f, 1).EmitKV()

	var chatType = ""
	var chatID int64 = 0

	for mv := range d.Stream() {
		kv, ok := mv.Value.(jstream.KV)
		if !ok {
			continue
		}

		if kv.Key == keyType {
			chatType = kv.Value.(string)
			if !strings.HasSuffix(chatType, supergroup) && !strings.HasSuffix(chatType, channel) {
				return "", 0, errors.New("chat type should be supergroup or channel")
			}
		}

		if kv.Key == keyID {
			chatID = -int64(kv.Value.(float64)) - 1e12
		}

		if chatType != "" && chatID != 0 {
			break
		}
	}

	if chatType == "" || chatID == 0 {
		return "", 0, errors.New("can not get chat type or chat id")
	}

	return chatType, chatID, nil
}
