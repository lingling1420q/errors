package errors

import (
	"testing"
)

func TestCauseFrom(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"cause from origin",
			args{
				New("origin"),
			},
			true,
		},
		{
			"cause from error",
			args{
				Annotate(New("origin"), "cause"),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CauseFrom(tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("CauseFrom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValueFrom(t *testing.T) {
	type fields struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"normal error",
			fields{
				WithCode(New("error reason"), &errorCode{1, "NumOneErr"}),
			},
			1,
		},
		{
			"code error",
			fields{
				CodeError(&errorCode{2, "NumTwoErr"}),
			},
			2,
		},
		{
			"no code error",
			fields{
				New("error"),
			},
			-1,
		},
		{
			"no error",
			fields{
				nil,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ce := tt.fields.err
			if got := ValueFrom(ce); got != tt.want {
				t.Errorf("ValueFrom error = %v, want %v", got, tt.want)
			}
		})
	}
}