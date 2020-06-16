package sharpen

import (
	"github.com/bbiao/nephele/command"
)

func init() {
	command.Register("sharpen", func() command.Command {
		return &Command{}
	})
}
