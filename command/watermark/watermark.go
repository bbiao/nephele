package watermark

import (
	"github.com/bbiao/nephele/command"
)

func init() {
	command.Register("watermark", func() command.Command {
		return &Command{}
	})
}
