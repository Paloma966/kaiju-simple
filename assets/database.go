package assets

// Database is the interface for loading game assets (textures, meshes, shaders, sounds, etc.).
//
// In the real kaiju, the assets package has:
//   - FileDatabase: reads assets from the filesystem
//   - ArchiveDatabase: reads assets from packed archive files
//   - Caching layers for textures, meshes, materials, and shaders
//   - A "Table of Contents" mapping asset paths to archive offsets
type Database interface {
	// Open returns a byte slice of the asset at the given path.
	// Returns an error if the asset doesn't exist or can't be read.
	Open(path string) ([]byte, error)

	// Exists returns true if an asset exists at the given path.
	Exists(path string) bool

	// List returns all asset paths under the given directory prefix.
	List(prefix string) []string

	// Close releases any resources held by the database.
	Close() error
}
