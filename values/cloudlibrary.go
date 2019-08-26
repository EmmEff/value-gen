package values

type CloudLibraryServer struct {
	Hostname string
}

func ConfigCloudLibrary(root *Values) error {
	vals := &root.CloudLibraryServer
	defaultHostname := "library.lvh.me"
	return Ask("CloudLibraryServer Hostname:", func() (err error) {
		vals.Hostname, err = ScanString(defaultHostname)
		return
	})
}
