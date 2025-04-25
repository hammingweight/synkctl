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
	chargeOn, _ := inv.GridChargeOn()
	if !chargeOn {
		t.Error("grid charging should report as on")
	}

	inv.Time1on = "false"
	inv.Time2on = "false"
	inv.Time3on = "false"
	inv.Time4on = "false"
	inv.Time5on = "false"
	inv.Time6on = "false"
	chargeOn, _ = inv.GridChargeOn()
	if chargeOn {
		t.Error("grid charging should report as off")
	}

	inv.Time6on = "true"
	chargeOn, _ = inv.GridChargeOn()
	if !chargeOn {
		t.Error("grid charging should report as on")
	}
}

func TestSetGridCharge(t *testing.T) {
	inv := Inverter{}

	inv.SetGridChargeOn(false)
	if inv.Time1on.(bool) {
		t.Error("grid charging is on")
	}
	chargeOn, _ := inv.GridChargeOn()
	if chargeOn {
		t.Error("grid charging is on")
	}

	inv.SetGridChargeOn(true)
	if !inv.Time6on.(bool) {
		t.Error("grid charging is off")
	}
	chargeOn, _ = inv.GridChargeOn()
	if !chargeOn {
		t.Error("grid charging should report as off")
	}
}

func TestBatteryFirst(t *testing.T) {
	inv := Inverter{}
	inv.EnergyMode = "1"

	inv.SetBatteryFirst(true)
	if inv.EnergyMode != "0" {
		t.Error("energy mode should be 0")
	}
	if !inv.BatteryFirst() {
		t.Error("battery-first is off")
	}

	inv.SetBatteryFirst(false)
	if inv.EnergyMode != "1" {
		t.Error("energy mode should be 1")
	}
	if inv.BatteryFirst() {
		t.Error("battery-first is on")
	}
}
