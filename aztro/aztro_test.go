package aztro_test

import (
	"testing"

	"github.com/irfansofyana/go-aztro-api-wrapper/aztro"
)

func Test_SignString(t *testing.T) {
	testCases := []struct {
		sign aztro.Sign
		str  string
	}{
		{sign: aztro.Aries, str: "aries"},
		{sign: aztro.Aquarius, str: "aquarius"},
		{sign: aztro.Cancer, str: "cancer"},
		{sign: aztro.Capricorn, str: "capricorn"},
		{sign: aztro.Gemini, str: "gemini"},
		{sign: aztro.Leo, str: "leo"},
		{sign: aztro.Libra, str: "libra"},
		{sign: aztro.Pisces, str: "pisces"},
		{sign: aztro.Sagittarius, str: "sagittarius"},
		{sign: aztro.Scorpio, str: "scorpio"},
		{sign: aztro.Taurus, str: "taurus"},
		{sign: aztro.Virgo, str: "virgo"},
	}

	for _, testCase := range testCases {
		got := testCase.sign.String()
		if got != testCase.str {
			t.Errorf("expected: %v, got: %v", testCase.str, got)
		}
	}
}
