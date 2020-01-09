package telegram

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	http "github.com/valyala/fasthttp"
	"golang.org/x/text/language"
)

func TestUserLanguage(t *testing.T) {
	for _, tc := range []struct {
		name      string
		user      User
		expResult language.Tag
	}{{
		name:      "russian",
		user:      User{LanguageCode: "ru"},
		expResult: language.Russian,
	}, {
		name:      "english",
		user:      User{LanguageCode: "en"},
		expResult: language.English,
	}, {
		name:      "other",
		user:      User{LanguageCode: "w-t-f"},
		expResult: language.Und,
	}, {
		name:      "empty",
		expResult: language.Und,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.user.Language(), tc.expResult)
		})
	}
}

func TestUserFullName(t *testing.T) {
	for _, tc := range []struct {
		name      string
		user      User
		expResult string
	}{{
		name:      "full",
		user:      User{FirstName: "Maxim", LastName: "Lebedev"},
		expResult: "Maxim Lebedev",
	}, {
		name:      "first",
		user:      User{FirstName: "Maxim"},
		expResult: "Maxim",
	}, {
		name:      "last",
		user:      User{LastName: "Lebedev"},
		expResult: "",
	}, {
		name:      "empty",
		expResult: "",
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.user.FullName(), tc.expResult)
		})
	}
}

func TestUserHasLastName(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		u := User{LastName: "Lebedev"}
		assert.True(t, u.HasLastName())
	})
	t.Run("false", func(t *testing.T) {
		u := User{FirstName: "Maxim"}
		assert.False(t, u.HasLastName())
	})
}

func TestUserHasUsername(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		u := User{Username: "toby3d", FirstName: "Maxim", LastName: "Lebedev"}
		assert.True(t, u.HasUsername())
	})
	t.Run("false", func(t *testing.T) {
		u := User{FirstName: "Maxim", LastName: "Lebedev"}
		assert.False(t, u.HasUsername())
	})
}

func TestChatIsPrivate(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		c := Chat{Type: ChatPrivate, FirstName: "Maxim", LastName: "Lebedev"}
		assert.True(t, c.IsPrivate())
	})
	t.Run("false", func(t *testing.T) {
		c := Chat{Type: ChatSuperGroup, FirstName: "Меня заставили создать эту группу"}
		assert.False(t, c.IsPrivate())
	})
}

func TestChatIsGroup(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		c := Chat{Type: ChatGroup, FirstName: "Меня заставили создать эту группу"}
		assert.True(t, c.IsGroup())
	})
	t.Run("false", func(t *testing.T) {
		c := Chat{Type: ChatPrivate, FirstName: "Maxim", LastName: "Lebedev"}
		assert.False(t, c.IsGroup())
	})
}

func TestChatIsSuperGroup(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		c := Chat{Type: ChatSuperGroup, FirstName: "Меня заставили создать эту группу"}
		assert.True(t, c.IsSuperGroup())
	})
	t.Run("false", func(t *testing.T) {
		c := Chat{Type: ChatPrivate, FirstName: "Maxim", LastName: "Lebedev"}
		assert.False(t, c.IsSuperGroup())
	})
}

func TestChatIsChannel(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		c := Chat{Type: ChatChannel, FirstName: "toby3d"}
		assert.True(t, c.IsChannel())
	})
	t.Run("false", func(t *testing.T) {
		c := Chat{Type: ChatPrivate, FirstName: "Меня заставили создать эту группу"}
		assert.False(t, c.IsChannel())
	})
}

func TestChatHasPinnedMessage(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		c := Chat{PinnedMessage: &Message{Text: "hello, world!"}}
		assert.True(t, c.HasPinnedMessage())
	})
	t.Run("false", func(t *testing.T) {
		c := Chat{}
		assert.False(t, c.HasPinnedMessage())
	})
}

func TestChatHasStickerSet(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		c := Chat{StickerSetName: "HotCherry"}
		assert.True(t, c.HasStickerSet())
	})
	t.Run("false", func(t *testing.T) {
		c := Chat{}
		assert.False(t, c.HasStickerSet())
	})
}

func TestChatFullName(t *testing.T) {
	for _, tc := range []struct {
		name      string
		chat      Chat
		expResult string
	}{{
		name:      "full",
		chat:      Chat{FirstName: "Maxim", LastName: "Lebedev"},
		expResult: "Maxim Lebedev",
	}, {
		name:      "first",
		chat:      Chat{FirstName: "Меня заставили создать эту группу"},
		expResult: "Меня заставили создать эту группу",
	}, {
		name:      "last",
		chat:      Chat{LastName: "WTF"},
		expResult: "",
	}, {
		name:      "empty",
		expResult: "",
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.chat.FullName(), tc.expResult)
		})
	}
}

func TestChatHasLastName(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		c := Chat{FirstName: "Maxim", LastName: "Lebedev"}
		assert.True(t, c.HasLastName())
	})
	t.Run("false", func(t *testing.T) {
		c := Chat{FirstName: "Меня заставили создать эту группу", LastName: ""}
		assert.False(t, c.HasLastName())
	})
}

func TestChatHasUsername(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		c := Chat{Username: "toby3dRu"}
		assert.True(t, c.HasUsername())
	})
	t.Run("false", func(t *testing.T) {
		c := Chat{}
		assert.False(t, c.HasUsername())
	})
}

func TestChatHasDescription(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		c := Chat{Description: "hello, world!"}
		assert.True(t, c.HasDescription())
	})
	t.Run("false", func(t *testing.T) {
		c := Chat{}
		assert.False(t, c.HasDescription())
	})
}

func TestChatHasInviteLink(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		c := Chat{InviteLink: "https://t.me/joinchat/AAAAAD84EWpWTRTMh4aahQ"}
		assert.True(t, c.HasInviteLink())
	})
	t.Run("false", func(t *testing.T) {
		c := Chat{}
		assert.False(t, c.HasInviteLink())
	})
}

func TestMessageIsCommand(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{Text: "/" + CommandStart, Entities: []*MessageEntity{{Type: EntityBotCommand}}}
		assert.True(t, m.IsCommand())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, @toby3d", Entities: []*MessageEntity{{Type: EntityMention}}}
		assert.False(t, m.IsCommand())
	})
}

func TestMessageIsCommandEqual(t *testing.T) {
	for _, tc := range []struct {
		name      string
		message   Message
		command   string
		expResult bool
	}{{
		name: "exact",
		message: Message{
			Text:     "/" + CommandStart,
			Entities: []*MessageEntity{{Type: EntityBotCommand, Length: len(CommandStart) + 1}},
		},
		command:   CommandStart,
		expResult: true,
	}, {
		name: "other",
		message: Message{
			Text:     "/" + CommandStart,
			Entities: []*MessageEntity{{Type: EntityBotCommand, Length: len(CommandStart) + 1}},
		},
		command:   CommandHelp,
		expResult: false,
	}, {
		name:      "empty",
		command:   CommandHelp,
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.message.IsCommandEqual(tc.command), tc.expResult)
		})
	}
}

func TestMessageCommand(t *testing.T) {
	for _, tc := range []struct {
		name      string
		message   Message
		expResult string
	}{{
		name: "command",
		message: Message{
			Text:     "/" + CommandStart,
			Entities: []*MessageEntity{{Type: EntityBotCommand, Length: len(CommandStart) + 1}},
		},
		expResult: CommandStart,
	}, {
		name: "other",
		message: Message{
			Text:     "hello world",
			Entities: []*MessageEntity{{Type: EntityBold, Offset: 6, Length: 5}},
		},
		expResult: "",
	}, {
		name:      "empty",
		expResult: "",
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.message.Command(), tc.expResult)
		})
	}
}

func TestMessageRawCommand(t *testing.T) {
	for _, tc := range []struct {
		name      string
		message   Message
		expResult string
	}{{
		name: "command",
		message: Message{
			Text:     "/" + CommandStart,
			Entities: []*MessageEntity{{Type: EntityBotCommand, Length: len(CommandStart) + 1}},
		},
		expResult: CommandStart,
	}, {
		name: "other",
		message: Message{
			Text:     "hello world",
			Entities: []*MessageEntity{{Type: EntityBold, Offset: 6, Length: 5}},
		},
		expResult: "",
	}, {
		name:      "empty",
		expResult: "",
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.message.RawCommand(), tc.expResult)
		})
	}
}

func TestMessageHasCommandArgument(t *testing.T) {
	for _, tc := range []struct {
		name      string
		message   Message
		expResult bool
	}{{
		name: "true",
		message: Message{
			Text:     "/start example",
			Entities: []*MessageEntity{{Type: EntityBotCommand, Length: len(CommandStart) + 1}},
		},
		expResult: true,
	}, {
		name: "other",
		message: Message{
			Text:     "/start",
			Entities: []*MessageEntity{{Type: EntityBold, Offset: 1, Length: len(CommandStart)}},
		},
		expResult: false,
	}, {
		name: "false",
		message: Message{
			Text:     "/start",
			Entities: []*MessageEntity{{Type: EntityBotCommand, Length: len(CommandStart) + 1}},
		},
		expResult: false,
	}, {
		name:      "empty",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.message.HasCommandArgument(), tc.expResult)
		})
	}
}

func TestMessageCommandArgument(t *testing.T) {
	for _, tc := range []struct {
		name      string
		message   Message
		expResult string
	}{{
		name: "true",
		message: Message{
			Text:     "/start example",
			Entities: []*MessageEntity{{Type: EntityBotCommand, Length: len(CommandStart) + 1}},
		},
		expResult: "example",
	}, {
		name: "false",
		message: Message{
			Text:     "/start",
			Entities: []*MessageEntity{{Type: EntityBotCommand, Length: len(CommandStart) + 1}},
		},
		expResult: "",
	}, {
		name:      "empty",
		expResult: "",
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.message.CommandArgument(), tc.expResult)
		})
	}
}

func TestMessageIsReply(t *testing.T) {
	for _, tc := range []struct {
		name      string
		message   Message
		expResult bool
	}{{
		name:      "true",
		message:   Message{Text: "hello!", ReplyToMessage: &Message{Text: "hello, world!"}},
		expResult: true,
	}, {
		name:      "false",
		message:   Message{Text: "hello!"},
		expResult: false,
	}, {
		name:      "empty",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.message.IsReply(), tc.expResult)
		})
	}
}

func TestMessageIsForward(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{Text: "hello, world!", ForwardDate: time.Now().Unix()}
		assert.True(t, m.IsForward())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.False(t, m.IsForward())
	})
}

func TestMessageTime(t *testing.T) {
	now := time.Now().Round(time.Second)
	zero := time.Time{}

	for _, tc := range []struct {
		name      string
		message   Message
		expResult time.Time
	}{{
		name:      "specific date",
		message:   Message{Date: now.Unix()},
		expResult: now,
	}, {
		name:      "zero",
		message:   Message{Date: 0},
		expResult: zero,
	}, {
		name:      "empty",
		expResult: zero,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, tc.message.Time().Equal(tc.expResult))
		})
	}
}

func TestMessageForwardTime(t *testing.T) {
	now := time.Now().Round(time.Second)
	zero := time.Time{}

	for _, tc := range []struct {
		name      string
		message   Message
		expResult time.Time
	}{{
		name:      "specific date",
		message:   Message{ForwardDate: now.Unix()},
		expResult: now,
	}, {
		name:      "false",
		message:   Message{ForwardDate: 0},
		expResult: zero,
	}, {
		name:      "empty",
		expResult: zero,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, tc.message.ForwardTime().Equal(tc.expResult))
		})
	}
}

func TestMessageEditTime(t *testing.T) {
	now := time.Now().Round(time.Second)
	zero := time.Time{}

	for _, tc := range []struct {
		name      string
		message   Message
		expResult time.Time
	}{{
		name:      "true",
		message:   Message{EditDate: now.Unix()},
		expResult: now,
	}, {
		name:      "false",
		message:   Message{EditDate: 0},
		expResult: zero,
	}, {
		name:      "empty",
		expResult: zero,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, tc.message.EditTime().Equal(tc.expResult))
		})
	}
}

func TestMessageHasBeenEdited(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{Text: "sorry, fixed", EditDate: time.Now().UTC().Unix()}
		assert.True(t, m.HasBeenEdited())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "sorry, fixed", EditDate: 0}
		assert.False(t, m.HasBeenEdited())
	})
}

func TestMessageIsText(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.True(t, m.IsText())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{}
		assert.False(t, m.IsText())
	})
}

func TestMessageIsAudio(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{Audio: &Audio{Performer: "One-Aloner", Title: "Monitor"}}
		assert.True(t, m.IsAudio())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{}
		assert.False(t, m.IsAudio())
	})
}

func TestMessageIsDocument(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{Document: &Document{FileName: "readme.txt"}}
		assert.True(t, m.IsDocument())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.False(t, m.IsDocument())
	})
}

func TestMessageIsGame(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{Game: &Game{Title: "MuSquare"}}
		assert.True(t, m.IsGame())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.False(t, m.IsGame())
	})
}

func TestMessageIsPhoto(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{Photo: []*PhotoSize{{FileID: "abc"}}}
		assert.True(t, m.IsPhoto())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.False(t, m.IsPhoto())
	})
}

func TestMessageIsSticker(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{Sticker: &Sticker{FileID: "abc"}}
		assert.True(t, m.IsSticker())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.False(t, m.IsSticker())
	})
}

func TestMessageIsVideo(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{Video: &Video{FileID: "abc"}}
		assert.True(t, m.IsVideo())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.False(t, m.IsVideo())
	})
}

func TestMessageIsVoice(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{Voice: &Voice{FileID: "abc"}}
		assert.True(t, m.IsVoice())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.False(t, m.IsVoice())
	})
}

func TestMessageIsVideoNote(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{VideoNote: &VideoNote{FileID: "abc"}}
		assert.True(t, m.IsVideoNote())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.False(t, m.IsVideoNote())
	})
}

func TestMessageIsContact(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{Contact: &Contact{FirstName: "Maxim", PhoneNumber: "1234567890"}}
		assert.True(t, m.IsContact())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.False(t, m.IsContact())
	})
}

func TestMessageIsLocation(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{Location: &Location{Latitude: 56.085180, Longitude: 60.735150}}
		assert.True(t, m.IsLocation())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.False(t, m.IsLocation())
	})
}

func TestMessageIsVenue(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{Venue: &Venue{
			Title:    "Памятник В. И. Ленину",
			Address:  "Россия, Челябинская область, Снежинск, улица Свердлова",
			Location: &Location{Latitude: 56.085180, Longitude: 60.735150},
		}}
		assert.True(t, m.IsVenue())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.False(t, m.IsVenue())
	})
}

func TestMessageIsAnimation(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{Animation: &Animation{FileID: "abc"}}
		assert.True(t, m.IsAnimation())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.False(t, m.IsAnimation())
	})
}

func TestMessageIsNewChatMembersEvent(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{NewChatMembers: []*User{{ID: 42, FirstName: "Maxim", LastName: "Lebedev"}}}
		assert.True(t, m.IsNewChatMembersEvent())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.False(t, m.IsNewChatMembersEvent())
	})
}

func TestMessageIsLeftChatMemberEvent(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{LeftChatMember: &User{ID: 42, FirstName: "Maxim", LastName: "Lebedev"}}
		assert.True(t, m.IsLeftChatMemberEvent())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.False(t, m.IsLeftChatMemberEvent())
	})
}

func TestMessageIsNewChatTitleEvent(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{NewChatTitle: "Меня заставили создать эту группу"}
		assert.True(t, m.IsNewChatTitleEvent())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.False(t, m.IsNewChatTitleEvent())
	})
}

func TestMessageIsNewChatPhotoEvent(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{NewChatPhoto: []*PhotoSize{{FileID: "abc"}}}
		assert.True(t, m.IsNewChatPhotoEvent())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.False(t, m.IsNewChatPhotoEvent())
	})
}

func TestMessageIsDeleteChatPhotoEvent(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{DeleteChatPhoto: true}
		assert.True(t, m.IsDeleteChatPhotoEvent())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.False(t, m.IsDeleteChatPhotoEvent())
	})
}

func TestMessageIsGroupChatCreatedEvent(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{GroupChatCreated: true}
		assert.True(t, m.IsGroupChatCreatedEvent())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.False(t, m.IsGroupChatCreatedEvent())
	})
}

func TestMessageIsSupergroupChatCreatedEvent(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{SupergroupChatCreated: true}
		assert.True(t, m.IsSupergroupChatCreatedEvent())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.False(t, m.IsSupergroupChatCreatedEvent())
	})
}

func TestMessageIsChannelChatCreatedEvent(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{ChannelChatCreated: true}
		assert.True(t, m.IsChannelChatCreatedEvent())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.False(t, m.IsChannelChatCreatedEvent())
	})
}

func TestMessageIsPinnedMessage(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{PinnedMessage: &Message{Text: "hello, world!"}}
		assert.True(t, m.IsPinnedMessage())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.False(t, m.IsPinnedMessage())
	})
}

func TestMessageIsInvoice(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{Invoice: &Invoice{Title: "Time Machine", TotalAmount: 1}}
		assert.True(t, m.IsInvoice())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.False(t, m.IsInvoice())
	})
}

func TestMessageIsSuccessfulPayment(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{SuccessfulPayment: &SuccessfulPayment{
			OrderInfo: &OrderInfo{Name: "Maxim Lebedev"}, TotalAmount: 1,
		}}
		assert.True(t, m.IsSuccessfulPayment())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.False(t, m.IsSuccessfulPayment())
	})
}

func TestMessageIsPoll(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{Poll: &Poll{
			Question: "What is the answer to life, the universe, and everything?",
			Options: []*PollOption{
				{Text: "42", VoterCount: 420},
				{Text: "24", VoterCount: 140},
			},
		}}
		assert.True(t, m.IsPoll())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.False(t, m.IsPoll())
	})
}

func TestMessageHasEntities(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m := Message{Text: "hello, world!", Entities: []*MessageEntity{
			{Type: EntityBold, Offset: 7, Length: 5},
		}}
		assert.True(t, m.HasEntities())
	})
	t.Run("false", func(t *testing.T) {
		m := Message{Text: "hello, world!"}
		assert.False(t, m.HasEntities())
	})
}

func TestMessageHasCaptionEntities(t *testing.T) {
	for _, tc := range []struct {
		name      string
		message   Message
		expResult bool
	}{{
		name: "true",
		message: Message{
			Caption: "hello, world!", CaptionEntities: []*MessageEntity{
				{Type: EntityBold, Offset: 7, Length: 5},
			},
		},
		expResult: true,
	}, {
		name:      "false",
		message:   Message{Text: "hello, world!"},
		expResult: false,
	}, {
		name:      "empty",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.message.HasCaptionEntities(), tc.expResult)
		})
	}
}

func TestMessageHasMentions(t *testing.T) {
	for _, tc := range []struct {
		name      string
		message   Message
		expResult bool
	}{{
		name: "true",
		message: Message{Text: "hello, @toby3d!", Entities: []*MessageEntity{
			{Type: EntityMention, Offset: 7, Length: 7},
		}},
		expResult: true,
	}, {
		name: "other",
		message: Message{Text: "hello, world!", Entities: []*MessageEntity{
			{Type: EntityBold, Offset: 7, Length: 5},
		}},
		expResult: false,
	}, {
		name:      "false",
		message:   Message{Text: "hello, world!"},
		expResult: false,
	}, {
		name:      "empty",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.message.HasMentions(), tc.expResult)
		})
	}
}

func TestMessageHasCaptionMentions(t *testing.T) {
	for _, tc := range []struct {
		name      string
		message   Message
		expResult bool
	}{{
		name: "true",
		message: Message{Caption: "hello, @toby3d!", CaptionEntities: []*MessageEntity{
			{Type: EntityMention, Offset: 7, Length: 7},
		}},
		expResult: true,
	}, {
		name: "other",
		message: Message{Text: "hello, world!", CaptionEntities: []*MessageEntity{
			{Type: EntityBold, Offset: 7, Length: 5},
		}},
		expResult: false,
	}, {
		name:      "false",
		message:   Message{Text: "hello, world!"},
		expResult: false,
	}, {
		name:      "empty",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.message.HasCaptionMentions(), tc.expResult)
		})
	}
}

func TestMessageHasCaption(t *testing.T) {
	for _, tc := range []struct {
		name      string
		message   Message
		expResult bool
	}{{
		name:      "true",
		message:   Message{Caption: "hello, world!"},
		expResult: true,
	}, {
		name:      "false",
		message:   Message{Text: "hello, world!"},
		expResult: false,
	}, {
		name:      "empty",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.message.HasCaption(), tc.expResult)
		})
	}
}

func TestMessageHasAuthorSignature(t *testing.T) {
	for _, tc := range []struct {
		name      string
		message   Message
		expResult bool
	}{{
		name:      "true",
		message:   Message{AuthorSignature: "Editor"},
		expResult: true,
	}, {
		name:      "false",
		message:   Message{Text: "hello, world!"},
		expResult: false,
	}, {
		name:      "empty",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.message.HasAuthorSignature(), tc.expResult)
		})
	}
}

func TestMessageIsEvent(t *testing.T) {
	for _, tc := range []struct {
		name      string
		message   Message
		expResult bool
	}{{
		name:      "ChannelChatCreated",
		message:   Message{ChannelChatCreated: true},
		expResult: true,
	}, {
		name:      "DeleteChatPhoto",
		message:   Message{DeleteChatPhoto: true},
		expResult: true,
	}, {
		name:      "GroupChatCreated",
		message:   Message{GroupChatCreated: true},
		expResult: true,
	}, {
		name:      "LeftChatMember",
		message:   Message{LeftChatMember: &User{ID: 42}},
		expResult: true,
	}, {
		name:      "NewChatMembers",
		message:   Message{NewChatMembers: []*User{{ID: 42}}},
		expResult: true,
	}, {
		name:      "NewChatTitle",
		message:   Message{NewChatTitle: "Меня заставили создать эту группу"},
		expResult: true,
	}, {
		name:      "SupergroupChatCreated",
		message:   Message{SupergroupChatCreated: true},
		expResult: true,
	}, {
		name:      "NewChatPhoto",
		message:   Message{NewChatPhoto: []*PhotoSize{{FileID: "abc"}}},
		expResult: true,
	}, {
		name:      "false",
		message:   Message{Text: "hello, world!"},
		expResult: false,
	}, {
		name:      "empty",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.message.IsEvent(), tc.expResult)
		})
	}
}

func TestMessageEntityIsBold(t *testing.T) {
	for _, tc := range []struct {
		name      string
		entity    MessageEntity
		expResult bool
	}{{
		name:      "true",
		entity:    MessageEntity{Type: EntityBold},
		expResult: true,
	}, {
		name:      "false",
		entity:    MessageEntity{Type: EntityMention},
		expResult: false,
	}, {
		name:      "empty",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.entity.IsBold(), tc.expResult)
		})
	}
}

func TestMessageEntityIsBotCommand(t *testing.T) {
	for _, tc := range []struct {
		name      string
		entity    MessageEntity
		expResult bool
	}{{
		name:      "true",
		entity:    MessageEntity{Type: EntityBotCommand},
		expResult: true,
	}, {
		name:      "false",
		entity:    MessageEntity{Type: EntityMention},
		expResult: false,
	}, {
		name:      "empty",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.entity.IsBotCommand(), tc.expResult)
		})
	}
}

func TestMessageEntityIsCashtag(t *testing.T) {
	for _, tc := range []struct {
		name      string
		entity    MessageEntity
		expResult bool
	}{{
		name:      "true",
		entity:    MessageEntity{Type: EntityCashtag},
		expResult: true,
	}, {
		name:      "false",
		entity:    MessageEntity{Type: EntityMention},
		expResult: false,
	}, {
		name:      "empty",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.entity.IsCashtag(), tc.expResult)
		})
	}
}

func TestMessageEntityIsCode(t *testing.T) {
	for _, tc := range []struct {
		name      string
		entity    MessageEntity
		expResult bool
	}{{
		name:      "true",
		entity:    MessageEntity{Type: EntityCode},
		expResult: true,
	}, {
		name:      "false",
		entity:    MessageEntity{Type: EntityMention},
		expResult: false,
	}, {
		name:      "empty",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.entity.IsCode(), tc.expResult)
		})
	}
}

func TestMessageEntityIsEmail(t *testing.T) {
	for _, tc := range []struct {
		name      string
		entity    MessageEntity
		expResult bool
	}{{
		name:      "true",
		entity:    MessageEntity{Type: EntityEmail},
		expResult: true,
	}, {
		name:      "false",
		entity:    MessageEntity{Type: EntityMention},
		expResult: false,
	}, {
		name:      "empty",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.entity.IsEmail(), tc.expResult)
		})
	}
}

func TestMessageEntityIsHashtag(t *testing.T) {
	for _, tc := range []struct {
		name      string
		entity    MessageEntity
		expResult bool
	}{{
		name:      "true",
		entity:    MessageEntity{Type: EntityHashtag},
		expResult: true,
	}, {
		name:      "false",
		entity:    MessageEntity{Type: EntityMention},
		expResult: false,
	}, {
		name:      "empty",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.entity.IsHashtag(), tc.expResult)
		})
	}
}

func TestMessageEntityIsItalic(t *testing.T) {
	for _, tc := range []struct {
		name      string
		entity    MessageEntity
		expResult bool
	}{{
		name:      "true",
		entity:    MessageEntity{Type: EntityItalic},
		expResult: true,
	}, {
		name:      "false",
		entity:    MessageEntity{Type: EntityMention},
		expResult: false,
	}, {
		name:      "empty",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.entity.IsItalic(), tc.expResult)
		})
	}
}

func TestMessageEntityIsMention(t *testing.T) {
	for _, tc := range []struct {
		name      string
		entity    MessageEntity
		expResult bool
	}{{
		name:      "true",
		entity:    MessageEntity{Type: EntityMention},
		expResult: true,
	}, {
		name:      "false",
		entity:    MessageEntity{Type: EntityURL},
		expResult: false,
	}, {
		name:      "empty",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.entity.IsMention(), tc.expResult)
		})
	}
}

func TestMessageEntityIsPhoneNumber(t *testing.T) {
	for _, tc := range []struct {
		name      string
		entity    MessageEntity
		expResult bool
	}{{
		name:      "true",
		entity:    MessageEntity{Type: EntityPhoneNumber},
		expResult: true,
	}, {
		name:      "false",
		entity:    MessageEntity{Type: EntityMention},
		expResult: false,
	}, {
		name:      "empty",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.entity.IsPhoneNumber(), tc.expResult)
		})
	}
}

func TestMessageEntityIsPre(t *testing.T) {
	for _, tc := range []struct {
		name      string
		entity    MessageEntity
		expResult bool
	}{{
		name:      "true",
		entity:    MessageEntity{Type: EntityPre},
		expResult: true,
	}, {
		name:      "false",
		entity:    MessageEntity{Type: EntityMention},
		expResult: false,
	}, {
		name:      "empty",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.entity.IsPre(), tc.expResult)
		})
	}
}

func TestMessageEntityIsStrikethrough(t *testing.T) {
	for _, tc := range []struct {
		name      string
		entity    MessageEntity
		expResult bool
	}{{
		name:      "true",
		entity:    MessageEntity{Type: EntityStrikethrough},
		expResult: true,
	}, {
		name:      "false",
		entity:    MessageEntity{Type: EntityMention},
		expResult: false,
	}, {
		name:      "empty",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.entity.IsStrikethrough(), tc.expResult)
		})
	}
}

func TestMessageEntityIsTextLink(t *testing.T) {
	for _, tc := range []struct {
		name      string
		entity    MessageEntity
		expResult bool
	}{{
		name:      "true",
		entity:    MessageEntity{Type: EntityTextLink},
		expResult: true,
	}, {
		name:      "false",
		entity:    MessageEntity{Type: EntityMention},
		expResult: false,
	}, {
		name:      "empty",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.entity.IsTextLink(), tc.expResult)
		})
	}
}

func TestMessageEntityIsTextMention(t *testing.T) {
	for _, tc := range []struct {
		name      string
		entity    MessageEntity
		expResult bool
	}{{
		name:      "true",
		entity:    MessageEntity{Type: EntityTextMention},
		expResult: true,
	}, {
		name:      "false",
		entity:    MessageEntity{Type: EntityMention},
		expResult: false,
	}, {
		name:      "empty",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.entity.IsTextMention(), tc.expResult)
		})
	}
}

func TestMessageEntityIsUnderline(t *testing.T) {
	for _, tc := range []struct {
		name      string
		entity    MessageEntity
		expResult bool
	}{{
		name:      "true",
		entity:    MessageEntity{Type: EntityUnderline},
		expResult: true,
	}, {
		name:      "false",
		entity:    MessageEntity{Type: EntityMention},
		expResult: false,
	}, {
		name:      "empty",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.entity.IsUnderline(), tc.expResult)
		})
	}
}

func TestMessageEntityIsURL(t *testing.T) {
	for _, tc := range []struct {
		name      string
		entity    MessageEntity
		expResult bool
	}{{
		name:      "true",
		entity:    MessageEntity{Type: EntityURL},
		expResult: true,
	}, {
		name:      "false",
		entity:    MessageEntity{Type: EntityMention},
		expResult: false,
	}, {
		name:      "empty",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.entity.IsURL(), tc.expResult)
		})
	}
}

func TestMessageEntityParseURL(t *testing.T) {
	link := http.AcquireURI()
	defer http.ReleaseURI(link)

	link.Update("https://toby3d.me")

	for _, tc := range []struct {
		name      string
		entity    MessageEntity
		text      string
		expResult *http.URI
	}{{
		name:      "valid",
		entity:    MessageEntity{Type: EntityURL, Length: len(link.String())},
		text:      link.String(),
		expResult: link,
	}, {
		name:      "other",
		entity:    MessageEntity{Type: EntityTextLink},
		text:      link.String(),
		expResult: nil,
	}, {
		name:      "wrong text",
		entity:    MessageEntity{Type: EntityURL, Length: len(link.String())},
		text:      "wtf",
		expResult: nil,
	}, {
		name:      "empty",
		text:      "",
		expResult: nil,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if tc.expResult == nil {
				assert.Nil(t, tc.entity.ParseURL(tc.text))
				return
			}

			assert.Equal(t, tc.entity.ParseURL(tc.text).String(), tc.expResult.String())
		})
	}
}

func TestMessageEntityTextLink(t *testing.T) {
	link := http.AcquireURI()
	defer http.ReleaseURI(link)

	link.Update("https://toby3d.me")

	for _, tc := range []struct {
		name      string
		entity    MessageEntity
		expResult *http.URI
	}{{
		name:      "valid",
		entity:    MessageEntity{Type: EntityTextLink, URL: link.String()},
		expResult: link,
	}, {
		name:      "other",
		entity:    MessageEntity{Type: EntityURL},
		expResult: nil,
	}, {
		name:      "empty",
		expResult: nil,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if tc.expResult == nil {
				assert.Nil(t, tc.entity.TextLink())
				return
			}

			assert.Equal(t, tc.entity.TextLink().String(), tc.expResult.String())
		})
	}
}

func TestAudioFullName(t *testing.T) {
	for _, tc := range []struct {
		name      string
		audio     Audio
		separator string
		expResult string
	}{{
		name:      "full",
		audio:     Audio{Performer: "One-Aloner", Title: "Monitor"},
		separator: " | ",
		expResult: "One-Aloner | Monitor",
	}, {
		name:      "performer",
		audio:     Audio{Performer: "One-Aloner"},
		separator: " | ",
		expResult: "One-Aloner | " + DefaultAudioTitle,
	}, {
		name:      "title",
		audio:     Audio{Title: "Monitor"},
		separator: " | ",
		expResult: "Monitor",
	}, {
		name:      "without separator",
		audio:     Audio{Performer: "One-Aloner", Title: "Monitor"},
		expResult: "One-Aloner" + DefaultAudioSeparator + "Monitor",
	}, {
		name:      "empty",
		expResult: DefaultAudioTitle,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.audio.FullName(tc.separator), tc.expResult)
		})
	}
}

func TestAudioHasPerformer(t *testing.T) {
	for _, tc := range []struct {
		name      string
		audio     Audio
		expResult bool
	}{{
		name:      "true",
		audio:     Audio{Performer: "One-Aloner", Title: "Monitor"},
		expResult: true,
	}, {
		name:      "false",
		audio:     Audio{Title: "Monitor"},
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.audio.HasPerformer(), tc.expResult)
		})
	}
}

func TestAudioHasTitle(t *testing.T) {
	for _, tc := range []struct {
		name      string
		audio     Audio
		expResult bool
	}{{
		name:      "true",
		audio:     Audio{Performer: "One-Aloner", Title: "Monitor"},
		expResult: true,
	}, {
		name:      "false",
		audio:     Audio{Performer: "One-Aloner"},
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.audio.HasTitle(), tc.expResult)
		})
	}
}

func TestAudioHasThumb(t *testing.T) {
	for _, tc := range []struct {
		name      string
		audio     Audio
		expResult bool
	}{{
		name:      "true",
		audio:     Audio{Thumb: &PhotoSize{FileID: "abc"}},
		expResult: true,
	}, {
		name:      "false",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.audio.HasThumb(), tc.expResult)
		})
	}
}

func TestAudioFile(t *testing.T) {
	for _, tc := range []struct {
		name      string
		audio     Audio
		expResult File
	}{{
		name:      "valid",
		audio:     Audio{FileID: "abc", FileUniqueID: "really_abc", FileSize: 42},
		expResult: File{FileID: "abc", FileUniqueID: "really_abc", FileSize: 42},
	}, {
		name:      "empty",
		expResult: File{},
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.audio.File(), tc.expResult)
		})
	}
}

func TestDocumentHasThumb(t *testing.T) {
	for _, tc := range []struct {
		name      string
		document  Document
		expResult bool
	}{{
		name:      "true",
		document:  Document{Thumb: &PhotoSize{FileID: "abc"}},
		expResult: true,
	}, {
		name:      "false",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.document.HasThumb(), tc.expResult)
		})
	}
}

func TestDocumentFile(t *testing.T) {
	for _, tc := range []struct {
		name      string
		document  Document
		expResult File
	}{{
		name:      "valid",
		document:  Document{FileID: "abc", FileUniqueID: "really_abc", FileSize: 42},
		expResult: File{FileID: "abc", FileUniqueID: "really_abc", FileSize: 42},
	}, {
		name:      "empty",
		expResult: File{},
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.document.File(), tc.expResult)
		})
	}
}

func TestVideoHasThumb(t *testing.T) {
	for _, tc := range []struct {
		name      string
		video     Video
		expResult bool
	}{{
		name:      "true",
		video:     Video{Thumb: &PhotoSize{FileID: "abc"}},
		expResult: true,
	}, {
		name:      "false",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.video.HasThumb(), tc.expResult)
		})
	}
}

func TestVideoFile(t *testing.T) {
	for _, tc := range []struct {
		name      string
		video     Video
		expResult File
	}{{
		name:      "valid",
		video:     Video{FileID: "abc", FileUniqueID: "really_abc", FileSize: 42},
		expResult: File{FileID: "abc", FileUniqueID: "really_abc", FileSize: 42},
	}, {
		name:      "empty",
		expResult: File{},
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.video.File(), tc.expResult)
		})
	}
}

func TestAnimationHasThumb(t *testing.T) {
	for _, tc := range []struct {
		name      string
		animation Animation
		expResult bool
	}{{
		name:      "true",
		animation: Animation{Thumb: &PhotoSize{FileID: "abc"}},
		expResult: true,
	}, {
		name:      "false",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.animation.HasThumb(), tc.expResult)
		})
	}
}

func TestAnimationFile(t *testing.T) {
	for _, tc := range []struct {
		name      string
		animation Animation
		expResult File
	}{{
		name:      "valid",
		animation: Animation{FileID: "abc", FileUniqueID: "really_abc", FileSize: 42},
		expResult: File{FileID: "abc", FileUniqueID: "really_abc", FileSize: 42},
	}, {
		name:      "empty",
		expResult: File{},
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.animation.File(), tc.expResult)
		})
	}
}

func TestVoiceFile(t *testing.T) {
	for _, tc := range []struct {
		name      string
		voice     Voice
		expResult File
	}{{
		name:      "valid",
		voice:     Voice{FileID: "abc", FileUniqueID: "really_abc", FileSize: 42},
		expResult: File{FileID: "abc", FileUniqueID: "really_abc", FileSize: 42},
	}, {
		name:      "empty",
		expResult: File{},
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.voice.File(), tc.expResult)
		})
	}
}

func TestVideoNoteHasThumb(t *testing.T) {
	for _, tc := range []struct {
		name      string
		videoNote VideoNote
		expResult bool
	}{{
		name:      "true",
		videoNote: VideoNote{Thumb: &PhotoSize{FileID: "abc"}},
		expResult: true,
	}, {
		name:      "false",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.videoNote.HasThumb(), tc.expResult)
		})
	}
}

func TestVideoNoteFile(t *testing.T) {
	for _, tc := range []struct {
		name      string
		videoNote VideoNote
		expResult File
	}{{
		name:      "valid",
		videoNote: VideoNote{FileID: "abc", FileUniqueID: "really_abc", FileSize: 42},
		expResult: File{FileID: "abc", FileUniqueID: "really_abc", FileSize: 42},
	}, {
		name:      "empty",
		expResult: File{},
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.videoNote.File(), tc.expResult)
		})
	}
}

func TestContactFullName(t *testing.T) {
	for _, tc := range []struct {
		name      string
		contact   Contact
		expResult string
	}{{
		name:      "full",
		contact:   Contact{FirstName: "Maxim", LastName: "Lebedev"},
		expResult: "Maxim Lebedev",
	}, {
		name:      "first",
		contact:   Contact{FirstName: "Maxim"},
		expResult: "Maxim",
	}, {
		name:      "last",
		contact:   Contact{LastName: "Lebedev"},
		expResult: "",
	}, {
		name:      "false",
		expResult: "",
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.contact.FullName(), tc.expResult)
		})
	}
}

func TestContactHasLastName(t *testing.T) {
	for _, tc := range []struct {
		name      string
		contact   Contact
		expResult bool
	}{{
		name:      "true",
		contact:   Contact{FirstName: "Maxim", LastName: "Lebedev"},
		expResult: true,
	}, {
		name:      "false",
		contact:   Contact{FirstName: "Maxim"},
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.contact.HasLastName(), tc.expResult)
		})
	}
}

func TestContactInTelegram(t *testing.T) {
	for _, tc := range []struct {
		name      string
		contact   Contact
		expResult bool
	}{{
		name:      "true",
		contact:   Contact{UserID: 42},
		expResult: true,
	}, {
		name:      "false",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.contact.InTelegram(), tc.expResult)
		})
	}
}

func TestContactHasVCard(t *testing.T) {
	for _, tc := range []struct {
		name      string
		contact   Contact
		expResult bool
	}{{
		name: "true",
		contact: Contact{
			VCard: "BEGIN:VCARD\nVERSION:3.0\nFN;CHARSET=UTF-8:Maxim Maksimovich Lebedev\n" +
				"N;CHARSET=UTF-8:Lebedev;Maxim;Maksimovich;;\nNICKNAME;CHARSET=UTF-8:toby3d\n" +
				"GENDER:M\nURL;CHARSET=UTF-8:https://toby3d.me\nREV:2020-01-08T01:31:36.277Z\n" +
				"END:VCARD",
		},
		expResult: true,
	}, {
		name:      "false",
		expResult: false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.contact.HasVCard(), tc.expResult)
		})
	}
}

func TestPollVotesCount(t *testing.T) {
	for _, tc := range []struct {
		name      string
		poll      Poll
		expResult int
	}{{
		name: "true",
		poll: Poll{Options: []*PollOption{
			{Text: "a", VoterCount: 24},
			{Text: "b", VoterCount: 42},
		}},
		expResult: 66,
	}, {
		name: "true",
		poll: Poll{Options: []*PollOption{
			{Text: "a", VoterCount: 10},
			{Text: "b", VoterCount: 0},
			{Text: "c", VoterCount: 120},
		}},
		expResult: 130,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.poll.VotesCount(), tc.expResult)
		})
	}
}

func TestChatPhotoSmallFile(t *testing.T) {
	for _, tc := range []struct {
		name      string
		chatPhoto ChatPhoto
		expResult File
	}{{
		name:      "valid",
		chatPhoto: ChatPhoto{SmallFileID: "abc", SmallFileUniqueID: "really_abc"},
		expResult: File{FileID: "abc", FileUniqueID: "really_abc"},
	}, {
		name:      "empty",
		expResult: File{},
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.chatPhoto.SmallFile(), tc.expResult)
		})
	}
}

func TestChatPhotoBigFile(t *testing.T) {
	for _, tc := range []struct {
		name      string
		chatPhoto ChatPhoto
		expResult File
	}{{
		name:      "valid",
		chatPhoto: ChatPhoto{BigFileID: "abc", BigFileUniqueID: "really_abc"},
		expResult: File{FileID: "abc", FileUniqueID: "really_abc"},
	}, {
		name:      "empty",
		expResult: File{},
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.chatPhoto.BigFile(), tc.expResult)
		})
	}
}

func TestChatMemberIsAdministrator(t *testing.T) {
	for _, tc := range []struct {
		name       string
		chatMember ChatMember
		expResult  bool
	}{{
		name:       "true",
		chatMember: ChatMember{Status: StatusAdministrator},
		expResult:  true,
	}, {
		name:       "false",
		chatMember: ChatMember{Status: StatusCreator},
		expResult:  false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.chatMember.IsAdministrator(), tc.expResult)
		})
	}
}

func TestChatMemberIsCreator(t *testing.T) {
	for _, tc := range []struct {
		name       string
		chatMember ChatMember
		expResult  bool
	}{{
		name:       "true",
		chatMember: ChatMember{Status: StatusCreator},
		expResult:  true,
	}, {
		name:       "false",
		chatMember: ChatMember{Status: StatusAdministrator},
		expResult:  false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.chatMember.IsCreator(), tc.expResult)
		})
	}
}

func TestChatMemberIsKicked(t *testing.T) {
	for _, tc := range []struct {
		name       string
		chatMember ChatMember
		expResult  bool
	}{{
		name:       "true",
		chatMember: ChatMember{Status: StatusKicked},
		expResult:  true,
	}, {
		name:       "false",
		chatMember: ChatMember{Status: StatusMember},
		expResult:  false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.chatMember.IsKicked(), tc.expResult)
		})
	}
}

func TestChatMemberIsLeft(t *testing.T) {
	for _, tc := range []struct {
		name       string
		chatMember ChatMember
		expResult  bool
	}{{
		name:       "true",
		chatMember: ChatMember{Status: StatusLeft},
		expResult:  true,
	}, {
		name:       "false",
		chatMember: ChatMember{Status: StatusMember},
		expResult:  false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.chatMember.IsLeft(), tc.expResult)
		})
	}
}

func TestChatMemberIsRestricted(t *testing.T) {
	for _, tc := range []struct {
		name       string
		chatMember ChatMember
		expResult  bool
	}{{
		name:       "true",
		chatMember: ChatMember{Status: StatusRestricted},
		expResult:  true,
	}, {
		name:       "false",
		chatMember: ChatMember{Status: StatusMember},
		expResult:  false,
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.chatMember.IsRestricted(), tc.expResult)
		})
	}
}

func TestChatMemberUntilTime(t *testing.T) {
	now := time.Now().AddDate(0, 1, 0).Round(time.Second)

	for _, tc := range []struct {
		name       string
		chatMember ChatMember
		expResult  time.Time
	}{{
		name:       "valid",
		chatMember: ChatMember{Status: StatusMember, UntilDate: now.Unix()},
		expResult:  now,
	}, {
		name:       "empty",
		chatMember: ChatMember{Status: StatusMember},
		expResult:  time.Unix(0, 0),
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.chatMember.UntilTime(), tc.expResult)
		})
	}
}

func TestInputFileIsFileID(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		f := InputFile{ID: "abc"}
		assert.True(t, f.IsFileID())
	})
	t.Run("false", func(t *testing.T) {
		f := InputFile{}
		assert.False(t, f.IsFileID())
	})
}

func TestInputFileIsURI(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		u := http.AcquireURI()
		defer http.ReleaseURI(u)
		u.Update("https://toby3d.me/image.jpeg")
		f := InputFile{URI: u}
		assert.True(t, f.IsURI())
	})
	t.Run("false", func(t *testing.T) {
		f := InputFile{}
		assert.False(t, f.IsURI())
	})
}

func TestInputFileIsAttachment(t *testing.T) {
	file, err := ioutil.TempFile(os.TempDir(), "photo_*.jpeg")
	assert.NoError(t, err)

	defer os.RemoveAll(file.Name())

	t.Run("true", func(t *testing.T) {
		f := InputFile{Attachment: file}
		assert.True(t, f.IsAttachment())
	})
	t.Run("false", func(t *testing.T) {
		f := InputFile{}
		assert.False(t, f.IsAttachment())
	})
}

func TestInputFileMarshalJSON(t *testing.T) {
	u := http.AcquireURI()
	defer http.ReleaseURI(u)
	u.Update("https://toby3d.me/image.jpeg")

	file, err := ioutil.TempFile(os.TempDir(), "photo_*.jpeg")
	assert.NoError(t, err)

	defer os.RemoveAll(file.Name())

	_, fileName := filepath.Split(file.Name())

	for _, tc := range []struct {
		name      string
		inputFile InputFile
		expResult string
	}{{
		name:      "id",
		inputFile: InputFile{ID: "abc"},
		expResult: "abc",
	}, {
		name:      "uri",
		inputFile: InputFile{URI: u},
		expResult: u.String(),
	}, {
		name:      "attach",
		inputFile: InputFile{Attachment: file},
		expResult: SchemeAttach + "://" + fileName,
	}, {
		name:      "empty",
		inputFile: InputFile{},
		expResult: "",
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			src, err := tc.inputFile.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tc.expResult, string(src))
		})
	}
}
