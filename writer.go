package gospin

// A Writer allows strings to be written to an output. Output may not appear
// until Sync is called.
type Writer interface {

	// Write len(b) bytes to the output or output buffer. It returns the number
	// of bytes written and an error, if any. Write returns a non-nil error when
	// n != len(b). The contents of Write may not appear on the output until
	// Sync is called.
	Write(b []byte) (int, error)

	// Sync flushes the writer to the output it is tied to.
	Sync() error
}
