package quality

import (
	"github.com/bbiao/nephele/command"
)

func init() {
	command.Register("quality", func() command.Command {
		return &Command{}
	})
}
