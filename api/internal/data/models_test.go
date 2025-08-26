package data

import "testing"

func Test_Ping(t *testing.T) {
	err := testDB.Ping()
	if err != nil {
		t.Fatal(err)
	}
}

func TestBook_GetAll(t *testing.T) {
	all, err := models.Book.GetAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(all) != 1 {
		t.Fatal("expected 1 book, got ", len(all))
	}
}

func TestBook_GetByID(t *testing.T) {
	b, err := models.Book.GetOneById(1)
	if err != nil {
		t.Fatal(err)
	}

	if b == nil {
		t.Fatal("expected book, got nil")
	}

	if b.Title != "My Book" {
		t.Errorf("expected 'My Book', got %s", b.Title)
	}
}

func TestBook_GetBySlug(t *testing.T) {
	b, err := models.Book.GetOneBySlug("my-book")
	if err != nil {
		t.Error("failed to get book by slug", err)
	}

	if b == nil {
		t.Fatal("expected book, got nil")
	}

	if b.Title != "My Book" {
		t.Errorf("expected 'My Book', got %s", b.Title)
	}
	_, err = models.Book.GetOneBySlug("bad-slug")
	if err == nil {
		t.Error("expected error, got nil")
	}
}
