// +build plan9 appengine wasm aix

package flags

func getTerminalColumns() int {
	return 80
}
