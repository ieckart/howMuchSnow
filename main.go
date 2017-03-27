package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/ieckart/howMuchSnow/mountain"
	"github.com/ieckart/howMuchSnow/multipass"
	"github.com/ieckart/howMuchSnow/twopass"
	"os"
	"time"
)

func main() {
	seedFlag := flag.Int64("s", int64(0), "int64 seed used to generate the mountain. Use 0 or unset for a random seed")
	multipassFlag := flag.Bool("m", false, "Solve using multipass")
	twoFlag := flag.Bool("t", false, "Solve using twopass")
	printFlag := flag.Bool("p", false, "Print the mountain as a string")
	demoFlag := flag.Int("d", 0, "Number of the demo to run, short:1, medium:2, long:3")

	// loop flags
	loopFlag := flag.Int("l", 0, "Number of times to run the tests in each step of the loop")
	fileFlag := flag.String("f", "run.csv", "File to dump output to when using -l")
	minWidthFlag := flag.Int("w", 10, "Min mountain width to start the loop at")
	maxWidthFlag := flag.Int("W", 1000, "Max mountain width to end the loop at")
	stepFlag := flag.Int("S", 1, "The number to increase the width by each time over the loop")

	flag.Parse()

	var mountMap, snowMap []int

	if *demoFlag != 0 {
		switch *demoFlag {
		case 1:
			demo1()
		case 2:
			demo2()
		case 3:
			demo3()
		default:
			fmt.Printf("No Demo %d", *demoFlag)
		}
	} else if *loopFlag != 0 { // run multipass and twopass x number of times, compare them, and dump output to a csv file
		runloops(*loopFlag, *minWidthFlag, *maxWidthFlag, *stepFlag, *fileFlag)
	} else {

		// generate the mountain
		seed := *seedFlag
		if seed == int64(0) {
			mountMap, seed = mountain.GetMountain()
		} else {
			mountMap, _ = mountain.GetMountainS(seed)
		}
		fmt.Printf("seed:%d\n", seed)

		// run multipass
		if *multipassFlag {
			snowMap = runThis("Muiltipass", multipass.HowMuchSnow, mountMap)
		}

		// run twopass
		if *twoFlag {
			snowMap = runThis("TwoPass", twopass.HowMuchSnow, mountMap)
		}

		// print the mountain
		if *printFlag {
			printMountWithSnow(mountMap, snowMap)
		}
	}
}

// runs any function that has the "func([]int)(int, []int)" signature and prints the runtime
func runThis(name string, f func([]int) (int, []int), mountMap []int) []int {
	start := time.Now()
	snow, snowMap := f(mountMap)
	fmt.Printf("%s: %d snow units | Runtime: %d Nanoseconds\n", name, snow, time.Since(start).Nanoseconds())
	return snowMap
}

// prints the mountain defined by mountMap with snow on it defined by snowMap along with a nicely formated array strings
func printMountWithSnow(mountMap, snowMap []int) {
	fmt.Println(mountain.GetMountainSnowString(mountMap, snowMap))
	fmt.Printf("Mountain: %s\n", mountain.FormatArr(mountMap))
	fmt.Printf("Snow: %s\n", mountain.FormatArr(snowMap))
}

// runs and compares twopass and multipass and dumps the output to a csv file
func runloops(loops, min, max, step int, fileName string) {

	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		fmt.Printf("Cannot create file: %s", err)
	}

	var (
		start     time.Time
		multiTime int64
		twoTime   int64
		records   [][]string
	)
	records = append(records, []string{"Width", "Multipass", "Twopass"})

	for width := min; width <= max; width += step {
		for i := 0; i < loops; i++ {
			mountMap, _ := mountain.GetMountainW(width)

			// multipass
			start = time.Now()
			multipass.HowMuchSnow(mountMap)
			multiTime += time.Since(start).Nanoseconds()

			// twoPass
			start = time.Now()
			twopass.HowMuchSnow(mountMap)
			twoTime += time.Since(start).Nanoseconds()
		}
		multiAv := fmt.Sprintf("%d", multiTime/int64(loops))
		twoAv := fmt.Sprintf("%d", twoTime/int64(loops))
		w := fmt.Sprintf("%d", width)
		records = append(records, []string{w, multiAv, twoAv})
	}

	w := csv.NewWriter(file)
	for _, record := range records {
		if err := w.Write(record); err != nil {
			fmt.Printf("error writing record to csv: %s", err)
		}
	}

	defer w.Flush()
}

// some demos I used in a presentation
func demo1() {
	mountMap := []int{1, 3, 4, 5, 6, 6, 5, 6, 8, 8, 7, 6, 5, 5, 6, 6, 5, 4, 3, 2, 2, 1}
	runThis("Muiltipass", multipass.HowMuchSnow, mountMap)
	snowMap := runThis("Twopass", twopass.HowMuchSnow, mountMap)
	printMountWithSnow(mountMap, snowMap)
}

func demo2() {
	mountMap := []int{
		21, 12, 11, 4, 23, 25, 30, 5, 5, 4, 6, 23, 30, 9, 21, 15,
		36, 2, 19, 22, 5, 11, 26, 16, 18, 27, 23, 25, 24, 13, 26,
		9, 12, 9, 28, 15, 27, 38, 9, 24, 10, 30, 8, 32, 11, 7, 37,
		8, 15, 12, 35, 36, 24, 10, 30, 8, 10, 37, 3, 9, 1, 6, 30,
		18, 22, 29, 35, 37, 4, 9, 29, 27, 18, 17, 30, 3, 19, 21,
		16, 25, 4, 17, 39, 19, 5, 37, 27, 30, 30, 26, 32, 16, 17,
		5, 25, 1, 39}
	runThis("Muiltipass", multipass.HowMuchSnow, mountMap)
	snowMap := runThis("Twopass", twopass.HowMuchSnow, mountMap)
	printMountWithSnow(mountMap, snowMap)
}

func demo3() {
	mountMap := []int{
		58, 83, 42, 65, 21, 44, 42, 40, 10, 5, 51, 36, 35, 49, 37,
		39, 11, 24, 26, 25, 42, 81, 5, 49, 7, 83, 96, 37, 15, 16,
		80, 21, 64, 60, 15, 48, 66, 96, 49, 50, 27, 41, 52, 81, 34,
		38, 82, 14, 43, 49, 37, 26, 88, 22, 28, 85, 27, 76, 36, 25,
		16, 12, 82, 92, 45, 57, 52, 30, 47, 43, 22, 49, 49, 14, 60,
		87, 79, 95, 25, 89, 30, 22, 39, 48, 15, 66, 62, 7, 18, 51,
		88, 7, 68, 93, 16, 14, 37, 16, 65, 57, 62, 15, 15, 99, 38,
		86, 22, 91, 78, 25, 97, 76, 86, 52, 88, 10, 10, 2, 25, 64,
		84, 23, 12, 53, 47, 62, 70, 95, 48, 42, 92, 54, 21, 67, 55,
		72, 39, 11, 23, 26, 9, 79, 94, 65, 84, 77, 90, 86, 11, 41,
		64, 71, 29, 30, 31, 8, 78, 19, 86, 50, 88, 72, 36, 72, 61,
		66, 11, 68, 56, 78, 70, 26, 4, 8, 19, 37, 23, 46, 55, 19, 90,
		52, 37, 57, 45, 67, 79, 53, 70, 91, 32, 73, 50, 10, 97, 71,
		51, 97, 17, 9, 51, 20, 64, 94, 26, 62, 91, 29, 84, 4, 43, 80}
	runThis("Muiltipass", multipass.HowMuchSnow, mountMap)
	snowMap := runThis("Twopass", twopass.HowMuchSnow, mountMap)
	printMountWithSnow(mountMap, snowMap)
}
