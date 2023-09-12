package database

import (
	"testing"
)

func TestConnectionDb(t *testing.T) {
	// check to See any err when create instance
	var client = DBinstance()
	OpenCollection(client, "menu")
}
