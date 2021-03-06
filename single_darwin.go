package single

import (
	"fmt"
	"os"
	"path/filepath"
)

// Filename returns an absolute filename, appropriate for the operating system
func (s *Single) Filename() string {
	if len(Lockfile) > 0 {
		return Lockfile
	}
	if s.Path != "" {
		return filepath.Join(s.Path, fmt.Sprintf("%s.lock", s.name))
	}
	return filepath.Join(os.TempDir(), fmt.Sprintf("%s.lock", s.name))
}
