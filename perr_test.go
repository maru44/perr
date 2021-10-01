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
		out              error
		lastMsgForClient string
	}{
		{
			errors.New("test01"),
			BadRequest,
			nil,
			errors.New("test01"),
			BadRequest,
			BadRequest.Error(),
		},
		{
			nil,
			BadGatewayWithUrgency,
			nil,
			nil,
			nil,
			"",
		},
		{
			errors.New("test03"),
			BadGateway,
			[]string{"message"},
			errors.New("test03"),
			BadGateway,
			strings.Repeat("message\n", 9) + "message",
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
					ps = append(ps, Wrap(ps[i-1], BadGateway, tt.msg...))
				}
			}

			if ps[9] != nil && !reflect.ValueOf(ps[9]).IsNil() {
				if tt.perrErr.Error() != ps[9].Error() {
					t.Errorf("want: %v\ngot: %v", tt.perrErr, ps[9].Unwrap())
				}
				if tt.lastMsgForClient != ps[9].Output().Error() {
					t.Errorf("want: %v\ngot: %v", tt.lastMsgForClient, ps[9].Output())
				}
				if ps[9].Traces().maxLayer() != 9 {
					t.Errorf("want: %v\ngot: %v", 9, ps[9].Traces().maxLayer())
				}
				if tt.lastMsgForClient != ps[9].msgForClient {
					t.Errorf("want: %v\ngot: %v", tt.lastMsgForClient, ps[9].msgForClient)
				}
			} else {
				if ps[9] != nil || !reflect.ValueOf(ps[9]).IsNil() {
					t.Errorf("want: nil\ngot: %v", ps[9])
				}
			}
		})
	}
}
