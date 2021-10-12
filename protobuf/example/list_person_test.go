package example_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	pb "github.com/gojustforfun/protobuf-grpc-gateway-example/protobuf/example"
	"github.com/stretchr/testify/assert"
)

func TestListPeople(t *testing.T) {
	want := `Person ID: 12345
 Name: Example Name
 E-mail address: name@example.com
 Home phone #: 123-456-7890
 Mobile phone #: 222-222-2222
 Work phone #: 111-111-1111
`
	testcases := map[string]struct {
		w    io.Writer
		book *pb.AddressBook
		want string
	}{
		"normal": {
			w: &bytes.Buffer{},
			book: &pb.AddressBook{
				People: []*pb.Person{
					{
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
						},
					},
				},
			},
			want: want,
		},
	}
	for name, tt := range testcases {
		t.Run(name, func(t *testing.T) {
			pb.ListPeople(tt.w, tt.book)
			s, ok := tt.w.(fmt.Stringer)
			assert.True(t, ok)
			assert.Equal(t, tt.want, s.String())
		})
	}
}
