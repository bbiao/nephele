package strip

import (
	"github.com/bbiao/nephele/command"
)

func init() {
	command.Register("strip", func() command.Command {
		return &Command{}
	})
}
