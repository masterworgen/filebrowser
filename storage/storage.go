package storage

import (
	"github.com/masterworgen/filebrowser/v2/auth"
	"github.com/masterworgen/filebrowser/v2/settings"
	"github.com/masterworgen/filebrowser/v2/share"
	"github.com/masterworgen/filebrowser/v2/users"
)

// Storage is a storage powered by a Backend which makes the necessary
// verifications when fetching and saving data to ensure consistency.
type Storage struct {
	Users    users.Store
	Share    *share.Storage
	Auth     *auth.Storage
	Settings *settings.Storage
}
