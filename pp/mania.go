package pp

import (
	"math"
	"fmt"
)

func calculateManiaPP(stars float64, hit300g int16, hit300 int16, hit200 int16, hit100 int16, hit50 int16, hitMiss int16) float64 {
	accPP := (320*float64(hit300g)+300*float64(hit300)+200*float64(hit200)+100*float64(hit100)+50*float64(hit50))/(320*float(hit300g+hit300+hit200+hit100+hit50+hitMiss))
	PPMax := calculateManiaPPMax(stars)
	PP := PPMax*math.Max(accPP-0.8,0)*5
	return PP
}
func calculateManiaPPMax(stars float64) float64 {
	return math.Pow(math.Max(stars-0.15,0.05),2.2)*8.8
}