package tests

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetApiKey(t *testing.T) {

	cases := []struct {
		want    string
		headers http.Header
		err     error
	}{
		{
			want:    "123456",
			headers: http.Header{"Authorization": []string{"ApiKey 123456"}},
			err:     nil,
		},
		{
			want:    "",
			headers: http.Header{"Authorization": []string{""}},
			err:     auth.ErrNoAuthHeaderIncluded,
		},
	}
    for _, c := range cases {
        got, err := auth.GetAPIKey(c.headers)
        if got != c.want {
            t.Fatalf("GetAPIKey(%v) == %v, want %v", c.headers, got, c.want)
        }
        if err != c.err {
            t.Fatalf("GetAPIKey(%v) == %v, want %v", c.headers, err, c.err)
        }
    }

}

