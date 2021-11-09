package secret

import "testing"

func TestGetSecretAnswer(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "Secret Test",
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSecretAnswer(); got != tt.want {
				t.Errorf("GetSecretAnswer() = %v, want %v", got, tt.want)
			}
		})
	}
}
