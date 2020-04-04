package main

// https://github.com/DragonMinded/bemaniutils/blob/master/bemani/common/card.py

import (
	"encoding/hex"
	"strings"
	"syscall/js"
)

var KEY = []int64{
	0x20d0d03c,
	0x868ecb41,
	0xbcd89c84,
	0x4c0e0d0d,
	0x84fc30ac,
	0x4cc1890e,
	0xfc5418a4,
	0x02c50f44,
	0x68acb4e0,
	0x06cd4a4e,
	0xcc28906c,
	0x4f0c8ac0,
	0xb03ca468,
	0x884ac7c4,
	0x389490d8,
	0xcf80c6c2,
	0x58d87404,
	0xc48ec444,
	0xb4e83c50,
	0x498d0147,
	0x64f454c0,
	0x4c4701c8,
	0xec302cc4,
	0xc6c949c1,
	0xc84c00f0,
	0xcdcc49cc,
	0x883c5cf4,
	0x8b0fcb80,
	0x703cc0b0,
	0xcb820a8d,
	0x78804c8c,
	0x4fca830e,
	0x80d0f03c,
	0x8ec84f8c,
	0x98c89c4c,
	0xc80d878f,
	0x54bc949c,
	0xc801c5ce,
	0x749078dc,
	0xc3c80d46,
	0x2c8070f0,
	0x0cce4dcf,
	0x8c3874e4,
	0x8d448ac3,
	0x987cac70,
	0xc0c20ac5,
	0x288cfc78,
	0xc28543c8,
	0x4c8c7434,
	0xc50e4f8d,
	0x8468f4b4,
	0xcb4a0307,
	0x2854dc98,
	0x48430b45,
	0x6858fce8,
	0x4681cd49,
	0xd04808ec,
	0x458d0fcb,
	0xe0a48ce4,
	0x880f8fce,
	0x7434b8fc,
	0xce080a8e,
	0x5860fc6c,
	0x46c886cc,
	0xd01098a4,
	0xce090b8c,
	0x1044cc2c,
	0x86898e0f,
	0xd0809c3c,
	0x4a05860f,
	0x54b4f80c,
	0x4008870e,
	0x1480b88c,
	0x0ac8854f,
	0x1c9034cc,
	0x08444c4e,
	0x0cb83c64,
	0x41c08cc6,
	0x1c083460,
	0xc0c603ce,
	0x2ca0645c,
	0x818246cb,
	0x0408e454,
	0xc5464487,
	0x88607c18,
	0xc1424187,
	0x284c7c90,
	0xc1030509,
	0x40486c94,
	0x4603494b,
	0xe0404ce4,
	0x4109094d,
	0x60443ce4,
	0x4c0b8b8d,
	0xe054e8bc,
	0x02008e89,
}
var LUT_A0 = []int64{
	0x02080008,
	0x02082000,
	0x00002008,
	0x00000000,
	0x02002000,
	0x00080008,
	0x02080000,
	0x02082008,
	0x00000008,
	0x02000000,
	0x00082000,
	0x00002008,
	0x00082008,
	0x02002008,
	0x02000008,
	0x02080000,
	0x00002000,
	0x00082008,
	0x00080008,
	0x02002000,
	0x02082008,
	0x02000008,
	0x00000000,
	0x00082000,
	0x02000000,
	0x00080000,
	0x02002008,
	0x02080008,
	0x00080000,
	0x00002000,
	0x02082000,
	0x00000008,
	0x00080000,
	0x00002000,
	0x02000008,
	0x02082008,
	0x00002008,
	0x02000000,
	0x00000000,
	0x00082000,
	0x02080008,
	0x02002008,
	0x02002000,
	0x00080008,
	0x02082000,
	0x00000008,
	0x00080008,
	0x02002000,
	0x02082008,
	0x00080000,
	0x02080000,
	0x02000008,
	0x00082000,
	0x00002008,
	0x02002008,
	0x02080000,
	0x00000008,
	0x02082000,
	0x00082008,
	0x00000000,
	0x02000000,
	0x02080008,
	0x00002000,
	0x00082008,
}
var LUT_A1 = []int64{
	0x08000004,
	0x00020004,
	0x00000000,
	0x08020200,
	0x00020004,
	0x00000200,
	0x08000204,
	0x00020000,
	0x00000204,
	0x08020204,
	0x00020200,
	0x08000000,
	0x08000200,
	0x08000004,
	0x08020000,
	0x00020204,
	0x00020000,
	0x08000204,
	0x08020004,
	0x00000000,
	0x00000200,
	0x00000004,
	0x08020200,
	0x08020004,
	0x08020204,
	0x08020000,
	0x08000000,
	0x00000204,
	0x00000004,
	0x00020200,
	0x00020204,
	0x08000200,
	0x00000204,
	0x08000000,
	0x08000200,
	0x00020204,
	0x08020200,
	0x00020004,
	0x00000000,
	0x08000200,
	0x08000000,
	0x00000200,
	0x08020004,
	0x00020000,
	0x00020004,
	0x08020204,
	0x00020200,
	0x00000004,
	0x08020204,
	0x00020200,
	0x00020000,
	0x08000204,
	0x08000004,
	0x08020000,
	0x00020204,
	0x00000000,
	0x00000200,
	0x08000004,
	0x08000204,
	0x08020200,
	0x08020000,
	0x00000204,
	0x00000004,
	0x08020004,
}
var LUT_A2 = []int64{
	0x80040100,
	0x01000100,
	0x80000000,
	0x81040100,
	0x00000000,
	0x01040000,
	0x81000100,
	0x80040000,
	0x01040100,
	0x81000000,
	0x01000000,
	0x80000100,
	0x81000000,
	0x80040100,
	0x00040000,
	0x01000000,
	0x81040000,
	0x00040100,
	0x00000100,
	0x80000000,
	0x00040100,
	0x81000100,
	0x01040000,
	0x00000100,
	0x80000100,
	0x00000000,
	0x80040000,
	0x01040100,
	0x01000100,
	0x81040000,
	0x81040100,
	0x00040000,
	0x81040000,
	0x80000100,
	0x00040000,
	0x81000000,
	0x00040100,
	0x01000100,
	0x80000000,
	0x01040000,
	0x81000100,
	0x00000000,
	0x00000100,
	0x80040000,
	0x00000000,
	0x81040000,
	0x01040100,
	0x00000100,
	0x01000000,
	0x81040100,
	0x80040100,
	0x00040000,
	0x81040100,
	0x80000000,
	0x01000100,
	0x80040100,
	0x80040000,
	0x00040100,
	0x01040000,
	0x81000100,
	0x80000100,
	0x01000000,
	0x81000000,
	0x01040100,
}
var LUT_A3 = []int64{
	0x04010801,
	0x00000000,
	0x00010800,
	0x04010000,
	0x04000001,
	0x00000801,
	0x04000800,
	0x00010800,
	0x00000800,
	0x04010001,
	0x00000001,
	0x04000800,
	0x00010001,
	0x04010800,
	0x04010000,
	0x00000001,
	0x00010000,
	0x04000801,
	0x04010001,
	0x00000800,
	0x00010801,
	0x04000000,
	0x00000000,
	0x00010001,
	0x04000801,
	0x00010801,
	0x04010800,
	0x04000001,
	0x04000000,
	0x00010000,
	0x00000801,
	0x04010801,
	0x00010001,
	0x04010800,
	0x04000800,
	0x00010801,
	0x04010801,
	0x00010001,
	0x04000001,
	0x00000000,
	0x04000000,
	0x00000801,
	0x00010000,
	0x04010001,
	0x00000800,
	0x04000000,
	0x00010801,
	0x04000801,
	0x04010800,
	0x00000800,
	0x00000000,
	0x04000001,
	0x00000001,
	0x04010801,
	0x00010800,
	0x04010000,
	0x04010001,
	0x00010000,
	0x00000801,
	0x04000800,
	0x04000801,
	0x00000001,
	0x04010000,
	0x00010800,
}
var LUT_B0 = []int64{
	0x00000400,
	0x00000020,
	0x00100020,
	0x40100000,
	0x40100420,
	0x40000400,
	0x00000420,
	0x00000000,
	0x00100000,
	0x40100020,
	0x40000020,
	0x00100400,
	0x40000000,
	0x00100420,
	0x00100400,
	0x40000020,
	0x40100020,
	0x00000400,
	0x40000400,
	0x40100420,
	0x00000000,
	0x00100020,
	0x40100000,
	0x00000420,
	0x40100400,
	0x40000420,
	0x00100420,
	0x40000000,
	0x40000420,
	0x40100400,
	0x00000020,
	0x00100000,
	0x40000420,
	0x00100400,
	0x40100400,
	0x40000020,
	0x00000400,
	0x00000020,
	0x00100000,
	0x40100400,
	0x40100020,
	0x40000420,
	0x00000420,
	0x00000000,
	0x00000020,
	0x40100000,
	0x40000000,
	0x00100020,
	0x00000000,
	0x40100020,
	0x00100020,
	0x00000420,
	0x40000020,
	0x00000400,
	0x40100420,
	0x00100000,
	0x00100420,
	0x40000000,
	0x40000400,
	0x40100420,
	0x40100000,
	0x00100420,
	0x00100400,
	0x40000400,
}
var LUT_B1 = []int64{
	0x00800000,
	0x00001000,
	0x00000040,
	0x00801042,
	0x00801002,
	0x00800040,
	0x00001042,
	0x00801000,
	0x00001000,
	0x00000002,
	0x00800002,
	0x00001040,
	0x00800042,
	0x00801002,
	0x00801040,
	0x00000000,
	0x00001040,
	0x00800000,
	0x00001002,
	0x00000042,
	0x00800040,
	0x00001042,
	0x00000000,
	0x00800002,
	0x00000002,
	0x00800042,
	0x00801042,
	0x00001002,
	0x00801000,
	0x00000040,
	0x00000042,
	0x00801040,
	0x00801040,
	0x00800042,
	0x00001002,
	0x00801000,
	0x00001000,
	0x00000002,
	0x00800002,
	0x00800040,
	0x00800000,
	0x00001040,
	0x00801042,
	0x00000000,
	0x00001042,
	0x00800000,
	0x00000040,
	0x00001002,
	0x00800042,
	0x00000040,
	0x00000000,
	0x00801042,
	0x00801002,
	0x00801040,
	0x00000042,
	0x00001000,
	0x00001040,
	0x00801002,
	0x00800040,
	0x00000042,
	0x00000002,
	0x00001042,
	0x00801000,
	0x00800002,
}
var LUT_B2 = []int64{
	0x10400000,
	0x00404010,
	0x00000010,
	0x10400010,
	0x10004000,
	0x00400000,
	0x10400010,
	0x00004010,
	0x00400010,
	0x00004000,
	0x00404000,
	0x10000000,
	0x10404010,
	0x10000010,
	0x10000000,
	0x10404000,
	0x00000000,
	0x10004000,
	0x00404010,
	0x00000010,
	0x10000010,
	0x10404010,
	0x00004000,
	0x10400000,
	0x10404000,
	0x00400010,
	0x10004010,
	0x00404000,
	0x00004010,
	0x00000000,
	0x00400000,
	0x10004010,
	0x00404010,
	0x00000010,
	0x10000000,
	0x00004000,
	0x10000010,
	0x10004000,
	0x00404000,
	0x10400010,
	0x00000000,
	0x00404010,
	0x00004010,
	0x10404000,
	0x10004000,
	0x00400000,
	0x10404010,
	0x10000000,
	0x10004010,
	0x10400000,
	0x00400000,
	0x10404010,
	0x00004000,
	0x00400010,
	0x10400010,
	0x00004010,
	0x00400010,
	0x00000000,
	0x10404000,
	0x10000010,
	0x10400000,
	0x10004010,
	0x00000010,
	0x00404000,
}
var LUT_B3 = []int64{
	0x00208080,
	0x00008000,
	0x20200000,
	0x20208080,
	0x00200000,
	0x20008080,
	0x20008000,
	0x20200000,
	0x20008080,
	0x00208080,
	0x00208000,
	0x20000080,
	0x20200080,
	0x00200000,
	0x00000000,
	0x20008000,
	0x00008000,
	0x20000000,
	0x00200080,
	0x00008080,
	0x20208080,
	0x00208000,
	0x20000080,
	0x00200080,
	0x20000000,
	0x00000080,
	0x00008080,
	0x20208000,
	0x00000080,
	0x20200080,
	0x20208000,
	0x00000000,
	0x00000000,
	0x20208080,
	0x00200080,
	0x20008000,
	0x00208080,
	0x00008000,
	0x20000080,
	0x00200080,
	0x20208000,
	0x00000080,
	0x00008080,
	0x20200000,
	0x20008080,
	0x20000000,
	0x20200000,
	0x00208000,
	0x20208080,
	0x00008080,
	0x00208000,
	0x20200080,
	0x00200000,
	0x20000080,
	0x20008000,
	0x00000000,
	0x00008000,
	0x00200000,
	0x20200080,
	0x00208080,
	0x20000000,
	0x20208000,
	0x00000080,
	0x20008080,
}

var VALID_CHARS = "0123456789ABCDEFGHJKLMNPRSTUWXYZ"
var CONV_CHARS = map[string]string{
	"I": "1",
	"O": "0",
}

func __typeFromCardId(cardid string) int64 {
	upper := strings.ToUpper(cardid[:2])
	if upper == "E0" {
		return 1
	}
	if upper == "01" {
		return 2
	}
	return 0
}
func __checksum(data []int64) int64 {
	checksum := int64(0)

	for i := int64(0); i < 15; i++ {
		checksum += ((i % 3) + 1) * data[i]
	}

	for checksum >= 0x20 {
		checksum = (checksum & 0x1f) + (checksum >> 5)
	}

	return checksum
}
func __ror(val, amount int64) int64 {
	return ((val << (32 - amount)) & 0xffffffff) | ((val >> amount) & 0xffffffff)
}
func __operatorA(off, state int64) int64 {
	v3 := (state >> 32) & 0xffffffff
	v4 := state & 0xffffffff

	for i := int64(0); i < 32; i += 4 {
		v20 := __ror(v3^KEY[off+i+1], 28)

		v4 ^=
			LUT_B0[(v20>>26)&0x3f] ^
				LUT_B1[(v20>>18)&0x3f] ^
				LUT_B2[(v20>>10)&0x3f] ^
				LUT_B3[(v20>>2)&0x3f] ^
				LUT_A0[((v3^KEY[off+i])>>26)&0x3f] ^
				LUT_A1[((v3^KEY[off+i])>>18)&0x3f] ^
				LUT_A2[((v3^KEY[off+i])>>10)&0x3f] ^
				LUT_A3[((v3^KEY[off+i])>>2)&0x3f]

		v21 := __ror(v4^KEY[off+i+3], 28)

		v3 ^=
			LUT_B0[(v21>>26)&0x3f] ^
				LUT_B1[(v21>>18)&0x3f] ^
				LUT_B2[(v21>>10)&0x3f] ^
				LUT_B3[(v21>>2)&0x3f] ^
				LUT_A0[((v4^KEY[off+i+2])>>26)&0x3f] ^
				LUT_A1[((v4^KEY[off+i+2])>>18)&0x3f] ^
				LUT_A2[((v4^KEY[off+i+2])>>10)&0x3f] ^
				LUT_A3[((v4^KEY[off+i+2])>>2)&0x3f]
	}

	return ((v3 & 0xffffffff) << 32) | (v4 & 0xffffffff)
}
func __operatorB(off, state int64) int64 {
	v3 := (state >> 32) & 0xffffffff
	v4 := state & 0xffffffff

	for i := int64(0); i < 32; i += 4 {
		v20 := __ror(v3^KEY[off+31-i], 28)

		v4 ^=
			LUT_A0[((v3^KEY[off+30-i])>>26)&0x3f] ^
				LUT_A1[((v3^KEY[off+30-i])>>18)&0x3f] ^
				LUT_A2[((v3^KEY[off+30-i])>>10)&0x3f] ^
				LUT_A3[((v3^KEY[off+30-i])>>2)&0x3f] ^
				LUT_B0[(v20>>26)&0x3f] ^
				LUT_B1[(v20>>18)&0x3f] ^
				LUT_B2[(v20>>10)&0x3f] ^
				LUT_B3[(v20>>2)&0x3f]

		v21 := __ror(v4^KEY[off+29-i], 28)

		v3 ^=
			LUT_A0[((v4^KEY[off+28-i])>>26)&0x3f] ^
				LUT_A1[((v4^KEY[off+28-i])>>18)&0x3f] ^
				LUT_A2[((v4^KEY[off+28-i])>>10)&0x3f] ^
				LUT_A3[((v4^KEY[off+28-i])>>2)&0x3f] ^
				LUT_B0[(v21>>26)&0x3f] ^
				LUT_B1[(v21>>18)&0x3f] ^
				LUT_B2[(v21>>10)&0x3f] ^
				LUT_B3[(v21>>2)&0x3f]
	}

	return ((v3 & 0xffffffff) << 32) | (v4 & 0xffffffff)
}
func __toInt64(data []int64) int64 {
	inX :=
		(data[0] & 0xff) |
			((data[1] & 0xff) << 8) |
			((data[2] & 0xff) << 16) |
			((data[3] & 0xff) << 24)

	inY :=
		(data[4] & 0xff) |
			((data[5] & 0xff) << 8) |
			((data[6] & 0xff) << 16) |
			((data[7] & 0xff) << 24)

	v7 := ((((inX ^ (inY >> 4)) & 0xf0f0f0f) << 4) ^ inY) & 0xffffffff
	v8 := (((inX ^ (inY >> 4)) & 0xf0f0f0f) ^ inX) & 0xffffffff
	v9 := (v7 ^ (v8 >> 16)) & 0x0000ffff
	v10 := (((v7 ^ (v8 >> 16)) << 16) ^ v8) & 0xffffffff
	v11 := (v9 ^ v7) & 0xffffffff
	v12 := (v10 ^ (v11 >> 2)) & 0x33333333
	v13 := (v11 ^ (v12 << 2)) & 0xffffffff
	v14 := (v12 ^ v10) & 0xffffffff
	v15 := (v13 ^ (v14 >> 8)) & 0x00ff00ff
	v16 := (v14 ^ (v15 << 8)) & 0xffffffff
	v17 := __ror(v15^v13, 1)
	v18 := (v16 ^ v17) & 0x55555555
	v3 := __ror(v18^v16, 1)
	v4 := (v18 ^ v17) & 0xffffffff

	return ((v3 & 0xffffffff) << (32)) | (v4 & 0xffffffff)
}
func __fromInt64(data []int64, state int64) []int64 {
	v3 := (state >> 32) & 0xffffffff
	v4 := state & 0xffffffff
	v22 := __ror(v4, 31)
	v23 := (v3 ^ v22) & 0x55555555
	v24 := (v23 ^ v22) & 0xffffffff
	v25 := __ror(v23^v3, 31)
	v26 := (v25 ^ (v24 >> 8)) & 0x00ff00ff
	v27 := (v24 ^ (v26 << 8)) & 0xffffffff
	v28 := (v26 ^ v25) & 0xffffffff
	v29 := ((v28 >> 2) ^ v27) & 0x33333333
	v30 := ((v29 << 2) ^ v28) & 0xffffffff
	v31 := (v29 ^ v27) & 0xffffffff
	v32 := (v30 ^ (v31 >> 16)) & 0x0000ffff
	v33 := (v31 ^ (v32 << 16)) & 0xffffffff
	v34 := (v32 ^ v30) & 0xffffffff
	v35 := (v33 ^ (v34 >> 4)) & 0xf0f0f0f
	outY := ((v35 << 4) ^ v34) & 0xffffffff
	outX := (v35 ^ v33) & 0xffffffff

	data[0] = outX & 0xff
	data[1] = (outX >> 8) & 0xff
	data[2] = (outX >> 16) & 0xff
	data[3] = (outX >> 24) & 0xff
	data[4] = outY & 0xff
	data[5] = (outY >> 8) & 0xff
	data[6] = (outY >> 16) & 0xff
	data[7] = (outY >> 24) & 0xff

	return data
}

func __int64ArrayToByteArray(inp []int64) []byte {
	out := make([]byte, len(inp))
	for i := 0; i < len(out); i++ {
		out[i] = byte(inp[i])
	}
	return out
}
func __byteArrayToInt64Array(inp []byte) []int64 {
	out := make([]int64, len(inp))
	for i := 0; i < len(out); i++ {
		out[i] = int64(inp[i])
	}
	return out
}
func __int64ArrayReverse(inp []int64) []int64 {
	for i, j := 0, len(inp)-1; i < j; i, j = i+1, j-1 {
		inp[i], inp[j] = inp[j], inp[i]
	}
	return inp
}

func __encode(intInt64s []int64) []int64 {
	if len(intInt64s) != 8 {
		return make([]int64, 8)
	}

	inp := intInt64s
	out := make([]int64, 8)

	out = __fromInt64(out, __operatorA(0x00, __toInt64(inp)))
	out = __fromInt64(out, __operatorB(0x20, __toInt64(out)))
	out = __fromInt64(out, __operatorA(0x40, __toInt64(out)))

	return out
}
func __decode(inInt64s []int64) []int64 {
	if len(inInt64s) != 8 {
		return make([]int64, 8)
	}

	inp := inInt64s
	out := make([]int64, 8)

	out = __fromInt64(out, __operatorB(0x40, __toInt64(inp)))
	out = __fromInt64(out, __operatorA(0x20, __toInt64(out)))
	out = __fromInt64(out, __operatorB(0x00, __toInt64(out)))

	return out
}

func encode(cardId string) string {
	if len(cardId) != 16 {
		return ""
	}

	cardBytes, err := hex.DecodeString(cardId)
	cardint64s := __byteArrayToInt64Array(cardBytes)

	if err != nil {
		return ""
	}

	reverse := __int64ArrayReverse(cardint64s)
	ciphered := __encode(reverse)

	bits := make([]int64, 65)
	for i := 0; i < 64; i++ {
		bits[i] = (ciphered[i>>3] >> (^i & 7)) & 1
	}

	groups := make([]int64, 16)
	for i := 0; i < 13; i++ {
		groups[i] =
			(bits[i*5] << 4) |
				(bits[i*5+1] << 3) |
				(bits[i*5+2] << 2) |
				(bits[i*5+3] << 1) |
				(bits[i*5+4] << 0)
	}

	groups[13] = 1
	groups[0] ^= __typeFromCardId(cardId)

	for i := 0; i < 14; i++ {
		index := i - 1
		if index < 0 {
			index = len(groups) + (i % len(groups)) - 1
		}
		groups[i] ^= groups[index]
	}

	groups[14] = __typeFromCardId(cardId)
	groups[15] = __checksum(groups)

	final := ""
	for _, i := range groups {
		final += string(VALID_CHARS[i])
	}

	return final
}
func decode(cardId string) string {
	cardId = strings.ReplaceAll(cardId, " ", "")
	cardId = strings.ReplaceAll(cardId, "-", "")
	cardId = strings.ToUpper(cardId)

	for o, n := range CONV_CHARS {
		cardId = strings.ReplaceAll(cardId, o, n)
	}

	if len(cardId) != 16 {
		return ""
	}

	for _, c := range cardId {
		if !strings.Contains(cardId, string(c)) {
			return ""
		}
	}

	groups := make([]int64, 16)
	for i := 0; i < 16; i++ {
		for j := int64(0); j < 32; j++ {
			if cardId[i] == VALID_CHARS[j] {
				groups[i] = j
				break
			}
		}
	}

	if groups[14] != 1 && groups[14] != 2 {
		return ""
	}

	if groups[15] != __checksum(groups) {
		return ""
	}

	for i := 13; i > 0; i-- {
		groups[i] ^= groups[i-1]
	}
	groups[0] ^= groups[14]

	bits := make([]int64, 64)
	for i := 0; i < 64; i++ {
		bits[i] = (groups[i/5] >> (4 - (i % 5))) & 1
	}

	ciphered := make([]int64, 8)
	for i := 0; i < 64; i++ {
		ciphered[i/8] |= bits[i] << (^i & 7)
	}

	deciphered := __decode(ciphered)
	reverse := __int64ArrayReverse(deciphered)
	final := strings.ToUpper(hex.EncodeToString(__int64ArrayToByteArray(reverse)))
	return final
}

func jsEncode(_ js.Value, inputs []js.Value) interface{} {
	callback := inputs[len(inputs)-1:][0]
	inp := inputs[0].String()
	return callback.Invoke(encode(inp))
}
func jsDecode(_ js.Value, inputs []js.Value) interface{} {
	callback := inputs[len(inputs)-1:][0]
	inp := inputs[0].String()
	return callback.Invoke(decode(inp))
}

func main() {
	c := make(chan bool)
	js.Global().Set("KONAMI_CARD_ENCODE", js.FuncOf(jsEncode))
	js.Global().Set("KONAMI_CARD_DECODE", js.FuncOf(jsDecode))
	<-c
}
