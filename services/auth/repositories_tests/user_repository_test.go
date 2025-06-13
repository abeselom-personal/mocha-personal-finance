package repositories_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/abeselom-personal/personal-finance/models"
	"github.com/abeselom-personal/personal-finance/repositories"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupUserRepo() (*repositories.UserRepository, sqlmock.Sqlmock, *gorm.DB) {
	db, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "postgres",
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})
	repo := repositories.NewUserRepository(gormDB)
	return repo, mock, gormDB
}

func TestUserRepository_Create(t *testing.T) {
	repo, mock, _ := setupUserRepo()

	user := &models.User{
		Username:     "testuser",
		Email:        "test@example.com",
		PasswordHash: "hashedpassword",
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users"`).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	err := repo.Create(user)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), user.ID)
}

func TestUserRepository_GetByUsername(t *testing.T) {
	repo, mock, _ := setupUserRepo()

	rows := sqlmock.NewRows([]string{"id", "username", "email"}).
		AddRow(1, "testuser", "test@example.com")

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE username = \$1 AND "users"\."deleted_at" IS NULL ORDER BY "users"\."id" LIMIT \$2`).
		WithArgs("testuser", 1).
		WillReturnRows(rows)

	mock.ExpectQuery(`SELECT \* FROM "user_permissions" WHERE "user_permissions"\."user_id" = \$1`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "permission"}))

	user, err := repo.GetByUsername("testuser")
	assert.NoError(t, err)
	assert.Equal(t, "testuser", user.Username)
}
