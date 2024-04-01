package errors_test

import (
	"reflect"
	"testing"

	"github.com/ferrysutanto/go-errors"
)

func TestError_Errors(t *testing.T) {
	type fields []struct {
		msg  string
		code errors.ErrCode
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"success: 1", fields{{"Test 1", 500}}, "Test 1"},
		{"success: 2", fields{{"Root cause", 500}, {"Client", 500}}, "Client\nRoot cause"},
		{"success: 3", fields{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var e *errors.Error
			for _, f := range tt.fields {
				e = errors.WrapWithCode(e, f.msg, 500)
			}

			if got := e.Errors(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Error.Errors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Client(t *testing.T) {
	type fields []struct {
		msg  string
		code errors.ErrCode
	}
	type want struct {
		msg  string
		code errors.ErrCode
	}
	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{"success: 1", fields{{"Test 1", 500}}, want{"Test 1", 500}},
		{"success: 2", fields{{"Root cause", 400}, {"Client", 500}}, want{"Client", 500}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var e *errors.Error
			for _, f := range tt.fields {
				e = errors.WrapWithCode(e, f.msg, f.code)
			}

			got := errors.Client(e)
			if got == nil {
				t.Errorf("Error.Client() = nil, want %v", tt.want)
			}

			if got.Error() != tt.want.msg {
				t.Errorf("Error.Client() = %v, want %v", got.Error(), tt.want.msg)
			}

			if got.Code() != tt.want.code {
				t.Errorf("Error.Client() = %v, want %v", got.Code(), tt.want.code)
			}
		})
	}
}

func TestError_RootCause(t *testing.T) {
	type fields []struct {
		msg  string
		code errors.ErrCode
	}
	type want struct {
		msg  string
		code errors.ErrCode
	}
	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{"success: 1", fields{{"Test 1", 500}}, want{"Test 1", 500}},
		{"success: 2", fields{{"Root cause", 400}, {"Client", 500}}, want{"Root cause", 400}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var e *errors.Error
			for _, f := range tt.fields {
				e = errors.WrapWithCode(e, f.msg, f.code)
			}

			got := errors.RootCause(e)
			if got == nil {
				t.Errorf("Error.RootCause() = nil, want %v", tt.want)
			}

			if got.Error() != tt.want.msg {
				t.Errorf("Error.RootCause() = %v, want %v", got.Error(), tt.want.msg)
			}

			if got.Code() != tt.want.code {
				t.Errorf("Error.RootCause() = %v, want %v", got.Code(), tt.want.code)
			}
		})
	}
}
