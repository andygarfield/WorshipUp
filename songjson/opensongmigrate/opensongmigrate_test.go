package opensongmigrate

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Args = []string{"cmd", "/Users/andy/Dropbox/Opensong/Songs"}
	main()
}
