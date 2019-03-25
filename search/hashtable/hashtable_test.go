package hashtable

import (
	"testing"
)

func TestLinearProbing(t *testing.T) {
	ht := NewLinearProbing(WithTableSize(15))
	keys := []string{
		"fizz", "bazz", "BBBBBB", "AaAaAa", "BBAaBB", "AaBBBB",
		"aoffckzdaoffckzdegjuqmpg", "aoffckzdaoffckzdelevflik",
		"aoffckzdaoffckzdmttyjiqf", "aoffckzdatafwjshdiiqutzn",
	}
	for i, s := range keys {
		ht.Put(s, i)
	}
	for i, s := range keys {
		got := ht.Get(s)
		if got != i {
			t.Errorf("Get(%q) = %d, want %d", s, got, i)
		}
	}
}
