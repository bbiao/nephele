package format

import (
	"github.com/bbiao/nephele/command"
)

func init() {
	command.Register("format", func() command.Command {
		return &Command{}
	})
}
