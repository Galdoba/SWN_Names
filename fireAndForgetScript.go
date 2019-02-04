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

type currentTask struct {
	input     string
	timeLen   string
	timeStart string
	result    string
}

func NewDuration(premiereCode string) *vidDuration {
	dur := vidDuration{}
	dur.timeCodePrem = premiereCode
	dur.timeCodeFF = premToFF(premiereCode)
	dur.frames = premToFrames(premiereCode)
	return &dur
}

func main() {
	video1 := ""
	audio1 := ""
	audio2 := ""
	shortNme := ""
	tasks := readTasks()
	for i := range tasks {
		///////////////////////////////////////////собираем информацию о задании
		taskArgs := strings.Split(tasks[i], " ")
		video1 = taskArgs[0]
		shortNme = shortName(video1)
		fmt.Println(shortNme)
		taskLen := taskArgs[1]
		taskStart := taskArgs[2]
		intendedResult := taskArgs[3]
		audioFiles := predictAudioNames(intendedResult)
		//audioFiles := pingTask(tasks[i])
		switch len(audioFiles) {
		case 1:
			audio1 = audioFiles[0]
		case 2:
			audio1 = audioFiles[0]
			audio2 = audioFiles[1]
		}
		////////////////////////////////////////проверяеи и режем видео если надо
		fmt.Println("Uncut Input Video:", video1)
		for !fileAvailable(video1) {
			time.Sleep(time.Second * 5)
		}
		fmt.Println("Input File Available")
		inputDuration := ffToPrem(videoDuration(video1))
		fmt.Println("Input Video Duration:", inputDuration)
		fmt.Println("Intended Duration:", taskLen)
		if inputDuration == taskLen {
			renameFile(video1, "cut_"+video1)
		} else {
			cutVideo(taskStart, taskLen, video1)
		}
		//////////////////////////////////////собираем все нужные файлы
		fmt.Println("cut_"+video1, fileAvailable("cut_"+video1))
		fmt.Println(audio1, fileAvailable(audio1))
		//fmt.Println("cut_"+video1, fileAvailable("cut_"+video1))

		//
		return
		vid1Found := false
		aud1Found := false
		aud2Found := false

		for i := 0; i < 1000; i++ {

			//fmt.Println("Loop", i)
			vid1Found, _ = fileExists(video1)
			if vid1Found {
				//fmt.Println(fileAvailable("video.mp4"))
			}
			aud1Found, _ = fileExists(audio1)
			aud2Found, _ = fileExists(audio2)

			fmt.Println("Video found:", vid1Found)
			fmt.Println("Audio 1 found:", aud1Found)
			fmt.Println("Audio 2 found:", aud2Found)
			if vid1Found == true && aud1Found == true && aud2Found == true {
				vid1Found = testTask()
				if vid1Found != true {
					i--
				} else {
					i = 2000
				}
			}
			time.Sleep(time.Second * 1)
		}
	}
}

func shortName(fullName string) string {
	validExtentions := []string{
		".mp4",
		".mpa",
		".ac3",
		".aac",
	}
	for i := range validExtentions {
		if strings.Contains(fullName, validExtentions[i]) {
			return strings.Split(fullName, validExtentions[i])[0]
		}
	}
	return ""
}

func pingTask(task string) []string {

	fmt.Println("Current task:", task)
	taskArgs := strings.Split(task, " ")
	taskLen := taskArgs[1]
	taskStart := taskArgs[2]
	taskResult := taskArgs[3]
	taskInput := taskArgs[0]
	if taskLen == ffToPrem(videoDuration(taskInput)) {
		fmt.Println("No need to cut")
	} else {
		fmt.Println("Need to cut:", taskStart, "from start")
	}
	names := predictAudioNames(taskResult)
	return names
}

func predictAudioNames(resultName string) []string {
	var expected []string
	//Pod_solncem_toskany_2003__hd_q0w1_ar2e2
	tagPool := strings.Split(resultName, "__")
	if len(tagPool) < 2 {
		fmt.Println(tagPool)
		return expected
	}
	allTags := tagPool[1]
	tags := strings.Split(allTags, "_")
	audioTag := tags[len(tags)-1]
	fmt.Println(audioTag)
	if strings.Contains(allTags, "sd_") {
		if strings.Contains(audioTag, "r2") {
			fmt.Println("case sd r2")
			audName := tagPool[0] + "_"
			for t := range tags {
				if t == len(tags)-1 {
					break
				}
				audName = audName + "_" + tags[t]
			}
			audName = audName + "_rus20.mpa"
			expected = append(expected, audName)
		}

		if strings.Contains(audioTag, "e2") {
			fmt.Println("case sd e2")
			audName := tagPool[0] + "_"
			for t := range tags {
				if t == len(tags)-1 {
					break
				}
				audName = audName + "_" + tags[t]
			}
			audName = audName + "_eng20.mpa"
			expected = append(expected, audName)
		}
	}

	if strings.Contains(allTags, "hd_") {
		if strings.Contains(audioTag, "r2") {
			fmt.Println("case hd r2")
			audName := tagPool[0] + "_"
			for t := range tags {
				if t == len(tags)-1 {
					break
				}
				audName = audName + "_" + tags[t]
			}
			audName = audName + "_rus20.aac"
			expected = append(expected, audName)
		}

		if strings.Contains(audioTag, "e2") {
			fmt.Println("case e2")
			audName := tagPool[0] + "_"
			for t := range tags {
				if t == len(tags)-1 {
					break
				}
				audName = audName + "_" + tags[t]
			}
			audName = audName + "_eng20.aac"
			expected = append(expected, audName)
		}

		if strings.Contains(audioTag, "r6") {
			audName := tagPool[0] + "_"
			for t := range tags {
				if t == len(tags)-1 {
					break
				}
				audName = audName + "_" + tags[t]
			}
			audName = audName + "_rus51.aac"
			expected = append(expected, audName)
		}
		if strings.Contains(audioTag, "e6") {
			audName := tagPool[0] + "_"
			for t := range tags {
				if t == len(tags)-1 {
					break
				}
				audName = audName + "_" + tags[t]
			}
			audName = audName + "_eng51.aac"
			expected = append(expected, audName)
		}
	}

	fmt.Println("done")

	return expected
}

func testTask() bool {
	tasks := readTasks()
	for i := range tasks {
		fmt.Println("Current task:", tasks[i])
		taskArgs := strings.Split(tasks[i], " ")
		taskLen := taskArgs[1]
		taskStart := taskArgs[2]
		taskInput := taskArgs[0]
		output := cutVideo(taskStart, taskLen, taskInput)
		//fmt.Println("OUTPUTFILE:", output)
		time.Sleep(time.Second * 1)
		if fileAvailable(output) {
			//fmt.Println(videoDuration(output))
		} else {
			time.Sleep(time.Second * 10)
			//	fmt.Println("Boom")
			return false
		}
	}
	return true
}

func cutVideo(timeStart, timeLen, file string) string {
	outputFile := "cut_" + file
	if fileAvailable(file) {
		cmd := exec.Command("ffmpeg", "-i", file, "-map", "0:0", "-vcodec", "copy", "-an", "-t", premToFF(timeLen), "-ss", premToFF(timeStart), outputFile)
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

	} else {
		fmt.Println("Warning:", file, "is not available")
		fmt.Println("Waiting...")
		fmt.Println("")
	}
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

func renameFile(old, new string) error {
	err := os.Rename(old, new)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func fileAvailable(file string) bool {
	err := os.Rename(file, "RENAMED_"+file)
	if err != nil {
		fmt.Println("Warning: file", file, "is not available...")
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
	if fileAvailable(file) {
		cmd := exec.Command("ffmpeg", "-i", file)

		output, _ := cmd.CombinedOutput()
		stringOUT := string(output)
		str1 := strings.Split(stringOUT, "Duration: ")
		time.Sleep(time.Second * 1)
		if len(str1) > 0 {
			durationSTR := strings.Split(str1[1], ", ")
			return durationSTR[0]
		}
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

/*
mux_hd_ruseng_names
set CURRENTDIR=%CD%
set cwd=%~dp0
for /f "delims="  %%i IN ( 'type names.txt' ) DO (
call :process "%%i"
)

pause
exit /b 0

:process
set f_fullname=%~1
set f_path=%~p1
set f_name=%~n1
set f_ext=%~x1

rem ffmpeg -i %f_name%_rus.mpa -map 0:0 -acodec ac3 -ab 320k %f_name%_rus.ac3
rem ffmpeg -i %f_name%_eng.mpa -map 0:0 -acodec ac3 -ab 320k %f_name%_eng.ac3

mkvmerge -o "%f_name%_ar2e2.mkv" -d 0 --language 0:rus --default-track 0:1 -A ="d:\Work\petr_proj\__done\OUT\%f_name%.mp4" -a 0 --language 0:rus --default-track 0:1 ="%f_name%_rus.mpa" -a 0 --language 0:eng --default-track 0:0 ="%f_name%_eng.mpa"

rem -s 0 --language 0:rus --default-track 0:0 ="%f_name%.srt"

exit /b 0




mux_mutevideo_1audio_names_51
set CURRENTDIR=%CD%
set cwd=%~dp0

for /f "delims="  %%i IN ( 'type names.txt' ) DO (
call :process "%%i"
)

pause
exit /b 0

:process
set f_fullname=%~1
set f_path=%~p1
set f_name=%~n1
set f_ext=%~x1
ffmpeg -i "%f_fullname%" -i "%f_name%.aac" -vcodec copy -acodec ac3 -ab 640k "OUT_%f_name%_ar6%f_ext%"
exit /b 0



mux_mutevideo_1audio_names_20
set CURRENTDIR=%CD%
set cwd=%~dp0

for /f "delims="  %%i IN ( 'type names.txt' ) DO (
call :process "%%i"
)

pause
exit /b 0

:process
set f_fullname=%~1
set f_path=%~p1
set f_name=%~n1
set f_ext=%~x1
ffmpeg -i "%f_fullname%" -i "%f_name%.aac" -vcodec copy -acodec ac3 -ab 320k "OUT_%f_name%_ar2%f_ext%"
exit /b 0




mux_mutevideo_1audio_names_mpa
set CURRENTDIR=%CD%
set cwd=%~dp0

for /f "delims="  %%i IN ( 'type names.txt' ) DO (
call :process "%%i"
)

pause
exit /b 0

:process
set f_fullname=%~1
set f_path=%~p1
set f_name=%~n1
set f_ext=%~x1
ffmpeg -i "%f_fullname%" -i "%f_name%.mpa" -vcodec copy -acodec copy -ab 320k "OUT_%f_name%_ar2%f_ext%"
exit /b 0


mutevideo_audio_engrus5ch
set aext=ac3
set name=shest_dney_sem_nochey_1998__hd_q0w0
set vsrc=f:\Work\petr_proj\__done\OUT\shest_dney_sem_nochey_1998__hd_q0w0.mp4

ffmpeg -i "%name%_rus.aac" -acodec ac3 -ab 640k %name%_rus.%aext%
ffmpeg -i "%name%_eng.aac" -acodec ac3 -ab 640k %name%_eng.%aext%

mkvmerge -o "%name%_ar6e6.mkv" -d 0 --language 0:rus --default-track 0:1 -A ="%vsrc%" -a 0 --language 0:rus --default-track 0:1 ="%name%_rus.%aext%" -a 0 --language 0:eng --default-track 0:0 ="%name%_eng.%aext%"

rem  -s 0 --language 0:rus --default-track 0:0 ="%1.srt"

pause



*/
