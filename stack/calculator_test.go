package stack

import "testing"

func Test_calculatorByStack(t *testing.T) {
	type args struct {
		formula string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "calculator_1", args: args{formula: "1+2-3*4"}, want: -9},
		{name: "calculator_2", args: args{formula: "1-2+3*4"}, want: 11},
		{name: "calculator_3", args: args{formula: "1-2-3*4-5/10"}, want: -13.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculatorByStack(tt.args.formula)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculatorByStack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calculatorByStack() got = %v, want %v", got, tt.want)
			}
		})
	}
}
