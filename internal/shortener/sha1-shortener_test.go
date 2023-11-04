package shortener

import "testing"

func TestSha1Shortener_Shrink(t *testing.T) {
	type args struct {
		originalURL string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"Test case 1: https://nomadesoft.io",
			args{originalURL: "https://nomadesoft.io"},
			"346ed1f",
			false,
		},
		{
			"Test case 1: https://google.com",
			args{originalURL: "https://google.com"},
			"72fe95c",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSha1Shortener()
			got, err := s.Shrink(tt.args.originalURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("Shrink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Shrink() got = %v, want %v", got, tt.want)
			}
		})
	}
}
