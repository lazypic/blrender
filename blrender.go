package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"log"
	"strings"
	"fmt"
	)
func path2renderpath(path string) string {
	//test.blend -> ./test/test.####.exr"
	_, file := filepath.Split(path)
	ext := filepath.Ext(file)
	if ext != ".blend" {
		fmt.Println("블랜더파일이 아닙니다.")
		os.Exit(1)
	}
	filename := strings.Split(file, ext)[0]
	return fmt.Sprintf("./%s/%s.####.exr", filename, filename)
}
func main() {
	cwdpath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	osx := "/Applications/Blender/blender.app/Contents/MacOS/blender"
	if len(os.Args) == 1 {
		fmt.Println("사용법 :\n $ blrender <blender filename>")
		os.Exit(1)
	}

	if _, err := os.Stat(osx); err == nil {
		bfile := cwdpath + "/" + os.Args[1]
		cmd := exec.Command(osx, "-b", bfile, "-o", path2renderpath(bfile), "-F", "EXR", "-a")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}
