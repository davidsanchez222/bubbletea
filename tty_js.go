//go:build js

package tea

import (
	"errors"
	"os"
)

// There is no TTY, resize signal or job control in a browser.
// The host feeds input through WithInput and reports size changes by sending
// WindowSizeMsg, so these hooks are no-ops.

func (p *Program) initInput() error { return nil }

func (p *Program) listenForResize(done chan struct{}) { close(done) }

func openInputTTY() (*os.File, error) {
	return nil, errors.New("bubbletea: no tty available in the browser")
}

const suspendSupported = false

func suspendProcess() {}
