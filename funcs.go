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
		return 0, fmt.Errorf("unsupport type")
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