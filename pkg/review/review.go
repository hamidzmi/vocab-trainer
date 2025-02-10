package review

import "time"

type Word struct {
	Term       string
	Definition string
	Example    string
	Level      int
	NextReview time.Time
}

func UpdateReview(w *Word, remembered bool) {
	if w.Level >= 3 {
		return
	}

	if remembered {
		w.Level++
		w.NextReview = time.Now().Add(time.Duration(w.Level) * 24 * time.Hour).Truncate(time.Second)
	} else {
		w.Level = 0
		w.NextReview = time.Now().Add(24 * time.Hour).Truncate(time.Second)
	}
}

func GetWordsDueForReview(words []Word) []Word {
	var result []Word
	for _, w := range words {
		if w.NextReview.Before(time.Now()) && w.Level < 3 {
			result = append(result, w)
		}
	}

	return result
}
