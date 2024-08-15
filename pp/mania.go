package pp

import (
	"math"
)

func calculateManiaPP(stars float64, noteCount float64, hit300g int16, hit300 int16, hit200 int16, hit100 int16, hit50 int16) float64 {
	accPP := float(320*hit300g+300*hit300+200*hit200+100*hit100+50*hit50)/(320*noteCount)
	return calculateManiaPPMax(stars,noteCount)*(accPP-0.8)*5
}
func calculateManiaPPMax(stars float64, noteCount float64) float64 {
	return math.Pow(math.Max(stars-0.15,0.05),2.2)*(1+0.1*math.Min(1,noteCount/1500))*8
}