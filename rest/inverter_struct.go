/*
Copyright 2025 Carl Meijer.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package rest

import (
	"encoding/json"

	"github.com/hammingweight/synkctl/types"
)

// Inverter is the complete SunSynk model of the inverter including fields that cannot be updated.
// For example, the batteryShutdownCap field is included in the response which can be useful but this
// REST API does not allow the field to be updated.
type Inverter struct {
	AbsorptionVolt               string `json:"absorptionVolt,omitempty"`
	AcCoupleFreqUpper            string `json:"acCoupleFreqUpper,omitempty"`
	AcCoupleOnGridSideEnable     string `json:"acCoupleOnGridSideEnable,omitempty"`
	AcCoupleOnLoadSideEnable     string `json:"acCoupleOnLoadSideEnable,omitempty"`
	AcCurrentUp                  string `json:"acCurrentUp,omitempty"`
	AcFreqLow                    string `json:"acFreqLow,omitempty"`
	AcFreqUp                     string `json:"acFreqUp,omitempty"`
	ActivePowerControl           string `json:"activePowerControl,omitempty"`
	AcType                       string `json:"acType,omitempty"`
	AcVoltLow                    string `json:"acVoltLow,omitempty"`
	AcVoltUp                     string `json:"acVoltUp,omitempty"`
	AllowRemoteControl           string `json:"allowRemoteControl,omitempty"`
	Ampm                         string `json:"ampm,omitempty"`
	ArcFactB                     string `json:"arcFactB,omitempty"`
	ArcFactC                     string `json:"arcFactC,omitempty"`
	ArcFactD                     string `json:"arcFactD,omitempty"`
	ArcFactFrz                   string `json:"arcFactFrz,omitempty"`
	ArcFactF                     string `json:"arcFactF,omitempty"`
	ArcFactI                     string `json:"arcFactI,omitempty"`
	ArcFactT                     string `json:"arcFactT,omitempty"`
	ArcFaultType                 string `json:"arcFaultType,omitempty"`
	AtsEnable                    string `json:"atsEnable,omitempty"`
	AtsSwitch                    string `json:"atsSwitch,omitempty"`
	AutoDim                      string `json:"autoDim,omitempty"`
	BackupDelay                  string `json:"backupDelay,omitempty"`
	BatErr                       string `json:"batErr,omitempty"`
	BatteryCap                   string `json:"batteryCap,omitempty"`
	BatteryChargeType            string `json:"batteryChargeType,omitempty"`
	BatteryEfficiency            string `json:"batteryEfficiency,omitempty"`
	BatteryEmptyVolt             string `json:"batteryEmptyVolt,omitempty"`
	BatteryEmptyV                string `json:"batteryEmptyV,omitempty"`
	BatteryImpedance             string `json:"batteryImpedance,omitempty"`
	BatteryLowCap                string `json:"batteryLowCap,omitempty"`
	BatteryLowVolt               string `json:"batteryLowVolt,omitempty"`
	BatteryMaxCurrentCharge      string `json:"batteryMaxCurrentCharge,omitempty"`
	BatteryMaxCurrentDischarge   string `json:"batteryMaxCurrentDischarge,omitempty"`
	BatteryOn                    string `json:"batteryOn,omitempty"`
	BatteryRestartCap            string `json:"batteryRestartCap,omitempty"`
	BatteryRestartVolt           string `json:"batteryRestartVolt,omitempty"`
	BatteryShutdownCap           string `json:"batteryShutdownCap,omitempty"`
	BatteryShutdownVolt          string `json:"batteryShutdownVolt,omitempty"`
	BatteryWorkStatus            string `json:"batteryWorkStatus,omitempty"`
	BattMode                     string `json:"battMode,omitempty"`
	BattType                     string `json:"battType,omitempty"`
	BatWarn                      string `json:"batWarn,omitempty"`
	Beep                         string `json:"beep,omitempty"`
	BmsErrStop                   string `json:"bmsErrStop,omitempty"`
	CaFStart                     string `json:"caFStart,omitempty"`
	CaFStop                      string `json:"caFStop,omitempty"`
	CaFwEnable                   string `json:"caFwEnable,omitempty"`
	CaliforniaFreqPressureEnable string `json:"californiaFreqPressureEnable,omitempty"`
	CaliforniaVoltPressureEnable string `json:"californiaVoltPressureEnable,omitempty"`
	CaLv3                        string `json:"caLv3,omitempty"`
	Cap1                         string `json:"cap1,omitempty"`
	Cap2                         string `json:"cap2,omitempty"`
	Cap3                         string `json:"cap3,omitempty"`
	Cap4                         string `json:"cap4,omitempty"`
	Cap5                         string `json:"cap5,omitempty"`
	Cap6                         string `json:"cap6,omitempty"`
	CaVoltPressureEnable         string `json:"caVoltPressureEnable,omitempty"`
	CaVStart                     string `json:"caVStart,omitempty"`
	CaVStop                      string `json:"caVStop,omitempty"`
	CaVwEnable                   string `json:"caVwEnable,omitempty"`
	ChargeCurrentLimit           string `json:"chargeCurrentLimit,omitempty"`
	ChargeCurrent                string `json:"chargeCurrent,omitempty"`
	ChargeLimit                  string `json:"chargeLimit,omitempty"`
	ChargeVolt                   string `json:"chargeVolt,omitempty"`
	CheckSelfTime                string `json:"checkSelfTime,omitempty"`
	CheckTime                    string `json:"checkTime,omitempty"`
	CommAddr                     string `json:"commAddr,omitempty"`
	CommBaudRate                 string `json:"commBaudRate,omitempty"`
	ComSet                       string `json:"comSet,omitempty"`
	Current10                    string `json:"current10,omitempty"`
	Current11                    string `json:"current11,omitempty"`
	Current12                    string `json:"current12,omitempty"`
	Current1                     string `json:"current1,omitempty"`
	Current2                     string `json:"current2,omitempty"`
	Current3                     string `json:"current3,omitempty"`
	Current4                     string `json:"current4,omitempty"`
	Current5                     string `json:"current5,omitempty"`
	Current6                     string `json:"current6,omitempty"`
	Current7                     string `json:"current7,omitempty"`
	Current8                     string `json:"current8,omitempty"`
	Current9                     string `json:"current9,omitempty"`
	DcVoltUp                     string `json:"dcVoltUp,omitempty"`
	DisableFloatCharge           string `json:"disableFloatCharge,omitempty"`
	DischargeCurrentLimit        string `json:"dischargeCurrentLimit,omitempty"`
	DischargeCurrent             string `json:"dischargeCurrent,omitempty"`
	DischargeLimit               string `json:"dischargeLimit,omitempty"`
	DischargeVolt                string `json:"dischargeVolt,omitempty"`
	DrmEnable                    string `json:"drmEnable,omitempty"`
	Eeprom                       string `json:"eeprom,omitempty"`
	EnergyMode                   string `json:"energyMode,omitempty"`
	EquChargeCycle               string `json:"equChargeCycle,omitempty"`
	EquChargeTime                string `json:"equChargeTime,omitempty"`
	EquipMode                    string `json:"equipMode,omitempty"`
	EquVoltCharge                string `json:"equVoltCharge,omitempty"`
	ExMeterCt                    string `json:"exMeterCt,omitempty"`
	ExMeterCtSwitch              string `json:"exMeterCtSwitch,omitempty"`
	ExternalCtRatio              string `json:"externalCtRatio,omitempty"`
	ExternalCurrent              string `json:"externalCurrent,omitempty"`
	FacHighProtect               string `json:"facHighProtect,omitempty"`
	FacLowProtect                string `json:"facLowProtect,omitempty"`
	Fac                          string `json:"fac,omitempty"`
	Flag2                        string `json:"flag2,omitempty"`
	FloatVolt                    string `json:"floatVolt,omitempty"`
	Four19                       string `json:"four19,omitempty"`
	FridayOn                     any    `json:"fridayOn,omitempty"`
	GenAndGridSignal             string `json:"genAndGridSignal,omitempty"`
	GenChargeOn                  string `json:"genChargeOn,omitempty"`
	GenConnectGrid               string `json:"genConnectGrid,omitempty"`
	GenCoolingTime               string `json:"genCoolingTime,omitempty"`
	GeneratorBatteryCurrent      string `json:"generatorBatteryCurrent,omitempty"`
	GeneratorForcedStart         string `json:"generatorForcedStart,omitempty"`
	GeneratorStartCap            string `json:"generatorStartCap,omitempty"`
	GeneratorStartVolt           string `json:"generatorStartVolt,omitempty"`
	GenMinSolar                  string `json:"genMinSolar,omitempty"`
	GenOffCap                    string `json:"genOffCap,omitempty"`
	GenOffVolt                   string `json:"genOffVolt,omitempty"`
	GenOnCap                     string `json:"genOnCap,omitempty"`
	GenOnVolt                    string `json:"genOnVolt,omitempty"`
	GenPeakPower                 string `json:"genPeakPower,omitempty"`
	GenPeakShaving               string `json:"genPeakShaving,omitempty"`
	GenSignal                    string `json:"genSignal,omitempty"`
	GenTime1on                   any    `json:"genTime1on,omitempty"`
	GenTime2on                   any    `json:"genTime2on,omitempty"`
	GenTime3on                   any    `json:"genTime3on,omitempty"`
	GenTime4on                   any    `json:"genTime4on,omitempty"`
	GenTime5on                   any    `json:"genTime5on,omitempty"`
	GenTime6on                   any    `json:"genTime6on,omitempty"`
	GenToLoadOn                  string `json:"genToLoadOn,omitempty"`
	GenToLoad                    string `json:"genToLoad,omitempty"`
	Gfdi                         string `json:"gfdi,omitempty"`
	GridAlwaysOn                 string `json:"gridAlwaysOn,omitempty"`
	GridMode                     string `json:"gridMode,omitempty"`
	GridPeakPower                string `json:"gridPeakPower,omitempty"`
	GridPeakShaving              string `json:"gridPeakShaving,omitempty"`
	GridSideINVMeter2            string `json:"gridSideINVMeter2,omitempty"`
	GridSignal                   string `json:"gridSignal,omitempty"`
	HighThrough                  string `json:"highThrough,omitempty"`
	ImpedanceLow                 string `json:"impedanceLow,omitempty"`
	ImportPower                  string `json:"importPower,omitempty"`
	InverterOutputVoltage        string `json:"inverterOutputVoltage,omitempty"`
	InverterType                 string `json:"inverterType,omitempty"`
	IsletProtect                 string `json:"isletProtect,omitempty"`
	Limit                        string `json:"limit,omitempty"`
	Limter                       string `json:"limter,omitempty"`
	LithiumMode                  string `json:"lithiumMode,omitempty"`
	LoadMode                     string `json:"loadMode,omitempty"`
	LockInVoltVar                string `json:"lockInVoltVar,omitempty"`
	LockInWattPF                 string `json:"lockInWattPF,omitempty"`
	LockOutChange                string `json:"lockOutChange,omitempty"`
	LockOutVoltVar               string `json:"lockOutVoltVar,omitempty"`
	LockOutWattPF                string `json:"lockOutWattPF,omitempty"`
	LowNoiseMode                 string `json:"lowNoiseMode,omitempty"`
	LowPowerMode                 string `json:"lowPowerMode,omitempty"`
	LowThrough                   string `json:"lowThrough,omitempty"`
	MaxOperatingTimeOfGen        string `json:"maxOperatingTimeOfGen,omitempty"`
	MeterA                       string `json:"meterA,omitempty"`
	MeterB                       string `json:"meterB,omitempty"`
	MeterC                       string `json:"meterC,omitempty"`
	MeterSelect                  string `json:"meterSelect,omitempty"`
	MicExportAll                 string `json:"micExportAll,omitempty"`
	MicExportGridOff             string `json:"micExportGridOff,omitempty"`
	ModbusSn                     string `json:"modbusSn,omitempty"`
	MondayOn                     any    `json:"mondayOn,omitempty"`
	MpptMulti                    string `json:"mpptMulti,omitempty"`
	MpptNum                      string `json:"mpptNum,omitempty"`
	MpptVoltLow                  string `json:"mpptVoltLow,omitempty"`
	NormalUpwardSlope            string `json:"normalUpwardSlope,omitempty"`
	OffGridImmediatelyOff        string `json:"offGridImmediatelyOff,omitempty"`
	Offset                       string `json:"offset,omitempty"`
	Open                         string `json:"open,omitempty"`
	OverFreq1Delay               string `json:"overFreq1Delay,omitempty"`
	OverFreq1                    string `json:"overFreq1,omitempty"`
	OverFreq2Delay               string `json:"overFreq2Delay,omitempty"`
	OverFreq2                    string `json:"overFreq2,omitempty"`
	OverLongVolt                 string `json:"overLongVolt,omitempty"`
	OverVolt1Delay               string `json:"overVolt1Delay,omitempty"`
	OverVolt1                    string `json:"overVolt1,omitempty"`
	OverVolt2Delay               string `json:"overVolt2Delay,omitempty"`
	OverVolt2                    string `json:"overVolt2,omitempty"`
	ParallelRegister1            string `json:"parallelRegister1,omitempty"`
	ParallelRegister2            string `json:"parallelRegister2,omitempty"`
	Parallel                     string `json:"parallel,omitempty"`
	PeakAndVallery               string `json:"peakAndVallery,omitempty"`
	Pf                           string `json:"pf,omitempty"`
	Phase                        string `json:"phase,omitempty"`
	PvLine                       string `json:"pvLine,omitempty"`
	PvMaxLimit                   string `json:"pvMaxLimit,omitempty"`
	Pwm                          string `json:"pwm,omitempty"`
	QvResponseTime               string `json:"qvResponseTime,omitempty"`
	Rcd                          string `json:"rcd,omitempty"`
	ReconnMaxFreq                string `json:"reconnMaxFreq,omitempty"`
	ReconnMaxVolt                string `json:"reconnMaxVolt,omitempty"`
	ReconnMinFreq                string `json:"reconnMinFreq,omitempty"`
	ReconnMinVolt                string `json:"reconnMinVolt,omitempty"`
	RecoveryTime                 string `json:"recoveryTime,omitempty"`
	RemoteLock                   string `json:"remoteLock,omitempty"`
	Reset                        string `json:"reset,omitempty"`
	Riso                         string `json:"riso,omitempty"`
	Rsd                          string `json:"rsd,omitempty"`
	SafetyType                   string `json:"safetyType,omitempty"`
	SaturdayOn                   any    `json:"saturdayOn,omitempty"`
	SdBatteryCurrent             string `json:"sdBatteryCurrent,omitempty"`
	SdChargeOn                   string `json:"sdChargeOn,omitempty"`
	SdStartCap                   string `json:"sdStartCap,omitempty"`
	SdStartVolt                  string `json:"sdStartVolt,omitempty"`
	SellTime1Pac                 string `json:"sellTime1Pac,omitempty"`
	SellTime1                    string `json:"sellTime1,omitempty"`
	SellTime1Volt                string `json:"sellTime1Volt,omitempty"`
	SellTime2Pac                 string `json:"sellTime2Pac,omitempty"`
	SellTime2                    string `json:"sellTime2,omitempty"`
	SellTime2Volt                string `json:"sellTime2Volt,omitempty"`
	SellTime3Pac                 string `json:"sellTime3Pac,omitempty"`
	SellTime3                    string `json:"sellTime3,omitempty"`
	SellTime3Volt                string `json:"sellTime3Volt,omitempty"`
	SellTime4Pac                 string `json:"sellTime4Pac,omitempty"`
	SellTime4                    string `json:"sellTime4,omitempty"`
	SellTime4Volt                string `json:"sellTime4Volt,omitempty"`
	SellTime5Pac                 string `json:"sellTime5Pac,omitempty"`
	SellTime5                    string `json:"sellTime5,omitempty"`
	SellTime5Volt                string `json:"sellTime5Volt,omitempty"`
	SellTime6Pac                 string `json:"sellTime6Pac,omitempty"`
	SellTime6                    string `json:"sellTime6,omitempty"`
	SellTime6Volt                string `json:"sellTime6Volt,omitempty"`
	SensorsCheck                 string `json:"sensorsCheck,omitempty"`
	SignalIslandModeEnable       string `json:"signalIslandModeEnable,omitempty"`
	SmartLoadOpenDelay           string `json:"smartLoadOpenDelay,omitempty"`
	Sn                           string `json:"sn,omitempty"`
	SoftStart                    string `json:"softStart,omitempty"`
	Solar1WindInputEnable        string `json:"solar1WindInputEnable,omitempty"`
	Solar2WindInputEnable        string `json:"solar2WindInputEnable,omitempty"`
	SolarMaxSellPower            string `json:"solarMaxSellPower,omitempty"`
	SolarPSU                     string `json:"solarPSU,omitempty"`
	SolarSell                    string `json:"solarSell,omitempty"`
	SpecialFunction              string `json:"specialFunction,omitempty"`
	Standard                     string `json:"standard,omitempty"`
	StartVoltLow                 string `json:"startVoltLow,omitempty"`
	StartVoltUp                  string `json:"startVoltUp,omitempty"`
	SundayOn                     any    `json:"sundayOn,omitempty"`
	SysWorkMode                  string `json:"sysWorkMode,omitempty"`
	Tempco                       string `json:"tempco,omitempty"`
	TestCommand                  string `json:"testCommand,omitempty"`
	Three41                      string `json:"three41,omitempty"`
	ThursdayOn                   any    `json:"thursdayOn,omitempty"`
	Time1on                      any    `json:"time1on,omitempty"`
	Time1On                      string `json:"time1On,omitempty"`
	Time2on                      any    `json:"time2on,omitempty"`
	Time2On                      string `json:"time2On,omitempty"`
	Time3on                      any    `json:"time3on,omitempty"`
	Time3On                      string `json:"time3On,omitempty"`
	Time4on                      any    `json:"time4on,omitempty"`
	Time4On                      string `json:"time4On,omitempty"`
	Time5on                      any    `json:"time5on,omitempty"`
	Time5On                      string `json:"time5On,omitempty"`
	Time6on                      any    `json:"time6on,omitempty"`
	Time6On                      string `json:"time6On,omitempty"`
	TimeSync                     string `json:"timeSync,omitempty"`
	TuesdayOn                    any    `json:"tuesdayOn,omitempty"`
	UnderFreq1Delay              string `json:"underFreq1Delay,omitempty"`
	UnderFreq1                   string `json:"underFreq1,omitempty"`
	UnderFreq2Delay              string `json:"underFreq2Delay,omitempty"`
	UnderFreq2                   string `json:"underFreq2,omitempty"`
	UnderVolt1Delay              string `json:"underVolt1Delay,omitempty"`
	UnderVolt1                   string `json:"underVolt1,omitempty"`
	UnderVolt2Delay              string `json:"underVolt2Delay,omitempty"`
	UnderVolt2                   string `json:"underVolt2,omitempty"`
	UpsStandard                  string `json:"upsStandard,omitempty"`
	VacHighProtect               string `json:"vacHighProtect,omitempty"`
	VacLowProtect                string `json:"vacLowProtect,omitempty"`
	VarQac1                      string `json:"varQac1,omitempty"`
	VarQac2                      string `json:"varQac2,omitempty"`
	VarQac3                      string `json:"varQac3,omitempty"`
	VarQac4                      string `json:"varQac4,omitempty"`
	VarVolt1                     string `json:"varVolt1,omitempty"`
	VarVolt2                     string `json:"varVolt2,omitempty"`
	VarVolt3                     string `json:"varVolt3,omitempty"`
	VarVolt4                     string `json:"varVolt4,omitempty"`
	VarVoltEnable                string `json:"varVoltEnable,omitempty"`
	VnResponseTime               string `json:"vnResponseTime,omitempty"`
	Volt10                       string `json:"volt10,omitempty"`
	Volt11                       string `json:"volt11,omitempty"`
	Volt12                       string `json:"volt12,omitempty"`
	Volt1                        string `json:"volt1,omitempty"`
	Volt2                        string `json:"volt2,omitempty"`
	Volt3                        string `json:"volt3,omitempty"`
	Volt4                        string `json:"volt4,omitempty"`
	Volt5                        string `json:"volt5,omitempty"`
	Volt6                        string `json:"volt6,omitempty"`
	Volt7                        string `json:"volt7,omitempty"`
	Volt8                        string `json:"volt8,omitempty"`
	Volt9                        string `json:"volt9,omitempty"`
	WattActivePf1                string `json:"wattActivePf1,omitempty"`
	WattActivePf2                string `json:"wattActivePf2,omitempty"`
	WattActivePf3                string `json:"wattActivePf3,omitempty"`
	WattActivePf4                string `json:"wattActivePf4,omitempty"`
	WattFreqEnable               string `json:"wattFreqEnable,omitempty"`
	WattOverExitFreqStopDelay    string `json:"wattOverExitFreqStopDelay,omitempty"`
	WattOverExitFreq             string `json:"wattOverExitFreq,omitempty"`
	WattOverFreq1StartDelay      string `json:"wattOverFreq1StartDelay,omitempty"`
	WattOverFreq1                string `json:"wattOverFreq1,omitempty"`
	WattOverWgralFreq            string `json:"wattOverWgralFreq,omitempty"`
	WattPf1                      string `json:"wattPf1,omitempty"`
	WattPf2                      string `json:"wattPf2,omitempty"`
	WattPf3                      string `json:"wattPf3,omitempty"`
	WattPf4                      string `json:"wattPf4,omitempty"`
	WattPfEnable                 string `json:"wattPfEnable,omitempty"`
	WattUnderExitFreqStopDelay   string `json:"wattUnderExitFreqStopDelay,omitempty"`
	WattUnderExitFreq            string `json:"wattUnderExitFreq,omitempty"`
	WattUnderFreq1StartDelay     string `json:"wattUnderFreq1StartDelay,omitempty"`
	WattUnderFreq1               string `json:"wattUnderFreq1,omitempty"`
	WattUnderWgalFreq            string `json:"wattUnderWgalFreq,omitempty"`
	WattV1                       string `json:"wattV1,omitempty"`
	WattV2                       string `json:"wattV2,omitempty"`
	WattV3                       string `json:"wattV3,omitempty"`
	WattV4                       string `json:"wattV4,omitempty"`
	WattVarActive1               string `json:"wattVarActive1,omitempty"`
	WattVarActive2               string `json:"wattVarActive2,omitempty"`
	WattVarActive3               string `json:"wattVarActive3,omitempty"`
	WattVarActive4               string `json:"wattVarActive4,omitempty"`
	WattVarEnable                string `json:"wattVarEnable,omitempty"`
	WattVarReactive1             string `json:"wattVarReactive1,omitempty"`
	WattVarReactive2             string `json:"wattVarReactive2,omitempty"`
	WattVarReactive3             string `json:"wattVarReactive3,omitempty"`
	WattVarReactive4             string `json:"wattVarReactive4,omitempty"`
	WattVoltEnable               string `json:"wattVoltEnable,omitempty"`
	WattW1                       string `json:"wattW1,omitempty"`
	WattW2                       string `json:"wattW2,omitempty"`
	WattW3                       string `json:"wattW3,omitempty"`
	WattW4                       string `json:"wattW4,omitempty"`
	WednesdayOn                  any    `json:"wednesdayOn,omitempty"`
	WorkState                    string `json:"workState,omitempty"`
	ZeroExportPower              string `json:"zeroExportPower,omitempty"`
}

// String() returns a pretty-printed JSON representation of an Inverter.
func (inverter *Inverter) String() string {
	m, err := json.MarshalIndent(inverter, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(m)
}

// InverterShortForm is a struct that includes only Inverter fields that can be updated.
type InverterShortForm struct {
	BattMode          string `json:"battMode,omitempty"`
	Cap1              string `json:"cap1,omitempty"`
	Cap2              string `json:"cap2,omitempty"`
	Cap3              string `json:"cap3,omitempty"`
	Cap4              string `json:"cap4,omitempty"`
	Cap5              string `json:"cap5,omitempty"`
	Cap6              string `json:"cap6,omitempty"`
	EnergyMode        string `json:"energyMode,omitempty"`
	FridayOn          any    `json:"fridayOn,omitempty"`
	GenTime1on        any    `json:"genTime1on,omitempty"`
	GenTime2on        any    `json:"genTime2on,omitempty"`
	GenTime3on        any    `json:"genTime3on,omitempty"`
	GenTime4on        any    `json:"genTime4on,omitempty"`
	GenTime5on        any    `json:"genTime5on,omitempty"`
	GenTime6on        any    `json:"genTime6on,omitempty"`
	MondayOn          any    `json:"mondayOn,omitempty"`
	PeakAndVallery    string `json:"peakAndVallery,omitempty"`
	PvMaxLimit        string `json:"pvMaxLimit,omitempty"`
	SafetyType        string `json:"safetyType,omitempty"`
	SaturdayOn        any    `json:"saturdayOn,omitempty"`
	SellTime1Pac      string `json:"sellTime1Pac,omitempty"`
	SellTime1         string `json:"sellTime1,omitempty"`
	SellTime1Volt     string `json:"sellTime1Volt,omitempty"`
	SellTime2Pac      string `json:"sellTime2Pac,omitempty"`
	SellTime2         string `json:"sellTime2,omitempty"`
	SellTime2Volt     string `json:"sellTime2Volt,omitempty"`
	SellTime3Pac      string `json:"sellTime3Pac,omitempty"`
	SellTime3         string `json:"sellTime3,omitempty"`
	SellTime3Volt     string `json:"sellTime3Volt,omitempty"`
	SellTime4Pac      string `json:"sellTime4Pac,omitempty"`
	SellTime4         string `json:"sellTime4,omitempty"`
	SellTime4Volt     string `json:"sellTime4Volt,omitempty"`
	SellTime5Pac      string `json:"sellTime5Pac,omitempty"`
	SellTime5         string `json:"sellTime5,omitempty"`
	SellTime5Volt     string `json:"sellTime5Volt,omitempty"`
	SellTime6Pac      string `json:"sellTime6Pac,omitempty"`
	SellTime6         string `json:"sellTime6,omitempty"`
	SellTime6Volt     string `json:"sellTime6Volt,omitempty"`
	SN                string `json:"sn,omitempty"`
	SolarMaxSellPower string `json:"solarMaxSellPower,omitempty"`
	SolarSell         string `json:"solarSell,omitempty"`
	SundayOn          any    `json:"sundayOn,omitempty"`
	SysWorkMode       string `json:"sysWorkMode,omitempty"`
	ThursdayOn        any    `json:"thursdayOn,omitempty"`
	Time1on           any    `json:"time1on,omitempty"`
	Time2on           any    `json:"time2on,omitempty"`
	Time3on           any    `json:"time3on,omitempty"`
	Time4on           any    `json:"time4on,omitempty"`
	Time5on           any    `json:"time5on,omitempty"`
	Time6on           any    `json:"time6on,omitempty"`
	TuesdayOn         any    `json:"tuesdayOn,omitempty"`
	WednesdayOn       any    `json:"wednesdayOn,omitempty"`
	ZeroExportPower   string `json:"zeroExportPower,omitempty"`
}

// InverterSettings contains the most important configurable settings for
// the inverter: battery-capacity, essential-only and grid-charge.
type InverterSettings struct {
	BatteryCapacity int         `json:"battery-capacity"`
	BatteryFirst    types.OnOff `json:"battery-first"`
	EssentialOnly   types.OnOff `json:"essential-only"`
	GridCharge      types.OnOff `json:"grid-charge"`
}

// String() returns a pretty-printed JSON representation of an InverterShortForm.
func (inverter *InverterShortForm) String() string {
	m, err := json.MarshalIndent(inverter, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(m)
}

// ToShortForm returns the short form of an Inverter.
func (inverter *Inverter) ToShortForm() (*InverterShortForm, error) {
	data, err := json.Marshal(inverter)
	if err != nil {
		return nil, err
	}
	is := &InverterShortForm{}
	err = json.Unmarshal(data, is)
	return is, err
}

// ToLongForm returns a struct with all Inverter fields. Fields that are not defined
// in the short form will be unpopulated.
func (inverter *InverterShortForm) ToLongForm() (*Inverter, error) {
	data, err := json.Marshal(inverter)
	if err != nil {
		return nil, err
	}
	il := &Inverter{}
	err = json.Unmarshal(data, il)
	return il, err
}

// ToSynkObject is an adapter that returns a SynkObject adaptation of an Inverter type.
func (inverter *Inverter) ToSynkObject() (*SynkObject, error) {
	data, err := json.Marshal(inverter)
	if err != nil {
		return nil, err
	}
	so := &SynkObject{}
	err = json.Unmarshal(data, so)
	return so, err
}
