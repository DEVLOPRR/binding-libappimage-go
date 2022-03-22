package libappimagego

import "C"

import (
	"fmt"
	"github.com/rainycape/dl"
)

const (
	APPIMAGE_FORMAT_INVALID = -1 // Not an AppImage file
	APPIMAGE_FORMAT_LEGACY = 0 // portable binaries that look and behave like AppImages but do not follow the standard
	APPIMAGE_FORMAT_1 = 1 // https://github.com/AppImage/AppImageSpec/blob/master/draft.md#type-1-image-format
	APPIMAGE_FORMAT_2 = 2 // https://github.com/AppImage/AppImageSpec/blob/master/draft.md#type-2-image-format
)

type libAppImageBind struct {
	lib *dl.DL

	// Returns boolean based on if the appimage should be registered or not.
	appimage_shall_not_be_integrated func(path *C.char) int

	// Returns boolean based on if the appimage got sucessfuly registered on not
	appimage_register_in_system      func(path *C.char, verbose int) int

	// Returns boolean based on if the appimage got sucessfuly unregistered on not
	appimage_unregister_in_system    func(path *C.char, verbose int) int

	// Returns AppImage Type (0, 1, 2): https://github.com/AppImage/AppImageSpec/blob/master/draft.md#image-format
	appimage_get_type                func(path *C.char, verbose int) int

	// Returns boolean based on if the appimage is a terminal app
	appimage_is_terminal_app         func(path *C.char) int
}

type LibAppImage interface {
	Register(filePath string) error
	Unregister(filePath string) error
	ShallNotBeIntegrated(filePath string) bool
	GetType(filePath string) int
	IsTerminalApp(filePath string) bool
	Close()
}

// Load up libappimage from the system, libappimage comes packed with the imageHub AppImage.
func loadLibAppImage() (*dl.DL, error) {
	// libappimage versions from latest to oldest, so that we can load the latest version
	sharedLibList := [17]string{
		".1.0.4", ".1.0.3", ".1.0.1", ".1.0.2",
		".1.0", ".0.1.9", ".0.1.8", ".0.1.7", ".0.1.6", ".0.1.5",
		".0.1.4", ".0.1.3", ".0.1.2", ".0.1.1", ".0.1.0", ".0", "",
	}

	for index := range sharedLibList {
		lib, err := dl.Open("libappimage.so" + sharedLibList[index], 0)
		if err == nil {
			return lib, nil
		}
	}

	return nil, fmt.Errorf("libappimage not found, desktop integration is disabled")
}

// Makes a new binding with libappimage, and returns a object with functions to register, unregister and other functions
func NewLibAppImageBindings() (LibAppImage, error) {
	bindings := libAppImageBind{}
	var err error
	bindings.lib, err = loadLibAppImage()

	if err != nil {
		return nil, err
	}

	err = bindings.lib.Sym("appimage_shall_not_be_integrated", &bindings.appimage_shall_not_be_integrated)
	if err != nil {
		return nil, err
	}

	err = bindings.lib.Sym("appimage_unregister_in_system", &bindings.appimage_unregister_in_system)
	if err != nil {
		return nil, err
	}

	err = bindings.lib.Sym("appimage_register_in_system", &bindings.appimage_register_in_system)
	if err != nil {
		return nil, err
	}

	err = bindings.lib.Sym("appimage_get_type", &bindings.appimage_get_type)
	if err != nil {
		return nil, err
	}

	err = bindings.lib.Sym("appimage_is_terminal_app", &bindings.appimage_is_terminal_app)
	if err != nil {
		return nil, err
	}

	return &bindings, nil
}

// Register a appimage
func (bind *libAppImageBind) Register(filePath string) error {
	if bind.appimage_register_in_system(C.CString(filePath), 1) != 0 {
		return fmt.Errorf("unregister failed")
	}

	return nil
}

// Unregister a appimage
func (bind *libAppImageBind) Unregister(filePath string) error {
	if bind.appimage_unregister_in_system(C.CString(filePath), 1) != 0 {
		return fmt.Errorf("unregister failed")
	}

	return nil
}

// Returns a boolean if a appimage should be integrated or not
func (bind *libAppImageBind) ShallNotBeIntegrated(filePath string) bool {
	return bind.appimage_shall_not_be_integrated(C.CString(filePath)) != 0
}

// Function to close the binding
func (bind *libAppImageBind) IsTerminalApp(filePath string) bool {
	return bind.appimage_is_terminal_app(C.CString(filePath)) != 0
}

// Function to close the binding
func (bind *libAppImageBind) Close() {
	_ = bind.lib.Close()
}

// Close the binding
func (bind *libAppImageBind) GetType(filePath string) int {
	return bind.appimage_get_type(C.CString(filePath), 1)
}
