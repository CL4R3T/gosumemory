package pp

import (
	"math"
)

func calculateManiaPP(stars float64, noteCount float64, hit300g int32, hit300 int32, hit200 int32, hit100 int32, hit50 int32, hitMiss int32) float64 {
	PPMax = calculateManiaPPMax(stars, noteCount)
	accPP = float64((320*hit300g+300*hit300+200*hit200+100*hit100+50*hit50))/(320*noteCount)
	return PPMax*math.Max((accPP-0.8),0)*5
}

func calculateManiaPPMax (stars float64, noteCount float64) float64 {
	return math.Pow(math.Max(stars-0.15,0.05),2.2)*(1+0.1*math.Min(1,noteCount/1500))*8
}