package shape

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/karamel29/go-courses/task005/shape/circle"
	"github.com/karamel29/go-courses/task005/shape/rectangle"
)

func TestDescribeShape(t *testing.T) {
	tests := []struct {
		name    string
		shape   Shape
		wantErr bool
	}{
		{
			name: "success for rectangle",
			shape: rectangle.Rectangle{
				Height: 1,
				Width:  1,
			},
			wantErr: false,
		},

		{
			name: "success for circle",
			shape: circle.Circle{
				Radius: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := DescribeShape(tt.shape)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
