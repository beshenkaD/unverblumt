package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) KickUser(user int, chat, until int64) error {
	_, err := b.tg.KickChatMember(tgbotapi.KickChatMemberConfig{
		UntilDate: until,
		ChatMemberConfig: tgbotapi.ChatMemberConfig{
			UserID: user,
			ChatID: chat,
		},
	})

	return err
}

func (b *Bot) DeleteMessage(chat int64, message int) error {
	_, err := b.tg.DeleteMessage(tgbotapi.DeleteMessageConfig{
		ChatID:    chat,
		MessageID: message,
	})

	return err
}

func (b *Bot) RestrictUser(chat, until int64, user int, messages, media, other, webpage bool) error {
	_, err := b.tg.RestrictChatMember(tgbotapi.RestrictChatMemberConfig{
		ChatMemberConfig: tgbotapi.ChatMemberConfig{
			ChatID: chat,
			UserID: user,
		},
		UntilDate:             until,
		CanSendMessages:       &messages,
		CanSendMediaMessages:  &media,
		CanSendOtherMessages:  &other,
		CanAddWebPagePreviews: &webpage,
	})

	return err
}
