package gospin

import (
	"bytes"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

// A Spinner outputs a spinner in place to stdout, allowing activity to be
// indicated. Run in porcelain mode a Spinner does nothing. A spinner will try
// to detect if it's not running in a terminal and set porcelain mode if it is.
type Spinner interface {
	// Advance the spinner forward one frame. The previous frame will be
	// overwritten by the new frame. On the first call to Advance the cursor
	// will also be hidden.
	Advance()

	// Done will remove the spinner, show the cursor and reset the spinner
	// ready for reuse.
	Done()

	// Width overrides the width of the spinner. By default the number of
	// characters in the first frame are used to determine the width of the
	// spinner, however this doesn't take into account things like control
	// characters.
	Width(w int)

	// Porcelain stops the spinner from producing any output and calls to
	// Advance and Done will do nothing.
	Porcelain()

	// Writer sets the writer to use when producing the spinner. By default the
	// writer is stdout. Setting the writer sets porcelain mode to false. If a
	// custom writer is used then setting porcelain needs to happen outside of
	// the Spinner and after the writer is set.
	Writer(writer Writer)
}

type spinner struct {
	porcelain bool
	clear     bool
	width     int
	frame     int
	frames    []string
	writer    Writer
}

var (
	back = []byte("\033[D")
	hide = []byte("\033[?25l")
	show = []byte("\033[?25h")
)

// Predefined spinners, most from https://github.com/sindresorhus/cli-spinners/
var (
	basic     = []string{"|", "/", "-", "\\"}
	dot       = []string{"⠈", "⠐", "⠠", "⢀", "⡀", "⠄", "⠂", "⠁"}
	dots      = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	ellipses  = []string{".  ", ".. ", "...", "   "}
	scrolling = []string{".  ", ".. ", "...", " ..", "  .", "   "}
)

// New returns a new Spinner that uses the given frames for the spinner
// animation. The Spinner will attempt to determine if it's running in the
// console, engaging porcelain mode if it isn't. The first frame of the
// animation is used to determine the width of the spinner. If no frames are
// provided then the basic frames ("|", "/", "-", and "\") are used.
func New(frames ...string) Spinner {
	s := &spinner{writer: os.Stdout}

	if len(frames) > 0 {
		s.frames = frames
	} else {
		s.frames = basic
	}

	s.width = len(s.frames[0])
	s.porcelain = !terminal.IsTerminal(int(os.Stdout.Fd()))
	return s
}

// Basic returns a basic Spinner.
func Basic() Spinner {
	return New(basic...)
}

// Dot returns a spinner that uses a single dot moving in a circle.
func Dot() Spinner {
	s := New(dot...)
	s.Width(1)
	return s
}

// Dots returns a Spinner that uses multiple dots moving in a circle.
func Dots() Spinner {
	s := New(dots...)
	s.Width(1)
	return s
}

// Ellipses returns a Spinner that repeatedly draws ellipses.
func Ellipses() Spinner {
	return New(ellipses...)
}

// Scrolling returns a Spinner that produces scrolling ellipses.
func Scrolling() Spinner {
	return New(scrolling...)
}

func (s *spinner) Advance() {
	if s.porcelain {
		return
	}

	if s.clear {
		s.writer.Write(bytes.Repeat(back, s.width))
	} else {
		s.writer.Write(hide)
		s.clear = true
	}

	s.writer.Write([]byte(s.frames[s.frame]))
	s.writer.Sync()

	s.frame++

	if s.frame == len(s.frames) {
		s.frame = 0
	}

}

func (s *spinner) Done() {
	if s.porcelain {
		return
	}

	s.frame = 0

	if s.clear {
		s.writer.Write(bytes.Repeat(back, s.width))
		s.writer.Write(show)
		s.writer.Sync()
	}
}

func (s *spinner) Width(w int) {
	s.width = w
}

func (s *spinner) Porcelain() {
	s.porcelain = true
}

func (s *spinner) Writer(writer Writer) {
	s.porcelain = false
	s.writer = writer
}
