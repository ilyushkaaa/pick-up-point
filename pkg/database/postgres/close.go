package database

func (p *PGDatabase) Close() error {
	p.Cluster.Close()
	return nil
}
