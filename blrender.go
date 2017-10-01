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
	//LIN = ""
	//WIN
)

// renderPath 함수는 블랜더파일을 받아서 렌더링에 필요한 렌더경로를 반환한다.
func renderPath(blenderPath string) string {
	ext := filepath.Ext(blenderPath)
	if ext != ".blend" {
		fmt.Fprintln(os.Stderr, "no blender file")
		os.Exit(1)
	}
	filename := filepath.Base(blenderPath)
	filename = filename[:len(filename)-len(ext)]
	return fmt.Sprintf("./%s/%s.####.exr", filename, filename)
}

func main() {
	_, err := os.Stat(MAC)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s 위치에 blender가 없습니다.\n", MAC)
		fmt.Fprintln(os.Stderr, "블렌더설치가 필요합니다.")
		os.Exit(1)
	}
	if len(os.Args) != 2 {
		fmt.Println("How to use :")
		fmt.Println("$ blrender <blender filename>")
		os.Exit(1)
	}
	cwdpath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	blenderPath := cwdpath + "/" + filepath.Base(os.Args[1])
	// blender파일에서 설정된 프레임값으로 exr렌더링 한다.
	cmd := exec.Command(MAC, "-b", blenderPath, "-o", renderPath(blenderPath), "-F", "EXR", "-a")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
