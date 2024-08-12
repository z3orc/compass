package repo

import (
	"testing"

	"github.com/z3orc/compass/internal/data"
)

func TestNewVersionRepositoryOneSource(t *testing.T) {
	src := data.NewPistonDataSource()

	repo := NewVersionRepository(src)
	if repo == nil {
		t.Fatal("Expected *VersionRepoistory got nil")
	}

	flavours := repo.GetFlavours()
	if len(flavours) != 1 {
		t.Fatalf("Expected 1 flavour got %v", len(flavours))
	}
}

func TestNewVersionRepositoryTwoSources(t *testing.T) {
	src := data.NewPistonDataSource()

	repo := NewVersionRepository(src, src)
	if repo == nil {
		t.Fatal("Expected *VersionRepoistory got nil")
	}

	flavours := repo.GetFlavours()
	if len(flavours) != 2 {
		t.Fatalf("Expected 2 flavours got %v", len(flavours))
	}
}

func TestNewVersionRepositoryNoSource(t *testing.T) {
	repo := NewVersionRepository()

	if repo != nil {
		t.Fatalf("Expected a nil pointer but got %T", repo)
	}
}
