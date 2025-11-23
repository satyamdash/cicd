package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr bool
	}{
		{
			name:    "Valid ApiKey",
			headers: http.Header{"Authorization": []string{"ApiKey mykey"}},
			want:    "mykey",
		},
		{
			name:    "missing header",
			headers: http.Header{},
			wantErr: true,
		}, {
			name:    "wrong scheme",
			headers: http.Header{"Authorization": []string{"ApKey xyz"}},
			wantErr: true,
		}, {
			name:    "empty key",
			headers: http.Header{"Authorization": []string{"ApiKey"}},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetAPIKey(tc.headers)
			if tc.wantErr && err == nil {
				t.Fatalf("expected error but got none")
			}
			if !tc.wantErr && got != tc.want {
				t.Fatalf("expected '%s', got '%s'", tc.want, got)
			}
		})
	}
}
