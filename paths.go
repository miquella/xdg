package xdg

import (
	"os"
	"path/filepath"
)

// IsValid returns whether the path can be used.
//
// A valid path is a non-empty, absolute path. The path need not exist.
func IsValid(path string) bool {
	return path != "" && filepath.IsAbs(string(path))
}

// Path is a representation of a single valid path.
type Path string

// PathWithDefault returns a Path representing path, if it is valid, or
// defaultPath otherwise.
func PathWithDefault(path string, defaultPath Path) Path {
	if IsValid(path) {
		return Path(path)
	}

	return defaultPath
}

// IsValid returns whether the path can be used.
//
// A valid path is a non-empty, absolute path. The path need not exist.
func (p Path) IsValid() bool {
	return IsValid(string(p))
}

// Find searches for an existent file or directory in the path. If no match is
// found, an empty string is returned.
func (p Path) Find(elem ...string) string {
	file := p.Join(elem...)
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return ""
	}

	return file
}

// Glob matches files or directories contained in the path. The filepath.Glob
// syntax is used for matching.
func (p Path) Glob(pattern string) ([]string, error) {
	return filepath.Glob(filepath.Join(string(p), pattern))
}

// Join returns a string representation of the path joined with additional
// path elements.
func (p Path) Join(elem ...string) string {
	return filepath.Join(append([]string{string(p)}, elem...)...)
}

// Paths is a representation of a set of valid paths.
type Paths []Path

// PathsWithDefault returns a set including all valid values from paths. If no
// values in paths are valid, defaultPaths is returned instead.
func PathsWithDefault(paths []string, defaultPaths Paths) Paths {
	var p Paths
	for _, path := range paths {
		if IsValid(path) {
			p = append(p, Path(path))
		}
	}

	if len(p) == 0 {
		for _, path := range defaultPaths {
			p = append(p, path)
		}
	}

	return p
}

// Find searches for existent files or directories in the paths. Existent files
// or directories are returned in the same order in which the paths are
// specified.
func (p Paths) Find(elem ...string) []string {
	var found []string
	for _, path := range p {
		file := path.Find(elem...)
		if file != "" {
			found = append(found, file)
		}
	}
	return found
}

// Glob matches files or directories contained in each of the paths. Matches
// are returned in the same order in which the paths are specified. The
// filepath.Glob syntax is used for matching.
func (p Paths) Glob(pattern string) ([]string, error) {
	var matches []string
	for _, path := range p {
		pathMatches, err := path.Glob(pattern)
		if err != nil {
			return nil, err
		}

		matches = append(matches, pathMatches...)
	}

	return matches, nil
}

// Join returns a set of string representations of each path joined with
// additional path elements.
func (p Paths) Join(elem ...string) []string {
	pathElem := append([]string{""}, elem...)

	var joined []string
	for _, path := range p {
		pathElem[0] = string(path)
		joined = append(joined, filepath.Join(pathElem...))
	}

	return joined
}
