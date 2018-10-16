// An implementation of the XDG Base Directory Specification.
//
// For more information, see: http://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html
package xdg

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	// Both user-specific and standard directories to search for data files.
	// (same as DATA_DIRS, but including preferring DATA_HOME)
	DATA Paths

	// Directory to write user-specific data files.
	DATA_HOME Path // $XDG_DATA_HOME (default: $HOME/.local/share)

	// Directories to search for data files, in preference order.
	DATA_DIRS Paths // $XDG_DATA_DIRS (default: /usr/local/share:/usr/share)

	// Both user-specific and standard directories to search for config files.
	// (same as CONFIG_DIRS, but including preferring CONFIG_HOME)
	CONFIG Paths

	// Directory to write user-specific config files.
	CONFIG_HOME Path // $XDG_CONFIG_HOME (default: $HOME/.config)

	// Directories to search for config files, in preference order.
	CONFIG_DIRS Paths // $XDG_CONFIG_DIRS (default: /etc/xdg)

	// Directory to write user-specific, non-essential (cached) data files.
	CACHE_HOME Path // $XDG_CACHE_HOME (default: $HOME/.cache)

	// Directory to write user-specific runtime files and other file objects.
	RUNTIME_DIR Path // $XDG_RUNTIME_DIR
)

func init() {
	home := os.Getenv("HOME")

	// defaults
	data_home_default := Path(filepath.Join(home, ".local", "share"))
	data_dirs_default := Paths{Path(filepath.FromSlash("/usr/local/share")), Path(filepath.FromSlash("/usr/share"))}

	config_home_default := Path(filepath.Join(home, ".config"))
	config_dirs_default := Paths{Path(filepath.FromSlash("/etc/xdg"))}

	cache_home_default := Path(filepath.Join(home, ".cache"))

	// DATA
	DATA_HOME = PathWithDefault(os.Getenv("XDG_DATA_HOME"), data_home_default)
	DATA_DIRS = PathsWithDefault(strings.Split(os.Getenv("XDG_DATA_DIRS"), ":"), data_dirs_default)
	DATA = append(Paths{DATA_HOME}, DATA_DIRS...)

	// CONFIG
	CONFIG_HOME = PathWithDefault(os.Getenv("XDG_CONFIG_HOME"), config_home_default)
	CONFIG_DIRS = PathsWithDefault(strings.Split(os.Getenv("XDG_CONFIG_DIRS"), ":"), config_dirs_default)
	CONFIG = append(Paths{CONFIG_HOME}, CONFIG_DIRS...)

	// CACHE
	CACHE_HOME = PathWithDefault(os.Getenv("XDG_CACHE_HOME"), cache_home_default)

	// RUNTIME
	// explicitly set to nothing if the XDG_RUNTIME_DIR is invalid
	RUNTIME_DIR = PathWithDefault(os.Getenv("XDG_RUNTIME_DIR"), "")
}

type XDG struct {
	// Both user-specific and standard directories to search for data files.
	// (same as DATA_DIRS, but including preferring DATA_HOME)
	DATA Paths

	// Directory to write user-specific data files.
	DATA_HOME Path // $XDG_DATA_HOME (default: $HOME/.local/share)

	// Directories to search for data files, in preference order.
	DATA_DIRS Paths // $XDG_DATA_DIRS (default: /usr/local/share:/usr/share)

	// Both user-specific and standard directories to search for config files.
	// (same as CONFIG_DIRS, but including preferring CONFIG_HOME)
	CONFIG Paths

	// Directory to write user-specific config files.
	CONFIG_HOME Path // $XDG_CONFIG_HOME (default: $HOME/.config)

	// Directories to search for config files, in preference order.
	CONFIG_DIRS Paths // $XDG_CONFIG_DIRS (default: /etc/xdg)

	// Directory to write user-specific, non-essential (cached) data files.
	CACHE_HOME Path // $XDG_CACHE_HOME (default: $HOME/.cache)

	// Directory to write user-specific runtime files and other file objects.
	RUNTIME_DIR Path // $XDG_RUNTIME_DIR
}

// WithSuffix joins a given name to each of the paths resolved by the package.
//
// This is useful for ensuring all XDG access is properly scoped to the calling
// application.
func WithSuffix(name string) *XDG {
	data_home := Path(DATA_HOME.Join(name))
	var data_dirs Paths
	for _, dir := range DATA_DIRS {
		data_dirs = append(data_dirs, Path(dir.Join(name)))
	}

	config_home := Path(CONFIG_HOME.Join(name))
	var config_dirs Paths
	for _, dir := range CONFIG_DIRS {
		config_dirs = append(config_dirs, Path(dir.Join(name)))
	}

	cache_home := Path(CACHE_HOME.Join(name))

	runtime_dir := Path(RUNTIME_DIR.Join(name))

	return &XDG{
		DATA_HOME: data_home,
		DATA_DIRS: data_dirs,
		DATA:      append(Paths{data_home}, data_dirs...),

		CONFIG_HOME: config_home,
		CONFIG_DIRS: config_dirs,
		CONFIG:      append(Paths{config_home}, config_dirs...),

		CACHE_HOME: cache_home,

		RUNTIME_DIR: runtime_dir,
	}
}
