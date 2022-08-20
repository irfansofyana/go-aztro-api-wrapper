package aztro_test

import (
	"net/http"
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

func Test_DayString(t *testing.T) {
	testCases := []struct {
		day aztro.Day
		str string
	}{
		{day: aztro.Yesterday, str: "yesterday"},
		{day: aztro.Today, str: "today"},
		{day: aztro.Tomorrow, str: "tomorrow"},
	}

	for _, testCase := range testCases {
		got := testCase.day.String()
		if got != testCase.str {
			t.Errorf("expected: %v, got: %v", testCase.str, got)
		}
	}
}

func Test_NewAztroClient(t *testing.T) {
	client, err := aztro.NewAztroClient()
	if err != nil {
		t.Errorf("expected to not have error when creating aztro client but received error: %v", err)
	}

	if client.HTTPReq.Method != http.MethodPost {
		t.Errorf("expected to use POST method in API request but got: %v", client.HTTPReq.Method)
	}

	if client.HTTPReq.URL.String() != aztro.AztroBaseURL {
		t.Errorf("expected to use Base URL: %v in API request but got: %v", client.HTTPReq.URL.String(), aztro.AztroBaseURL)
	}
}
