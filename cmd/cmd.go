package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "werd",
	Short: "A CLI tool to help organize Quran revisioning.",
}

func init() {
	RootCmd.AddCommand(readCommand())
}

func readCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "read",
		Short: "Calculate daily reading plan based on specified range.",
		Args:  cobra.MinimumNArgs(3), // [start_verse, end_verse, days]
		Run: func(cmd *cobra.Command, args []string) {
			startVerse, _ := strconv.Atoi(args[0])
			endVerse, _ := strconv.Atoi(args[1])
			days, _ := strconv.Atoi(args[2])

			if startVerse > endVerse || days <= 0 {
				fmt.Println("Invalid input: Start verse must be less than or equal to end verse and days must be greater than zero.")
				return
			}

			calculateReadingPlan(startVerse, endVerse, days)
		},
	}
	return cmd
}

func calculateReadingPlan(startVerse, endVerse, days int) {
	ayahs := fetchAyahsFromAPI(startVerse, endVerse)

	totalAyahs := len(ayahs)

	dailyReadings := divideIntoDailyReadings(totalAyahs, days)

	fmt.Println("Daily Reading Plan:")
	for i, dailyReading := range dailyReadings {
		fmt.Printf("Day %d: Read ayahs %d to %d\n", i+1, dailyReading.start.StartSurah.Order, dailyReading.end.EndSurah.Order)
	}
}

func divideIntoDailyReadings(totalAyahs, days int) []DailyReading {
	dailyAyahs := totalAyahs / days

	var readings []DailyReading
	currentStart := 1
	for i := 0; i < days-1; i++ {
		reading := DailyReading{start: Thumn{Surah{"fatiha", 1}, currentStart, Surah{"fatiha", 1}, currentStart}, end: Thumn{Surah{"fatiha", 1}, currentStart + dailyAyahs - 1, Surah{"fatiha", 1}, currentStart + dailyAyahs - 1}}
		readings = append(readings, reading)
		currentStart += dailyAyahs
	}
	lastReading := DailyReading{start: Thumn{Surah{"fatiha", 1}, currentStart, Surah{"fatiha", 1}, currentStart}, end: Thumn{Surah{"fatiha", 1}, totalAyahs, Surah{"fatiha", 1}, totalAyahs}}
	readings = append(readings, lastReading)

	return readings
}

func fetchAyahsFromAPI(startVerse, endVerse int) []int {
	ayahs := make([]int, 0)
	for i := startVerse; i <= endVerse; i++ {
		ayahs = append(ayahs, i)
	}
	return ayahs
}
