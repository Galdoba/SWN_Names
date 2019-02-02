package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

/*
TZ:
1 прочитать задание
2 определить файлы которые надо ждать
3 найти нужные файлы (если их нет - ждать 2 часа с запросом каждую минуту)
4 проверить нужно ли резать видео


*/

var inputVideoDuration vidDuration

type vidDuration struct {
	frames       int
	timeCodePrem string
	timeCodeFF   string
}

func NewDuration(premiereCode string) *vidDuration {
	dur := vidDuration{}
	dur.timeCodePrem = premiereCode
	dur.timeCodeFF = premToFF(premiereCode)
	dur.frames = premToFrames(premiereCode)
	return &dur
}

func main() {
	// durationFFMPEG := videoDuration("video.mp4")
	// fmt.Println("---------------", durationFFMPEG)
	// premiereDur := ffToPrem(durationFFMPEG)
	// fmt.Println("---------------", premiereDur)
	// VideoDuration := NewDuration(premiereDur)
	// fmt.Println(VideoDuration)
	//durFF := premToFF(premiereDur)

	vid1Found := false
	aud1Found := false
	aud2Found := false

	for i := 0; i < 1000; i++ {

		fmt.Println("Loop", i)
		vid1Found, _ = fileExists("video.mp4")
		if vid1Found {
			fmt.Println(fileAvailable("video.mp4"))
		}
		aud1Found, _ = fileExists("audio1.ac3")
		aud2Found, _ = fileExists("audio2.ac3")

		fmt.Println("Video found:", vid1Found)
		fmt.Println("Audio 1 found:", aud1Found)
		fmt.Println("Audio 2 found:", aud2Found)
		if vid1Found == true && aud1Found == true && aud2Found == true {
			testTask()
			i = 2000
		}
		time.Sleep(time.Second * 1)
	}
}

func testTask() {
	tasks := readTasks()
	for i := range tasks {
		fmt.Println(tasks[i])
		taskArgs := strings.Split(tasks[i], " ")
		taskLen := taskArgs[1]
		taskStart := taskArgs[2]
		taskInput := taskArgs[0]
		output := cutVideo(taskStart, taskLen, taskInput)
		fmt.Println("OUTPUTFILE:", output)
		time.Sleep(time.Second * 1)
		if fileAvailable(output) {
			fmt.Println(videoDuration(output))
		} else {
			time.Sleep(time.Second * 1)
			fmt.Println("Boom")
		}
	}
}

func cutVideo(timeStart, timeLen, file string) string {
	outputFile := "cut_" + file
	cmd := exec.Command("ffmpeg", "-ss", premToFF(timeStart), "-i", file, "-c", "copy", "-t", premToFF(timeLen), "-q", "0", "-loglevel", "verbose", outputFile)
	// stdout, errOut := cmd.Output()
	// fmt.Print(errOut)
	// if err := cmd.Run(); err != nil {
	// 	fmt.Println(err)
	// 	log.Fatal(err)
	// }
	// fmt.Print(string(stdout))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	return outputFile
}

// exists returns whether the given file or directory exists
func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	fmt.Println(err)
	return true, err
}

func fileAvailable(file string) bool {
	err := os.Rename(file, "RENAMED_"+file)
	if err != nil {
		fmt.Println(err)
		return false
	}
	os.Rename("RENAMED_"+file, file)
	return true
}

func readTasks() []string {
	var tasks []string
	file, err := os.Open("FFtask.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tasks = append(tasks, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return tasks
}

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}

func videoDuration(file string) string {
	cmd := exec.Command("ffmpeg", "-i", "video.mp4")

	output, _ := cmd.CombinedOutput()
	stringOUT := string(output)
	str1 := strings.Split(stringOUT, "Duration: ")
	time.Sleep(time.Second * 1)
	if len(str1) > 0 {
		durationSTR := strings.Split(str1[1], ", ")
		return durationSTR[0]
	}
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	//cmd.Run()
	return "00:00:00.00"
}

func ffToPrem(duration string) string {
	premiereDur := ""
	timeArgs := strings.Split(duration, ":")
	hours, errHour := strconv.Atoi(timeArgs[0])
	if errHour != nil {
		panic(errHour)
	}
	minutes, errMin := strconv.Atoi(timeArgs[1])
	if errMin != nil {
		panic(errMin)
	}
	secs, err := strconv.ParseFloat(timeArgs[2], 64)
	if err != nil {
		panic(err)
	}
	secsInt := int(secs)
	strHour := strconv.Itoa(hours)
	if hours < 10 {
		strHour = "0" + strHour
	}
	strMin := strconv.Itoa(minutes)
	if minutes < 10 {
		strMin = "0" + strMin
	}
	strSec := strconv.Itoa(secsInt)
	if secsInt < 10 {
		strSec = "0" + strSec
	}
	part := secs
	for part > 1 {
		part = part - 1
	}
	part = toFixed(part, 3)
	part = part / 0.04
	frame := int(part)
	strFrame := strconv.Itoa(frame)
	if frame < 10 {
		strFrame = "0" + strFrame
	}
	premiereDur = strHour + ":" + strMin + ":" + strSec + ":" + strFrame
	return premiereDur
}

func floatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 2, 64)
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func premToFF(duration string) string {
	parts := strings.Split(duration, ":")
	secsInt, _ := strconv.Atoi(parts[2])
	partsInt, _ := strconv.Atoi(parts[3])
	partsFl := float64(partsInt)*40/1000 + float64(secsInt)
	sec := floatToString(partsFl)
	return parts[0] + ":" + parts[1] + ":" + sec
}

func premToFrames(duration string) int {
	parts := strings.Split(duration, ":")
	hour, _ := strconv.Atoi(parts[0])
	min, _ := strconv.Atoi(parts[1])
	sec, _ := strconv.Atoi(parts[2])
	frm, _ := strconv.Atoi(parts[3])
	return frm + 25*sec + 1500*min + 90000*hour
}

func exe_cmd(cmd string, wg *sync.WaitGroup) {
	fmt.Println("command is ", cmd)
	// splitting head => g++ parts => rest of the command
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
	wg.Done() // Need to signal to waitgroup that this goroutine is done
}
