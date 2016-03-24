package filo

// A function used internally by a File to write.
type FileWrite func(Header, []byte) (int, error)

func borderColorWrite(h Header, in []byte) (int, error) {
	return 0, nil
}

func borderWidthWrite(h Header, in []byte) (int, error) {
	return 0, nil
}

func widthWrite(h Header, in []byte) (int, error) {
	return 0, nil
}

func heightWrite(h Header, in []byte) (int, error) {
	return 0, nil
}

func xWrite(h Header, in []byte) (int, error) {
	return 0, nil
}

func yWrite(h Header, in []byte) (int, error) {
	return 0, nil
}
