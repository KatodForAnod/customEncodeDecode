package main

type SPI int

const (
	OFF SPI = 0
	ON  SPI = 1
)

type BL int

const (
	BlockOFF BL = 0
	BlockOn  BL = 1
)

type SB int

const (
	SubstitutionNo  SB = 0
	SubstitutionYes SB = 1
)

type BS int // warning

const (
	SectionReadyForLoad    BS = 0
	SectionNotReadyForLoad BS = 1
)

type NT int

const (
	ActualVal     NT = 0
	IrrelevantVal NT = 1
)

type IV int

const (
	ValidVal   IV = 0
	InvalidVal IV = 1
)

type EI int

const (
	TimeIntervalValValid   EI = 0
	TimeIntervalValInvalid EI = 1
)

type VTI int

const (
	EquipNotTransitionState VTI = 0
	EquipTransitionState    VTI = 1
)

type CY int

const (
	NoOverflowDurIntegrationPeriod CY = 0
	OverflowDurIntegrationPeriod   CY = 1
)

type CA int

const (
	CounterNotSetSinceLastRead CA = 0
	CounterSetSinceLastRead    CA = 1
)

type GS int

const (
	GeneralStartWorkNotHappen GS = 0
	GeneralStartWork          GS = 1
)

type SL1 int

const (
	StartPhaseANotHappen SL1 = 0
	StartPhaseA          SL1 = 1
)

type SL2 int

const (
	StartPhaseBNotHappen SL2 = 0
	StartPhaseB          SL2 = 1
)

type SL3 int

const (
	IEStartNotHappen SL3 = 0
	IEStartHappen    SL3 = 1
)

type SIE int

const (
	StartPhaseCNotHappen SIE = 0
	StartPhaseC          SIE = 1
)

type SRD int

const (
	BodyNormal_SRD  SRD = 0
	BodyReverse_SRD SRD = 1
)

type GC int

const (
	NoGenCommOutChain GC = 0
	GenCommOutChain   GC = 1
)

type Cl1 int

const (
	NoCommPhaseOutputChain Cl1 = 0
	CommPhaseOutputChain   Cl1 = 1
)

type Cl2 int

const (
	NoCommPhaseBOutChain Cl2 = 0
	CommPhaseBOutChain   Cl2 = 1
)

type Cl3 int

const (
	NoCommandPhaseCOutputChain Cl3 = 0
	CommandPhaseCOutputChain   Cl3 = 1
)

type DCS int

const (
	NotAllowedDCS  DCS = 0
	TurnOFF        DCS = 1
	TurnOn         DCS = 2
	NotAllowed_DCS DCS = 3
)

type RCS int

const (
	NotAllowedRCS  RCS = 0
	NextStepDown   RCS = 1
	NextStepUp     RCS = 2
	NotAllowed_RCS RCS = 3
)

type RES1 int

const (
	TrueTime    RES1 = 0
	ChangedTime RES1 = 1
)

type UI7 int

const (
	LocalPowerUp UI7 = iota
	localManualReset
	RemoteReset
)

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

type RQT int

const (
	NoCounterRequested RQT = iota
	GroupCounterQuery1
	GroupCounterQuery2
	GroupCounterQuery3
	GroupCounterQuery4
	GeneralCounterRequest
)

type FRZ int

const (
	Reading FRZ = iota
	FixCountWithoutReset
	LockCountWithReset
	ResetCount
)

type QPM int

const (
	NotUsing_QPM QPM = iota
	ThresholdVal
	SmoothingFactor
	LowLimMeasuredValTransfer
	UpLimMeasuredValTransfer
)

type LPC int

const (
	NoChange LPC = 0
	Change   LPC = 1
)

type POP int

const (
	InWork    POP = 0
	NotInWork POP = 1
)

type QPA int

const (
	NotUsing_QPA               QPA = 0
	OnOFFPreloadedParam        QPA = 1
	OnOFFParamAddrObj          QPA = 2
	OnOFFPermanentTransAddrObj QPA = 3
)

type QU int

const (
	NoAdditionalDefine     QU = 0
	ShortImpulseSysParamCP QU = 1
	LongImpulseSysParamCP  QU = 2
	ConstantOutput         QU = 3
)

type QRP int

const (
	NotUsing_QRP               QRP = 0
	GeneralSetProcessInitState QRP = 1
	DelTimestampEventBuffer    QRP = 2
)

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

type SCQ5to8 int

const (
	Default_SCQ5to8 SCQ5to8 = iota
	ReqMemAreaNotAvail_SCQ5to8
	ChecksumError_SCQ5to8
	UnCommService_SCQ5to8
	NoneExisFileName_SCQ5to8
	NonExisSectionName_SCQ5to8
)

type AFQ1to4 int

const (
	Default_AFQ1to4 AFQ1to4 = iota
	PosHandshakeFileTransfer
	NegHandshakeFileTransfer
	PosHandshakeSectTransfer
	NegHandshakeSectTransfer
)

type AFQ5to8 int

const (
	NotUsing_AFQ5to8 AFQ5to8 = iota
	ReqMemAreaNotAvail
	ChecksumError
	UnCommunicationServ
	NoneExistFileName
	NonExistSectionName
)

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

type FOR int

const (
	NameDefFile   FOR = 0
	NameDefSubDir FOR = 1
)

type FA int

const (
	FileWaitTransferred FA = 0
	TransferFileActive  FA = 1
)

type SOrE int

const (
	Execution_SOrE            SOrE = 0
	PreliminarySelection_SOrE SOrE = 1
)

const (
	C_IC_NA_1 byte = 100 // survey command
	C_CI_NA_1 byte = 101 // counter polling command
	C_CS_NA_1 byte = 103 // clock synchronization command
	C_RP_NA_1 byte = 105 // reset process command
)
