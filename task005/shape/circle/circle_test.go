package circle

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

// circle perimeter tests
func TestCircle_Perimeter(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name: "when value equal 0",
			fields: fields{
				radius: 0,
			},
			want:    float64(0),
			wantErr: true,
		},

		{
			name: "when value equal -1",
			fields: fields{
				radius: -1,
			},
			want:    float64(0),
			wantErr: true,
		},

		{
			name: "when value equal 1",
			fields: fields{
				radius: 1,
			},
			want:    float64(6.283185307179586),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				Radius: tt.fields.radius,
			}
			got, err := c.Perimeter()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

// circle area tests
func TestCircle_Area(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name: "when value equal 0",
			fields: fields{
				radius: 0,
			},
			want:    float64(0),
			wantErr: true,
		},

		{
			name: "when value equal -1",
			fields: fields{
				radius: -1,
			},
			want:    float64(0),
			wantErr: true,
		},

		{
			name: "when value equal 1",
			fields: fields{
				radius: 1,
			},
			want:    float64(math.Pi),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				Radius: tt.fields.radius,
			}
			got, err := c.Area()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

// test for string when output radius circle
func TestCircle_String(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "whether there is a value radius in string",
			fields: fields{
				radius: 0,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				Radius: tt.fields.radius,
			}
			got := c.String()
			s := fmt.Sprintf("%.2f", c.Radius)
			require.Contains(t, got, s)
		})
	}
}