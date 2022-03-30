# LibAppImage Go
Go Bindings For LibAppImage

---
### Usage

Make sure to first make a new binding:
```go
myBinding, err := libappimagego.NewLibAppImageBindings()
```
And this is some extra stuff for storing appimage & debug flag
```go
appImagePath, debug := "/home/aditya/test.appimage", false
```

#### API
##### `(bind *libAppImageBind) Register(filePath string, debug bool) (error)`
Register the appimage from the given path to system, registering here is a term which can be used interchangeably with integrating.

```go
err := myBinding.Register(appImagePath, debug)
if err != nil {
    panic(err)
}
```

##### `(bind *libAppImageBind) UnRegister(filePath string, debug bool) (error)`
UnRegister/De-Integrate the appimage from the given path from system.

```go
err := myBinding.UnRegister(appImagePath, debug)
if err != nil {
    panic(err)
}
```

##### `(bind *libAppImageBind) ShallAppImageBeRegistered(filePath string) (bool)`
Returns a boolean representing if the AppImage Distributor/Author wants the appimage to be integrated to system or not.

```go
shallBeRegistered := myBinding.ShallAppImageBeRegistered(appImagePath)
if shallBeRegistered {
    doSomething()
}
```

##### `(bind *libAppImageBind) IsRegistered(filePath string) (bool)`
Returns a boolean representing if the appimage is registered/integrated or not.

```go
isRegistered := myBinding.IsRegistered(appImagePath)
if isRegistered {
    doSomething()
}
```

##### `(bind *libAppImageBind) IsTerminalApp(filePath string) (bool)`
Returns a boolean representing if the appimage is a terminal app or not.

```go
isTerminalApp := myBinding.IsTerminalApp(appImagePath)
if isTerminalApp {
    doSomething()
}
```

##### `(bind *libAppImageBind) GetType(filePath string, debug bool) (int)`
Returns a integer representing the type of the appimage, 0 means Legacy, 1 means type 1, 2 means type 2 & -1 means invalid appimage. [Read More About AppImage Types](https://github.com/AppImage/AppImageSpec/blob/master/draft.md#image-format). LibAppImage Go provides constants to make to code more readable.

```go
appimageType := myBinding.GetType(appImagePath, debug)
if appimageType == libappimagego.APPIMAGE_TYPE_LEGACY {
    doSomething()
} else if appimageType == APPIMAGE_TYPE_1 {
    doSomething()
} else if appimageType == APPIMAGE_TYPE_2 {
    doSomething()
} else if appimageType == APPIMAGE_TYPE_INVALID {
    doSomething()
}
```

##### `(bind *libAppImageBind) Close()`
Make sure to call the `Close()` method so that it closes the bindings.

```go
myBinding.Close()
```

---

## Thanks
