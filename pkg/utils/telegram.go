package utils

type telegram struct{}

var Telegram = telegram{}

func (t telegram) GetSenderName(first, last string) string {
	if last == "" {
		return first
	}
	if first == "" {
		return last
	}

	return first + " " + last
}

func (t telegram) GetDeepLink(bot string, code string) string {
	return "https://t.me/" + bot + "?start=" + code
}
