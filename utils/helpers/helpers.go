// (c) Jisin0
//
// utils/helpers/helpers.go contains miscellaneous helper methods.

package helpers

import (
	"encoding/json"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

// IDFromUsername creates a custom reuqest to make a getchat request with a username and returns the target chatID.
func IDFromUsername(bot *gotgbot.Bot, username string) (int64, *gotgbot.Chat, error) {
	r, err := bot.Request("getChat", map[string]string{"chat_id": "@" + username}, nil, nil)
	if err != nil {
		return 0, nil, err
	}

	var c gotgbot.Chat

	err = json.Unmarshal(r, &c)
	if err != nil {
		return 0, nil, err
	}

	return c.Id, &c, nil
}

// MergeMaps just concatenates two maps.
func MergeMaps(dest, src map[string]any) {
	for key, value := range src {
		dest[key] = value
	}
}

// Contains Checks if a string slice Contains an item.
func Contains(l []string, v string) bool {
	for _, i := range l {
		if i == v {
			return true
		}
	}

	return false
}
