package main

import (
	"fmt"
	"libappimage-go/src"
)

func main() {
	appimageFile, debug := "/home/adityam/Applications/balenaEtcher-1.7.8-x64.AppImage", false
	bindings, err := libappimagego.NewLibAppImageBindings()
	if err != nil {
		panic(err)
	}

	fmt.Println("AppImage Type:", bindings.GetType(appimageFile, debug))
	fmt.Println("Is Terminal App:", bindings.IsTerminalApp(appimageFile))
	fmt.Println("Should Be Integrated:", bindings.ShallAppImageBeRegistered(appimageFile))
	fmt.Println("Is Registered:", bindings.IsRegistered(appimageFile))

	fmt.Println("------- Running Integration Test -------")
	shallBeIntegrated := bindings.ShallAppImageBeRegistered(appimageFile)
	if shallBeIntegrated {
		fmt.Println("Integrating The AppImage...")
		err := bindings.Register(appimageFile + "", debug)
		if err != nil {
			fmt.Println("Error When Integrating: " + err.Error())
			fmt.Println("------- Test Ended With Errors -------")
		} else {
			fmt.Println("Successfully integrated the appimage...\nRemoving the integration now...")
			err := bindings.Unregister(appimageFile, debug)
			if err != nil {
				fmt.Println("Error When Removing Integration: " + err.Error())
				fmt.Println("------- Test Ended With Errors -------")
				} else {
				fmt.Println("Successfully Removed The Integration")
				fmt.Println("-------- Test Ended Sucessfully --------")
			}
		}
	}
}