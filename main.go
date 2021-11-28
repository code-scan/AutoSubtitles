package main

import (
	"fmt"
	"log"
	"os"
	"video/config"
	"video/pkg"
)

func OSS(audio string) string {
	oss := pkg.NewOSS(config.AccessKeyID, config.AccessKey, config.EndPoint, config.BucketName)
	uri := oss.UploadFile(audio)
	return uri
}
func Zimu(uri string) {
	zimu := pkg.NewZimu(config.AccessKeyID, config.AccessKey, config.Appkey)
	zimu.StartConvert(uri)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("video.exe path/of/xxx.mp4")
		os.Exit(1)
	}
	log.Println("[*] GetAudio ", os.Args[1])
	audio := pkg.GetAudio(os.Args[1])
	log.Println("[*] Upload Audio ", audio)
	uri := OSS(audio)
	log.Println("[*] Upload Link : ", uri)
	log.Println("[*] Convert Start")
	Zimu(uri)
	log.Println("[*] Convert End")
	log.Println("[*] Merge Video Start")
	result := pkg.MergeText(os.Args[1], "output/video.srt")
	log.Println("[*] Merge Video End")
	log.Println("[*] Result: ", result)

}

// ffmpeg -i test.mkv -vf subtitles=video.srt output.mkv
