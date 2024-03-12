package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"homework/internal/pick-up_point/model"
)

const fileName = "storage_pp.json"

type PPStorage interface {
	AddPickUpPoint(point model.PickUpPoint) error
	GetPickUpPoints() ([]model.PickUpPoint, error)
	GetPickUpPointByName(name string) (*model.PickUpPoint, error)
	UpdatePickUpPoint(point model.PickUpPoint) error
	Close() error
}

type FilePPStorage struct {
	file *os.File
	cash []model.PickUpPoint
	mu   *sync.RWMutex
}

func New(logChan chan<- string) (*FilePPStorage, error) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, fmt.Errorf("error in opening file")
	}
	filePPStorage := &FilePPStorage{
		file: file,
		mu:   &sync.RWMutex{},
		cash: make([]model.PickUpPoint, 0),
	}
	err = filePPStorage.getCash()
	go filePPStorage.processSavingCash(logChan)
	return filePPStorage, err

}

func (fs *FilePPStorage) getCash() error {
	decoder := json.NewDecoder(fs.file)
	var pickUpPoints []model.PickUpPoint
	if err := decoder.Decode(&pickUpPoints); err != nil {
		if !errors.Is(err, io.EOF) {
			return err
		}
	}
	_, err := fs.file.Seek(0, 0)
	if err != nil {
		return err
	}
	fs.cash = pickUpPoints
	return nil
}

func (fs *FilePPStorage) SaveCashToFile() error {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	_, err := fs.file.Seek(0, 0)
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(fs.file)
	return encoder.Encode(fs.cash)
}

func (fs *FilePPStorage) processSavingCash(logChan chan<- string) {
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		err := fs.SaveCashToFile()
		var logMessage string
		if err != nil {
			logMessage = fmt.Sprintf("Log info: error in saving cash into file: %s", err)
		} else {
			logMessage = "Log info: cash was successfully saved into file"
		}
		logChan <- logMessage
	}
}
