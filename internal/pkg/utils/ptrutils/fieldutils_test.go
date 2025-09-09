// Basic data type conversion and pointer type conversion
package ptrutils

import (
	"fmt"
	"testing"
)

func TestIntToPtr(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal process",
			args: args{
				i: int(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IntToPtr(tt.args.i)
		})
	}
}

func TestInt8ToPtr(t *testing.T) {
	type args struct {
		i int8
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal process",
			args: args{
				i: int8(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Int8ToPtr(tt.args.i)
		})
	}
}

func TestInt16ToPtr(t *testing.T) {
	type args struct {
		i int16
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal process",
			args: args{
				i: int16(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Int16ToPtr(tt.args.i)
		})
	}
}

func TestInt32ToPtr(t *testing.T) {
	type args struct {
		i int32
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal process",
			args: args{
				i: int32(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Int32ToPtr(tt.args.i)
		})
	}
}

func TestInt64ToPtr(t *testing.T) {
	type args struct {
		i int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal process",
			args: args{
				i: int64(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Int64ToPtr(tt.args.i)
		})
	}
}

func TestUintToPtr(t *testing.T) {
	type args struct {
		i uint
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal process",
			args: args{
				i: uint(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UintToPtr(tt.args.i)
		})
	}
}

func TestUint8ToPtr(t *testing.T) {
	type args struct {
		i uint8
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal process",
			args: args{
				i: uint8(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Uint8ToPtr(tt.args.i)
		})
	}
}

func TestUint16ToPtr(t *testing.T) {
	type args struct {
		i uint16
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal process",
			args: args{
				i: uint16(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Uint16ToPtr(tt.args.i)
		})
	}
}

func TestUint32ToPtr(t *testing.T) {
	type args struct {
		i uint32
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal process",
			args: args{
				i: uint32(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Uint32ToPtr(tt.args.i)
		})
	}
}

func TestUint64ToPtr(t *testing.T) {
	type args struct {
		i uint64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal process",
			args: args{
				i: uint64(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Uint64ToPtr(tt.args.i)
		})
	}
}

func TestFloat32ToPtr(t *testing.T) {
	type args struct {
		f float32
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal process",
			args: args{
				f: float32(1.0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Float32ToPtr(tt.args.f)
		})
	}
}

func TestFloat64ToPtr(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal process",
			args: args{
				f: float64(1.0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Float64ToPtr(tt.args.f)
		})
	}
}

func TestStringToPtr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal process",
			args: args{
				s: "s",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StringToPtr(tt.args.s)
		})
	}
}

func TestBoolToPtr(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal process",
			args: args{
				b: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BoolToPtr(tt.args.b)
		})
	}
}

func TestToInt(t *testing.T) {
	type args struct {
		i *int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal process",
			args: args{
				i: IntToPtr(1),
			},
			want: int(1),
		}, {
			name: "parameter nil",
			args: args{
				i: nil,
			},
			want: int(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt(tt.args.i); got != tt.want {
				t.Errorf("ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt8(t *testing.T) {
	type args struct {
		i *int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{
			name: "normal process",
			args: args{
				i: Int8ToPtr(1),
			},
			want: int8(1),
		}, {
			name: "parameter nil",
			args: args{
				i: nil,
			},
			want: int8(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt8(tt.args.i); got != tt.want {
				t.Errorf("ToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt16(t *testing.T) {
	type args struct {
		i *int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			name: "normal process",
			args: args{
				i: Int16ToPtr(1),
			},
			want: int16(1),
		}, {
			name: "parameter nil",
			args: args{
				i: nil,
			},
			want: int16(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt16(tt.args.i); got != tt.want {
				t.Errorf("ToInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt32(t *testing.T) {
	type args struct {
		i *int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "normal process",
			args: args{
				i: Int32ToPtr(1),
			},
			want: int32(1),
		}, {
			name: "parameter nil",
			args: args{
				i: nil,
			},
			want: int32(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt32(tt.args.i); got != tt.want {
				t.Errorf("ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt64(t *testing.T) {
	type args struct {
		i *int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "normal process",
			args: args{
				i: Int64ToPtr(1),
			},
			want: int64(1),
		}, {
			name: "parameter nil",
			args: args{
				i: nil,
			},
			want: int64(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt64(tt.args.i); got != tt.want {
				t.Errorf("ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint(t *testing.T) {
	type args struct {
		i *uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "normal process",
			args: args{
				i: UintToPtr(1),
			},
			want: uint(1),
		}, {
			name: "parameter nil",
			args: args{
				i: nil,
			},
			want: uint(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint(tt.args.i); got != tt.want {
				t.Errorf("ToUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint8(t *testing.T) {
	type args struct {
		i *uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			name: "normal process",
			args: args{
				i: Uint8ToPtr(1),
			},
			want: uint8(1),
		}, {
			name: "parameter nil",
			args: args{
				i: nil,
			},
			want: uint8(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint8(tt.args.i); got != tt.want {
				t.Errorf("ToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint16(t *testing.T) {
	type args struct {
		i *uint16
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "normal process",
			args: args{
				i: Uint16ToPtr(1),
			},
			want: uint16(1),
		}, {
			name: "parameter nil",
			args: args{
				i: nil,
			},
			want: uint16(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint16(tt.args.i); got != tt.want {
				t.Errorf("ToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint32(t *testing.T) {
	type args struct {
		i *uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "normal process",
			args: args{
				i: Uint32ToPtr(1),
			},
			want: uint32(1),
		}, {
			name: "parameter nil",
			args: args{
				i: nil,
			},
			want: uint32(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint32(tt.args.i); got != tt.want {
				t.Errorf("ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint64(t *testing.T) {
	type args struct {
		i *uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "normal process",
			args: args{
				i: Uint64ToPtr(1),
			},
			want: uint64(1),
		}, {
			name: "parameter nil",
			args: args{
				i: nil,
			},
			want: uint64(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint64(tt.args.i); got != tt.want {
				t.Errorf("ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat32(t *testing.T) {
	var num float32
	fmt.Println(num)
	type args struct {
		f *float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "normal process",
			args: args{
				f: Float32ToPtr(1.0),
			},
			want: float32(1.0),
		}, {
			name: "parameter nil",
			args: args{
				f: nil,
			},
			want: float32(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFloat32(tt.args.f); got != tt.want {
				t.Errorf("ToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat64(t *testing.T) {
	type args struct {
		f *float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "normal process",
			args: args{
				f: Float64ToPtr(1.0),
			},
			want: float64(1.0),
		}, {
			name: "parameter nil",
			args: args{
				f: nil,
			},
			want: float64(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFloat64(tt.args.f); got != tt.want {
				t.Errorf("ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToString(t *testing.T) {
	type args struct {
		s *string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal process",
			args: args{
				s: StringToPtr("s"),
			},
			want: "s",
		}, {
			name: "parameter nil",
			args: args{
				s: nil,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToString(tt.args.s); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBool(t *testing.T) {
	type args struct {
		b *bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal process",
			args: args{
				b: BoolToPtr(true),
			},
			want: true,
		}, {
			name: "parameter nil",
			args: args{
				b: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBool(tt.args.b); got != tt.want {
				t.Errorf("ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
