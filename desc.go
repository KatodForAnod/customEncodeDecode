package main

func (s SPI) String() (result string) {
	switch s {
	case OFF:
		result = "SPI: OFF"
	case ON:
		result = "SPI: ON"
	default:
		result = "SPI: UNDEFINED"
	}
	return result
}

func (s BL) String() (result string) {
	switch s {
	case BlockOFF:
		result = "BL: OFF"
	case BlockOn:
		result = "BL: ON"
	default:
		result = "BL: UNDEFINED"
	}
	return result
}
func (s SB) String() (result string) {
	switch s {
	case SubstitutionNo:
		result = "SB: Substitution No"
	case SubstitutionYes:
		result = "SB: Substitution YES"
	default:
		result = "SB: UNDEFINED"
	}
	return result
}
func (s NT) String() (result string) {
	switch s {
	case ActualVal:
		result = "NT: Actual value"
	case IrrelevantVal:
		result = "NT: Irrelevant value"
	default:
		result = "NT: UNDEFINED"
	}
	return result
}
func (s IV) String() (result string) {
	switch s {
	case ValidVal:
		result = "IV: Valid value"
	case InvalidVal:
		result = "IV: Invalid value"
	default:
		result = "IV: UNDEFINED"
	}
	return result
}

func (s EI) String() (result string) {
	switch s {
	case TimeIntervalValValid:
		result = "EI: Time interval value valid"
	case TimeIntervalValInvalid:
		result = "EI: Time interval value invalid"
	default:
		result = "EI: UNDEFINED"
	}
	return result
}
func (s VTI) String() (result string) {
	switch s {
	case EquipNotTransitionState:
		result = "VTI: Equipment not in transition state"
	case EquipTransitionState:
		result = "VTI: Equipment in transition state"
	default:
		result = "VTI: UNDEFINED"
	}
	return result
}

func (s CY) String() (result string) {
	switch s {
	case NoOverflowDurIntegrationPeriod:
		result = "CY: No overflow during integration period"
	case OverflowDurIntegrationPeriod:
		result = "CY: Overflow during integration period"
	default:
		result = "CY: UNDEFINED"
	}
	return result
}

func (s CA) String() (result string) {
	switch s {
	case CounterNotSetSinceLastRead:
		result = "CA: Counter not set since last read"
	case CounterSetSinceLastRead:
		result = "CA: Counter set since last read"
	default:
		result = "CA: UNDEFINED"
	}
	return result
}
func (s GS) String() (result string) {
	switch s {
	case GeneralStartWorkNotHappen:
		result = "GS: General start work has not happened"
	case GeneralStartWork:
		result = "GS: General start work"
	default:
		result = "GS: UNDEFINED"
	}
	return result
}

func (s SL1) String() (result string) {
	switch s {
	case StartPhaseANotHappen:
		result = "SL1: Start phaseA has not happened"
	case StartPhaseA:
		result = "SL1: Start phaseA"
	default:
		result = "SL1: UNDEFINED"
	}
	return result
}

func (s SL2) String() (result string) {
	switch s {
	case StartPhaseBNotHappen:
		result = "SL2: Start phaseB has not happened"
	case StartPhaseB:
		result = "SL2: Start phaseB"
	default:
		result = "SL2: UNDEFINED"
	}
	return result
}
func (s SL3) String() (result string) {
	switch s {
	case IEStartNotHappen:
		result = "SL3: Start IE has not happened"
	case IEStartHappen:
		result = "SL3: Start IE"
	default:
		result = "SL3: UNDEFINED"
	}
	return result
}
func (s SIE) String() (result string) {
	switch s {
	case StartPhaseCNotHappen:
		result = "SIE: Start phaseC has not happened"
	case StartPhaseC:
		result = "SIE: Start phaseC"
	default:
		result = "SIE: UNDEFINED"
	}
	return result
}

func (s SRD) String() (result string) {
	switch s {
	case BodyNormal_SRD:
		result = "SRD: Start work device of reverse sequence has not happened"
	case BodyReverse_SRD:
		result = "SRD: Start work device of reverse sequence"
	default:
		result = "SRD: UNDEFINED"
	}
	return result
}
func (s GC) String() (result string) {
	switch s {
	case NoGenCommOutChain:
		result = "GC: No general command for out chain"
	case GenCommOutChain:
		result = "GC: General command for out chain"
	default:
		result = "GC: UNDEFINED"
	}
	return result
}
func (s Cl1) String() (result string) {
	switch s {
	case NoCommPhaseOutputChain:
		result = "Cl1: No general command for out chain of phaseA"
	case CommPhaseOutputChain:
		result = "Cl1: General command for out chain of phaseA"
	default:
		result = "Cl1: UNDEFINED"
	}
	return result
}
func (s Cl2) String() (result string) {
	switch s {
	case NoCommPhaseBOutChain:
		result = "Cl2: No general command for out chain of phaseB"
	case CommPhaseBOutChain:
		result = "Cl2: General command for out chain of phaseB"
	default:
		result = "Cl2: UNDEFINED"
	}
	return result
}

func (s Cl3) String() (result string) {
	switch s {
	case NoCommandPhaseCOutputChain:
		result = "Cl3: No general command for out chain of phaseC"
	case CommandPhaseCOutputChain:
		result = "Cl3: General command for out chain of phaseC"
	default:
		result = "Cl3: UNDEFINED"
	}
	return result
}

func (s DCS) String() (result string) {
	switch s {
	case NotAllowedDCS:
		result = "DCS: No allowed"
	case TurnOFF:
		result = "DCS: Turn OFF"
	case TurnOn:
		result = "DCS: Turn ON"
	case NotAllowed_DCS:
		result = "DCS: No allowed"
	default:
		result = "DCS: UNDEFINED"
	}
	return result
}

func (s RCS) String() (result string) {
	switch s {
	case NotAllowedRCS:
		result = "RCS: No allowed"
	case NextStepDown:
		result = "RCS: Next step down"
	case NextStepUp:
		result = "RCS: Next step up"
	case NotAllowed_RCS:
		result = "RCS: No allowed"
	default:
		result = "RCS: UNDEFINED"
	}
	return result
}
func (s RES1) String() (result string) {
	switch s {
	case TrueTime:
		result = "RES1: True time"
	case ChangedTime:
		result = "RES1: Changed time"
	default:
		result = "RES1: UNDEFINED"
	}
	return result
}
func (s UI7) String() (result string) {
	switch s {
	case LocalPowerUp:
		result = "UI7: Local power up"
	case LocalManualReset:
		result = "UI7: Local manual reset"
	case RemoteReset:
		result = "UI7: Remote reset"
	default:
		result = "UI7: UNDEFINED"
	}
	return result
}

func (s RES) String() (result string) {
	switch s {
	case NotUsing_RES:
		result = "RES: Local power up"
	case StationInterrogation:
		result = "RES: Local manual reset"
	case GroupSurvey1:
		result = "RES: Group survey1"
	case GroupSurvey2:
		result = "RES: Group survey2"
	case GroupSurvey3:
		result = "RES: Group survey3"
	case GroupSurvey4:
		result = "RES: Group survey4"
	case GroupSurvey5:
		result = "RES: Group survey5"
	case GroupSurvey6:
		result = "RES: Group survey6"
	case GroupSurvey7:
		result = "RES: Group survey7"
	case GroupSurvey8:
		result = "RES: Group survey8"
	case GroupSurvey9:
		result = "RES: Group survey9"
	case GroupSurvey10:
		result = "RES: Group survey10"
	case GroupSurvey11:
		result = "RES: Group survey11"
	case GroupSurvey12:
		result = "RES: Group GroupSurvey12"
	case GroupSurvey13:
		result = "RES: Group GroupSurvey13"
	case GroupSurvey14:
		result = "RES: Group GroupSurvey14"
	case GroupSurvey15:
		result = "RES: Group GroupSurvey15"
	case GroupSurvey16:
		result = "RES: Group GroupSurvey16"
	default:
		result = "RES: UNDEFINED"
	}
	return result
}

func (s RQT) String() (result string) {
	switch s {
	case NoCounterRequested:
		result = "RQT: No counter requested"
	case GroupCounterQuery1:
		result = "RQT: Request group counter 1"
	case GroupCounterQuery2:
		result = "RQT: Request group counter 2"
	case GroupCounterQuery3:
		result = "RQT: Request group counter 3"
	case GroupCounterQuery4:
		result = "RQT: Request group counter 4"
	case GeneralCounterRequest:
		result = "RQT: General counter request"
	default:
		result = "RQT: UNDEFINED"
	}
	return result
}

func (s FRZ) String() (result string) {
	switch s {
	case Reading:
		result = "FRZ: Reading"
	case FixCountWithoutReset:
		result = "FRZ: Fix count without reset"
	case LockCountWithReset:
		result = "FRZ: Lock count with reset"
	case ResetCount:
		result = "FRZ: Reset count"
	default:
		result = "FRZ: UNDEFINED"
	}
	return result
}

func (s QPM) String() (result string) {
	switch s {
	case NotUsing_QPM:
		result = "QPM: Not using"
	case ThresholdVal:
		result = "QPM: Threshold value"
	case SmoothingFactor:
		result = "QPM: Smoothing factor"
	case LowLimMeasuredValTransfer:
		result = "QPM: Low limit measured value transfer"
	case UpLimMeasuredValTransfer:
		result = "QPM: Up Limit measured value transfer"
	default:
		result = "QPM: UNDEFINED"
	}
	return result
}

func (s LPC) String() (result string) {
	switch s {
	case NoChange:
		result = "LPC: No changed"
	case Change:
		result = "LPC: Changed"
	default:
		result = "LPC: UNDEFINED"
	}
	return result
}

func (s POP) String() (result string) {
	switch s {
	case InWork:
		result = "POP: In work"
	case NotInWork:
		result = "POP: Not in work"
	default:
		result = "POP: UNDEFINED"
	}
	return result
}

func (s QPA) String() (result string) {
	switch s {
	case NotUsing_QPA:
		result = "QPA: Not using"
	case OnOFFPreloadedParam:
		result = "QPA: On/OFF preloaded parameters"
	case OnOFFParamAddrObj:
		result = "QPA: On/OFF parameters address object"
	case OnOFFPermanentTransAddrObj:
		result = "QPA: On/OFF permanent transition address object"
	default:
		result = "QPA: UNDEFINED"
	}
	return result
}
func (s QU) String() (result string) {
	switch s {
	case NoAdditionalDefine:
		result = "QU: No additional define"
	case ShortImpulseSysParamCP:
		result = "QU: Short impulse system parametersKP"
	case LongImpulseSysParamCP:
		result = "QU: Long impulse system parametersKP"
	case ConstantOutput:
		result = "QU: Constant output"
	default:
		result = "QU: UNDEFINED"
	}
	return result
}
func (s QRP) String() (result string) {
	switch s {
	case NotUsing_QRP:
		result = "QRP: Not using"
	case GeneralSetProcessInitState:
		result = "QRP: General set process init state"
	case DelTimestampEventBuffer:
		result = "QRP: Delete timestamp event buffer"
	default:
		result = "QRP: UNDEFINED"
	}
	return result
}

func (s SCQ1to4) String() (result string) {
	switch s {
	case Default_SCQ1to4:
		result = "SCQ1to4: Default"
	case ChoosingFile:
		result = "SCQ1to4: Choosing file"
	case RequestFile:
		result = "SCQ1to4: Request file"
	case DeactivatingFile:
		result = "SCQ1to4: Deactivating filet"
	case DeletingFile:
		result = "SCQ1to4: Deleting file"
	case ChoosingSection:
		result = "SCQ1to4: Choosing section"
	case RequestSection:
		result = "SCQ1to4: Request section"
	case DeactivatingSection:
		result = "SCQ1to4: Deactivating section"
	default:
		result = "SCQ1to4: UNDEFINED"
	}
	return result
}

func (s SCQ5to8) String() (result string) {
	switch s {
	case Default_SCQ5to8:
		result = "SCQ5to8: Default"
	case ReqMemAreaNotAvail_SCQ5to8:
		result = "SCQ5to8: Request memory area not available"
	case ChecksumError_SCQ5to8:
		result = "SCQ5to8: Checksum error"
	case UnCommService_SCQ5to8:
		result = "SCQ5to8: Unforeseen service"
	case NoneExisFileName_SCQ5to8:
		result = "SCQ5to8: None exist file name"
	case NonExisSectionName_SCQ5to8:
		result = "SCQ5to8: Non exist section name"
	default:
		result = "SCQ5to8: UNDEFINED"
	}
	return result
}
func (s AFQ1to4) String() (result string) {
	switch s {
	case Default_AFQ1to4:
		result = "AFQ1to4: Default"
	case PosHandshakeFileTransfer:
		result = "AFQ1to4: Positive handshake file transfer"
	case NegHandshakeFileTransfer:
		result = "AFQ1to4: Negative handshake file transfer"
	case PosHandshakeSectTransfer:
		result = "AFQ1to4: Positive handshake section transfer"
	case NegHandshakeSectTransfer:
		result = "AFQ1to4: Negative handshake section transfer"
	default:
		result = "AFQ1to4: UNDEFINED"
	}
	return result
}
func (s AFQ5to8) String() (result string) {
	switch s {
	case NotUsing_AFQ5to8:
		result = "AFQ5to8: Not using"
	case ReqMemAreaNotAvail:
		result = "AFQ5to8: Request memory area not available"
	case ChecksumError:
		result = "AFQ5to8: Checksum error"
	case UnCommunicationServ:
		result = "AFQ5to8: Uncommunicative service"
	case NoneExistFileName:
		result = "AFQ5to8: None exist file name"
	case NonExistSectionName:
		result = "AFQ5to8: None exist section name"
	default:
		result = "AFQ5to8: UNDEFINED"
	}
	return result
}
func (s LFD) String() (result string) {
	switch s {
	case NextFileSameDirectory:
		result = "LFD: Next file same directory"
	case LastFileOfDirectory:
		result = "LFD: Last file of directory"
	default:
		result = "LFD: UNDEFINED"
	}
	return result
}

func (s FOR) String() (result string) {
	switch s {
	case NameDefFile:
		result = "FOR: Name define file"
	case NameDefSubDir:
		result = "FOR: Name define sub directory"
	default:
		result = "FOR: UNDEFINED"
	}
	return result
}

func (s SOrE) String() (result string) {
	switch s {
	case Execution_SOrE:
		result = "SOrE: Execution"
	case PreliminarySelection_SOrE:
		result = "SOrE: Preliminary selection"
	default:
		result = "SOrE: UNDEFINED"
	}
	return result
}
