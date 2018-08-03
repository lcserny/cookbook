package confformat

import "testing"

func TestUnmarshalAll(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UnMarshallAll(); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
