package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)
const (
	MAC = "/Applications/Blender/blender.app/Contents/MacOS/blender"
	//LIN = "" // Needs discussion.
	//WIN = "" // Needs discussion.
)

// renderPath function receives the Blender filepath and
// returns the relative path required for rendering.
func renderPath(blenderPath string) string {
	dir := filepath.Dir(blenderPath)
	ext := filepath.Ext(blenderPath)
	if ext != ".blend" {
		fmt.Fprintln(os.Stderr, "not blender file type.")
		os.Exit(1)
	}
	filename := filepath.Base(blenderPath)
	filename = filename[:len(filename)-len(ext)]
	return fmt.Sprintf("%s/%s/%s.####.exr", dir, filename, filename)
}

func main() {
	_, err := os.Stat(MAC)
	if err != nil {
		fmt.Fprintf(os.Stderr, "No Blender App at %s.\n", MAC)
		fmt.Fprintln(os.Stderr, "Need Install : https://www.blender.org/download/")
		os.Exit(1)
	}
	if len(os.Args) != 2 {
		fmt.Println("How to use :")
		fmt.Println(" $ blrender <renderfile.blend>")
		os.Exit(1)
	}
	cwdpath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	// check path Absolute or Relative
	blenderPath := os.Args[1]
	if os.Args[1] == filepath.Base(os.Args[1]) {
		blenderPath = cwdpath + "/" + filepath.Base(os.Args[1])
	}
	_, err = os.Stat(blenderPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "blender file does not exist.")
		os.Exit(1)
	}
	// render
	cmd := exec.Command(MAC, "-b", blenderPath, "-o", renderPath(blenderPath), "-F", "EXR", "-a")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
