package errorlist

import (
	"errors"
	"fmt"
)

var (
	// GENERAL
	ErrAccessDenied   = errors.New("Process stopped")
	ErrNotExist       = errors.New("file does not exist")
	ErrForbiddenName  = errors.New("Forbidden name")
	ErrForbiddenChars = errors.New("Forbidden characters")

	// EVENTS
	ErrNoSubEvent     = errors.New("SubEvent is not set")
	ErrUnknownCommand = errors.New("Unknown command")

	// PARSING
	ErrEmptyCommand = errors.New("Empty command")
	ErrTooFewArgs   = errors.New("Not enough arguments")
	ErrNoPayload    = errors.New("Payload not found")

	// IN DEVELOPMENT
	ErrorInwork = errors.New("In development!")

	// INITIALIZATION
	ErrTempDirEmpty = errors.New("os.TempDir() returned an empty path")
	ErrInitRoot     = errors.New("Initialization error!")
)

// InvalidName builds an invalid name error,
// adding the problematic name and wrapping the cause err.
func InvalidName(name string, err error) error {
	return fmt.Errorf("Invalid name %q: %w", name, err)
}

// Create — error creating an object with the name.
func Create(name string, err error) error {
	return fmt.Errorf("Create %s: %w", name, err)
}

// Run — error running an object with the name.
func Run(name string, err error) error {
	return fmt.Errorf("Run %s: %w", name, err)
}

// Remove — error removing an object with the name.
func Remove(name string, err error) error {
	return fmt.Errorf("Remove %s: %w", name, err)
}

// Read — error reading an object with the name.
func Read(name string, err error) error {
	return fmt.Errorf("Read %s: %w", name, err)
}

// Open — error opening an object with the name.
func Open(name string, err error) error {
	return fmt.Errorf("Open %s: %w", name, err)
}

// ClearLog — error clearing the log.
func ClearLog(err error) error {
	return fmt.Errorf("Clear log: %w", err)
}

// RunCommand — error running the command.
func RunCommand(command string, err error) error {
	return fmt.Errorf("Run command %q: %w", command, err)
}

// ClearTerminal — error clearing the terminal.
func ClearTerminal(err error) error {
	return fmt.Errorf("Clear terminal: %w", err)
}

// OpenFile — error opening the file at path.
func OpenFile(path string, err error) error {
	return fmt.Errorf("Open %s: %w", path, err)
}

// ReadFile — error reading the file/directory at path.
func ReadFile(path string, err error) error {
	return fmt.Errorf("Read %s: %w", path, err)
}

// CreateFile — error creating the file at path.
func CreateFile(path string, err error) error {
	return fmt.Errorf("Create %s: %w; Возможно необходимо -init", path, err)
}

// WriteFile — error writing to the file at path.
func WriteFile(path string, err error) error {
	return fmt.Errorf("Write %s: %w", path, err)
}

// RemoveFile — error removing the file at path.
func RemoveFile(path string, err error) error {
	return fmt.Errorf("Remove %s: %w", path, err)
}
