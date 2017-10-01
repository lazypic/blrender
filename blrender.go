package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	MAC = "/Applications/Blender/blender.app/Contents/MacOS/blender"
)

func renderPath(blenderPath string) string {
	ext := filepath.Ext(blenderPath)
	if ext != ".blend" {
		fmt.Fprintln(os.Stderr, "no blender file")
		os.Exit(1)
	}
	filename := filepath.Base(blenderPath)
	filename := filename[:len(filename)-len(ext)]
	return fmt.Sprintf("./%s/%s.####.exr", filename, filename)
}

func main() {
	cwdpath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) != 2 {
		fmt.Println("How to use :\n $ blrender <blender filename>")
		os.Exit(1)
	}

	_, err := os.Stat(MAC)
	if err != nil {
		fmt.Fprintln(os.Stderr, "블렌더설치가 필요함.")
		os.Exit(1)
	}
	bfile := cwdpath + "/" + os.Args[1]
	cmd := exec.Command(MAC, "-b", bfile, "-o", renderPath(blenderPath), "-F", "EXR", "-a")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
