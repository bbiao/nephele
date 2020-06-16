package resize

import (
	"github.com/bbiao/nephele/command"
)

func init() {
	command.Register("resize", func() command.Command {
		return &Command{}
	})
}
