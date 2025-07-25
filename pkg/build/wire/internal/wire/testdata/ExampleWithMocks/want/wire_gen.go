// Code generated by Wire. DO NOT EDIT.

//go:generate go run ./pkg/build/wire/cmd/wire/main.go
//go:build !wireinject

package main

// Injectors from wire.go:

// initApp returns a real app.
func initApp() *app {
	mainTimer := _wireRealTimeValue
	mainGreeter := greeter{
		T: mainTimer,
	}
	mainApp := &app{
		g: mainGreeter,
	}
	return mainApp
}

var (
	_wireRealTimeValue = realTime{}
)

// initMockedAppFromArgs returns an app with mocked dependencies provided via
// arguments (Approach A). Note that the argument's type is the interface
// type (timer), but the concrete mock type should be passed.
func initMockedAppFromArgs(mt timer) *app {
	mainGreeter := greeter{
		T: mt,
	}
	mainApp := &app{
		g: mainGreeter,
	}
	return mainApp
}

// initMockedApp returns an app with its mocked dependencies, created
// via providers (Approach B).
func initMockedApp() *appWithMocks {
	mainMockTimer := newMockTimer()
	mainGreeter := greeter{
		T: mainMockTimer,
	}
	mainApp := app{
		g: mainGreeter,
	}
	mainAppWithMocks := &appWithMocks{
		app: mainApp,
		mt:  mainMockTimer,
	}
	return mainAppWithMocks
}
