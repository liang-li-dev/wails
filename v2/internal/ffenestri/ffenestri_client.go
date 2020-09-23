package ffenestri

/*

#cgo linux CFLAGS: -DFFENESTRI_LINUX=1
#cgo linux pkg-config: gtk+-3.0 webkit2gtk-4.0

#include <stdlib.h>
#include "ffenestri.h"

*/
import "C"

import (
	"strconv"
	"unsafe"

	"github.com/wailsapp/wails/v2/internal/logger"
)

// Client is our implentation of messageDispatcher.Client
type Client struct {
	app    *Application
	logger logger.CustomLogger
}

func newClient(app *Application) *Client {
	return &Client{
		app:    app,
		logger: app.logger,
	}
}

// Quit the application
func (c *Client) Quit() {
	c.app.logger.Trace("Got shutdown message")
	C.Quit(c.app.app)
}

// NotifyEvent will pass on the event message to the frontend
func (c *Client) NotifyEvent(message string) {
	eventMessage := `window.wails._.Notify(` + strconv.Quote(message) + `);`
	c.app.logger.Trace("eventMessage = %+v", eventMessage)
	C.ExecJS(c.app.app, c.app.string2CString(eventMessage))
}

// CallResult contains the result of the call from JS
func (c *Client) CallResult(message string) {
	callbackMessage := `window.wails._.Callback(` + strconv.Quote(message) + `);`
	c.app.logger.Trace("callbackMessage = %+v", callbackMessage)
	C.ExecJS(c.app.app, c.app.string2CString(callbackMessage))
}

// WindowSetTitle sets the window title to the given string
func (c *Client) WindowSetTitle(title string) {
	C.SetTitle(c.app.app, c.app.string2CString(title))
}

// WindowFullscreen will set the window to be fullscreen
func (c *Client) WindowFullscreen() {
	C.Fullscreen(c.app.app)
}

// WindowUnFullscreen will unfullscreen the window
func (c *Client) WindowUnFullscreen() {
	C.UnFullscreen(c.app.app)
}

// WindowShow will show the window
func (c *Client) WindowShow() {
	C.Show(c.app.app)
}

// WindowHide will hide the window
func (c *Client) WindowHide() {
	C.Hide(c.app.app)
}

// WindowCenter will hide the window
func (c *Client) WindowCenter() {
	C.Center(c.app.app)
}

// WindowMaximise will maximise the window
func (c *Client) WindowMaximise() {
	C.Maximise(c.app.app)
}

// WindowMinimise will minimise the window
func (c *Client) WindowMinimise() {
	C.Minimise(c.app.app)
}

// WindowUnmaximise will unmaximise the window
func (c *Client) WindowUnmaximise() {
	C.Unmaximise(c.app.app)
}

// WindowUnminimise will unminimise the window
func (c *Client) WindowUnminimise() {
	C.Unminimise(c.app.app)
}

// WindowPosition will position the window to x,y on the
// monitor that the window is mostly on
func (c *Client) WindowPosition(x int, y int) {
	C.SetPosition(c.app.app, C.int(x), C.int(y))
}

// WindowSize will resize the window to the given
// width and height
func (c *Client) WindowSize(width int, height int) {
	C.SetSize(c.app.app, C.int(width), C.int(height))
}

// WindowSetColour sets the window colour
func (c *Client) WindowSetColour(colour int) {
	r, g, b, a := intToColour(colour)
	C.SetColour(c.app.app, r, g, b, a)
}

// OpenFileDialog will open a file dialog with the given title
func (c *Client) OpenFileDialog(title string, filter string) string {

	cstring := C.OpenFileDialog(c.app.app, c.app.string2CString(title), c.app.string2CString(filter))
	var result string
	if cstring != nil {
		result = C.GoString(cstring)
		// Free the C string that was allocated by the dialog
		C.free(unsafe.Pointer(cstring))
	}
	return result
}

// SaveFileDialog will open a save file dialog with the given title
func (c *Client) SaveFileDialog(title string, filter string) string {

	cstring := C.SaveFileDialog(c.app.app, c.app.string2CString(title), c.app.string2CString(filter))
	var result string
	if cstring != nil {
		result = C.GoString(cstring)
		// Free the C string that was allocated by the dialog
		C.free(unsafe.Pointer(cstring))
	}
	return result
}

// OpenDirectoryDialog will open a directory dialog with the given title
func (c *Client) OpenDirectoryDialog(title string, filter string) string {

	cstring := C.OpenDirectoryDialog(c.app.app, c.app.string2CString(title), c.app.string2CString(filter))
	var result string
	if cstring != nil {
		result = C.GoString(cstring)
		// Free the C string that was allocated by the dialog
		C.free(unsafe.Pointer(cstring))
	}
	return result
}