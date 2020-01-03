// Code generated by ragel. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

//line conf.rl:1
// Copyright 2018 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package hba

import (
	"net"
	"strings"
	"unicode/utf8"

	"github.com/pkg/errors"
)

func Parse(input string) (*Conf, error) {
	if !utf8.ValidString(input) {
		return nil, errors.New("invalid UTF-8")
	}
	// To ease parsing, ensure a newline at EOF.
	data := []rune(input + "\n")

//line conf.rl:29

//line conf.rl:30

//line conf.go:36
	var _scanner_actions []byte = []byte{
		0, 1, 0, 1, 3, 1, 10, 1, 14,
		1, 15, 1, 16, 2, 1, 4, 2,
		1, 5, 2, 2, 4, 2, 2, 5,
		2, 2, 9, 2, 2, 13, 2, 6,
		8, 2, 7, 8, 2, 16, 17, 3,
		1, 4, 11, 3, 1, 4, 12, 3,
		1, 5, 11, 3, 1, 5, 12, 3,
		2, 4, 11, 3, 2, 4, 12, 3,
		2, 5, 11, 3, 2, 5, 12, 3,
		2, 13, 14, 4, 7, 8, 2, 13,
	}

	var _scanner_key_offsets []int16 = []int16{
		0, 0, 1, 2, 3, 4, 6, 13,
		19, 26, 32, 46, 53, 60, 67, 80,
		81, 91, 100, 113, 122, 126, 130, 132,
		136, 150, 159, 168, 185, 202, 218, 234,
		240, 246, 248, 251, 253, 256, 262, 268,
		270, 273, 275, 278, 283,
	}

	var _scanner_trans_keys []int32 = []int32{
		10, 111, 115, 116, 9, 32, 9, 32,
		34, 35, 44, 10, 13, 9, 32, 35,
		44, 10, 13, 9, 32, 34, 35, 44,
		10, 13, 9, 32, 35, 44, 10, 13,
		9, 32, 45, 46, 48, 58, 65, 70,
		71, 90, 97, 102, 103, 122, 9, 32,
		45, 65, 90, 97, 122, 9, 32, 44,
		10, 13, 34, 35, 9, 10, 32, 35,
		44, 11, 13, 9, 10, 32, 35, 95,
		45, 46, 48, 57, 65, 90, 97, 122,
		10, 61, 95, 45, 46, 48, 57, 65,
		90, 97, 122, 95, 45, 46, 48, 57,
		65, 90, 97, 122, 9, 10, 32, 35,
		95, 45, 46, 48, 57, 65, 90, 97,
		122, 9, 32, 47, 46, 58, 65, 70,
		97, 102, 9, 32, 48, 57, 9, 32,
		48, 57, 48, 57, 9, 32, 48, 57,
		9, 32, 45, 47, 46, 58, 65, 70,
		71, 90, 97, 102, 103, 122, 9, 32,
		44, 10, 13, 34, 35, 48, 57, 9,
		10, 32, 35, 44, 11, 13, 48, 57,
		9, 10, 32, 34, 35, 44, 95, 11,
		13, 45, 46, 48, 57, 65, 90, 97,
		122, 9, 10, 32, 35, 44, 61, 95,
		11, 13, 45, 46, 48, 57, 65, 90,
		97, 122, 9, 10, 32, 35, 44, 95,
		11, 13, 45, 46, 48, 57, 65, 90,
		97, 122, 9, 10, 32, 35, 44, 95,
		11, 13, 45, 46, 48, 57, 65, 90,
		97, 122, 32, 34, 35, 44, 9, 13,
		9, 32, 35, 44, 10, 13, 10, 34,
		9, 32, 44, 10, 34, 9, 32, 44,
		32, 34, 35, 44, 9, 13, 9, 32,
		35, 44, 10, 13, 10, 34, 9, 32,
		44, 10, 34, 9, 32, 44, 32, 35,
		104, 9, 13, 32, 35, 104, 9, 13,
	}

	var _scanner_single_lengths []byte = []byte{
		0, 1, 1, 1, 1, 2, 5, 4,
		5, 4, 4, 3, 3, 5, 5, 1,
		2, 1, 5, 3, 2, 2, 0, 2,
		4, 3, 5, 7, 7, 6, 6, 4,
		4, 2, 3, 2, 3, 4, 4, 2,
		3, 2, 3, 3, 3,
	}

	var _scanner_range_lengths []byte = []byte{
		0, 0, 0, 0, 0, 0, 1, 1,
		1, 1, 5, 2, 2, 1, 4, 0,
		4, 4, 4, 3, 1, 1, 1, 1,
		5, 3, 2, 5, 5, 5, 5, 1,
		1, 0, 0, 0, 0, 1, 1, 0,
		0, 0, 0, 1, 1,
	}

	var _scanner_index_offsets []int16 = []int16{
		0, 0, 2, 4, 6, 8, 11, 18,
		24, 31, 37, 47, 53, 59, 66, 76,
		78, 85, 91, 101, 108, 112, 116, 118,
		122, 132, 139, 147, 160, 173, 185, 197,
		203, 209, 212, 216, 219, 223, 229, 235,
		238, 242, 245, 249, 254,
	}

	var _scanner_indicies []byte = []byte{
		1, 0, 3, 2, 4, 2, 5, 2,
		6, 6, 2, 8, 8, 9, 2, 2,
		2, 7, 11, 11, 2, 12, 2, 10,
		14, 14, 15, 2, 2, 2, 13, 17,
		17, 2, 18, 2, 16, 19, 19, 20,
		21, 21, 22, 20, 22, 20, 2, 23,
		23, 24, 24, 24, 2, 26, 26, 2,
		2, 2, 25, 28, 29, 28, 30, 2,
		2, 27, 31, 32, 31, 33, 34, 34,
		34, 34, 34, 2, 32, 33, 36, 35,
		35, 35, 35, 35, 2, 37, 37, 37,
		37, 37, 2, 38, 39, 38, 40, 37,
		37, 37, 37, 37, 2, 41, 41, 43,
		42, 42, 42, 2, 41, 41, 44, 2,
		45, 45, 44, 2, 46, 2, 47, 47,
		46, 2, 48, 48, 24, 43, 42, 49,
		24, 49, 24, 2, 50, 50, 2, 2,
		2, 51, 25, 52, 29, 52, 30, 2,
		2, 53, 27, 54, 32, 54, 2, 33,
		2, 55, 2, 55, 55, 55, 55, 25,
		28, 29, 28, 30, 2, 57, 56, 2,
		56, 56, 56, 56, 27, 28, 29, 28,
		30, 2, 58, 2, 58, 58, 58, 58,
		27, 59, 60, 59, 61, 2, 58, 2,
		58, 58, 58, 58, 27, 2, 63, 2,
		2, 2, 62, 65, 65, 2, 66, 2,
		64, 2, 68, 67, 69, 69, 70, 2,
		2, 72, 71, 73, 73, 74, 2, 2,
		76, 2, 2, 2, 75, 78, 78, 2,
		79, 2, 77, 2, 81, 80, 82, 82,
		83, 2, 2, 85, 84, 86, 86, 87,
		2, 1, 0, 89, 1, 88, 90, 91,
		92, 90, 88,
	}

	var _scanner_trans_targs []byte = []byte{
		1, 43, 0, 3, 4, 5, 6, 7,
		6, 41, 7, 8, 37, 9, 8, 35,
		9, 10, 31, 10, 11, 19, 24, 12,
		11, 13, 12, 13, 14, 44, 15, 14,
		44, 15, 16, 16, 17, 18, 14, 44,
		15, 20, 19, 22, 21, 12, 23, 12,
		25, 24, 25, 26, 27, 26, 27, 28,
		28, 29, 30, 14, 44, 15, 32, 33,
		32, 10, 31, 33, 34, 10, 31, 35,
		36, 10, 31, 38, 39, 38, 8, 37,
		39, 40, 8, 37, 41, 42, 8, 37,
		0, 2, 43, 1, 2,
	}

	var _scanner_trans_actions []byte = []byte{
		0, 0, 11, 0, 0, 0, 5, 1,
		0, 3, 0, 56, 19, 1, 0, 3,
		0, 60, 19, 0, 1, 1, 1, 25,
		0, 1, 0, 0, 28, 28, 28, 0,
		0, 0, 1, 0, 0, 0, 7, 7,
		7, 0, 0, 0, 0, 34, 0, 31,
		25, 0, 0, 1, 76, 0, 0, 1,
		0, 0, 0, 72, 72, 72, 1, 3,
		0, 68, 22, 0, 0, 52, 16, 0,
		0, 44, 13, 1, 3, 0, 64, 22,
		0, 0, 48, 16, 0, 0, 40, 13,
		37, 0, 9, 9, 9,
	}

	var _scanner_eof_actions []byte = []byte{
		0, 0, 11, 11, 11, 11, 11, 11,
		11, 11, 11, 11, 11, 11, 11, 11,
		11, 11, 11, 11, 11, 11, 11, 11,
		11, 11, 11, 11, 11, 11, 11, 11,
		11, 11, 11, 11, 11, 11, 11, 11,
		11, 11, 11, 0, 9,
	}

	const scanner_start int = 43
	const scanner_first_final int = 43
	const scanner_error int = 0

	const scanner_en_main int = 43

//line conf.rl:31

	// These are generated by ragel. Reference them to avoid unused lint errors.
	_, _, _ = scanner_first_final, scanner_error, scanner_en_main

	cs, p, pe, eof := 0, 0, len(data), len(data)

	var (
		mark   int
		ms     []String
		s      String
		ipn    *net.IPNet
		e      Entry
		err    error
		d      string
		option [2]string
		conf   Conf
	)

//line conf.go:228
	{
		cs = scanner_start
	}

//line conf.go:233
	{
		var _klen int
		var _trans int
		var _acts int
		var _nacts uint
		var _keys int
		if p == pe {
			goto _test_eof
		}
		if cs == 0 {
			goto _out
		}
	_resume:
		_keys = int(_scanner_key_offsets[cs])
		_trans = int(_scanner_index_offsets[cs])

		_klen = int(_scanner_single_lengths[cs])
		if _klen > 0 {
			_lower := int(_keys)
			var _mid int
			_upper := int(_keys + _klen - 1)
			for {
				if _upper < _lower {
					break
				}

				_mid = _lower + ((_upper - _lower) >> 1)
				switch {
				case data[p] < _scanner_trans_keys[_mid]:
					_upper = _mid - 1
				case data[p] > _scanner_trans_keys[_mid]:
					_lower = _mid + 1
				default:
					_trans += int(_mid - int(_keys))
					goto _match
				}
			}
			_keys += _klen
			_trans += _klen
		}

		_klen = int(_scanner_range_lengths[cs])
		if _klen > 0 {
			_lower := int(_keys)
			var _mid int
			_upper := int(_keys + (_klen << 1) - 2)
			for {
				if _upper < _lower {
					break
				}

				_mid = _lower + (((_upper - _lower) >> 1) & ^1)
				switch {
				case data[p] < _scanner_trans_keys[_mid]:
					_upper = _mid - 2
				case data[p] > _scanner_trans_keys[_mid+1]:
					_lower = _mid + 2
				default:
					_trans += int((_mid - int(_keys)) >> 1)
					goto _match
				}
			}
			_trans += _klen
		}

	_match:
		_trans = int(_scanner_indicies[_trans])
		cs = int(_scanner_trans_targs[_trans])

		if _scanner_trans_actions[_trans] == 0 {
			goto _again
		}

		_acts = int(_scanner_trans_actions[_trans])
		_nacts = uint(_scanner_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _scanner_actions[_acts-1] {
			case 0:
//line conf.rl:53
				mark = p
			case 1:
//line conf.rl:55

				s = String{
					Value:  string(data[mark : p-1]),
					Quoted: true,
				}

			case 2:
//line conf.rl:61

				s = String{Value: string(data[mark:p])}

			case 3:
//line conf.rl:65
				mark = p + 1
			case 4:
//line conf.rl:76
				ms = []String{s}
			case 5:
//line conf.rl:77
				ms = append(ms, s)
			case 6:
//line conf.rl:80

				d = string(data[mark:p])

			case 7:
//line conf.rl:83

				d = strings.Join(strings.Fields(string(data[mark:p])), "/")

			case 8:
//line conf.rl:86

				_, ipn, err = net.ParseCIDR(d)
				if err != nil {
					return nil, err
				}
				e.Address = ipn

			case 9:
//line conf.rl:93

				e.Address = s

			case 10:
//line conf.rl:109

				e = Entry{Type: "host"}

			case 11:
//line conf.rl:112

				e.Database = ms

			case 12:
//line conf.rl:115

				e.User = ms

			case 13:
//line conf.rl:118

				e.Method = string(data[mark:p])

			case 14:
//line conf.rl:121

				copy(option[:], strings.Split(string(data[mark:p]), "="))
				e.Options = append(e.Options, option)

			case 15:
//line conf.rl:131

				conf.Entries = append(conf.Entries, e)

			case 16:
//line conf.rl:146
				return nil, errors.Errorf("entry %d invalid", len(conf.Entries)+1)
			case 17:
//line conf.rl:152
				return nil, errors.New("invalid")
//line conf.go:398
			}
		}

	_again:
		if cs == 0 {
			goto _out
		}
		p++
		if p != pe {
			goto _resume
		}
	_test_eof:
		{
		}
		if p == eof {
			__acts := _scanner_eof_actions[cs]
			__nacts := uint(_scanner_actions[__acts])
			__acts++
			for ; __nacts > 0; __nacts-- {
				__acts++
				switch _scanner_actions[__acts-1] {
				case 15:
//line conf.rl:131

					conf.Entries = append(conf.Entries, e)

				case 16:
//line conf.rl:146
					return nil, errors.Errorf("entry %d invalid", len(conf.Entries)+1)
//line conf.go:425
				}
			}
		}

	_out:
		{
		}
	}

//line conf.rl:160

	if len(conf.Entries) == 0 {
		return nil, errors.New("no entries")
	}

	return &conf, nil
}
