package libappimagego

import (
	"C"
	"errors"
)

// Unregister a appimage
func (bind *libAppImageBind) Unregister(filePath string, debug bool) error {
	if bind.appimage_unregister_in_system(C.CString(filePath), boolToInt(debug)) != 0 {
		return errors.New("unregister failed")
	}

	return nil
}
