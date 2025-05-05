package csv

import (
	encsv "encoding/csv"
	"errors"
	"layersapi/entities"
	"os"
	"path/filepath"
	"time"
)

type UserRepository struct {
	filePath string
}

func NewUserRepository(path ...string) *UserRepository {
	fp := "data/data.csv"
	if len(path) > 0 && path[0] != "" {
		fp = path[0]
	}
	fp = filepath.Clean(fp)
	return &UserRepository{filePath: fp}
}

func (u *UserRepository) GetAll() ([]entities.User, error) {
	file, err := os.Open(u.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := encsv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var result []entities.User
	for i, record := range records {
		if i == 0 {
			continue // skip header
		}
		createdAt, _ := time.Parse(time.RFC3339, record[3])
		updatedAt, _ := time.Parse(time.RFC3339, record[4])
		meta := entities.Metadata{
			CreatedAt: createdAt.String(),
			UpdatedAt: updatedAt.String(),
			CreatedBy: record[5],
			UpdatedBy: record[6],
		}
		result = append(result, entities.NewUser(record[0], record[1], record[2], meta))
	}
	return result, nil
}

func (u *UserRepository) GetById(id string) (entities.User, error) {
	file, err := os.Open(u.filePath)
	if err != nil {
		return entities.User{}, err
	}
	defer file.Close()

	reader := encsv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return entities.User{}, err
	}

	for i, record := range records {
		if i == 0 {
			continue
		}
		if record[0] == id {
			createdAt, _ := time.Parse(time.RFC3339, record[3])
			updatedAt, _ := time.Parse(time.RFC3339, record[4])
			meta := entities.Metadata{
				CreatedAt: createdAt.String(),
				UpdatedAt: updatedAt.String(),
				CreatedBy: record[5],
				UpdatedBy: record[6],
			}
			return entities.NewUser(record[0], record[1], record[2], meta), nil
		}
	}
	return entities.User{}, errors.New("user not found")
}

func (u *UserRepository) Create(user entities.User) error {
	file, err := os.OpenFile(u.filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := encsv.NewWriter(file)
	defer writer.Flush()

	newUser := []string{
		user.Id,
		user.Name,
		user.Email,
		user.Metadata.CreatedAt,
		user.Metadata.UpdatedAt,
		user.Metadata.CreatedBy,
		user.Metadata.UpdatedBy,
	}
	if err := writer.Write(newUser); err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) Update(id, name, email string) error {
	records, err := ReadAllFromFile(u.filePath)
	if err != nil {
		return err
	}

	updated := false
	for i, record := range records {
		if i == 0 {
			continue
		}
		if record[0] == id {
			record[1] = name
			record[2] = email
			record[4] = time.Now().Format(time.RFC3339) // update UpdatedAt
			records[i] = record
			updated = true
			break
		}
	}
	if !updated {
		return errors.New("user not found")
	}
	return WriteAllToFile(u.filePath, records)
}

func (u *UserRepository) Delete(id string) error {
	records, err := ReadAllFromFile(u.filePath)
	if err != nil {
		return err
	}

	var updatedRecords [][]string
	deleted := false
	for i, record := range records {
		if i == 0 {
			updatedRecords = append(updatedRecords, record) // keep header
			continue
		}
		if record[0] == id {
			deleted = true
			continue
		}
		updatedRecords = append(updatedRecords, record)
	}
	if !deleted {
		return errors.New("user not found")
	}
	return WriteAllToFile(u.filePath, updatedRecords)
}
