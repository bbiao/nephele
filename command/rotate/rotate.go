package rotate

import (
	"github.com/bbiao/nephele/command"
)

func init() {
	command.Register("rotate", func() command.Command {
		return &Command{}
	})
}
