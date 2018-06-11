package path

import (
	"os"
	"path/filepath"
)

var Dir = filepath.Join(os.Getenv("HOME"), ".config", "twg")
var dir = Dir
var DirVerify = filepath.Join(dir, "verify.json")
var DirUser = filepath.Join(dir, "user.json")
var DirImg = filepath.Join(dir, "img")
