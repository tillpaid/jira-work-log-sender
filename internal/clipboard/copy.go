package clipboard

import (
	"bytes"
	"os/exec"
)

func CopyToClipboard(text string) error {
	cmd := exec.Command("pbcopy")
	cmd.Stdin = bytes.NewReader([]byte(text))

	return cmd.Run()
}
