package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
)

func GetShellType() string {
	var shell = os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/bash"
	}
	return shell
}

func Run(shell string) string {
	var cmd *exec.Cmd
	if runtime.GOOS == "linux" {
		cmd = exec.Command(GetShellType(), "-c", shell)
	} else {
		cmd = exec.Command("cmd.exe ", "/c", shell)
	}
	buf, _ := cmd.Output()
	buf2 := fmt.Sprintf("%s", buf)
	if buf2 == "" {
		buf2 = "empty"
	}
	return buf2
}
func GetAudio(video string) string {

	name := path.Base(video)
	cmd := "ffmpeg -y -i %s -ar 16000 output/%s_audio.wav"
	Run(fmt.Sprintf(cmd, video, name))
	return fmt.Sprintf("output/%s_audio.wav", name)
}
func MergeText(video string, textpath string) string {
	name := path.Base(video)
	cmd := "ffmpeg -y -i %s -vf subtitles=output/video.srt output/merge_%s.mp4"
	Run(fmt.Sprintf(cmd, video, name))
	return fmt.Sprintf("output/merge_%s.mp4", name)
}
