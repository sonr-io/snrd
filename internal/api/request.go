package api

import (
	"os"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/internal/device"
)

// IsDev returns true if the node is running in development mode.
func (ir *InitializeRequest) IsDev() bool {
	return ir.GetEnvironment().IsDev()
}

// fsOpts returns the device.FSOptions
func (ir *InitializeRequest) fsOpts() []device.FSOption {
	fsOpts := make([]device.FSOption, 0)
	if ir.GetDeviceOptions() != nil {
		// Set Temporary Path
		fsOpts = append(fsOpts, device.FSOption{
			Path: ir.GetDeviceOptions().GetCacheDir(),
			Type: device.Temporary,
		}, device.FSOption{
			Path: ir.GetDeviceOptions().GetDownloadsDir(),
			Type: device.Downloads,
		}, device.FSOption{
			Path: ir.GetDeviceOptions().GetDocumentsDir(),
			Type: device.Documents,
		}, device.FSOption{
			Path: ir.GetDeviceOptions().GetSupportDir(),
			Type: device.Support,
		}, device.FSOption{
			Path: ir.GetDeviceOptions().GetDatabaseDir(),
			Type: device.Database,
		}, device.FSOption{
			Path: ir.GetDeviceOptions().GetTextileDir(),
			Type: device.Textile,
		})
	}
	return fsOpts
}

// ParseOpts parses the device options and returns the device.FSOptions
func (ir *InitializeRequest) ParseOpts() []device.FSOption {
	logger.Info("Parsing Initialize Request...")

	// Check DeviceID
	ir.SetEnvVars()
	ir.SetDeviceID()
	return ir.fsOpts()
}

// SetEnvVars sets the environment variables
func (ir *InitializeRequest) SetEnvVars() {
	vars := ir.GetVariables()
	count := len(vars)

	// Set Env Variables
	if count > 0 {
		for k, v := range vars {
			os.Setenv(k, v)
		}

		golog.Info("Added Enviornment Variable(s)", golog.Fields{
			"Total": count,
		})
	} else {
		golog.Warn("No Enviornment Variable(s) passed")
	}
}

// SetDeviceID sets the device id
func (ir *InitializeRequest) SetDeviceID() {
	did := ir.GetDeviceOptions().GetId()
	if did != "" {
		logger.Info("Device ID Passed: " + did)
		device.SetDeviceID(did)
	} else {
		golog.Warn("No Device ID Passed")
	}
}

// IsDelete returns true if the request is a delete request
func (er *EditRequest) IsDelete() bool {
	return er.GetType() == EditRequest_DELETE
}
