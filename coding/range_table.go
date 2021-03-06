package coding

import (
	. "unicode"

	. "golang.org/x/text/unicode/rangetable"
)

//goland:noinspection GoSnakeCaseUsage
var (
	_ASCII     = &RangeTable{R16: []Range16{{0x00, 0x7F, 1}}}
	_Latin1    = Merge(_ASCII, Latin)
	_Cyrillic  = Merge(_ASCII, Cyrillic)
	_Hebrew    = Merge(_ASCII, Hebrew)
	_EUC_KR    = Merge(_ASCII, _EUC_KR_Definition)
	_Shift_JIS = Merge(_ASCII, _Shift_JIS_Definition)
)

//goland:noinspection GoSnakeCaseUsage
var (
	_Shift_JIS_Definition = &RangeTable{R16: []Range16{
		{0x00A1, 0x0460, 1}, {0x2010, 0x2670, 1}, {0x3000, 0x33CE, 1},
		{0x4E00, 0x9FA6, 1}, {0xF929, 0xFA2E, 1}, {0xFF01, 0xFFE6, 1},
	}}
	_EUC_KR_Definition = &RangeTable{R16: []Range16{
		{0x00A1, 0x0452, 1}, {0x2015, 0x266E, 1}, {0x3000, 0x33DE, 1},
		{0x4E00, 0x9F9D, 1}, {0xAC00, 0xD7A4, 1}, {0xF900, 0xFA0C, 1},
		{0xFF01, 0xFFE7, 1},
	}}
)
