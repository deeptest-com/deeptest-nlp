package agentHandler

import (
	"fmt"
	_domain "github.com/utlai/utl/internal/pkg/domain"
	"golang.org/x/net/context"
)

type ArithCtrl struct {
}

func NewArithCtrl() *ArithCtrl {
	return &ArithCtrl{}
}

func (c *ArithCtrl) Add(ctx context.Context, args *_domain.Args, reply *_domain.Reply) error {
	reply.C = args.A + args.B
	fmt.Printf("call: %d + %d = %d\n", args.A, args.B, reply.C)
	return nil
}
