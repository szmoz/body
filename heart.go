package main

type Heart struct {
	cycleLength    int
	diastoleLength int
	systoleLength  int
	strokeVolumen  int
	left           Vessel
	right          Vessel
}

func (h *Heart) createHeart(newCycleLength, newStrokeVolumen int, leftBlood, rightBlood Blood) {
	h.setCycleLength(newCycleLength)
	h.setStrokeVolumen(newStrokeVolumen)
	h.left = Vessel{}
	h.left.createVessel("Left ventricule", h.strokeVolumen, leftBlood)
	h.right = Vessel{}
	h.right.createVessel("Right ventricule", h.strokeVolumen, rightBlood)
}

func (h *Heart) setCycleLength(newCycleLength int) {
	h.cycleLength = newCycleLength
	h.diastoleLength = h.cycleLength / 8 * 5
	h.systoleLength = h.cycleLength - h.diastoleLength
}

func (h *Heart) setStrokeVolumen(newStrokeVolumen int) {
	h.strokeVolumen = newStrokeVolumen
}
