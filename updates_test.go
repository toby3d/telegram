package telegram

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	http "github.com/valyala/fasthttp"
)

func TestUpdateIsMessage(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		u := Update{Message: &Message{ID: 42}}
		assert.True(t, u.IsMessage())
	})
	t.Run("false", func(t *testing.T) {
		u := Update{}
		assert.False(t, u.IsMessage())
	})
}

func TestUpdateIsEditedMessage(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		u := Update{EditedMessage: &Message{ID: 42}}
		assert.True(t, u.IsEditedMessage())
	})
	t.Run("false", func(t *testing.T) {
		u := Update{}
		assert.False(t, u.IsEditedMessage())
	})
}

func TestUpdateIsChannelPost(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		u := Update{ChannelPost: &Message{ID: 42}}
		assert.True(t, u.IsChannelPost())
	})
	t.Run("false", func(t *testing.T) {
		u := Update{}
		assert.False(t, u.IsChannelPost())
	})
}

func TestUpdateIsEditedChannelPost(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		u := Update{EditedChannelPost: &Message{ID: 42}}
		assert.True(t, u.IsEditedChannelPost())
	})
	t.Run("false", func(t *testing.T) {
		u := Update{}
		assert.False(t, u.IsEditedChannelPost())
	})
}

func TestUpdateIsInlineQuery(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		u := Update{InlineQuery: &InlineQuery{Query: "abc"}}
		assert.True(t, u.IsInlineQuery())
	})
	t.Run("false", func(t *testing.T) {
		u := Update{}
		assert.False(t, u.IsInlineQuery())
	})
}

func TestUpdateIsChosenInlineResult(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		u := Update{ChosenInlineResult: &ChosenInlineResult{ResultID: "abc"}}
		assert.True(t, u.IsChosenInlineResult())
	})
	t.Run("false", func(t *testing.T) {
		u := Update{}
		assert.False(t, u.IsChosenInlineResult())
	})
}

func TestUpdateIsCallbackQuery(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		u := Update{CallbackQuery: &CallbackQuery{ID: "abc"}}
		assert.True(t, u.IsCallbackQuery())
	})
	t.Run("false", func(t *testing.T) {
		u := Update{}
		assert.False(t, u.IsCallbackQuery())
	})
}

func TestUpdateIsShippingQuery(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		u := Update{ShippingQuery: &ShippingQuery{ID: "abc"}}
		assert.True(t, u.IsShippingQuery())
	})
	t.Run("false", func(t *testing.T) {
		u := Update{}
		assert.False(t, u.IsShippingQuery())
	})
}

func TestUpdateIsPreCheckoutQuery(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		u := Update{PreCheckoutQuery: &PreCheckoutQuery{ID: "abc"}}
		assert.True(t, u.IsPreCheckoutQuery())
	})
	t.Run("false", func(t *testing.T) {
		u := Update{}
		assert.False(t, u.IsPreCheckoutQuery())
	})
}

func TestUpdateIsPoll(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		u := Update{Poll: &Poll{ID: "abc"}}
		assert.True(t, u.IsPoll())
	})
	t.Run("false", func(t *testing.T) {
		u := Update{}
		assert.False(t, u.IsPoll())
	})
}

func TestUpdateType(t *testing.T) {
	for _, tc := range []struct {
		name      string
		update    Update
		expResult string
	}{{
		name:      UpdateCallbackQuery,
		update:    Update{CallbackQuery: &CallbackQuery{ID: "abc"}},
		expResult: UpdateCallbackQuery,
	}, {
		name:      UpdateChannelPost,
		update:    Update{ChannelPost: &Message{ID: 42}},
		expResult: UpdateChannelPost,
	}, {
		name:      UpdateChosenInlineResult,
		update:    Update{ChosenInlineResult: &ChosenInlineResult{Query: "query"}},
		expResult: UpdateChosenInlineResult,
	}, {
		name:      UpdateEditedChannelPost,
		update:    Update{EditedChannelPost: &Message{ID: 42}},
		expResult: UpdateEditedChannelPost,
	}, {
		name:      UpdateEditedMessage,
		update:    Update{EditedMessage: &Message{ID: 42}},
		expResult: UpdateEditedMessage,
	}, {
		name:      UpdateInlineQuery,
		update:    Update{InlineQuery: &InlineQuery{ID: "abc"}},
		expResult: UpdateInlineQuery,
	}, {
		name:      UpdateMessage,
		update:    Update{Message: &Message{ID: 42}},
		expResult: UpdateMessage,
	}, {
		name:      UpdatePoll,
		update:    Update{Poll: &Poll{ID: "abc"}},
		expResult: UpdatePoll,
	}, {
		name:      UpdatePreCheckoutQuery,
		update:    Update{PreCheckoutQuery: &PreCheckoutQuery{ID: "abc"}},
		expResult: UpdatePreCheckoutQuery,
	}, {
		name:      UpdateShippingQuery,
		update:    Update{ShippingQuery: &ShippingQuery{ID: "abc"}},
		expResult: UpdateShippingQuery,
	}, {
		name:      "other",
		update:    Update{},
		expResult: "",
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.update.Type(), tc.expResult)
		})
	}
}

func TestWebhookInfoLastErrorTime(t *testing.T) {
	now := time.Now().Round(time.Second)
	wi := WebhookInfo{LastErrorDate: now.Unix()}
	assert.Equal(t, wi.LastErrorTime(), now)
}

func TestWebhookInfoHasURL(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		wi := WebhookInfo{URL: "https://bot.toby3d.me"}
		assert.True(t, wi.HasURL())
	})
	t.Run("false", func(t *testing.T) {
		wi := WebhookInfo{}
		assert.False(t, wi.HasURL())
	})
}

func TestWebhookInfoURI(t *testing.T) {
	u := http.AcquireURI()
	defer http.ReleaseURI(u)

	u.Update("https://bot.toby3d.me")
	t.Run("true", func(t *testing.T) {
		wi := WebhookInfo{URL: u.String()}
		assert.Equal(t, wi.URI().String(), u.String())
	})
	t.Run("false", func(t *testing.T) {
		wi := WebhookInfo{}
		assert.Nil(t, wi.URI())
	})
}
