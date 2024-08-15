package repo

import (
	"testing"

	"github.com/z3orc/compass/internal/data"
	"github.com/z3orc/compass/internal/model"
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

func TestGetVersion(t *testing.T) {
	src := data.NewPistonDataSource()
	repo := NewVersionRepository(src)
	selectedVersion := "1.21"

	version, err := repo.GetVersion(model.FlavourPiston, selectedVersion)
	if err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}

	if version == nil {
		t.Fatal("Expected *model.Version got nil")
	}

	if version.Flavour != model.FlavourPiston {
		t.Fatalf("Expected flavour to be %v but got %v", model.FlavourPiston, version.Flavour)
	}

	if version.Id != selectedVersion {
		t.Fatalf("Expected id to be %v but got %v", selectedVersion, version.Id)
	}

	// if version.Url != "https://piston-meta.mojang.com/v1/packages/1.2112312/1.2112312.json" {
	// 	t.Fatalf("Expected url to be %v but got %v", "https://piston-meta.mojang.com/v1/packages/1.2112312/1.2112312.json", version.Url)
	// }

}

func TestGetVersionError(t *testing.T) {
	src := data.NewPistonDataSource()
	repo := NewVersionRepository(src)
	selectedVersion := "1.21"

	version, err := repo.GetVersion(model.FlavourPaper, selectedVersion)
	if err == nil {
		t.Fatal("Expected an error but got nil")
	}

	if version != nil {
		t.Fatalf("Expected a nil pointer but got %T", version)
	}
}
