package all

import (
	"github.com/josh/restic-api/api/backend/azure"
	"github.com/josh/restic-api/api/backend/b2"
	"github.com/josh/restic-api/api/backend/gs"
	"github.com/josh/restic-api/api/backend/local"
	"github.com/josh/restic-api/api/backend/location"
	"github.com/josh/restic-api/api/backend/rclone"
	"github.com/josh/restic-api/api/backend/rest"
	"github.com/josh/restic-api/api/backend/s3"
	"github.com/josh/restic-api/api/backend/sftp"
	"github.com/josh/restic-api/api/backend/swift"
)

func Backends() *location.Registry {
	backends := location.NewRegistry()
	backends.Register(azure.NewFactory())
	backends.Register(b2.NewFactory())
	backends.Register(gs.NewFactory())
	backends.Register(local.NewFactory())
	backends.Register(rclone.NewFactory())
	backends.Register(rest.NewFactory())
	backends.Register(s3.NewFactory())
	backends.Register(sftp.NewFactory())
	backends.Register(swift.NewFactory())
	return backends
}
