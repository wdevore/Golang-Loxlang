package src

type configJSON struct {
	BinaryName string
	Generate   string // "Binary", "Ascii"
}

type Properties struct {
	Config configJSON
}
