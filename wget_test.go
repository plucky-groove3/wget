package wget_test

import (
	"os"
	"path"
	"testing"

	"github.com/plucky-groove3/wget"
)

func TestWgetDownloadFile(t *testing.T) {
	wget.Download("https://github.com/plucky-groove3/fap-demo/blob/master/LICENSE")
	wd, _ := os.Getwd()
	_, err := os.Open(path.Join(wd, "LICENSE"))
	if err != nil {
		t.Errorf("A file with name  was unable to be opened", "LICENSE")
	}

	wget.Download("https://github.com/plucky-groove3/fap-demo/blob/master", "README.rst")
	wd, _ = os.Getwd()
	_, err = os.Open(path.Join(wd, "README.rst"))
	if err != nil {
		t.Errorf("A file with name  was unable to be opened", "README.rst")
	}
}
