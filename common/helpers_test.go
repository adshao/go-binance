package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAmountToLotSize(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		lot       float64
		precision int
		amount    float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test with lot of zero and invalid amount",
			args: args{
				lot:       0.00100000,
				precision: 8,
				amount:    0.00010000,
			},
			want: 0,
		},
		{
			name: "test with lot",
			args: args{
				lot:       0.00100000,
				precision: 3,
				amount:    1.39,
			},
			want: 1.389,
		},
		{
			name: "test with big decimal",
			args: args{
				lot:       0.00100000,
				precision: 8,
				amount:    11.31232419283240912834434,
			},
			want: 11.312,
		},
		{
			name: "test with big number",
			args: args{
				lot:       0.0010000,
				precision: 8,
				amount:    11232821093480213.31232419283240912834434,
			},
			want: 11232821093480213.3123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(tt.want, AmountToLotSize(tt.args.lot, tt.args.precision, tt.args.amount))
		})
	}
}
