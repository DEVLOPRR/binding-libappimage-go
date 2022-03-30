package libappimagego

import (
	"C"
	"errors"
)

// Register a appimage
func (bind *libAppImageBind) Register(filePath string, debug bool) error {
	if bind.appimage_register_in_system(C.CString(filePath), boolToInt(debug)) != 0 {
		return errors.New("register failed")
	}

	return nil
}
