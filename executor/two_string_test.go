package executor

import (
	abstract_mock "example/abstract/mock"
	"example/concrete"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestTwoStringLengthCounter_Count(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name    string
		prepare func(m *abstract_mock.MockStringLengthCounter)
		args    args
		want    uint
	}{
		{
			name: "should call Count correctly",
			prepare: func(m *abstract_mock.MockStringLengthCounter) {
				m.EXPECT().Count("HELLO").Times(1).Return(uint(5))
				m.EXPECT().Count("WORLD").Times(1).Return(uint(5))
			},
			args: args{
				a: "HELLO",
				b: "WORLD",
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mocked := abstract_mock.NewMockStringLengthCounter(ctrl)
			c := TwoStringLengthCounter{
				Counter: mocked,
			}

			tt.prepare(mocked)
			if got := c.Count(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("TwoStringLengthCounter.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoStringLengthCounter_Count_2(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "should count correctly",
			args: args{
				a: "HELLO",
				b: "WORLD",
			},
			want: 10,
		},
		{
			name: "should count correctly on empty case",
			args: args{
				a: "",
				b: "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := TwoStringLengthCounter{
				Counter: concrete.NativeStringLengthCounter{},
			}
			if got := c.Count(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("TwoStringLengthCounter.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
