package pp

import (
	"math"
)

func calculateManiaPP(stars float64, noteCount float64, hit300g int16, hit300 int16, hit200 int16, hit100 int16, hit50 int16) float64 {
	accPP := (320*float64(hit300g)+300*float64(hit300)+200*float64(hit200)+100*float64(hit100)+50*float64(hit50))/(320*noteCount)
	return calculateManiaPPMax(stars,noteCount)*math.Max(accPP-0.8,0)*5
}
func calculateManiaPPMax(stars float64, noteCount float64) float64 {
	return math.Pow(math.Max(stars-0.15,0.05),2.2)*(1+0.1*math.Min(1,noteCount/1500))*8
}