package storage

func (fs *FileOrderStorage) Close() error {
	return fs.file.Close()
}
