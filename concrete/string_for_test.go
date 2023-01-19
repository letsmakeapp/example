package concrete

import "testing"

func TestForLoopStringLengthCounter_Count(t *testing.T) {
	const (
		InputString    = "HELLO WORLD!"
		ExpectedLength = 12
	)

	counter := ForLoopStringLengthCounter{}
	got := counter.Count(InputString)

	if got != ExpectedLength {
		t.Errorf("expect a string with length %v but got %v", ExpectedLength, got)
	}
}

func TestForLoopStringLengthCounter_Count_Empty(t *testing.T) {
	const (
		InputString    = ""
		ExpectedLength = 0
	)

	counter := ForLoopStringLengthCounter{}
	got := counter.Count(InputString)

	if got != ExpectedLength {
		t.Errorf("expect a string with length %v but got %v", ExpectedLength, got)
	}
}

func TestForLoopStringLengthCounter_Count_TableTest(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		f    ForLoopStringLengthCounter
		args args
		want uint
	}{
		{
			name: "non empty string",
			f:    ForLoopStringLengthCounter{},
			args: args{
				s: "HELLO WORLD!",
			},
			want: 12,
		},
		{
			name: "empty string",
			f:    ForLoopStringLengthCounter{},
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "only whitespace",
			f:    ForLoopStringLengthCounter{},
			args: args{
				s: "    ",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := ForLoopStringLengthCounter{}
			if got := f.Count(tt.args.s); got != tt.want {
				t.Errorf("ForLoopStringLengthCounter.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
