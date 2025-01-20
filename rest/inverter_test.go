package rest

import "testing"

func TestReadGridCharge(t *testing.T) {
	inv := Inverter{}

	inv.Time1on = false
	inv.Time2on = true
	inv.Time3on = false
	inv.Time4on = false
	inv.Time5on = false
	inv.Time6on = false
	if !inv.GridChargeOn() {
		t.Error("grid charging should report as on")
	}

	inv.Time1on = "false"
	inv.Time2on = "false"
	inv.Time3on = "false"
	inv.Time4on = "false"
	inv.Time5on = "false"
	inv.Time6on = "false"
	if inv.GridChargeOn() {
		t.Error("grid charging should report as off")
	}

	inv.Time6on = "true"
	if !inv.GridChargeOn() {
		t.Error("grid charging should report as on")
	}
}

func TestSetGridCharge(t *testing.T) {
	inv := Inverter{}

	inv.SetGridChargeOn(false)
	if inv.Time1on.(bool) {
		t.Error("grid charging is on")
	}
	if inv.GridChargeOn() {
		t.Error("grid charging is on")
	}

	inv.SetGridChargeOn(true)
	if !inv.Time6on.(bool) {
		t.Error("grid charging is off")
	}
	if !inv.GridChargeOn() {
		t.Error("grid charging should report as off")
	}
}
