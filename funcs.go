package main

import "fmt"

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

func QPMEncode() {

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
