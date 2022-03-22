# LibAppImage Go
Go Bindings For LibAppImage

---
### Usage

Make sure to first make a new binding:
```go
myBinding, err := libappimagego.NewLibAppImageBindings()
appImagePath := "/home/aditya/test.appimage"
```

Now Access All The Functions You Want:
```go
shallBeIntegrated := myBinding.ShallNotBeIntegrated(appImagePath)

if shallBeIntegrated == true {
    err = myBinding.Register(appImagePath)
    err = myBinding.Unregister(appImagePath)
}

isATerminalApp := myBinding.IsTerminalApp(appImagePath)

/*
 TYPE CAN BE:
  libappimagego.APPIMAGE_TYPE_INVALID
  libappimagego.APPIMAGE_TYPE_LEGACY
  libappimagego.APPIMAGE_TYPE_1
  libappimagego.APPIMAGE_TYPE_2
*/
appImageType := myBinding.GetType(appImagePath)

// Make sure to close the bindings after using it.
myBinding.Close()
```

---

## Thanks