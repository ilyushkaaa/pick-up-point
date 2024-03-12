package storage

func (fs *FilePPStorage) Close() error {
	err := fs.SaveCashToFile()
	if err != nil {
		return err
	}
	return fs.file.Close()
}
