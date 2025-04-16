package main

import (
	"snippetbox.kakani21.net/internal/assert"
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {

	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2024, 12, 21, 12, 21, 0, 0, time.UTC),
			want: "21 Dec 2024 at 12:21",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2024, 12, 21, 12, 21, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "21 Dec 2024 at 11:21",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			assert.Equal(t, hd, tt.want)

		})
	}

}
