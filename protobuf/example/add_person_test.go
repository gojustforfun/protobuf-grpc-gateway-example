package example_test

import (
	"bufio"
	"io"
	"strings"
	"testing"

	pb "github.com/gojustforfun/protobuf-grpc-gateway-example/protobuf/example"
	"github.com/stretchr/testify/assert"
)

func Test_PromptForAddress(t *testing.T) {

	normal := `12345
Example Name
name@example.com
123-456-7890
home
222-222-2222
mobile
111-111-1111
work
777-777-7777
unknown

`

	testcases := map[string]struct {
		r    io.Reader
		want *pb.Person
	}{
		"normal": {
			r: bufio.NewReader(strings.NewReader(normal)),
			want: &pb.Person{
				Id:    int32(12345),
				Name:  "Example Name",
				Email: "name@example.com",
				Phones: []*pb.Person_PhoneNumber{
					{
						Number: "123-456-7890",
						Type:   pb.Person_HOME,
					},
					{
						Number: "222-222-2222",
						Type:   pb.Person_MOBILE,
					},
					{
						Number: "111-111-1111",
						Type:   pb.Person_WORK,
					},
					{
						Number: "777-777-7777",
						Type:   pb.Person_MOBILE,
					},
				},
			},
		},
	}
	for name, tt := range testcases {
		t.Run(name, func(t *testing.T) {
			got, err := pb.PromptForAddress(tt.r)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
