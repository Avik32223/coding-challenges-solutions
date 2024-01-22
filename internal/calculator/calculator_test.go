package calculator

import "testing"

func TestCalculate(t *testing.T) {
	tests := []struct {
		name    string
		source  string
		want    string
		wantErr bool
	}{
		{"test 1", "1 + 2", "3", false},
		{"test 2", "2 - 1", "1", false},
		{"test 3", "2 * 3", "6", false},
		{"test 4", "3 / 2", "1.50", false},
		{"test 5", "1 + 1 * 5", "6", false},
		{"test 6", "( 1 + 1 ) * 5", "10", false},
		{"test 7", "1 + ( 4  * 5 )", "21", false},
		{"test 8", "1 + 2 ^ 2 ^ 3", "257", false},
		{"test 9", "1 + 2 ^ 2 * 3", "13", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
