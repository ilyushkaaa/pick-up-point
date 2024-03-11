package storage

func (fs *FilePPStorage) Close() error {
	return fs.file.Close()
}
