package awssecrets

import "github.com/alecthomas/kingpin"

// Type list of secrets, used to allow giving multiple secrets in the command line
// This needs to implement the kingpin.Value interface
type secretListValue []string

// Set needs to be implemented to satisfy kingpins Value interface
func (list *secretListValue) Set(value string) error {
	*list = append(*list, value)
	return nil
}

// String needs to be implemented to satisfy kingpins Value interface
func (list secretListValue) String() string {
	return ""
}

// This is used by kingpin package to decide if the argument can be used multiple times
func (list secretListValue) IsCumulative() bool {
	return true
}

// SecretList takes kingpin settings and sets value to be of secretListValue type
func SecretList(s kingpin.Settings) *[]string {
	var target []string
	s.SetValue((*secretListValue)(&target))
	return &target
}
