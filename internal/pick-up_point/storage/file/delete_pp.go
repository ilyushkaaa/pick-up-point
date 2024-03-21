package storage

import "context"

// DeletePickUpPoint Добавил чтобы удовлетворять обновленному интерфейсу, чтобы до сих пор можно было пользоваться cli режимом
func (fs *FilePPStorage) DeletePickUpPoint(_ context.Context, _ uint64) (bool, error) {
	return false, nil
}
