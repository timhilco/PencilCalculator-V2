// Code generated from HilcoPencilGrammarLexer.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type HilcoPencilGrammarLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var HilcoPencilGrammarLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func hilcopencilgrammarlexerLexerInit() {
	staticData := &HilcoPencilGrammarLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "':='", "'case'", "'endcase'", "'is'", "'switch'", "'endswitch'",
		"'if'", "'then'", "'else'", "'::'", "';'", "'NOT'", "'^'", "'*'", "'/'",
		"'+'", "'-'", "'>'", "'>='", "'<'", "'<='", "'='", "'~='", "'AND'",
		"'OR'", "':'", "'['", "']'", "'{'", "'}'", "'('", "')'", "'_'", "'%'",
		"'@'", "','", "'.'", "'true'", "'false'", "'nil'",
	}
	staticData.SymbolicNames = []string{
		"", "ASSIGNMENT", "CASE", "END_CASE", "IS", "SWITCH", "END_SWITCH",
		"IF", "THEN", "ELSE", "DOUBLE_COLON", "SEMI_COLON", "NOT", "EXPONENTIAL",
		"MULTIPLY", "DIVIDE", "ADD", "MINUS", "GREATER_THAN", "GREATER_THAN_EQUAL",
		"LESS_THAN", "LESS_THAN_EQUAL", "EQUAL", "NOT_EQUAL", "AND", "OR", "COLON",
		"LBRACKET", "RBRACKET", "CURLYLBRACKET", "CURLYRBRACKET", "LPAREN",
		"RPAREN", "UNDERSCORE", "PERCENT_SIGN", "AT_SIGN", "COMMA", "DOT", "KEYWORD_TRUE",
		"KEYWORD_FALSE", "KEYWORD_NIL", "CLASSNAME", "ID", "ATFUNCTION", "INT",
		"STRING", "FLOAT", "PERCENT", "DATE", "NEWLINE", "WS",
	}
	staticData.RuleNames = []string{
		"ASSIGNMENT", "CASE", "END_CASE", "IS", "SWITCH", "END_SWITCH", "IF",
		"THEN", "ELSE", "DOUBLE_COLON", "SEMI_COLON", "NOT", "EXPONENTIAL",
		"MULTIPLY", "DIVIDE", "ADD", "MINUS", "GREATER_THAN", "GREATER_THAN_EQUAL",
		"LESS_THAN", "LESS_THAN_EQUAL", "EQUAL", "NOT_EQUAL", "AND", "OR", "COLON",
		"LBRACKET", "RBRACKET", "CURLYLBRACKET", "CURLYRBRACKET", "LPAREN",
		"RPAREN", "UNDERSCORE", "PERCENT_SIGN", "AT_SIGN", "COMMA", "DOT", "KEYWORD_TRUE",
		"KEYWORD_FALSE", "KEYWORD_NIL", "DIGIT", "LOWERCASELETTERS", "UPPERCASELETTERS",
		"CLASSNAME", "ID", "ATFUNCTION", "INT", "STRING", "FLOAT", "PERCENT",
		"DATE", "EXPONENT", "HEX_DIGIT", "ESC_SEQ", "OCTAL_ESC", "UNICODE_ESC",
		"NEWLINE", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 50, 401, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21,
		1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1,
		24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29,
		1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1,
		35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38,
		1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1,
		41, 1, 41, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 5, 43, 258,
		8, 43, 10, 43, 12, 43, 261, 9, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1,
		44, 1, 44, 5, 44, 270, 8, 44, 10, 44, 12, 44, 273, 9, 44, 1, 45, 1, 45,
		1, 45, 3, 45, 278, 8, 45, 1, 45, 1, 45, 1, 45, 5, 45, 283, 8, 45, 10, 45,
		12, 45, 286, 9, 45, 1, 46, 4, 46, 289, 8, 46, 11, 46, 12, 46, 290, 1, 47,
		1, 47, 1, 47, 5, 47, 296, 8, 47, 10, 47, 12, 47, 299, 9, 47, 1, 47, 1,
		47, 1, 48, 4, 48, 304, 8, 48, 11, 48, 12, 48, 305, 1, 48, 1, 48, 5, 48,
		310, 8, 48, 10, 48, 12, 48, 313, 9, 48, 1, 49, 1, 49, 3, 49, 317, 8, 49,
		1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1,
		50, 3, 50, 330, 8, 50, 1, 50, 1, 50, 3, 50, 334, 8, 50, 1, 50, 1, 50, 3,
		50, 338, 8, 50, 3, 50, 340, 8, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1,
		50, 3, 50, 348, 8, 50, 1, 50, 1, 50, 1, 51, 1, 51, 3, 51, 354, 8, 51, 1,
		51, 4, 51, 357, 8, 51, 11, 51, 12, 51, 358, 1, 52, 1, 52, 1, 53, 1, 53,
		1, 53, 1, 53, 3, 53, 367, 8, 53, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1,
		54, 1, 54, 1, 54, 1, 54, 3, 54, 378, 8, 54, 1, 55, 1, 55, 1, 55, 1, 55,
		1, 55, 1, 55, 1, 55, 1, 56, 3, 56, 388, 8, 56, 1, 56, 1, 56, 1, 56, 1,
		56, 1, 57, 1, 57, 1, 57, 1, 57, 3, 57, 398, 8, 57, 1, 57, 1, 57, 0, 0,
		58, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21,
		11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39,
		20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57,
		29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75,
		38, 77, 39, 79, 40, 81, 0, 83, 0, 85, 0, 87, 41, 89, 42, 91, 43, 93, 44,
		95, 45, 97, 46, 99, 47, 101, 48, 103, 0, 105, 0, 107, 0, 109, 0, 111, 0,
		113, 49, 115, 50, 1, 0, 7, 2, 0, 39, 39, 92, 92, 2, 0, 45, 45, 47, 47,
		2, 0, 69, 69, 101, 101, 2, 0, 43, 43, 45, 45, 3, 0, 48, 57, 65, 70, 97,
		102, 8, 0, 34, 34, 39, 39, 92, 92, 98, 98, 102, 102, 110, 110, 114, 114,
		116, 116, 3, 0, 9, 9, 12, 12, 32, 32, 426, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1,
		0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27,
		1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0,
		35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0,
		0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0,
		0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0,
		0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1,
		0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73,
		1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0,
		87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0,
		0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0,
		0, 0, 0, 113, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0, 1, 117, 1, 0, 0, 0, 3, 120,
		1, 0, 0, 0, 5, 125, 1, 0, 0, 0, 7, 133, 1, 0, 0, 0, 9, 136, 1, 0, 0, 0,
		11, 143, 1, 0, 0, 0, 13, 153, 1, 0, 0, 0, 15, 156, 1, 0, 0, 0, 17, 161,
		1, 0, 0, 0, 19, 166, 1, 0, 0, 0, 21, 169, 1, 0, 0, 0, 23, 171, 1, 0, 0,
		0, 25, 175, 1, 0, 0, 0, 27, 177, 1, 0, 0, 0, 29, 179, 1, 0, 0, 0, 31, 181,
		1, 0, 0, 0, 33, 183, 1, 0, 0, 0, 35, 185, 1, 0, 0, 0, 37, 187, 1, 0, 0,
		0, 39, 190, 1, 0, 0, 0, 41, 192, 1, 0, 0, 0, 43, 195, 1, 0, 0, 0, 45, 197,
		1, 0, 0, 0, 47, 200, 1, 0, 0, 0, 49, 204, 1, 0, 0, 0, 51, 207, 1, 0, 0,
		0, 53, 209, 1, 0, 0, 0, 55, 211, 1, 0, 0, 0, 57, 213, 1, 0, 0, 0, 59, 215,
		1, 0, 0, 0, 61, 217, 1, 0, 0, 0, 63, 219, 1, 0, 0, 0, 65, 221, 1, 0, 0,
		0, 67, 223, 1, 0, 0, 0, 69, 225, 1, 0, 0, 0, 71, 227, 1, 0, 0, 0, 73, 229,
		1, 0, 0, 0, 75, 231, 1, 0, 0, 0, 77, 236, 1, 0, 0, 0, 79, 242, 1, 0, 0,
		0, 81, 246, 1, 0, 0, 0, 83, 248, 1, 0, 0, 0, 85, 250, 1, 0, 0, 0, 87, 252,
		1, 0, 0, 0, 89, 262, 1, 0, 0, 0, 91, 274, 1, 0, 0, 0, 93, 288, 1, 0, 0,
		0, 95, 292, 1, 0, 0, 0, 97, 303, 1, 0, 0, 0, 99, 316, 1, 0, 0, 0, 101,
		320, 1, 0, 0, 0, 103, 351, 1, 0, 0, 0, 105, 360, 1, 0, 0, 0, 107, 366,
		1, 0, 0, 0, 109, 377, 1, 0, 0, 0, 111, 379, 1, 0, 0, 0, 113, 387, 1, 0,
		0, 0, 115, 397, 1, 0, 0, 0, 117, 118, 5, 58, 0, 0, 118, 119, 5, 61, 0,
		0, 119, 2, 1, 0, 0, 0, 120, 121, 5, 99, 0, 0, 121, 122, 5, 97, 0, 0, 122,
		123, 5, 115, 0, 0, 123, 124, 5, 101, 0, 0, 124, 4, 1, 0, 0, 0, 125, 126,
		5, 101, 0, 0, 126, 127, 5, 110, 0, 0, 127, 128, 5, 100, 0, 0, 128, 129,
		5, 99, 0, 0, 129, 130, 5, 97, 0, 0, 130, 131, 5, 115, 0, 0, 131, 132, 5,
		101, 0, 0, 132, 6, 1, 0, 0, 0, 133, 134, 5, 105, 0, 0, 134, 135, 5, 115,
		0, 0, 135, 8, 1, 0, 0, 0, 136, 137, 5, 115, 0, 0, 137, 138, 5, 119, 0,
		0, 138, 139, 5, 105, 0, 0, 139, 140, 5, 116, 0, 0, 140, 141, 5, 99, 0,
		0, 141, 142, 5, 104, 0, 0, 142, 10, 1, 0, 0, 0, 143, 144, 5, 101, 0, 0,
		144, 145, 5, 110, 0, 0, 145, 146, 5, 100, 0, 0, 146, 147, 5, 115, 0, 0,
		147, 148, 5, 119, 0, 0, 148, 149, 5, 105, 0, 0, 149, 150, 5, 116, 0, 0,
		150, 151, 5, 99, 0, 0, 151, 152, 5, 104, 0, 0, 152, 12, 1, 0, 0, 0, 153,
		154, 5, 105, 0, 0, 154, 155, 5, 102, 0, 0, 155, 14, 1, 0, 0, 0, 156, 157,
		5, 116, 0, 0, 157, 158, 5, 104, 0, 0, 158, 159, 5, 101, 0, 0, 159, 160,
		5, 110, 0, 0, 160, 16, 1, 0, 0, 0, 161, 162, 5, 101, 0, 0, 162, 163, 5,
		108, 0, 0, 163, 164, 5, 115, 0, 0, 164, 165, 5, 101, 0, 0, 165, 18, 1,
		0, 0, 0, 166, 167, 5, 58, 0, 0, 167, 168, 5, 58, 0, 0, 168, 20, 1, 0, 0,
		0, 169, 170, 5, 59, 0, 0, 170, 22, 1, 0, 0, 0, 171, 172, 5, 78, 0, 0, 172,
		173, 5, 79, 0, 0, 173, 174, 5, 84, 0, 0, 174, 24, 1, 0, 0, 0, 175, 176,
		5, 94, 0, 0, 176, 26, 1, 0, 0, 0, 177, 178, 5, 42, 0, 0, 178, 28, 1, 0,
		0, 0, 179, 180, 5, 47, 0, 0, 180, 30, 1, 0, 0, 0, 181, 182, 5, 43, 0, 0,
		182, 32, 1, 0, 0, 0, 183, 184, 5, 45, 0, 0, 184, 34, 1, 0, 0, 0, 185, 186,
		5, 62, 0, 0, 186, 36, 1, 0, 0, 0, 187, 188, 5, 62, 0, 0, 188, 189, 5, 61,
		0, 0, 189, 38, 1, 0, 0, 0, 190, 191, 5, 60, 0, 0, 191, 40, 1, 0, 0, 0,
		192, 193, 5, 60, 0, 0, 193, 194, 5, 61, 0, 0, 194, 42, 1, 0, 0, 0, 195,
		196, 5, 61, 0, 0, 196, 44, 1, 0, 0, 0, 197, 198, 5, 126, 0, 0, 198, 199,
		5, 61, 0, 0, 199, 46, 1, 0, 0, 0, 200, 201, 5, 65, 0, 0, 201, 202, 5, 78,
		0, 0, 202, 203, 5, 68, 0, 0, 203, 48, 1, 0, 0, 0, 204, 205, 5, 79, 0, 0,
		205, 206, 5, 82, 0, 0, 206, 50, 1, 0, 0, 0, 207, 208, 5, 58, 0, 0, 208,
		52, 1, 0, 0, 0, 209, 210, 5, 91, 0, 0, 210, 54, 1, 0, 0, 0, 211, 212, 5,
		93, 0, 0, 212, 56, 1, 0, 0, 0, 213, 214, 5, 123, 0, 0, 214, 58, 1, 0, 0,
		0, 215, 216, 5, 125, 0, 0, 216, 60, 1, 0, 0, 0, 217, 218, 5, 40, 0, 0,
		218, 62, 1, 0, 0, 0, 219, 220, 5, 41, 0, 0, 220, 64, 1, 0, 0, 0, 221, 222,
		5, 95, 0, 0, 222, 66, 1, 0, 0, 0, 223, 224, 5, 37, 0, 0, 224, 68, 1, 0,
		0, 0, 225, 226, 5, 64, 0, 0, 226, 70, 1, 0, 0, 0, 227, 228, 5, 44, 0, 0,
		228, 72, 1, 0, 0, 0, 229, 230, 5, 46, 0, 0, 230, 74, 1, 0, 0, 0, 231, 232,
		5, 116, 0, 0, 232, 233, 5, 114, 0, 0, 233, 234, 5, 117, 0, 0, 234, 235,
		5, 101, 0, 0, 235, 76, 1, 0, 0, 0, 236, 237, 5, 102, 0, 0, 237, 238, 5,
		97, 0, 0, 238, 239, 5, 108, 0, 0, 239, 240, 5, 115, 0, 0, 240, 241, 5,
		101, 0, 0, 241, 78, 1, 0, 0, 0, 242, 243, 5, 110, 0, 0, 243, 244, 5, 105,
		0, 0, 244, 245, 5, 108, 0, 0, 245, 80, 1, 0, 0, 0, 246, 247, 2, 48, 57,
		0, 247, 82, 1, 0, 0, 0, 248, 249, 2, 97, 122, 0, 249, 84, 1, 0, 0, 0, 250,
		251, 2, 65, 90, 0, 251, 86, 1, 0, 0, 0, 252, 259, 3, 85, 42, 0, 253, 258,
		3, 83, 41, 0, 254, 258, 3, 85, 42, 0, 255, 258, 3, 81, 40, 0, 256, 258,
		3, 65, 32, 0, 257, 253, 1, 0, 0, 0, 257, 254, 1, 0, 0, 0, 257, 255, 1,
		0, 0, 0, 257, 256, 1, 0, 0, 0, 258, 261, 1, 0, 0, 0, 259, 257, 1, 0, 0,
		0, 259, 260, 1, 0, 0, 0, 260, 88, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 262,
		271, 3, 83, 41, 0, 263, 270, 3, 83, 41, 0, 264, 270, 3, 85, 42, 0, 265,
		270, 3, 81, 40, 0, 266, 270, 3, 65, 32, 0, 267, 270, 3, 51, 25, 0, 268,
		270, 5, 92, 0, 0, 269, 263, 1, 0, 0, 0, 269, 264, 1, 0, 0, 0, 269, 265,
		1, 0, 0, 0, 269, 266, 1, 0, 0, 0, 269, 267, 1, 0, 0, 0, 269, 268, 1, 0,
		0, 0, 270, 273, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0,
		272, 90, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 274, 277, 3, 69, 34, 0, 275,
		278, 3, 83, 41, 0, 276, 278, 3, 85, 42, 0, 277, 275, 1, 0, 0, 0, 277, 276,
		1, 0, 0, 0, 278, 284, 1, 0, 0, 0, 279, 283, 3, 83, 41, 0, 280, 283, 3,
		85, 42, 0, 281, 283, 3, 81, 40, 0, 282, 279, 1, 0, 0, 0, 282, 280, 1, 0,
		0, 0, 282, 281, 1, 0, 0, 0, 283, 286, 1, 0, 0, 0, 284, 282, 1, 0, 0, 0,
		284, 285, 1, 0, 0, 0, 285, 92, 1, 0, 0, 0, 286, 284, 1, 0, 0, 0, 287, 289,
		3, 81, 40, 0, 288, 287, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 288, 1,
		0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 94, 1, 0, 0, 0, 292, 297, 5, 39, 0,
		0, 293, 296, 3, 107, 53, 0, 294, 296, 8, 0, 0, 0, 295, 293, 1, 0, 0, 0,
		295, 294, 1, 0, 0, 0, 296, 299, 1, 0, 0, 0, 297, 295, 1, 0, 0, 0, 297,
		298, 1, 0, 0, 0, 298, 300, 1, 0, 0, 0, 299, 297, 1, 0, 0, 0, 300, 301,
		5, 39, 0, 0, 301, 96, 1, 0, 0, 0, 302, 304, 3, 81, 40, 0, 303, 302, 1,
		0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 303, 1, 0, 0, 0, 305, 306, 1, 0, 0,
		0, 306, 307, 1, 0, 0, 0, 307, 311, 5, 46, 0, 0, 308, 310, 3, 81, 40, 0,
		309, 308, 1, 0, 0, 0, 310, 313, 1, 0, 0, 0, 311, 309, 1, 0, 0, 0, 311,
		312, 1, 0, 0, 0, 312, 98, 1, 0, 0, 0, 313, 311, 1, 0, 0, 0, 314, 317, 3,
		93, 46, 0, 315, 317, 3, 97, 48, 0, 316, 314, 1, 0, 0, 0, 316, 315, 1, 0,
		0, 0, 317, 318, 1, 0, 0, 0, 318, 319, 3, 67, 33, 0, 319, 100, 1, 0, 0,
		0, 320, 321, 3, 57, 28, 0, 321, 322, 3, 81, 40, 0, 322, 323, 3, 81, 40,
		0, 323, 339, 7, 1, 0, 0, 324, 325, 3, 81, 40, 0, 325, 326, 3, 81, 40, 0,
		326, 340, 1, 0, 0, 0, 327, 330, 3, 83, 41, 0, 328, 330, 3, 85, 42, 0, 329,
		327, 1, 0, 0, 0, 329, 328, 1, 0, 0, 0, 330, 333, 1, 0, 0, 0, 331, 334,
		3, 83, 41, 0, 332, 334, 3, 85, 42, 0, 333, 331, 1, 0, 0, 0, 333, 332, 1,
		0, 0, 0, 334, 337, 1, 0, 0, 0, 335, 338, 3, 83, 41, 0, 336, 338, 3, 85,
		42, 0, 337, 335, 1, 0, 0, 0, 337, 336, 1, 0, 0, 0, 338, 340, 1, 0, 0, 0,
		339, 324, 1, 0, 0, 0, 339, 329, 1, 0, 0, 0, 340, 341, 1, 0, 0, 0, 341,
		342, 7, 1, 0, 0, 342, 343, 3, 81, 40, 0, 343, 347, 3, 81, 40, 0, 344, 345,
		3, 81, 40, 0, 345, 346, 3, 81, 40, 0, 346, 348, 1, 0, 0, 0, 347, 344, 1,
		0, 0, 0, 347, 348, 1, 0, 0, 0, 348, 349, 1, 0, 0, 0, 349, 350, 3, 59, 29,
		0, 350, 102, 1, 0, 0, 0, 351, 353, 7, 2, 0, 0, 352, 354, 7, 3, 0, 0, 353,
		352, 1, 0, 0, 0, 353, 354, 1, 0, 0, 0, 354, 356, 1, 0, 0, 0, 355, 357,
		2, 48, 57, 0, 356, 355, 1, 0, 0, 0, 357, 358, 1, 0, 0, 0, 358, 356, 1,
		0, 0, 0, 358, 359, 1, 0, 0, 0, 359, 104, 1, 0, 0, 0, 360, 361, 7, 4, 0,
		0, 361, 106, 1, 0, 0, 0, 362, 363, 5, 92, 0, 0, 363, 367, 7, 5, 0, 0, 364,
		367, 3, 111, 55, 0, 365, 367, 3, 109, 54, 0, 366, 362, 1, 0, 0, 0, 366,
		364, 1, 0, 0, 0, 366, 365, 1, 0, 0, 0, 367, 108, 1, 0, 0, 0, 368, 369,
		5, 92, 0, 0, 369, 370, 2, 48, 51, 0, 370, 371, 2, 48, 55, 0, 371, 378,
		2, 48, 55, 0, 372, 373, 5, 92, 0, 0, 373, 374, 2, 48, 55, 0, 374, 378,
		2, 48, 55, 0, 375, 376, 5, 92, 0, 0, 376, 378, 2, 48, 55, 0, 377, 368,
		1, 0, 0, 0, 377, 372, 1, 0, 0, 0, 377, 375, 1, 0, 0, 0, 378, 110, 1, 0,
		0, 0, 379, 380, 5, 92, 0, 0, 380, 381, 5, 117, 0, 0, 381, 382, 3, 105,
		52, 0, 382, 383, 3, 105, 52, 0, 383, 384, 3, 105, 52, 0, 384, 385, 3, 105,
		52, 0, 385, 112, 1, 0, 0, 0, 386, 388, 5, 13, 0, 0, 387, 386, 1, 0, 0,
		0, 387, 388, 1, 0, 0, 0, 388, 389, 1, 0, 0, 0, 389, 390, 5, 10, 0, 0, 390,
		391, 1, 0, 0, 0, 391, 392, 6, 56, 0, 0, 392, 114, 1, 0, 0, 0, 393, 398,
		7, 6, 0, 0, 394, 395, 5, 13, 0, 0, 395, 398, 5, 10, 0, 0, 396, 398, 5,
		3, 0, 0, 397, 393, 1, 0, 0, 0, 397, 394, 1, 0, 0, 0, 397, 396, 1, 0, 0,
		0, 398, 399, 1, 0, 0, 0, 399, 400, 6, 57, 0, 0, 400, 116, 1, 0, 0, 0, 25,
		0, 257, 259, 269, 271, 277, 282, 284, 290, 295, 297, 305, 311, 316, 329,
		333, 337, 339, 347, 353, 358, 366, 377, 387, 397, 1, 0, 1, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// HilcoPencilGrammarLexerInit initializes any static state used to implement HilcoPencilGrammarLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewHilcoPencilGrammarLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func HilcoPencilGrammarLexerInit() {
	staticData := &HilcoPencilGrammarLexerLexerStaticData
	staticData.once.Do(hilcopencilgrammarlexerLexerInit)
}

// NewHilcoPencilGrammarLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewHilcoPencilGrammarLexer(input antlr.CharStream) *HilcoPencilGrammarLexer {
	HilcoPencilGrammarLexerInit()
	l := new(HilcoPencilGrammarLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &HilcoPencilGrammarLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "HilcoPencilGrammarLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// HilcoPencilGrammarLexer tokens.
const (
	HilcoPencilGrammarLexerASSIGNMENT         = 1
	HilcoPencilGrammarLexerCASE               = 2
	HilcoPencilGrammarLexerEND_CASE           = 3
	HilcoPencilGrammarLexerIS                 = 4
	HilcoPencilGrammarLexerSWITCH             = 5
	HilcoPencilGrammarLexerEND_SWITCH         = 6
	HilcoPencilGrammarLexerIF                 = 7
	HilcoPencilGrammarLexerTHEN               = 8
	HilcoPencilGrammarLexerELSE               = 9
	HilcoPencilGrammarLexerDOUBLE_COLON       = 10
	HilcoPencilGrammarLexerSEMI_COLON         = 11
	HilcoPencilGrammarLexerNOT                = 12
	HilcoPencilGrammarLexerEXPONENTIAL        = 13
	HilcoPencilGrammarLexerMULTIPLY           = 14
	HilcoPencilGrammarLexerDIVIDE             = 15
	HilcoPencilGrammarLexerADD                = 16
	HilcoPencilGrammarLexerMINUS              = 17
	HilcoPencilGrammarLexerGREATER_THAN       = 18
	HilcoPencilGrammarLexerGREATER_THAN_EQUAL = 19
	HilcoPencilGrammarLexerLESS_THAN          = 20
	HilcoPencilGrammarLexerLESS_THAN_EQUAL    = 21
	HilcoPencilGrammarLexerEQUAL              = 22
	HilcoPencilGrammarLexerNOT_EQUAL          = 23
	HilcoPencilGrammarLexerAND                = 24
	HilcoPencilGrammarLexerOR                 = 25
	HilcoPencilGrammarLexerCOLON              = 26
	HilcoPencilGrammarLexerLBRACKET           = 27
	HilcoPencilGrammarLexerRBRACKET           = 28
	HilcoPencilGrammarLexerCURLYLBRACKET      = 29
	HilcoPencilGrammarLexerCURLYRBRACKET      = 30
	HilcoPencilGrammarLexerLPAREN             = 31
	HilcoPencilGrammarLexerRPAREN             = 32
	HilcoPencilGrammarLexerUNDERSCORE         = 33
	HilcoPencilGrammarLexerPERCENT_SIGN       = 34
	HilcoPencilGrammarLexerAT_SIGN            = 35
	HilcoPencilGrammarLexerCOMMA              = 36
	HilcoPencilGrammarLexerDOT                = 37
	HilcoPencilGrammarLexerKEYWORD_TRUE       = 38
	HilcoPencilGrammarLexerKEYWORD_FALSE      = 39
	HilcoPencilGrammarLexerKEYWORD_NIL        = 40
	HilcoPencilGrammarLexerCLASSNAME          = 41
	HilcoPencilGrammarLexerID                 = 42
	HilcoPencilGrammarLexerATFUNCTION         = 43
	HilcoPencilGrammarLexerINT                = 44
	HilcoPencilGrammarLexerSTRING             = 45
	HilcoPencilGrammarLexerFLOAT              = 46
	HilcoPencilGrammarLexerPERCENT            = 47
	HilcoPencilGrammarLexerDATE               = 48
	HilcoPencilGrammarLexerNEWLINE            = 49
	HilcoPencilGrammarLexerWS                 = 50
)