package storage

func (fs *FilePPStorage) Close() error {
	err := fs.SaveCacheToFile()
	if err != nil {
		return err
	}
	return fs.file.Close()
}
