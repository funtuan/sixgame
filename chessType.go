package main

import "strings"

type ChessType struct {
	Name      string
	Structure []string
}

var chessTypeList = []ChessType{
	ChessType{
		"6L",
		[]string{"AAAAAA"},
	},
	ChessType{
		"5L:11",
		[]string{"?AAAAA?"},
	},
	ChessType{
		"5L:01",
		[]string{"AAAAA?", "?AAAAA"},
	},
	ChessType{
		"5J1",
		[]string{"A?AAAA", "AAAA?A"},
	},
	ChessType{
		"5P1",
		[]string{"AA?AAA", "AAA?AA"},
	},
	ChessType{
		"4L:22",
		[]string{"??AAAA??"},
	},
	ChessType{
		"4L:12",
		[]string{"?AAAA??", "??AAAA?"},
	},
	ChessType{
		"4L:02",
		[]string{"AAAA??", "??AAAA"},
	},
	ChessType{
		"4L:11",
		[]string{"?AAAA?"},
	},
	ChessType{
		"4J1:11",
		[]string{"?A?AAA?", "?AAA?A?"},
	},
	ChessType{
		"4J1:01",
		[]string{"?A?AAA", "?AAA?A", "A?AAA?", "AAA?A?"},
	},
	ChessType{
		"4J2",
		[]string{"A??AAA", "AAA??A"},
	},
	ChessType{
		"4J1P1",
		[]string{"A?A?AA", "AA?A?A"},
	},
	ChessType{
		"4P2",
		[]string{"AA??AA"},
	},
	ChessType{
		"4J1S1",
		[]string{"A?AA?A"},
	},
	ChessType{
		"3L:33",
		[]string{"???AAA???"},
	},
	ChessType{
		"3L:23",
		[]string{"??AAA???", "???AAA??"},
	},
	ChessType{
		"3L:13",
		[]string{"?AAA???", "???AAA?"},
	},
	ChessType{
		"3L:03",
		[]string{"AAA???", "???AAA"},
	},
	ChessType{
		"3L:22",
		[]string{"??AAA??"},
	},
	ChessType{
		"3L:12",
		[]string{"?AAA??", "??AAA?"},
	},
	ChessType{
		"3J1:22",
		[]string{"??A?AA??", "??AA?A??"},
	},
	ChessType{
		"3J1:12",
		[]string{"?A?AA??", "??A?AA?", "?AA?A??", "??AA?A?"},
	},
	ChessType{
		"3J1:02",
		[]string{"A?AA??", "??A?AA", "AA?A??", "??AA?A"},
	},
	ChessType{
		"3J1:11",
		[]string{"?A?AA?", "?AA?A?"},
	},
	ChessType{
		"3J2:11",
		[]string{"?A??AA?", "?AA??A?"},
	},
	ChessType{
		"3J2:01",
		[]string{"A??AA?", "AA??A?", "?A??AA", "?AA??A"},
	},
	ChessType{
		"3J3",
		[]string{"A???AA", "AA???A"},
	},
	ChessType{
		"3J1P1:11",
		[]string{"?A?A?A?"},
	},
	ChessType{
		"3J1P1:01",
		[]string{"A?A?A?", "?A?A?A"},
	},
	ChessType{
		"3J2P1",
		[]string{"A??A?A", "A?A??A"},
	},
	ChessType{
		"2L:44",
		[]string{"????AA????"},
	},
	ChessType{
		"2L:34",
		[]string{"???AA????", "????AA???"},
	},
	ChessType{
		"2L:24",
		[]string{"??AA????", "????AA??"},
	},
	ChessType{
		"2L:14",
		[]string{"?AA????", "????AA?"},
	},
	ChessType{
		"2L:04",
		[]string{"AA????", "????AA"},
	},
	ChessType{
		"2L:33",
		[]string{"???AA???"},
	},
	ChessType{
		"2L:23",
		[]string{"??AA???", "???AA??"},
	},
	ChessType{
		"2L:13",
		[]string{"?AA???", "???AA?"},
	},
	ChessType{
		"2L:22",
		[]string{"??AA??"},
	},
	ChessType{
		"2J1:33",
		[]string{"???A?A???"},
	},
	ChessType{
		"2J1:23",
		[]string{"???A?A??", "??A?A???"},
	},
	ChessType{
		"2J1:13",
		[]string{"?A?A???", "???A?A?"},
	},
	ChessType{
		"2J1:03",
		[]string{"A?A???", "???A?A"},
	},
	ChessType{
		"2J1:22",
		[]string{"??A?A??"},
	},
	ChessType{
		"2J1:12",
		[]string{"?A?A??", "??A?A?"},
	},
	ChessType{
		"2J2:22",
		[]string{"??A??A??"},
	},
	ChessType{
		"2J2:12",
		[]string{"?A??A??", "??A??A?"},
	},
	ChessType{
		"2J2:02",
		[]string{"A??A??", "??A??A"},
	},
	ChessType{
		"2J2:11",
		[]string{"?A??A?"},
	},
	ChessType{
		"2J3:11",
		[]string{"?A???A?"},
	},
	ChessType{
		"2J3:01",
		[]string{"?A???A", "A???A?"},
	},
	ChessType{
		"2J4",
		[]string{"A????A"},
	},
}

func getChessTypes(chessStructure string) []string {
	matchType := []string{}
	for _, chessType := range chessTypeList {
	matchLoop:
		for _, structure := range chessType.Structure {
			matchIndex := strings.Index(chessStructure, structure)
			if matchIndex != -1 {
				matchType = append(matchType, chessType.Name)
				break matchLoop
			}
		}
	}

	return matchType
}
