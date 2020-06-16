package autoorient

import (
	"github.com/bbiao/nephele/command"
)

func init() {
	command.Register("autoorient", func() command.Command {
		return &Command{}
	})
}
