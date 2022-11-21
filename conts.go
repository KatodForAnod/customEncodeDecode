package main

type SPI int

const (
	OFF SPI = 0
	ON  SPI = 1
)

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

type BL int

const (
	BlockOFF BL = 0
	BlockOn  BL = 1
)

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

type SB int

const (
	SubstitutionNo  SB = 0
	SubstitutionYes SB = 1
)

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

type BS int // warning

const (
	SectionReadyForLoad    BS = 0
	SectionNotReadyForLoad BS = 1
)

func (s BS) String() (result string) {
	switch s {
	case SectionReadyForLoad:
		result = "BS: Section ready for load"
	case SectionNotReadyForLoad:
		result = "BS: Section not ready for load"
	default:
		result = "BS: UNDEFINED"
	}
	return result
}

type NT int

const (
	ActualVal     NT = 0
	IrrelevantVal NT = 1
)

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

type IV int

const (
	ValidVal   IV = 0
	InvalidVal IV = 1
)

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

type EI int

const (
	TimeIntervalValValid   EI = 0
	TimeIntervalValInvalid EI = 1
)

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

type VTI int

const (
	EquipNotTransitionState VTI = 0
	EquipTransitionState    VTI = 1
)

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

type CY int

const (
	NoOverflowDurIntegrationPeriod CY = 0
	OverflowDurIntegrationPeriod   CY = 1
)

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

type CA int

const (
	CounterNotSetSinceLastRead CA = 0
	CounterSetSinceLastRead    CA = 1
)

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

type GS int

const (
	GeneralStartWorkNotHappen GS = 0
	GeneralStartWork          GS = 1
)

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

type SL1 int

const (
	StartPhaseANotHappen SL1 = 0
	StartPhaseA          SL1 = 1
)

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

type SL2 int

const (
	StartPhaseBNotHappen SL2 = 0
	StartPhaseB          SL2 = 1
)

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

type SL3 int

const (
	IEStartNotHappen SL3 = 0
	IEStartHappen    SL3 = 1
)

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

type SIE int

const (
	StartPhaseCNotHappen SIE = 0
	StartPhaseC          SIE = 1
)

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

type SRD int

const (
	BodyNormal_SRD  SRD = 0
	BodyReverse_SRD SRD = 1
)

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

type GC int

const (
	NoGenCommOutChain GC = 0
	GenCommOutChain   GC = 1
)

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

type Cl1 int

const (
	NoCommPhaseOutputChain Cl1 = 0
	CommPhaseOutputChain   Cl1 = 1
)

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

type Cl2 int

const (
	NoCommPhaseBOutChain Cl2 = 0
	CommPhaseBOutChain   Cl2 = 1
)

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

type Cl3 int

const (
	NoCommandPhaseCOutputChain Cl3 = 0
	CommandPhaseCOutputChain   Cl3 = 1
)

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

type DCS int

const (
	NotAllowedDCS  DCS = 0
	TurnOFF        DCS = 1
	TurnOn         DCS = 2
	NotAllowed_DCS DCS = 3
)

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

type RCS int

const (
	NotAllowedRCS  RCS = 0
	NextStepDown   RCS = 1
	NextStepUp     RCS = 2
	NotAllowed_RCS RCS = 3
)

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

type RES1 int

const (
	TrueTime    RES1 = 0
	ChangedTime RES1 = 1
)

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

type UI7 int

const (
	LocalPowerUp UI7 = iota
	LocalManualReset
	RemoteReset
)

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

type RES int

const (
	NotUsing_RES         RES = iota
	StationInterrogation RES = iota + 19
	GroupSurvey1
	GroupSurvey2
	GroupSurvey3
	GroupSurvey4
	GroupSurvey5
	GroupSurvey6
	GroupSurvey7
	GroupSurvey8
	GroupSurvey9
	GroupSurvey10
	GroupSurvey11
	GroupSurvey12
	GroupSurvey13
	GroupSurvey14
	GroupSurvey15
	GroupSurvey16
)

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

type RQT int

const (
	NoCounterRequested RQT = iota
	GroupCounterQuery1
	GroupCounterQuery2
	GroupCounterQuery3
	GroupCounterQuery4
	GeneralCounterRequest
)

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

type FRZ int

const (
	Reading FRZ = iota
	FixCountWithoutReset
	LockCountWithReset
	ResetCount
)

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

type QPM int

const (
	NotUsing_QPM QPM = iota
	ThresholdVal
	SmoothingFactor
	LowLimMeasuredValTransfer
	UpLimMeasuredValTransfer
)

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

type LPC int

const (
	NoChange LPC = 0
	Change   LPC = 1
)

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

type POP int

const (
	InWork    POP = 0
	NotInWork POP = 1
)

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

type QPA int

const (
	NotUsing_QPA               QPA = 0
	OnOFFPreloadedParam        QPA = 1
	OnOFFParamAddrObj          QPA = 2
	OnOFFPermanentTransAddrObj QPA = 3
)

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

type QU int

const (
	NoAdditionalDefine     QU = 0
	ShortImpulseSysParamCP QU = 1
	LongImpulseSysParamCP  QU = 2
	ConstantOutput         QU = 3
)

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

type QRP int

const (
	NotUsing_QRP               QRP = 0
	GeneralSetProcessInitState QRP = 1
	DelTimestampEventBuffer    QRP = 2
)

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

type SCQ1to4 int

const (
	Default_SCQ1to4 SCQ1to4 = iota
	ChoosingFile
	RequestFile
	DeactivatingFile
	DeletingFile
	ChoosingSection
	RequestSection
	DeactivatingSection
)

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

type SCQ5to8 int

const (
	Default_SCQ5to8 SCQ5to8 = iota
	ReqMemAreaNotAvail_SCQ5to8
	ChecksumError_SCQ5to8
	UnCommService_SCQ5to8
	NoneExisFileName_SCQ5to8
	NonExisSectionName_SCQ5to8
)

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

type AFQ1to4 int

const (
	Default_AFQ1to4 AFQ1to4 = iota
	PosHandshakeFileTransfer
	NegHandshakeFileTransfer
	PosHandshakeSectTransfer
	NegHandshakeSectTransfer
)

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

type AFQ5to8 int

const (
	NotUsing_AFQ5to8 AFQ5to8 = iota
	ReqMemAreaNotAvail
	ChecksumError
	UnCommunicationServ
	NoneExistFileName
	NonExistSectionName
)

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

// ???
const (
	NotUsing AFQ5to8 = iota
	FileTransferNoDeact_AFQ5to8
	FileTransferDeact_AFQ5to8
	SectionTransferNoDeact_AFQ5to8
	SectionTransferDeact_AFQ5to8
)

type LFD int

const (
	NextFileSameDirectory LFD = 0
	LastFileOfDirectory   LFD = 1
)

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

type FOR int

const (
	NameDefFile   FOR = 0
	NameDefSubDir FOR = 1
)

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

type FA int

const (
	FileWaitTransferred FA = 0
	TransferFileActive  FA = 1
)

func (s FA) String() (result string) {
	switch s {
	case FileWaitTransferred:
		result = "FA: File waiting for transfer"
	case TransferFileActive:
		result = "FA: Transferring file is active"
	default:
		result = "FA: UNDEFINED"
	}
	return result
}

type SOrE int

const (
	Execution_SOrE            SOrE = 0
	PreliminarySelection_SOrE SOrE = 1
)

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

const (
	C_IC_NA_1 byte = 100 // survey command
	C_CI_NA_1 byte = 101 // counter polling command
	C_CS_NA_1 byte = 103 // clock synchronization command
	C_RP_NA_1 byte = 105 // reset process command
)
