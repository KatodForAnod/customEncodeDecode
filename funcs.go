package main

import (
	"encoding/binary"
	"fmt"
)

// QOSEncode 7.2.6.39
func QOSEncode(ql int, e SOrE) [1]byte {
	qlByte := byte(ql) << 1
	eByte := byte(e)
	union := qlByte | eByte
	return [1]byte{union}
}

// QOSDecode 7.2.6.39
func QOSDecode(b [1]byte) (ql int, e SOrE) {
	ql = int(b[0] >> 1)
	e = SOrE(b[0] & 0x01)
	return ql, e
}

// SOFEncode 7.2.6.38
func SOFEncode(status int, lfd LFD, p FOR, fa FA) [1]byte {
	statusByte := byte(status) << 3
	lfdByte := byte(lfd) << 2
	forByte := byte(p) << 1
	faByte := byte(fa)

	union := statusByte | lfdByte | forByte | faByte
	return [1]byte{union}
}

// SOFDecode 7.2.6.38
func SOFDecode(b [1]byte) (int, LFD, FOR, FA) {
	status := b[0] >> 3
	lfd := (b[0] & 0x04) >> 2
	p := (b[0] & 0x02) >> 1
	fa := b[0] & 0x01
	return int(status), LFD(lfd), FOR(p), FA(fa)
}

// CHSEncode 7.2.6.37
func CHSEncode() {
	/// ???
}

// CHSDecode 7.2.6.37
func CHSDecode() {
	/// ???
}

// LOSEncode 7.2.6.36
func LOSEncode() {
	/// ???
}

// LOSDecode 7.2.6.36
func LOSDecode(b [1]byte) uint8 {
	return uint8(binary.BigEndian.Uint32(b[:]))
}

// LOFEncode 7.2.6.35
func LOFEncode() {
	/// ???
}

// LOFDecode 7.2.6.35
func LOFDecode(b [3]byte) uint32 {
	return binary.BigEndian.Uint32(b[:])
}

// NOFEncode 7.2.6.33
func NOFEncode(name uint16) ([2]byte, error) {
	data := make([]byte, 2)
	binary.BigEndian.PutUint16(data, name)
	return [2]byte{data[0], data[1]}, nil
}

// NOFDecode 7.2.6.33
func NOFDecode(b [2]byte) uint16 {
	out := binary.BigEndian.Uint16(b[:])
	return out
}

// AFQEncode 7.2.6.32
func AFQEncode(to4 AFQ1to4, to8 AFQ5to8) [1]byte {
	to4Byte := byte(to4) << 4
	to8Byte := byte(to8)
	union := to4Byte | to8Byte
	return [1]byte{union}
}

// AFQDecode 7.2.6.32
func AFQDecode(b [1]byte) (to4 AFQ1to4, to8 AFQ5to8) {
	to4 = AFQ1to4(b[0] >> 4)
	to8 = AFQ5to8(b[0] & 0xF)
	return to4, to8
}

// LSQEncode 7.2.6.31
func LSQEncode(to8 AFQ5to8) [1]byte {
	return [1]byte{byte(to8)}
}

// LSQDecode 7.2.6.31
func LSQDecode(b [1]byte) (to8 AFQ5to8, err error) {
	intConvert := int(b[0])

	switch intConvert {
	case int(NotUsing):
		return NotUsing, nil
	case int(FileTransferNoDeact_AFQ5to8):
		return FileTransferNoDeact_AFQ5to8, nil
	case int(FileTransferDeact_AFQ5to8):
		return FileTransferDeact_AFQ5to8, nil
	case int(SectionTransferNoDeact_AFQ5to8):
		return SectionTransferNoDeact_AFQ5to8, nil
	case int(SectionTransferDeact_AFQ5to8):
		return SectionTransferDeact_AFQ5to8, nil
	default:
		return 0, fmt.Errorf("LSQDecode err: unsupport type")
	}
}

// SCQEncode 7.2.6.30
func SCQEncode(to4 SCQ1to4, to8 SCQ5to8) [1]byte {
	to4Byte := byte(to4) << 4
	to8Byte := byte(to8)

	union := to4Byte | to8Byte
	return [1]byte{union}
}

// SCQDecode 7.2.6.30
func SCQDecode(b [1]byte) (to4 SCQ1to4, to8 SCQ5to8) {
	to4 = SCQ1to4(b[0] >> 4)
	to8 = SCQ5to8(b[0] & 0xF)

	return to4, to8
}

// SRQEncode 7.2.6.29
func SRQEncode(ui7 UI7, bs BS) [1]byte { //BS warning
	ui7Byte := byte(ui7) << 1
	bsByte := byte(bs)
	union := ui7Byte | bsByte
	return [1]byte{union}
}

// SRQDecode 7.2.6.29
func SRQDecode(b [1]byte) (ui7 UI7, bs BS) { //BS warning
	ui7 = UI7(b[0] >> 1)
	bs = BS(b[0] & 0x01)
	return ui7, bs
}

// FRQEncode 7.2.6.28
func FRQEncode(ui7 UI7, bs int) [1]byte {
	ui7Byte := byte(ui7) << 1
	bsByte := byte(bs)
	union := ui7Byte | bsByte
	return [1]byte{union}
}

// FRQDecode 7.2.6.28
func FRQDecode(b [1]byte) (ui7 UI7, bs int) {
	ui7 = UI7(b[0] >> 1)
	bs = int(b[0] & 0x01)
	return ui7, bs
}

// QRPEncode 7.2.6.27
func QRPEncode(q QRP) [1]byte {
	return [1]byte{byte(q)}
}

// QRPDecode 7.2.6.27
func QRPDecode(b [1]byte) (QRP, error) {
	intConvert := int(b[0])

	switch intConvert {
	case int(NotUsing_QRP):
		return NotUsing_QRP, nil
	case int(DelTimestampEventBuffer):
		return DelTimestampEventBuffer, nil
	case int(GeneralSetProcessInitState):
		return GeneralSetProcessInitState, nil
	default:
		return 0, fmt.Errorf("QRPDecode err: unsupport type")
	}
}

// QOCEncode 7.2.6.26
func QOCEncode(qu QU, sore SOrE) [1]byte {
	a := byte(qu) << 1 //qu max = 31;  00 01 11 11 => 00 11 11 10
	b := byte(sore)    //sore max = 1; 00 00 00 01    00 00 00 01

	c := a | b //                   result => 00 11 11 11 from 3 to 7 is QU and 8 is SOrE
	return [1]byte{c}
}

// QOCDecode 7.2.6.26
func QOCDecode(b [1]byte) (QU, SOrE) {
	tmp := b[0] & 0x3F // xx 11 11 11 and 00 11 11 11
	qu := tmp >> 1     // 00 11 11 1x >> 00 01 11 11
	sore := tmp & 0x01 // 00 xx xx x1 and 00 00 00 01
	return QU(qu), SOrE(sore)
}

// QPAEncode 7.2.6.25
func QPAEncode(qpa QPA) [1]byte {
	return [1]byte{byte(qpa)}
}

// QPADecode 7.2.6.25
func QPADecode(b [1]byte) (QPA, error) {
	intConvert := int(b[0])

	switch intConvert {
	case int(NotUsing_QPA):
		return NotUsing_QPA, nil
	case int(OnOFFPreloadedParam):
		return OnOFFPreloadedParam, nil
	case int(OnOFFParamAddrObj):
		return OnOFFParamAddrObj, nil
	case int(OnOFFPermanentTransAddrObj):
		return OnOFFPermanentTransAddrObj, nil
	default:
		return 0, fmt.Errorf("QPADecode err: unsupport type")
	}
}

// QPMEncode 7.2.6.24
func QPMEncode(kpa int, lpc LPC, pop POP) [1]byte {
	kpaByte := byte(kpa) << 2
	lpcByte := byte(lpc) << 1
	popByte := byte(pop)

	union := kpaByte | lpcByte | popByte
	return [1]byte{union}
}

// QPMDecode 7.2.6.24
func QPMDecode(b [1]byte) (kpa int, lpc LPC, pop POP) {
	kpa = int(b[0] >> 2)
	lpc = LPC((b[0] & 0x02) >> 1)
	pop = POP(b[0] & 0x01)
	return kpa, lpc, pop
}

// QCCEncode 7.2.6.23
func QCCEncode(rqt RQT, frz FRZ) [1]byte {
	rqtByte := byte(rqt) << 2
	frzByte := byte(frz)

	union := rqtByte | frzByte
	return [1]byte{union}
}

// QCCDecode 7.2.6.23
func QCCDecode(b [1]byte) (rqt RQT, frz FRZ) {
	rqt = RQT(b[0] >> 2)
	frz = FRZ(b[0] & 0x03)
	return rqt, frz
}

// QOIEncode 7.2.6.22
func QOIEncode(res RES) [1]byte {
	return [1]byte{byte(res)}
}

// QOIDecode 7.2.6.22
func QOIDecode(b [1]byte) (RES, error) {
	intConvert := int(b[0])

	switch intConvert {
	case int(NotUsing_RES):
		return NotUsing_RES, nil
	case int(StationInterrogation):
		return StationInterrogation, nil
	case int(GroupSurvey1):
		return GroupSurvey1, nil
	case int(GroupSurvey2):
		return GroupSurvey2, nil
	case int(GroupSurvey3):
		return GroupSurvey3, nil
	case int(GroupSurvey4):
		return GroupSurvey4, nil
	case int(GroupSurvey5):
		return GroupSurvey5, nil
	case int(GroupSurvey6):
		return GroupSurvey6, nil
	case int(GroupSurvey7):
		return GroupSurvey7, nil
	case int(GroupSurvey8):
		return GroupSurvey8, nil
	case int(GroupSurvey9):
		return GroupSurvey9, nil
	case int(GroupSurvey10):
		return GroupSurvey10, nil
	case int(GroupSurvey11):
		return GroupSurvey11, nil
	case int(GroupSurvey12):
		return GroupSurvey12, nil
	case int(GroupSurvey13):
		return GroupSurvey13, nil
	case int(GroupSurvey14):
		return GroupSurvey14, nil
	case int(GroupSurvey15):
		return GroupSurvey15, nil
	case int(GroupSurvey16):
		return GroupSurvey16, nil
	default:
		return 0, fmt.Errorf("QOIDecode err: unsupport type")
	}
}

// COIEncode 7.2.6.21
func COIEncode(ui7 UI7, bs int) [1]byte {
	ui7Byte := byte(ui7) << 1
	bsByte := byte(bs)
	union := ui7Byte | bsByte
	return [1]byte{union}
}

// COIDecode 7.2.6.21
func COIDecode(b [1]byte) (ui7 UI7, bs int) {
	ui7 = UI7(b[0] >> 1)
	bs = int(b[0] & 0x01)
	return ui7, bs
}
