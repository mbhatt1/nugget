// Generated from C:/Users/drew/GoglandProjects/nugget\Nugget2.g4 by ANTLR 4.7.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 23, 167,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 119, 10, 15, 3, 16, 3, 16, 3,
	16, 3, 17, 6, 17, 125, 10, 17, 13, 17, 14, 17, 126, 3, 18, 6, 18, 130,
	10, 18, 13, 18, 14, 18, 131, 3, 19, 3, 19, 3, 19, 3, 19, 7, 19, 138, 10,
	19, 12, 19, 14, 19, 141, 11, 19, 3, 19, 3, 19, 3, 20, 6, 20, 146, 10, 20,
	13, 20, 14, 20, 147, 3, 20, 3, 20, 3, 21, 5, 21, 153, 10, 21, 3, 21, 3,
	21, 3, 22, 3, 22, 3, 22, 3, 22, 7, 22, 161, 10, 22, 12, 22, 14, 22, 164,
	11, 22, 3, 22, 3, 22, 2, 2, 23, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15,
	9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33,
	18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 3, 2, 8, 4, 2, 62, 62, 64,
	64, 3, 2, 50, 59, 4, 2, 67, 92, 99, 124, 3, 2, 36, 36, 5, 2, 11, 12, 15,
	15, 34, 34, 4, 2, 12, 12, 15, 15, 2, 176, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3,
	2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21,
	3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2,
	29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2,
	2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2,
	2, 3, 45, 3, 2, 2, 2, 5, 47, 3, 2, 2, 2, 7, 49, 3, 2, 2, 2, 9, 52, 3, 2,
	2, 2, 11, 59, 3, 2, 2, 2, 13, 64, 3, 2, 2, 2, 15, 68, 3, 2, 2, 2, 17, 73,
	3, 2, 2, 2, 19, 78, 3, 2, 2, 2, 21, 85, 3, 2, 2, 2, 23, 94, 3, 2, 2, 2,
	25, 102, 3, 2, 2, 2, 27, 109, 3, 2, 2, 2, 29, 118, 3, 2, 2, 2, 31, 120,
	3, 2, 2, 2, 33, 124, 3, 2, 2, 2, 35, 129, 3, 2, 2, 2, 37, 133, 3, 2, 2,
	2, 39, 145, 3, 2, 2, 2, 41, 152, 3, 2, 2, 2, 43, 156, 3, 2, 2, 2, 45, 46,
	7, 63, 2, 2, 46, 4, 3, 2, 2, 2, 47, 48, 7, 126, 2, 2, 48, 6, 3, 2, 2, 2,
	49, 50, 7, 99, 2, 2, 50, 51, 7, 117, 2, 2, 51, 8, 3, 2, 2, 2, 52, 53, 7,
	117, 2, 2, 53, 54, 7, 118, 2, 2, 54, 55, 7, 116, 2, 2, 55, 56, 7, 107,
	2, 2, 56, 57, 7, 112, 2, 2, 57, 58, 7, 105, 2, 2, 58, 10, 3, 2, 2, 2, 59,
	60, 7, 117, 2, 2, 60, 61, 7, 106, 2, 2, 61, 62, 7, 99, 2, 2, 62, 63, 7,
	51, 2, 2, 63, 12, 3, 2, 2, 2, 64, 65, 7, 111, 2, 2, 65, 66, 7, 102, 2,
	2, 66, 67, 7, 55, 2, 2, 67, 14, 3, 2, 2, 2, 68, 69, 7, 112, 2, 2, 69, 70,
	7, 118, 2, 2, 70, 71, 7, 104, 2, 2, 71, 72, 7, 117, 2, 2, 72, 16, 3, 2,
	2, 2, 73, 74, 7, 104, 2, 2, 74, 75, 7, 107, 2, 2, 75, 76, 7, 110, 2, 2,
	76, 77, 7, 103, 2, 2, 77, 18, 3, 2, 2, 2, 78, 79, 7, 114, 2, 2, 79, 80,
	7, 99, 2, 2, 80, 81, 7, 101, 2, 2, 81, 82, 7, 109, 2, 2, 82, 83, 7, 103,
	2, 2, 83, 84, 7, 118, 2, 2, 84, 20, 3, 2, 2, 2, 85, 86, 7, 103, 2, 2, 86,
	87, 7, 122, 2, 2, 87, 88, 7, 107, 2, 2, 88, 89, 7, 104, 2, 2, 89, 90, 7,
	107, 2, 2, 90, 91, 7, 112, 2, 2, 91, 92, 7, 104, 2, 2, 92, 93, 7, 113,
	2, 2, 93, 22, 3, 2, 2, 2, 94, 95, 7, 103, 2, 2, 95, 96, 7, 122, 2, 2, 96,
	97, 7, 118, 2, 2, 97, 98, 7, 116, 2, 2, 98, 99, 7, 99, 2, 2, 99, 100, 7,
	101, 2, 2, 100, 101, 7, 118, 2, 2, 101, 24, 3, 2, 2, 2, 102, 103, 7, 104,
	2, 2, 103, 104, 7, 107, 2, 2, 104, 105, 7, 110, 2, 2, 105, 106, 7, 118,
	2, 2, 106, 107, 7, 103, 2, 2, 107, 108, 7, 116, 2, 2, 108, 26, 3, 2, 2,
	2, 109, 110, 7, 46, 2, 2, 110, 28, 3, 2, 2, 2, 111, 119, 9, 2, 2, 2, 112,
	113, 7, 64, 2, 2, 113, 119, 7, 63, 2, 2, 114, 115, 7, 62, 2, 2, 115, 119,
	7, 63, 2, 2, 116, 117, 7, 63, 2, 2, 117, 119, 7, 63, 2, 2, 118, 111, 3,
	2, 2, 2, 118, 112, 3, 2, 2, 2, 118, 114, 3, 2, 2, 2, 118, 116, 3, 2, 2,
	2, 119, 30, 3, 2, 2, 2, 120, 121, 7, 93, 2, 2, 121, 122, 7, 95, 2, 2, 122,
	32, 3, 2, 2, 2, 123, 125, 9, 3, 2, 2, 124, 123, 3, 2, 2, 2, 125, 126, 3,
	2, 2, 2, 126, 124, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 34, 3, 2, 2,
	2, 128, 130, 9, 4, 2, 2, 129, 128, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131,
	129, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 36, 3, 2, 2, 2, 133, 139, 7,
	36, 2, 2, 134, 135, 7, 36, 2, 2, 135, 138, 7, 36, 2, 2, 136, 138, 10, 5,
	2, 2, 137, 134, 3, 2, 2, 2, 137, 136, 3, 2, 2, 2, 138, 141, 3, 2, 2, 2,
	139, 137, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 142, 3, 2, 2, 2, 141,
	139, 3, 2, 2, 2, 142, 143, 7, 36, 2, 2, 143, 38, 3, 2, 2, 2, 144, 146,
	9, 6, 2, 2, 145, 144, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147, 145, 3, 2,
	2, 2, 147, 148, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 150, 8, 20, 2, 2,
	150, 40, 3, 2, 2, 2, 151, 153, 7, 15, 2, 2, 152, 151, 3, 2, 2, 2, 152,
	153, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 155, 7, 12, 2, 2, 155, 42,
	3, 2, 2, 2, 156, 157, 7, 49, 2, 2, 157, 158, 7, 49, 2, 2, 158, 162, 3,
	2, 2, 2, 159, 161, 10, 7, 2, 2, 160, 159, 3, 2, 2, 2, 161, 164, 3, 2, 2,
	2, 162, 160, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 165, 3, 2, 2, 2, 164,
	162, 3, 2, 2, 2, 165, 166, 8, 22, 2, 2, 166, 44, 3, 2, 2, 2, 11, 2, 118,
	126, 131, 137, 139, 147, 152, 162, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "'|'", "'as'", "'string'", "'sha1'", "'md5'", "'ntfs'", "'file'",
	"'packet'", "'exifinfo'", "'extract'", "'filter'", "','", "", "'[]'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "COMPOP", "LISTOP",
	"INT", "ID", "STRING", "WS", "NL", "LINE_COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "COMPOP", "LISTOP", "INT", "ID", "STRING",
	"WS", "NL", "LINE_COMMENT",
}

type Nugget2Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewNugget2Lexer(input antlr.CharStream) *Nugget2Lexer {

	l := new(Nugget2Lexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Nugget2.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Nugget2Lexer tokens.
const (
	Nugget2LexerT__0         = 1
	Nugget2LexerT__1         = 2
	Nugget2LexerT__2         = 3
	Nugget2LexerT__3         = 4
	Nugget2LexerT__4         = 5
	Nugget2LexerT__5         = 6
	Nugget2LexerT__6         = 7
	Nugget2LexerT__7         = 8
	Nugget2LexerT__8         = 9
	Nugget2LexerT__9         = 10
	Nugget2LexerT__10        = 11
	Nugget2LexerT__11        = 12
	Nugget2LexerT__12        = 13
	Nugget2LexerCOMPOP       = 14
	Nugget2LexerLISTOP       = 15
	Nugget2LexerINT          = 16
	Nugget2LexerID           = 17
	Nugget2LexerSTRING       = 18
	Nugget2LexerWS           = 19
	Nugget2LexerNL           = 20
	Nugget2LexerLINE_COMMENT = 21
)
