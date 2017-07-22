package query_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/erichnascimento/go-sqlquery"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var (
	db   *sql.DB
	mock sqlmock.Sqlmock
)

type pet struct {
	id       int64
	name     string
	isPoodle bool
	birthday time.Time
}

func TestQueryWithValidData(t *testing.T) {
	expects := mockPetsForSuccess()

	result, err := query.Query(db, "SELECT * FROM pets WHERE id >= ?", 1)
	if err != nil {
		t.Errorf(`err = %#v, want = %#v`, err, nil)
	}
	defer result.Close()

	for i := 0; result.Next(); i++ {
		row, err := result.Read()
		if err != nil {
			t.Errorf(`err = %#v, want = %#v`, err, nil)
		}

		pet := expects[i]

		if id := row["id"].AsInt64().Int64; id != pet.id {
			t.Errorf("id = %#v, want %#v", id, pet.id)
		}

		if name := row["name"].AsString().String; name != pet.name {
			t.Errorf("name = %#v, want %#v", name, pet.name)
		}

		if isPoodle := row["is_poodle"].AsBool().Bool; isPoodle != pet.isPoodle {
			t.Errorf("is_poodle = %#v, want %#v", isPoodle, pet.isPoodle)
		}

		if bithday := row["bithday"].AsTime().Time; !pet.birthday.Equal(bithday) {
			t.Errorf("bithday = %#v, want %#v", bithday, pet.birthday)
		}
	}
}

func mockPetsForSuccess() []*pet {
	dj, _ := time.Parse("2006-01-02 15:04:05", "2011-01-03 12:20:00")
	dz, _ := time.Parse("2006-01-02", "2006-11-23")

	pets := []*pet{
		&pet{1, "Jacob", true, dj},
		&pet{2, "Zuca", false, dz},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "is_poodle", "bithday"})
	for _, p := range pets {
		rows.AddRow(p.id, p.name, p.isPoodle, p.birthday)
	}
	mock.ExpectQuery(`^SELECT \* FROM pets.*`).WillReturnRows(rows)

	return pets
}

func TestMain(m *testing.M) {
	setup()
	defer os.Exit(m.Run())
	tearDown()
}

func setup() {
	var err error
	db, mock, err = sqlmock.New()
	if err != nil {
		panic(fmt.Sprintf("an error '%s' was not expected when opening a stub database connection", err))
	}
}

func tearDown() {
	db.Close()
}
