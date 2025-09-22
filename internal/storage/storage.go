package storage

import "github.com/nithulj2/students-api/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentBYId(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
}
