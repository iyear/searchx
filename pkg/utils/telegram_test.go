package utils

import "testing"

func TestGetMsgLink(t *testing.T) {
	tt := []struct {
		chat int64
		msg  int
		out  string
	}{
		{1697797156, 4, "https://t.me/c/1697797156/4"},
		{1231494493, 148, "https://t.me/c/1231494493/148"},
	}
	for _, n := range tt {
		if out := Telegram.GetMsgLink(n.chat, n.msg); out != n.out {
			t.Errorf("expected %s, got %s", n.out, out)
		}
	}
}
