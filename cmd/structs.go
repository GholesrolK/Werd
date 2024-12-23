package cmd

type DailyReading struct {
	start Thumn
	end   Thumn
}

type Thumn struct {
	StartSurah Surah
	StartAyah  int
	EndSurah   Surah
	EndAyah    int
}

type Surah struct {
	Name  string
	Order int16
}
