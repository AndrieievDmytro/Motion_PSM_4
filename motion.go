package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
)

func convertToRadians(angle float64) float64 {
	return angle * math.Pi / 180
}

var (
	mass   = 1.0
	radius = 2.0
	gVal   = 10.0
	alfa   = convertToRadians(45) // 45 degree in radians
	deltaT = 0.05
	hight  = 20.0 //hight
)

func csvExport(data [][]string, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		if err := writer.Write(value); err != nil {
			return err
		}
	}
	return nil
}

func convertToStringArray(posX float64, posY float64) []string {
	sx := fmt.Sprintf("%f", posX)
	sy := fmt.Sprintf("%f", posY)
	conv := []string{sx, sy}
	return conv
}

// When object gets to the bottom
func stopCondition() float64 {
	posX := hight / math.Tan(alfa)
	d := math.Sqrt(hight*hight + posX*posX)
	return d
}

func midPointBall() {
	data := [][]string{}
	data_linear := [][]string{}
	energy := [][]string{}
	posX := 0.0
	posY := radius
	velocity := 0.0
	inertiaBall := 2.0 / 5.0 * mass * math.Pow(radius, 2)
	accelerationBall := gVal * math.Sin(alfa) / (1.0 + inertiaBall/(mass*radius*radius))
	epsilon := accelerationBall / radius
	beta := 0.0
	omega := 0.0
	var deltaPosX float64
	var deltaVelocity float64
	var massCenterX float64
	var massCenterY float64
	var deltaOmega float64
	var deltaBeta float64

	for posX < stopCondition() {

		posX += deltaPosX
		velocity += deltaVelocity

		circleX := radius*math.Sin(beta) + massCenterX
		circleY := radius*math.Cos(beta) + massCenterY

		velocity_2 := velocity + accelerationBall*deltaT/2

		deltaPosX = velocity_2 * deltaT
		deltaVelocity = accelerationBall * deltaT

		beta += deltaBeta
		omega += deltaOmega

		omega_2 := epsilon * deltaT / 2
		deltaBeta = omega_2 * deltaT
		deltaOmega = epsilon * deltaT

		massCenterX = posX*math.Cos(-alfa) - posY*math.Sin(-alfa)
		massCenterY = posX*(math.Sin(-alfa)) + posY*(math.Cos(-alfa)) + hight

		eKinetic := mass * gVal * massCenterY
		ePotential := mass*math.Pow(velocity, 2)/2 + inertiaBall*math.Pow(omega, 2)/2

		data = append(data, convertToStringArray(circleX, circleY))
		data_linear = append(data_linear, convertToStringArray(massCenterX, massCenterY))
		energy = append(energy, convertToStringArray(eKinetic, ePotential))

	}
	csvExport(data, "angleMidPointBall.csv")
	csvExport(data_linear, "linearMidPointBall.csv")
	csvExport(energy, "energyBall.csv")
}

func midPointSphere() {
	data := [][]string{}
	data_linear := [][]string{}
	energy := [][]string{}
	posX := 0.0
	posY := radius
	velocity := 0.0
	inertiaSphere := 2.0 / 3.0 * mass * math.Pow(radius, 2)
	accelerationSphere := (gVal * math.Sin(alfa)) / (1 + inertiaSphere/(mass*math.Pow(radius, 2)))
	epsilon := accelerationSphere / radius
	beta := 0.0
	omega := 0.0
	var deltaPosX float64
	var deltaVelocity float64
	var massCenterX float64
	var massCenterY float64
	var deltaOmega float64
	var deltaBeta float64

	for posX < stopCondition() {

		posX += deltaPosX
		velocity += deltaVelocity

		circleX := radius*math.Sin(beta) + massCenterX
		circleY := radius*math.Cos(beta) + massCenterY

		velocity_2 := velocity + accelerationSphere*deltaT/2

		deltaPosX = velocity_2 * deltaT
		deltaVelocity = accelerationSphere * deltaT

		beta += deltaBeta
		omega += deltaOmega

		omega_2 := epsilon * deltaT / 2
		deltaBeta = omega_2 * deltaT
		deltaOmega = epsilon * deltaT

		massCenterX = posX*math.Cos(-alfa) - posY*math.Sin(-alfa)
		massCenterY = posX*(math.Sin(-alfa)) + posY*(math.Cos(-alfa)) + hight

		eKinetic := mass * gVal * massCenterY
		ePotential := mass*math.Pow(velocity, 2)/2 + inertiaSphere*math.Pow(omega, 2)/2

		data = append(data, convertToStringArray(circleX, circleY))
		data_linear = append(data_linear, convertToStringArray(massCenterX, massCenterY))
		energy = append(energy, convertToStringArray(eKinetic, ePotential))

	}
	csvExport(data, "angleMidPointSphere.csv")
	csvExport(data_linear, "linearMidPointSphere.csv")
	csvExport(energy, "energySphere.csv")
}

func main() {
	midPointBall()
	midPointSphere()
}
