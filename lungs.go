package main

type Lungs struct {
	blood Vessel
}

func (l *Lungs) createLungs(fullBloodVolume int, newBlood Blood) {
	l.blood = Vessel{}
	l.blood.createVessel("Lung blood", fullBloodVolume, newBlood)
}
