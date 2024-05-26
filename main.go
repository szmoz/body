package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

const bloodAmount int = 10

func main() {
	log.Println("program start")
	programStartTime := time.Now()
	// Create body
	// Create circulation
	totalBloodAmount := 4500
	arteriesBloodAmount :=
		totalBloodAmount / 100 * 12 / bloodAmount * bloodAmount
	capillariesBloodAmount :=
		totalBloodAmount / 100 * 7 / bloodAmount * bloodAmount
	veinsBloodAmount :=
		totalBloodAmount / 100 * 60 / bloodAmount * bloodAmount
	heartBloodAmount := 70 * 2
	lungBloodAmount := 500
	pulmVeinBloodAmount := (totalBloodAmount -
		arteriesBloodAmount -
		capillariesBloodAmount -
		veinsBloodAmount -
		heartBloodAmount -
		lungBloodAmount) / 2 / bloodAmount * bloodAmount
	pulmArteryBloodAmount := pulmVeinBloodAmount
	standardArteryBlood := Blood{PO2: 100, PCO2: 40}
	standardVeinBlood := Blood{PO2: 40, PCO2: 45}
	// Create heart
	heartRate := 72
	heart := Heart{}
	heart.createHeart(
		60*1000/heartRate,
		heartBloodAmount/2,
		standardArteryBlood,
		standardVeinBlood,
	)
	// Create lungs
	lungs := Lungs{}
	lungs.createLungs(
		lungBloodAmount,
		standardArteryBlood,
	)
	// Create vessels
	arteries := Vessel{}
	arteries.createVessel(
		"Arteries",
		arteriesBloodAmount,
		standardArteryBlood,
	)
	capillaries := Vessel{}
	capillaries.createVessel(
		"Capillaries",
		capillariesBloodAmount,
		standardVeinBlood,
	)
	veins := Vessel{}
	veins.createVessel(
		"Veins",
		veinsBloodAmount,
		standardVeinBlood,
	)
	pulmArtery := Vessel{}
	pulmArtery.createVessel(
		"Pulm Artery",
		pulmArteryBloodAmount,
		standardVeinBlood,
	)
	pulmVein := Vessel{}
	pulmVein.createVessel(
		"Pulm Vein",
		pulmVeinBloodAmount,
		standardArteryBlood,
	)
	// Connect vessels
	bloodStreamElements := []*Vessel{
		&arteries,
		&capillaries,
		&veins,
		&heart.right,
		&pulmArtery,
		&lungs.blood,
		&pulmVein,
		&heart.left,
	}
	bloodStreamElements[0].connectToPrev(
		bloodStreamElements[len(bloodStreamElements)-1].bloodChannel,
	)
	for i := 1; i < len(bloodStreamElements); i++ {
		bloodStreamElements[i].connectToPrev(
			bloodStreamElements[i-1].bloodChannel,
		)
	}
	for i := 0; i < len(bloodStreamElements); i++ {
		bloodStreamElements[i].bloodProcess =
			bloodStreamElements[i].standardBloodProcess
	}
	capillaries.bloodProcess = capillaries.capillariesBloodProcess
	lungs.blood.bloodProcess = lungs.blood.lungsBloodProcess

	// Concurrency setup
	var wg sync.WaitGroup

	// Run loop
	running := true
	for running {
		cycleStartTime := time.Now()
		cycleLength := time.Duration(heart.cycleLength * int(time.Millisecond))
		for _, element := range bloodStreamElements {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for strokeBloodAmount := 0; strokeBloodAmount <
					heart.strokeVolumen/bloodAmount; strokeBloodAmount++ {
					element.bloodProcess()
				}
			}()
		}
		wg.Wait()
		// print
		bloodStreamPrinter(bloodStreamElements)
		// Get remaining time
		actTime := time.Now()
		fmt.Println(actTime.Sub(cycleStartTime), cycleLength)
		cycleEnd := cycleStartTime.Add(cycleLength)
		remaining := time.Duration(cycleEnd.Sub(actTime))
		if remaining < 0 {
			log.Printf(
				"Cycle runtime '%v' longer than cycle time '%v'.\n",
				actTime.Sub(cycleStartTime),
				cycleLength,
			)
			continue
		}
		time.Sleep(remaining)

	}

	log.Printf("program end. Full runtime: %v", time.Since(programStartTime))
}

func bloodStreamPrinter(bloodStreamElements []*Vessel) {
	fmt.Print("\n\n")
	for _, element := range bloodStreamElements {
		fmt.Printf(
			"O2:%3d CO2:%3d blood amount:%4d   %s\n",
			element.lastProcessedBlood.PO2,
			element.lastProcessedBlood.PCO2,
			len(element.bloodChannel)*bloodAmount,
			element.name,
		)
	}
}
