package crop

import (
	"github.com/bbiao/nephele/command"
)

func init() {
	command.Register("crop", func() command.Command {
		return &Command{}
	})
}
