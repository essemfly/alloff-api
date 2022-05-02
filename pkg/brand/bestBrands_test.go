package brand

import (
	"github.com/lessbutter/alloff-api/cmd"
	"testing"
)

func TestMakeSnapshot(t *testing.T) {
	cmd.SetBaseConfig("dev")
	t.Run("test make snapshot", func(t *testing.T) {
		MakeSnapshot()
	})
}
