package perr

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

var (
	wrap_table = []struct {
		// input
		in  error
		as  error
		msg []string
		// output
		perrErr          error
		outAs            error
		lastMsgForClient string
	}{
		{
			errors.New("test01"),
			ErrBadRequest,
			nil,
			errors.New("test01"),
			ErrBadRequest,
			ErrBadRequest.Error(),
		},
		{
			nil,
			ErrBadGatewayWithUrgency,
			nil,
			nil,
			nil,
			"",
		},
		{
			errors.New("test03"),
			ErrBadGateway,
			[]string{"message"},
			errors.New("test03"),
			ErrBadGateway,
			strings.Repeat("message\n", 9) + "message",
		},
	}

	newTable = []struct {
		// input
		text         string
		as           error
		msgForClient []string
		// output
		perrErr          error
		outAs            error
		lastMsgForClient string
	}{
		{
			"new error(developer)",
			ErrInternalServerError,
			nil,
			errors.New("new error(developer)"),
			ErrInternalServerError,
			ErrInternalServerError.Error(),
		},
		{
			"",
			ErrNotFound,
			nil,
			ErrNotFound,
			ErrNotFound,
			ErrNotFound.Error(),
		},
	}
)

func TestWrapPerr(t *testing.T) {
	for i, tt := range wrap_table {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var ps []*Err
			for i := 0; i < 10; i++ {
				if len(ps) == 0 {
					ps = append(ps, Wrap(tt.in, tt.as, tt.msg...))
				} else {
					ps = append(ps, Wrap(ps[i-1], ErrBadGateway, tt.msg...))
				}
			}

			if ps[9] != nil && !reflect.ValueOf(ps[9]).IsNil() {
				// whether error for developer is wrapped correctly
				if tt.perrErr.Error() != ps[9].Error() {
					t.Errorf("want: %v\ngot: %v", tt.perrErr, ps[9].Unwrap())
				}
				// if tt.outAs != ps[9].as {
				// 	t.Errorf("want: %v\ngot: %v", tt.outAs, ps[9].as)
				// }

				// whether output message is wrapped correctly (under condition with msgForClient)
				if tt.lastMsgForClient != ps[9].Output().Error() {
					t.Errorf("want: %v\ngot: %v", tt.lastMsgForClient, ps[9].Output())
				}
				if ps[9].Traces().maxLayer() != 9 {
					t.Errorf("want: %v\ngot: %v", 9, ps[9].Traces().maxLayer())
				}
				// whether output error is wrapped correctly (under condition witout msgForClient)
				if !ps[9].IsOutput(tt.outAs) {
					t.Errorf("want: %v\ngot: %v", tt.outAs, ps[9].as)
				}
			} else {
				if ps[9] != nil || !reflect.ValueOf(ps[9]).IsNil() {
					t.Errorf("want: nil\ngot: %v", ps[9])
				}
			}
		})
	}
}

func TestNewPerr(t *testing.T) {
	for i, tt := range newTable {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			p := New(tt.text, tt.as)
			// dev
			if tt.perrErr.Error() != p.Unwrap().Error() {
				t.Errorf("want: %v\ngot: %v", tt.perrErr, p.Unwrap())
			}
			// client
			if tt.outAs.Error() != p.as.Error() {
				t.Errorf("want: %v\ngot: %v", tt.outAs, p.as)
			}
			if tt.lastMsgForClient != p.Output().Error() {
				t.Errorf("want: %v\ngot: %v", tt.lastMsgForClient, p.msgForClient)
			}
		})
	}
}
