package all

import (
	"github.com/restic/restic/api/backend/azure"
	"github.com/restic/restic/api/backend/b2"
	"github.com/restic/restic/api/backend/gs"
	"github.com/restic/restic/api/backend/local"
	"github.com/restic/restic/api/backend/location"
	"github.com/restic/restic/api/backend/rclone"
	"github.com/restic/restic/api/backend/rest"
	"github.com/restic/restic/api/backend/s3"
	"github.com/restic/restic/api/backend/sftp"
	"github.com/restic/restic/api/backend/swift"
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
