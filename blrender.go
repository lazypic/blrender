package main

import (
	"os"
	"os/exec"
	)

func file2renderpath(filename string) string {
	//test.blend -> ./test/test.####.exr"
	filename := strings.Split(filename, ".blend")[0]
}

func main() {
	osx := "/Applications/Blender/blender.app/Contents/MacOS/blender"
	if _, err := os.Stat(osx); err == nil {
		cmd := exec.Command(osx, "-b", os.Args[1], "-o", "./test/test.####.exr", "-F", "EXR", "-a")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}
