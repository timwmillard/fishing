package fishing

import (
	"strconv"
	"testing"
)

func TestNewHashID(t *testing.T) {

	tests := []struct {
		id   HashID
		want string
	}{
		{id: 0, want: "1xv6o52lpj"},
		{id: 1, want: "m3oq3e687p"},
		{id: 123, want: "l95q8jy6nm"},
		{id: 100000, want: "xv6o30vk6l"},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(int(tt.id)), func(t *testing.T) {
			hid := HashID(tt.id)
			got, _ := hid.Hash()
			if got != tt.want {
				t.Errorf("NewHashID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashID_ID(t *testing.T) {
	tests := []struct {
		hash string
		want HashID
	}{
		{hash: "1xv6o52lpj", want: 0},
		{hash: "m3oq3e687p", want: 1},
		{hash: "l95q8jy6nm", want: 123},
		{hash: "xv6o30vk6l", want: 100000},
	}
	for _, tt := range tests {
		t.Run(string(tt.hash), func(t *testing.T) {
			got, _ := NewHashID(tt.hash)
			if got != tt.want {
				t.Errorf("HashID.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}
