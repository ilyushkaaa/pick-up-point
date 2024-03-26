package storage

func (s *PPStorageDB) Close() error {
	s.cluster.Close()
	return nil
}
