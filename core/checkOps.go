package core

import "runtime"

func checkOps() string {
	ops := runtime.GOOS
	sep := ""
	switch ops {
	case "windows":
		sep = "\r\n"
	case "linux":
		sep = "\n"
	}
	return sep
}
