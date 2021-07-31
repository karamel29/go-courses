package rectangle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// rectangle area tests
func TestRectangle_Area(t *testing.T) {
	type fields struct {
		height float64
		width  float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name: "when value of height equal 0",
			fields: fields{
				height: 0,
				width:  1,
			},
			want:    float64(0),
			wantErr: true,
		},

		{
			name: "when value of width equal 0",
			fields: fields{
				height: 1,
				width:  0,
			},
			want:    float64(0),
			wantErr: true,
		},

		{
			name: "when value of height equal -1",
			fields: fields{
				height: -1,
				width:  1,
			},
			want:    float64(0),
			wantErr: true,
		},

		{
			name: "when value of width equal -1",
			fields: fields{
				height: 1,
				width:  -1,
			},
			want:    float64(0),
			wantErr: true,
		},

		{
			name: "when value equal 1",
			fields: fields{
				height: 1,
				width:  1,
			},
			want:    float64(1),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				Height: tt.fields.height,
				Width:  tt.fields.width,
			}
			got, err := r.Area()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

// rectangle perimeter tests
func TestRectangle_Perimeter(t *testing.T) {
	type fields struct {
		height float64
		width  float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name: "when value of height equal 0",
			fields: fields{
				height: 0,
				width:  1,
			},
			want:    float64(0),
			wantErr: true,
		},

		{
			name: "when value of width equal 0",
			fields: fields{
				height: 1,
				width:  0,
			},
			want:    float64(0),
			wantErr: true,
		},

		{
			name: "when value of height equal -1",
			fields: fields{
				height: -1,
				width:  1,
			},
			want:    float64(0),
			wantErr: true,
		},

		{
			name: "when value of width equal -1",
			fields: fields{
				height: 1,
				width:  -1,
			},
			want:    float64(0),
			wantErr: true,
		},

		{
			name: "when value equal 1",
			fields: fields{
				height: 1,
				width:  1,
			},
			want:    float64(4),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				Height: tt.fields.height,
				Width:  tt.fields.width,
			}
			got, err := r.Perimeter()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

// test for string when output height and width rectangle
func TestRectangle_String(t *testing.T) {
	type fields struct {
		height float64
		width  float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "whether there is a value height in string",
			fields: fields{
				height: 1,
			},
			want: "",
		},

		{
			name: "whether there is a value width in string",
			fields: fields{
				width: 1,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				Height: tt.fields.height,
				Width:  tt.fields.width,
			}
			got := r.String()
			h := fmt.Sprintf("%.2f", r.Height)
			require.Contains(t, got, h)
			w := fmt.Sprintf("%.2f", r.Width)
			require.Contains(t, got, w)
		})
	}
}