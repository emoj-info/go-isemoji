package isemoji

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmoji(t *testing.T) {
	assert.Equal(t, true, IsEmoji("🤗"), "U+1F917 is an emoji")
	assert.Equal(t, true, IsEmoji("👨🏼‍🦰"), "1F468 F3FC 200D F9B0 is an emoji")
	assert.Equal(t, true, IsEmoji("🏻"), "U+1F3FB is a component emoji")
	assert.Equal(t, false, IsEmoji("👱‍♀"), "1F471 200D 2640 is a minimally-qualified emoji")
	assert.Equal(t, false, IsEmoji("test"), "the string \"test\" is not an emoji")
	assert.Equal(t, true, IsEmoji("🫩"), "Face with Bags Under Eyes")
	assert.Equal(t, true, IsEmoji("🫆"), "Fingerprint")
	assert.Equal(t, true, IsEmoji("🪾"), "Leafless Tree")
	assert.Equal(t, true, IsEmoji("🫜"), "Root Vegetable")
	assert.Equal(t, true, IsEmoji("🪉"), "Harp")
	assert.Equal(t, true, IsEmoji("🪏"), "Shovel")
	assert.Equal(t, true, IsEmoji("🫟"), "Splatter")
	assert.Equal(t, true, IsEmoji("🇨🇶"), "Flag: Sark")
}

func TestIsEmojiNonStrict(t *testing.T) {
	assert.Equal(t, true, IsEmojiNonStrict("🤗"), "U+1F917 is an emoji")
	assert.Equal(t, true, IsEmojiNonStrict("👨🏼‍🦰"), "1F468 F3FC 200D F9B0 is an emoji")
	assert.Equal(t, true, IsEmoji("🏻"), "U+1F3FB is a component emoji")
	assert.Equal(t, true, IsEmojiNonStrict("👱‍♀"), "1F471 200D 2640 is a minimally-qualified emoji")
	assert.Equal(t, true, IsEmojiNonStrict("☠"), "U+2620 is an unqualified emoji")
	assert.Equal(t, false, IsEmojiNonStrict("test"), "the string \"test\" is not an emoji")
}

func TestName(t *testing.T) {
	assert.Equal(t, "man: medium-light skin tone, red hair", Name("👨🏼‍🦰"), "that is the name")
	assert.Equal(t, "", Name("test"), "\"test\" is not an emoji with a name")
}

// TestLatest tests a newly-added emoji to ensure the proper version is being used.
func TestLatest(t *testing.T) {
	assert.Equal(t, true, IsEmoji("🫥"), "1FAE5 was added in Emoji 14.0")
	assert.Equal(t, "dotted line face", Name("🫥"), "that is the name of an emoji added in 14.0")
}
