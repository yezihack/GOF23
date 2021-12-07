package simple_factory

import (
	"reflect"
	"testing"
)

func TestNewTransport(t *testing.T) {
	type args struct {
		tool Tool
	}
	tests := []struct {
		name string
		args args
		want Transport
	}{
		{
			name: "truck",
			args: args{tool: TruckTool},
			want: &Truck{},
		},
		{
			name: "ship",
			args: args{tool: ShipTool},
			want: &Ship{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTransport(tt.args.tool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTransport() = %v, want %v", got, tt.want)
			}
		})
	}
}
