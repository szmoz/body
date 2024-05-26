package main

import "fmt"

type Vessel struct {
	name               string
	capacity           int
	bloodChannel       chan *Blood
	prevChannel        chan *Blood
	bloodProcess       func()
	lastProcessedBlood *Blood
}

func (v *Vessel) createVessel(name string, fullVolume int, newBlood Blood) {
	v.name = name
	v.capacity = fullVolume / bloodAmount
	v.bloodChannel = make(chan *Blood, v.capacity+(140/bloodAmount))
	v.fillWithBlood(newBlood)
	v.bloodProcess = v.standardBloodProcess
}

func (v *Vessel) fillWithBlood(newBlood Blood) {
	for i := 0; i < v.capacity; i++ {
		v.bloodChannel <- &Blood{PO2: newBlood.PO2, PCO2: newBlood.PCO2}
	}
}

func (v *Vessel) connectToPrev(prevChannel chan *Blood) {
	v.prevChannel = prevChannel
}

func (v *Vessel) standardBloodProcess() {
	v.lastProcessedBlood = <-v.prevChannel
	v.bloodChannel <- v.lastProcessedBlood
}

func (v *Vessel) capillariesBloodProcess() {
	v.lastProcessedBlood = <-v.prevChannel
	v.lastProcessedBlood.PO2 = 40
	v.lastProcessedBlood.PCO2 = 45
	v.bloodChannel <- v.lastProcessedBlood
}

func (v *Vessel) lungsBloodProcess() {
	v.lastProcessedBlood = <-v.prevChannel
	v.lastProcessedBlood.PO2 = 100
	v.lastProcessedBlood.PCO2 = 40
	fmt.Println(v.lastProcessedBlood)
	v.bloodChannel <- v.lastProcessedBlood
}
