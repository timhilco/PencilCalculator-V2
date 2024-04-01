// Code generated from HilcoPencilGrammarParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // HilcoPencilGrammarParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type HilcoPencilGrammarParser struct {
	*antlr.BaseParser
}

var HilcoPencilGrammarParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func hilcopencilgrammarparserParserInit() {
	staticData := &HilcoPencilGrammarParserParserStaticData
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
		"program", "statement", "expression", "caseStatement", "caseList", "caseItem",
		"ifStatement", "name", "worksheetVariable", "atFunction", "argList",
		"dataAccessor", "accessorList", "accessorObjectOrArray", "accessorObject",
		"accessorArray", "specialKeyword",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 50, 187, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 1, 0, 1, 0, 5, 0, 37, 8, 0, 10, 0, 12, 0, 40, 9, 0, 1, 0,
		3, 0, 43, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 71, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 112,
		8, 2, 10, 2, 12, 2, 115, 9, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 3, 4, 127, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 140, 8, 6, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 3, 9, 151, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 3, 10, 158, 8, 10, 5, 10, 160, 8, 10, 10, 10, 12, 10, 163, 9, 10,
		1, 11, 1, 11, 4, 11, 167, 8, 11, 11, 11, 12, 11, 168, 1, 12, 1, 12, 1,
		12, 1, 13, 1, 13, 3, 13, 176, 8, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 0, 1, 4, 17, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 0, 1, 1, 0, 38, 40, 206, 0, 42,
		1, 0, 0, 0, 2, 44, 1, 0, 0, 0, 4, 70, 1, 0, 0, 0, 6, 116, 1, 0, 0, 0, 8,
		126, 1, 0, 0, 0, 10, 128, 1, 0, 0, 0, 12, 133, 1, 0, 0, 0, 14, 141, 1,
		0, 0, 0, 16, 143, 1, 0, 0, 0, 18, 147, 1, 0, 0, 0, 20, 154, 1, 0, 0, 0,
		22, 164, 1, 0, 0, 0, 24, 170, 1, 0, 0, 0, 26, 175, 1, 0, 0, 0, 28, 177,
		1, 0, 0, 0, 30, 179, 1, 0, 0, 0, 32, 184, 1, 0, 0, 0, 34, 43, 3, 4, 2,
		0, 35, 37, 3, 2, 1, 0, 36, 35, 1, 0, 0, 0, 37, 40, 1, 0, 0, 0, 38, 36,
		1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 43, 1, 0, 0, 0, 40, 38, 1, 0, 0, 0,
		41, 43, 1, 0, 0, 0, 42, 34, 1, 0, 0, 0, 42, 38, 1, 0, 0, 0, 42, 41, 1,
		0, 0, 0, 43, 1, 1, 0, 0, 0, 44, 45, 3, 4, 2, 0, 45, 46, 5, 1, 0, 0, 46,
		47, 3, 4, 2, 0, 47, 48, 5, 11, 0, 0, 48, 3, 1, 0, 0, 0, 49, 50, 6, 2, -1,
		0, 50, 71, 3, 6, 3, 0, 51, 71, 3, 12, 6, 0, 52, 53, 5, 31, 0, 0, 53, 54,
		3, 4, 2, 0, 54, 55, 5, 32, 0, 0, 55, 71, 1, 0, 0, 0, 56, 57, 5, 17, 0,
		0, 57, 71, 3, 4, 2, 25, 58, 59, 5, 12, 0, 0, 59, 71, 3, 4, 2, 24, 60, 71,
		3, 14, 7, 0, 61, 71, 3, 16, 8, 0, 62, 71, 3, 18, 9, 0, 63, 71, 3, 22, 11,
		0, 64, 71, 3, 32, 16, 0, 65, 71, 5, 47, 0, 0, 66, 71, 5, 44, 0, 0, 67,
		71, 5, 46, 0, 0, 68, 71, 5, 45, 0, 0, 69, 71, 5, 48, 0, 0, 70, 49, 1, 0,
		0, 0, 70, 51, 1, 0, 0, 0, 70, 52, 1, 0, 0, 0, 70, 56, 1, 0, 0, 0, 70, 58,
		1, 0, 0, 0, 70, 60, 1, 0, 0, 0, 70, 61, 1, 0, 0, 0, 70, 62, 1, 0, 0, 0,
		70, 63, 1, 0, 0, 0, 70, 64, 1, 0, 0, 0, 70, 65, 1, 0, 0, 0, 70, 66, 1,
		0, 0, 0, 70, 67, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 70, 69, 1, 0, 0, 0, 71,
		113, 1, 0, 0, 0, 72, 73, 10, 23, 0, 0, 73, 74, 5, 13, 0, 0, 74, 112, 3,
		4, 2, 23, 75, 76, 10, 22, 0, 0, 76, 77, 5, 14, 0, 0, 77, 112, 3, 4, 2,
		23, 78, 79, 10, 21, 0, 0, 79, 80, 5, 15, 0, 0, 80, 112, 3, 4, 2, 22, 81,
		82, 10, 20, 0, 0, 82, 83, 5, 16, 0, 0, 83, 112, 3, 4, 2, 21, 84, 85, 10,
		19, 0, 0, 85, 86, 5, 17, 0, 0, 86, 112, 3, 4, 2, 20, 87, 88, 10, 18, 0,
		0, 88, 89, 5, 18, 0, 0, 89, 112, 3, 4, 2, 19, 90, 91, 10, 17, 0, 0, 91,
		92, 5, 19, 0, 0, 92, 112, 3, 4, 2, 18, 93, 94, 10, 16, 0, 0, 94, 95, 5,
		20, 0, 0, 95, 112, 3, 4, 2, 17, 96, 97, 10, 15, 0, 0, 97, 98, 5, 21, 0,
		0, 98, 112, 3, 4, 2, 16, 99, 100, 10, 14, 0, 0, 100, 101, 5, 22, 0, 0,
		101, 112, 3, 4, 2, 15, 102, 103, 10, 13, 0, 0, 103, 104, 5, 23, 0, 0, 104,
		112, 3, 4, 2, 14, 105, 106, 10, 12, 0, 0, 106, 107, 5, 24, 0, 0, 107, 112,
		3, 4, 2, 13, 108, 109, 10, 11, 0, 0, 109, 110, 5, 25, 0, 0, 110, 112, 3,
		4, 2, 12, 111, 72, 1, 0, 0, 0, 111, 75, 1, 0, 0, 0, 111, 78, 1, 0, 0, 0,
		111, 81, 1, 0, 0, 0, 111, 84, 1, 0, 0, 0, 111, 87, 1, 0, 0, 0, 111, 90,
		1, 0, 0, 0, 111, 93, 1, 0, 0, 0, 111, 96, 1, 0, 0, 0, 111, 99, 1, 0, 0,
		0, 111, 102, 1, 0, 0, 0, 111, 105, 1, 0, 0, 0, 111, 108, 1, 0, 0, 0, 112,
		115, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 5, 1,
		0, 0, 0, 115, 113, 1, 0, 0, 0, 116, 117, 5, 2, 0, 0, 117, 118, 3, 4, 2,
		0, 118, 119, 5, 4, 0, 0, 119, 120, 3, 8, 4, 0, 120, 121, 5, 3, 0, 0, 121,
		7, 1, 0, 0, 0, 122, 123, 3, 10, 5, 0, 123, 124, 3, 8, 4, 0, 124, 127, 1,
		0, 0, 0, 125, 127, 1, 0, 0, 0, 126, 122, 1, 0, 0, 0, 126, 125, 1, 0, 0,
		0, 127, 9, 1, 0, 0, 0, 128, 129, 3, 4, 2, 0, 129, 130, 5, 10, 0, 0, 130,
		131, 3, 4, 2, 0, 131, 132, 5, 11, 0, 0, 132, 11, 1, 0, 0, 0, 133, 134,
		5, 7, 0, 0, 134, 135, 3, 4, 2, 0, 135, 136, 5, 8, 0, 0, 136, 139, 3, 4,
		2, 0, 137, 138, 5, 9, 0, 0, 138, 140, 3, 4, 2, 0, 139, 137, 1, 0, 0, 0,
		139, 140, 1, 0, 0, 0, 140, 13, 1, 0, 0, 0, 141, 142, 5, 42, 0, 0, 142,
		15, 1, 0, 0, 0, 143, 144, 5, 41, 0, 0, 144, 145, 5, 26, 0, 0, 145, 146,
		5, 42, 0, 0, 146, 17, 1, 0, 0, 0, 147, 148, 5, 43, 0, 0, 148, 150, 5, 31,
		0, 0, 149, 151, 3, 20, 10, 0, 150, 149, 1, 0, 0, 0, 150, 151, 1, 0, 0,
		0, 151, 152, 1, 0, 0, 0, 152, 153, 5, 32, 0, 0, 153, 19, 1, 0, 0, 0, 154,
		161, 3, 4, 2, 0, 155, 157, 5, 36, 0, 0, 156, 158, 3, 4, 2, 0, 157, 156,
		1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 160, 1, 0, 0, 0, 159, 155, 1, 0,
		0, 0, 160, 163, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0,
		162, 21, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 164, 166, 5, 41, 0, 0, 165,
		167, 3, 24, 12, 0, 166, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 166,
		1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 23, 1, 0, 0, 0, 170, 171, 5, 37,
		0, 0, 171, 172, 3, 26, 13, 0, 172, 25, 1, 0, 0, 0, 173, 176, 3, 28, 14,
		0, 174, 176, 3, 30, 15, 0, 175, 173, 1, 0, 0, 0, 175, 174, 1, 0, 0, 0,
		176, 27, 1, 0, 0, 0, 177, 178, 5, 42, 0, 0, 178, 29, 1, 0, 0, 0, 179, 180,
		5, 42, 0, 0, 180, 181, 5, 31, 0, 0, 181, 182, 3, 20, 10, 0, 182, 183, 5,
		32, 0, 0, 183, 31, 1, 0, 0, 0, 184, 185, 7, 0, 0, 0, 185, 33, 1, 0, 0,
		0, 12, 38, 42, 70, 111, 113, 126, 139, 150, 157, 161, 168, 175,
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

// HilcoPencilGrammarParserInit initializes any static state used to implement HilcoPencilGrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewHilcoPencilGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func HilcoPencilGrammarParserInit() {
	staticData := &HilcoPencilGrammarParserParserStaticData
	staticData.once.Do(hilcopencilgrammarparserParserInit)
}

// NewHilcoPencilGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewHilcoPencilGrammarParser(input antlr.TokenStream) *HilcoPencilGrammarParser {
	HilcoPencilGrammarParserInit()
	this := new(HilcoPencilGrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &HilcoPencilGrammarParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "HilcoPencilGrammarParser.g4"

	return this
}

// HilcoPencilGrammarParser tokens.
const (
	HilcoPencilGrammarParserEOF                = antlr.TokenEOF
	HilcoPencilGrammarParserASSIGNMENT         = 1
	HilcoPencilGrammarParserCASE               = 2
	HilcoPencilGrammarParserEND_CASE           = 3
	HilcoPencilGrammarParserIS                 = 4
	HilcoPencilGrammarParserSWITCH             = 5
	HilcoPencilGrammarParserEND_SWITCH         = 6
	HilcoPencilGrammarParserIF                 = 7
	HilcoPencilGrammarParserTHEN               = 8
	HilcoPencilGrammarParserELSE               = 9
	HilcoPencilGrammarParserDOUBLE_COLON       = 10
	HilcoPencilGrammarParserSEMI_COLON         = 11
	HilcoPencilGrammarParserNOT                = 12
	HilcoPencilGrammarParserEXPONENTIAL        = 13
	HilcoPencilGrammarParserMULTIPLY           = 14
	HilcoPencilGrammarParserDIVIDE             = 15
	HilcoPencilGrammarParserADD                = 16
	HilcoPencilGrammarParserMINUS              = 17
	HilcoPencilGrammarParserGREATER_THAN       = 18
	HilcoPencilGrammarParserGREATER_THAN_EQUAL = 19
	HilcoPencilGrammarParserLESS_THAN          = 20
	HilcoPencilGrammarParserLESS_THAN_EQUAL    = 21
	HilcoPencilGrammarParserEQUAL              = 22
	HilcoPencilGrammarParserNOT_EQUAL          = 23
	HilcoPencilGrammarParserAND                = 24
	HilcoPencilGrammarParserOR                 = 25
	HilcoPencilGrammarParserCOLON              = 26
	HilcoPencilGrammarParserLBRACKET           = 27
	HilcoPencilGrammarParserRBRACKET           = 28
	HilcoPencilGrammarParserCURLYLBRACKET      = 29
	HilcoPencilGrammarParserCURLYRBRACKET      = 30
	HilcoPencilGrammarParserLPAREN             = 31
	HilcoPencilGrammarParserRPAREN             = 32
	HilcoPencilGrammarParserUNDERSCORE         = 33
	HilcoPencilGrammarParserPERCENT_SIGN       = 34
	HilcoPencilGrammarParserAT_SIGN            = 35
	HilcoPencilGrammarParserCOMMA              = 36
	HilcoPencilGrammarParserDOT                = 37
	HilcoPencilGrammarParserKEYWORD_TRUE       = 38
	HilcoPencilGrammarParserKEYWORD_FALSE      = 39
	HilcoPencilGrammarParserKEYWORD_NIL        = 40
	HilcoPencilGrammarParserCLASSNAME          = 41
	HilcoPencilGrammarParserID                 = 42
	HilcoPencilGrammarParserATFUNCTION         = 43
	HilcoPencilGrammarParserINT                = 44
	HilcoPencilGrammarParserSTRING             = 45
	HilcoPencilGrammarParserFLOAT              = 46
	HilcoPencilGrammarParserPERCENT            = 47
	HilcoPencilGrammarParserDATE               = 48
	HilcoPencilGrammarParserNEWLINE            = 49
	HilcoPencilGrammarParserWS                 = 50
)

// HilcoPencilGrammarParser rules.
const (
	HilcoPencilGrammarParserRULE_program               = 0
	HilcoPencilGrammarParserRULE_statement             = 1
	HilcoPencilGrammarParserRULE_expression            = 2
	HilcoPencilGrammarParserRULE_caseStatement         = 3
	HilcoPencilGrammarParserRULE_caseList              = 4
	HilcoPencilGrammarParserRULE_caseItem              = 5
	HilcoPencilGrammarParserRULE_ifStatement           = 6
	HilcoPencilGrammarParserRULE_name                  = 7
	HilcoPencilGrammarParserRULE_worksheetVariable     = 8
	HilcoPencilGrammarParserRULE_atFunction            = 9
	HilcoPencilGrammarParserRULE_argList               = 10
	HilcoPencilGrammarParserRULE_dataAccessor          = 11
	HilcoPencilGrammarParserRULE_accessorList          = 12
	HilcoPencilGrammarParserRULE_accessorObjectOrArray = 13
	HilcoPencilGrammarParserRULE_accessorObject        = 14
	HilcoPencilGrammarParserRULE_accessorArray         = 15
	HilcoPencilGrammarParserRULE_specialKeyword        = 16
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *HilcoPencilGrammarParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HilcoPencilGrammarParserRULE_program)
	var _la int

	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&562677223133316) != 0 {
			{
				p.SetState(35)
				p.Statement()
			}

			p.SetState(40)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	ASSIGNMENT() antlr.TerminalNode
	SEMI_COLON() antlr.TerminalNode

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *StatementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserASSIGNMENT, 0)
}

func (s *StatementContext) SEMI_COLON() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserSEMI_COLON, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *HilcoPencilGrammarParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HilcoPencilGrammarParserRULE_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.expression(0)
	}
	{
		p.SetState(45)
		p.Match(HilcoPencilGrammarParserASSIGNMENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(46)
		p.expression(0)
	}
	{
		p.SetState(47)
		p.Match(HilcoPencilGrammarParserSEMI_COLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NameCalculatorContext struct {
	ExpressionContext
}

func NewNameCalculatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NameCalculatorContext {
	var p = new(NameCalculatorContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NameCalculatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameCalculatorContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *NameCalculatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterNameCalculator(s)
	}
}

func (s *NameCalculatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitNameCalculator(s)
	}
}

type UnaryNotCalculatorContext struct {
	ExpressionContext
}

func NewUnaryNotCalculatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryNotCalculatorContext {
	var p = new(UnaryNotCalculatorContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryNotCalculatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryNotCalculatorContext) NOT() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserNOT, 0)
}

func (s *UnaryNotCalculatorContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryNotCalculatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterUnaryNotCalculator(s)
	}
}

func (s *UnaryNotCalculatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitUnaryNotCalculator(s)
	}
}

type PercentContext struct {
	ExpressionContext
}

func NewPercentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PercentContext {
	var p = new(PercentContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *PercentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PercentContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserPERCENT, 0)
}

func (s *PercentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterPercent(s)
	}
}

func (s *PercentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitPercent(s)
	}
}

type ParensContext struct {
	ExpressionContext
}

func NewParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensContext {
	var p = new(ParensContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserLPAREN, 0)
}

func (s *ParensContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParensContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserRPAREN, 0)
}

func (s *ParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterParens(s)
	}
}

func (s *ParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitParens(s)
	}
}

type BinaryRelationalCalculatorContext struct {
	ExpressionContext
}

func NewBinaryRelationalCalculatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryRelationalCalculatorContext {
	var p = new(BinaryRelationalCalculatorContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryRelationalCalculatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryRelationalCalculatorContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BinaryRelationalCalculatorContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinaryRelationalCalculatorContext) GREATER_THAN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserGREATER_THAN, 0)
}

func (s *BinaryRelationalCalculatorContext) GREATER_THAN_EQUAL() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserGREATER_THAN_EQUAL, 0)
}

func (s *BinaryRelationalCalculatorContext) LESS_THAN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserLESS_THAN, 0)
}

func (s *BinaryRelationalCalculatorContext) LESS_THAN_EQUAL() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserLESS_THAN_EQUAL, 0)
}

func (s *BinaryRelationalCalculatorContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserEQUAL, 0)
}

func (s *BinaryRelationalCalculatorContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserNOT_EQUAL, 0)
}

func (s *BinaryRelationalCalculatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterBinaryRelationalCalculator(s)
	}
}

func (s *BinaryRelationalCalculatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitBinaryRelationalCalculator(s)
	}
}

type WorkSheetVariableCalculatorContext struct {
	ExpressionContext
}

func NewWorkSheetVariableCalculatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WorkSheetVariableCalculatorContext {
	var p = new(WorkSheetVariableCalculatorContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *WorkSheetVariableCalculatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WorkSheetVariableCalculatorContext) WorksheetVariable() IWorksheetVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWorksheetVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWorksheetVariableContext)
}

func (s *WorkSheetVariableCalculatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterWorkSheetVariableCalculator(s)
	}
}

func (s *WorkSheetVariableCalculatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitWorkSheetVariableCalculator(s)
	}
}

type StringContext struct {
	ExpressionContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitString(s)
	}
}

type ExpressionAtFunctionContext struct {
	ExpressionContext
}

func NewExpressionAtFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionAtFunctionContext {
	var p = new(ExpressionAtFunctionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionAtFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionAtFunctionContext) AtFunction() IAtFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtFunctionContext)
}

func (s *ExpressionAtFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterExpressionAtFunction(s)
	}
}

func (s *ExpressionAtFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitExpressionAtFunction(s)
	}
}

type DateContext struct {
	ExpressionContext
}

func NewDateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateContext {
	var p = new(DateContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *DateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateContext) DATE() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserDATE, 0)
}

func (s *DateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterDate(s)
	}
}

func (s *DateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitDate(s)
	}
}

type CaseContext struct {
	ExpressionContext
}

func NewCaseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CaseContext {
	var p = new(CaseContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseContext) CaseStatement() ICaseStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseStatementContext)
}

func (s *CaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterCase(s)
	}
}

func (s *CaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitCase(s)
	}
}

type IntegerContext struct {
	ExpressionContext
}

func NewIntegerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerContext {
	var p = new(IntegerContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) INT() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserINT, 0)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitInteger(s)
	}
}

type BinaryArithmeticCalculatorContext struct {
	ExpressionContext
}

func NewBinaryArithmeticCalculatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryArithmeticCalculatorContext {
	var p = new(BinaryArithmeticCalculatorContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryArithmeticCalculatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryArithmeticCalculatorContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BinaryArithmeticCalculatorContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinaryArithmeticCalculatorContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserMULTIPLY, 0)
}

func (s *BinaryArithmeticCalculatorContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserDIVIDE, 0)
}

func (s *BinaryArithmeticCalculatorContext) ADD() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserADD, 0)
}

func (s *BinaryArithmeticCalculatorContext) MINUS() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserMINUS, 0)
}

func (s *BinaryArithmeticCalculatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterBinaryArithmeticCalculator(s)
	}
}

func (s *BinaryArithmeticCalculatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitBinaryArithmeticCalculator(s)
	}
}

type BinaryLogicalCalculatorContext struct {
	ExpressionContext
}

func NewBinaryLogicalCalculatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryLogicalCalculatorContext {
	var p = new(BinaryLogicalCalculatorContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryLogicalCalculatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryLogicalCalculatorContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BinaryLogicalCalculatorContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinaryLogicalCalculatorContext) AND() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserAND, 0)
}

func (s *BinaryLogicalCalculatorContext) OR() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserOR, 0)
}

func (s *BinaryLogicalCalculatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterBinaryLogicalCalculator(s)
	}
}

func (s *BinaryLogicalCalculatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitBinaryLogicalCalculator(s)
	}
}

type ExpressionKeywordContext struct {
	ExpressionContext
}

func NewExpressionKeywordContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionKeywordContext {
	var p = new(ExpressionKeywordContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionKeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionKeywordContext) SpecialKeyword() ISpecialKeywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpecialKeywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpecialKeywordContext)
}

func (s *ExpressionKeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterExpressionKeyword(s)
	}
}

func (s *ExpressionKeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitExpressionKeyword(s)
	}
}

type FloatContext struct {
	ExpressionContext
}

func NewFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatContext {
	var p = new(FloatContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserFLOAT, 0)
}

func (s *FloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterFloat(s)
	}
}

func (s *FloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitFloat(s)
	}
}

type UnaryMinusCalculatorContext struct {
	ExpressionContext
}

func NewUnaryMinusCalculatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryMinusCalculatorContext {
	var p = new(UnaryMinusCalculatorContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryMinusCalculatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryMinusCalculatorContext) MINUS() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserMINUS, 0)
}

func (s *UnaryMinusCalculatorContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryMinusCalculatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterUnaryMinusCalculator(s)
	}
}

func (s *UnaryMinusCalculatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitUnaryMinusCalculator(s)
	}
}

type BinaryExponentialCalculatorContext struct {
	ExpressionContext
}

func NewBinaryExponentialCalculatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryExponentialCalculatorContext {
	var p = new(BinaryExponentialCalculatorContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryExponentialCalculatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExponentialCalculatorContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BinaryExponentialCalculatorContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinaryExponentialCalculatorContext) EXPONENTIAL() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserEXPONENTIAL, 0)
}

func (s *BinaryExponentialCalculatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterBinaryExponentialCalculator(s)
	}
}

func (s *BinaryExponentialCalculatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitBinaryExponentialCalculator(s)
	}
}

type ExpressionDataAccessContext struct {
	ExpressionContext
}

func NewExpressionDataAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionDataAccessContext {
	var p = new(ExpressionDataAccessContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionDataAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionDataAccessContext) DataAccessor() IDataAccessorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataAccessorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataAccessorContext)
}

func (s *ExpressionDataAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterExpressionDataAccess(s)
	}
}

func (s *ExpressionDataAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitExpressionDataAccess(s)
	}
}

type IfContext struct {
	ExpressionContext
}

func NewIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfContext {
	var p = new(IfContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *IfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterIf(s)
	}
}

func (s *IfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitIf(s)
	}
}

func (p *HilcoPencilGrammarParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *HilcoPencilGrammarParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, HilcoPencilGrammarParserRULE_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCaseContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(50)
			p.CaseStatement()
		}

	case 2:
		localctx = NewIfContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(51)
			p.IfStatement()
		}

	case 3:
		localctx = NewParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(52)
			p.Match(HilcoPencilGrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.expression(0)
		}
		{
			p.SetState(54)
			p.Match(HilcoPencilGrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewUnaryMinusCalculatorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(56)
			p.Match(HilcoPencilGrammarParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)
			p.expression(25)
		}

	case 5:
		localctx = NewUnaryNotCalculatorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(58)
			p.Match(HilcoPencilGrammarParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)
			p.expression(24)
		}

	case 6:
		localctx = NewNameCalculatorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(60)
			p.Name()
		}

	case 7:
		localctx = NewWorkSheetVariableCalculatorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(61)
			p.WorksheetVariable()
		}

	case 8:
		localctx = NewExpressionAtFunctionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(62)
			p.AtFunction()
		}

	case 9:
		localctx = NewExpressionDataAccessContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(63)
			p.DataAccessor()
		}

	case 10:
		localctx = NewExpressionKeywordContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(64)
			p.SpecialKeyword()
		}

	case 11:
		localctx = NewPercentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(65)
			p.Match(HilcoPencilGrammarParserPERCENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		localctx = NewIntegerContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(66)
			p.Match(HilcoPencilGrammarParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		localctx = NewFloatContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(67)
			p.Match(HilcoPencilGrammarParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		localctx = NewStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(68)
			p.Match(HilcoPencilGrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 15:
		localctx = NewDateContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(69)
			p.Match(HilcoPencilGrammarParserDATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(111)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryExponentialCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(72)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
					goto errorExit
				}
				{
					p.SetState(73)
					p.Match(HilcoPencilGrammarParserEXPONENTIAL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(74)
					p.expression(23)
				}

			case 2:
				localctx = NewBinaryArithmeticCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(75)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
					goto errorExit
				}
				{
					p.SetState(76)
					p.Match(HilcoPencilGrammarParserMULTIPLY)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(77)
					p.expression(23)
				}

			case 3:
				localctx = NewBinaryArithmeticCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(78)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
				{
					p.SetState(79)
					p.Match(HilcoPencilGrammarParserDIVIDE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(80)
					p.expression(22)
				}

			case 4:
				localctx = NewBinaryArithmeticCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(81)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(82)
					p.Match(HilcoPencilGrammarParserADD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(83)
					p.expression(21)
				}

			case 5:
				localctx = NewBinaryArithmeticCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(84)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(85)
					p.Match(HilcoPencilGrammarParserMINUS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(86)
					p.expression(20)
				}

			case 6:
				localctx = NewBinaryRelationalCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(87)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(88)
					p.Match(HilcoPencilGrammarParserGREATER_THAN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(89)
					p.expression(19)
				}

			case 7:
				localctx = NewBinaryRelationalCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(90)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(91)
					p.Match(HilcoPencilGrammarParserGREATER_THAN_EQUAL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(92)
					p.expression(18)
				}

			case 8:
				localctx = NewBinaryRelationalCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(93)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(94)
					p.Match(HilcoPencilGrammarParserLESS_THAN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(95)
					p.expression(17)
				}

			case 9:
				localctx = NewBinaryRelationalCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(96)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(97)
					p.Match(HilcoPencilGrammarParserLESS_THAN_EQUAL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(98)
					p.expression(16)
				}

			case 10:
				localctx = NewBinaryRelationalCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(99)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(100)
					p.Match(HilcoPencilGrammarParserEQUAL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(101)
					p.expression(15)
				}

			case 11:
				localctx = NewBinaryRelationalCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(102)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(103)
					p.Match(HilcoPencilGrammarParserNOT_EQUAL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(104)
					p.expression(14)
				}

			case 12:
				localctx = NewBinaryLogicalCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(105)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(106)
					p.Match(HilcoPencilGrammarParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(107)
					p.expression(13)
				}

			case 13:
				localctx = NewBinaryLogicalCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(108)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(109)
					p.Match(HilcoPencilGrammarParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(110)
					p.expression(12)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICaseStatementContext is an interface to support dynamic dispatch.
type ICaseStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	Expression() IExpressionContext
	IS() antlr.TerminalNode
	CaseList() ICaseListContext
	END_CASE() antlr.TerminalNode

	// IsCaseStatementContext differentiates from other interfaces.
	IsCaseStatementContext()
}

type CaseStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseStatementContext() *CaseStatementContext {
	var p = new(CaseStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_caseStatement
	return p
}

func InitEmptyCaseStatementContext(p *CaseStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_caseStatement
}

func (*CaseStatementContext) IsCaseStatementContext() {}

func NewCaseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseStatementContext {
	var p = new(CaseStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_caseStatement

	return p
}

func (s *CaseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseStatementContext) CASE() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserCASE, 0)
}

func (s *CaseStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CaseStatementContext) IS() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserIS, 0)
}

func (s *CaseStatementContext) CaseList() ICaseListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseListContext)
}

func (s *CaseStatementContext) END_CASE() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserEND_CASE, 0)
}

func (s *CaseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterCaseStatement(s)
	}
}

func (s *CaseStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitCaseStatement(s)
	}
}

func (p *HilcoPencilGrammarParser) CaseStatement() (localctx ICaseStatementContext) {
	localctx = NewCaseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, HilcoPencilGrammarParserRULE_caseStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(HilcoPencilGrammarParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.expression(0)
	}
	{
		p.SetState(118)
		p.Match(HilcoPencilGrammarParserIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.CaseList()
	}
	{
		p.SetState(120)
		p.Match(HilcoPencilGrammarParserEND_CASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICaseListContext is an interface to support dynamic dispatch.
type ICaseListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CaseItem() ICaseItemContext
	CaseList() ICaseListContext

	// IsCaseListContext differentiates from other interfaces.
	IsCaseListContext()
}

type CaseListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseListContext() *CaseListContext {
	var p = new(CaseListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_caseList
	return p
}

func InitEmptyCaseListContext(p *CaseListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_caseList
}

func (*CaseListContext) IsCaseListContext() {}

func NewCaseListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseListContext {
	var p = new(CaseListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_caseList

	return p
}

func (s *CaseListContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseListContext) CaseItem() ICaseItemContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseItemContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseItemContext)
}

func (s *CaseListContext) CaseList() ICaseListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseListContext)
}

func (s *CaseListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterCaseList(s)
	}
}

func (s *CaseListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitCaseList(s)
	}
}

func (p *HilcoPencilGrammarParser) CaseList() (localctx ICaseListContext) {
	localctx = NewCaseListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, HilcoPencilGrammarParserRULE_caseList)
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case HilcoPencilGrammarParserCASE, HilcoPencilGrammarParserIF, HilcoPencilGrammarParserNOT, HilcoPencilGrammarParserMINUS, HilcoPencilGrammarParserLPAREN, HilcoPencilGrammarParserKEYWORD_TRUE, HilcoPencilGrammarParserKEYWORD_FALSE, HilcoPencilGrammarParserKEYWORD_NIL, HilcoPencilGrammarParserCLASSNAME, HilcoPencilGrammarParserID, HilcoPencilGrammarParserATFUNCTION, HilcoPencilGrammarParserINT, HilcoPencilGrammarParserSTRING, HilcoPencilGrammarParserFLOAT, HilcoPencilGrammarParserPERCENT, HilcoPencilGrammarParserDATE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.CaseItem()
		}
		{
			p.SetState(123)
			p.CaseList()
		}

	case HilcoPencilGrammarParserEND_CASE:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICaseItemContext is an interface to support dynamic dispatch.
type ICaseItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	DOUBLE_COLON() antlr.TerminalNode
	SEMI_COLON() antlr.TerminalNode

	// IsCaseItemContext differentiates from other interfaces.
	IsCaseItemContext()
}

type CaseItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseItemContext() *CaseItemContext {
	var p = new(CaseItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_caseItem
	return p
}

func InitEmptyCaseItemContext(p *CaseItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_caseItem
}

func (*CaseItemContext) IsCaseItemContext() {}

func NewCaseItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseItemContext {
	var p = new(CaseItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_caseItem

	return p
}

func (s *CaseItemContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseItemContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *CaseItemContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CaseItemContext) DOUBLE_COLON() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserDOUBLE_COLON, 0)
}

func (s *CaseItemContext) SEMI_COLON() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserSEMI_COLON, 0)
}

func (s *CaseItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterCaseItem(s)
	}
}

func (s *CaseItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitCaseItem(s)
	}
}

func (p *HilcoPencilGrammarParser) CaseItem() (localctx ICaseItemContext) {
	localctx = NewCaseItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, HilcoPencilGrammarParserRULE_caseItem)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.expression(0)
	}
	{
		p.SetState(129)
		p.Match(HilcoPencilGrammarParserDOUBLE_COLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(130)
		p.expression(0)
	}
	{
		p.SetState(131)
		p.Match(HilcoPencilGrammarParserSEMI_COLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	THEN() antlr.TerminalNode
	ELSE() antlr.TerminalNode

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserIF, 0)
}

func (s *IfStatementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) THEN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserTHEN, 0)
}

func (s *IfStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserELSE, 0)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (p *HilcoPencilGrammarParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, HilcoPencilGrammarParserRULE_ifStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Match(HilcoPencilGrammarParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.expression(0)
	}
	{
		p.SetState(135)
		p.Match(HilcoPencilGrammarParserTHEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(136)
		p.expression(0)
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(137)
			p.Match(HilcoPencilGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.expression(0)
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_name
	return p
}

func InitEmptyNameContext(p *NameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_name
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) ID() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserID, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *HilcoPencilGrammarParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, HilcoPencilGrammarParserRULE_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(HilcoPencilGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWorksheetVariableContext is an interface to support dynamic dispatch.
type IWorksheetVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CLASSNAME() antlr.TerminalNode
	COLON() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsWorksheetVariableContext differentiates from other interfaces.
	IsWorksheetVariableContext()
}

type WorksheetVariableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWorksheetVariableContext() *WorksheetVariableContext {
	var p = new(WorksheetVariableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_worksheetVariable
	return p
}

func InitEmptyWorksheetVariableContext(p *WorksheetVariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_worksheetVariable
}

func (*WorksheetVariableContext) IsWorksheetVariableContext() {}

func NewWorksheetVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WorksheetVariableContext {
	var p = new(WorksheetVariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_worksheetVariable

	return p
}

func (s *WorksheetVariableContext) GetParser() antlr.Parser { return s.parser }

func (s *WorksheetVariableContext) CLASSNAME() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserCLASSNAME, 0)
}

func (s *WorksheetVariableContext) COLON() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserCOLON, 0)
}

func (s *WorksheetVariableContext) ID() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserID, 0)
}

func (s *WorksheetVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WorksheetVariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WorksheetVariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterWorksheetVariable(s)
	}
}

func (s *WorksheetVariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitWorksheetVariable(s)
	}
}

func (p *HilcoPencilGrammarParser) WorksheetVariable() (localctx IWorksheetVariableContext) {
	localctx = NewWorksheetVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, HilcoPencilGrammarParserRULE_worksheetVariable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(HilcoPencilGrammarParserCLASSNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Match(HilcoPencilGrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Match(HilcoPencilGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAtFunctionContext is an interface to support dynamic dispatch.
type IAtFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ATFUNCTION() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ArgList() IArgListContext

	// IsAtFunctionContext differentiates from other interfaces.
	IsAtFunctionContext()
}

type AtFunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtFunctionContext() *AtFunctionContext {
	var p = new(AtFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_atFunction
	return p
}

func InitEmptyAtFunctionContext(p *AtFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_atFunction
}

func (*AtFunctionContext) IsAtFunctionContext() {}

func NewAtFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtFunctionContext {
	var p = new(AtFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_atFunction

	return p
}

func (s *AtFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *AtFunctionContext) ATFUNCTION() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserATFUNCTION, 0)
}

func (s *AtFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserLPAREN, 0)
}

func (s *AtFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserRPAREN, 0)
}

func (s *AtFunctionContext) ArgList() IArgListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *AtFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterAtFunction(s)
	}
}

func (s *AtFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitAtFunction(s)
	}
}

func (p *HilcoPencilGrammarParser) AtFunction() (localctx IAtFunctionContext) {
	localctx = NewAtFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, HilcoPencilGrammarParserRULE_atFunction)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Match(HilcoPencilGrammarParserATFUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.Match(HilcoPencilGrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&562677223133316) != 0 {
		{
			p.SetState(149)
			p.ArgList()
		}

	}
	{
		p.SetState(152)
		p.Match(HilcoPencilGrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgListContext is an interface to support dynamic dispatch.
type IArgListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgListContext differentiates from other interfaces.
	IsArgListContext()
}

type ArgListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgListContext() *ArgListContext {
	var p = new(ArgListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_argList
	return p
}

func InitEmptyArgListContext(p *ArgListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_argList
}

func (*ArgListContext) IsArgListContext() {}

func NewArgListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgListContext {
	var p = new(ArgListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_argList

	return p
}

func (s *ArgListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(HilcoPencilGrammarParserCOMMA)
}

func (s *ArgListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserCOMMA, i)
}

func (s *ArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterArgList(s)
	}
}

func (s *ArgListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitArgList(s)
	}
}

func (p *HilcoPencilGrammarParser) ArgList() (localctx IArgListContext) {
	localctx = NewArgListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, HilcoPencilGrammarParserRULE_argList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.expression(0)
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == HilcoPencilGrammarParserCOMMA {
		{
			p.SetState(155)
			p.Match(HilcoPencilGrammarParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&562677223133316) != 0 {
			{
				p.SetState(156)
				p.expression(0)
			}

		}

		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDataAccessorContext is an interface to support dynamic dispatch.
type IDataAccessorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CLASSNAME() antlr.TerminalNode
	AllAccessorList() []IAccessorListContext
	AccessorList(i int) IAccessorListContext

	// IsDataAccessorContext differentiates from other interfaces.
	IsDataAccessorContext()
}

type DataAccessorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataAccessorContext() *DataAccessorContext {
	var p = new(DataAccessorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_dataAccessor
	return p
}

func InitEmptyDataAccessorContext(p *DataAccessorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_dataAccessor
}

func (*DataAccessorContext) IsDataAccessorContext() {}

func NewDataAccessorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataAccessorContext {
	var p = new(DataAccessorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_dataAccessor

	return p
}

func (s *DataAccessorContext) GetParser() antlr.Parser { return s.parser }

func (s *DataAccessorContext) CLASSNAME() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserCLASSNAME, 0)
}

func (s *DataAccessorContext) AllAccessorList() []IAccessorListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAccessorListContext); ok {
			len++
		}
	}

	tst := make([]IAccessorListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAccessorListContext); ok {
			tst[i] = t.(IAccessorListContext)
			i++
		}
	}

	return tst
}

func (s *DataAccessorContext) AccessorList(i int) IAccessorListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessorListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccessorListContext)
}

func (s *DataAccessorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataAccessorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataAccessorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterDataAccessor(s)
	}
}

func (s *DataAccessorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitDataAccessor(s)
	}
}

func (p *HilcoPencilGrammarParser) DataAccessor() (localctx IDataAccessorContext) {
	localctx = NewDataAccessorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, HilcoPencilGrammarParserRULE_dataAccessor)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.Match(HilcoPencilGrammarParserCLASSNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(165)
				p.AccessorList()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAccessorListContext is an interface to support dynamic dispatch.
type IAccessorListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOT() antlr.TerminalNode
	AccessorObjectOrArray() IAccessorObjectOrArrayContext

	// IsAccessorListContext differentiates from other interfaces.
	IsAccessorListContext()
}

type AccessorListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessorListContext() *AccessorListContext {
	var p = new(AccessorListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_accessorList
	return p
}

func InitEmptyAccessorListContext(p *AccessorListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_accessorList
}

func (*AccessorListContext) IsAccessorListContext() {}

func NewAccessorListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessorListContext {
	var p = new(AccessorListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_accessorList

	return p
}

func (s *AccessorListContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessorListContext) DOT() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserDOT, 0)
}

func (s *AccessorListContext) AccessorObjectOrArray() IAccessorObjectOrArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessorObjectOrArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccessorObjectOrArrayContext)
}

func (s *AccessorListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessorListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessorListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterAccessorList(s)
	}
}

func (s *AccessorListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitAccessorList(s)
	}
}

func (p *HilcoPencilGrammarParser) AccessorList() (localctx IAccessorListContext) {
	localctx = NewAccessorListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, HilcoPencilGrammarParserRULE_accessorList)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(HilcoPencilGrammarParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(171)
		p.AccessorObjectOrArray()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAccessorObjectOrArrayContext is an interface to support dynamic dispatch.
type IAccessorObjectOrArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AccessorObject() IAccessorObjectContext
	AccessorArray() IAccessorArrayContext

	// IsAccessorObjectOrArrayContext differentiates from other interfaces.
	IsAccessorObjectOrArrayContext()
}

type AccessorObjectOrArrayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessorObjectOrArrayContext() *AccessorObjectOrArrayContext {
	var p = new(AccessorObjectOrArrayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_accessorObjectOrArray
	return p
}

func InitEmptyAccessorObjectOrArrayContext(p *AccessorObjectOrArrayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_accessorObjectOrArray
}

func (*AccessorObjectOrArrayContext) IsAccessorObjectOrArrayContext() {}

func NewAccessorObjectOrArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessorObjectOrArrayContext {
	var p = new(AccessorObjectOrArrayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_accessorObjectOrArray

	return p
}

func (s *AccessorObjectOrArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessorObjectOrArrayContext) AccessorObject() IAccessorObjectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessorObjectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccessorObjectContext)
}

func (s *AccessorObjectOrArrayContext) AccessorArray() IAccessorArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessorArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccessorArrayContext)
}

func (s *AccessorObjectOrArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessorObjectOrArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessorObjectOrArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterAccessorObjectOrArray(s)
	}
}

func (s *AccessorObjectOrArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitAccessorObjectOrArray(s)
	}
}

func (p *HilcoPencilGrammarParser) AccessorObjectOrArray() (localctx IAccessorObjectOrArrayContext) {
	localctx = NewAccessorObjectOrArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, HilcoPencilGrammarParserRULE_accessorObjectOrArray)
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(173)
			p.AccessorObject()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(174)
			p.AccessorArray()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAccessorObjectContext is an interface to support dynamic dispatch.
type IAccessorObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsAccessorObjectContext differentiates from other interfaces.
	IsAccessorObjectContext()
}

type AccessorObjectContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessorObjectContext() *AccessorObjectContext {
	var p = new(AccessorObjectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_accessorObject
	return p
}

func InitEmptyAccessorObjectContext(p *AccessorObjectContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_accessorObject
}

func (*AccessorObjectContext) IsAccessorObjectContext() {}

func NewAccessorObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessorObjectContext {
	var p = new(AccessorObjectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_accessorObject

	return p
}

func (s *AccessorObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessorObjectContext) ID() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserID, 0)
}

func (s *AccessorObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessorObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessorObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterAccessorObject(s)
	}
}

func (s *AccessorObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitAccessorObject(s)
	}
}

func (p *HilcoPencilGrammarParser) AccessorObject() (localctx IAccessorObjectContext) {
	localctx = NewAccessorObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, HilcoPencilGrammarParserRULE_accessorObject)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.Match(HilcoPencilGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAccessorArrayContext is an interface to support dynamic dispatch.
type IAccessorArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	ArgList() IArgListContext
	RPAREN() antlr.TerminalNode

	// IsAccessorArrayContext differentiates from other interfaces.
	IsAccessorArrayContext()
}

type AccessorArrayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessorArrayContext() *AccessorArrayContext {
	var p = new(AccessorArrayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_accessorArray
	return p
}

func InitEmptyAccessorArrayContext(p *AccessorArrayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_accessorArray
}

func (*AccessorArrayContext) IsAccessorArrayContext() {}

func NewAccessorArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessorArrayContext {
	var p = new(AccessorArrayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_accessorArray

	return p
}

func (s *AccessorArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessorArrayContext) ID() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserID, 0)
}

func (s *AccessorArrayContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserLPAREN, 0)
}

func (s *AccessorArrayContext) ArgList() IArgListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *AccessorArrayContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserRPAREN, 0)
}

func (s *AccessorArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessorArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessorArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterAccessorArray(s)
	}
}

func (s *AccessorArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitAccessorArray(s)
	}
}

func (p *HilcoPencilGrammarParser) AccessorArray() (localctx IAccessorArrayContext) {
	localctx = NewAccessorArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, HilcoPencilGrammarParserRULE_accessorArray)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Match(HilcoPencilGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)
		p.Match(HilcoPencilGrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
		p.ArgList()
	}
	{
		p.SetState(182)
		p.Match(HilcoPencilGrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISpecialKeywordContext is an interface to support dynamic dispatch.
type ISpecialKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KEYWORD_TRUE() antlr.TerminalNode
	KEYWORD_FALSE() antlr.TerminalNode
	KEYWORD_NIL() antlr.TerminalNode

	// IsSpecialKeywordContext differentiates from other interfaces.
	IsSpecialKeywordContext()
}

type SpecialKeywordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecialKeywordContext() *SpecialKeywordContext {
	var p = new(SpecialKeywordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_specialKeyword
	return p
}

func InitEmptySpecialKeywordContext(p *SpecialKeywordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_specialKeyword
}

func (*SpecialKeywordContext) IsSpecialKeywordContext() {}

func NewSpecialKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecialKeywordContext {
	var p = new(SpecialKeywordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_specialKeyword

	return p
}

func (s *SpecialKeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecialKeywordContext) KEYWORD_TRUE() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserKEYWORD_TRUE, 0)
}

func (s *SpecialKeywordContext) KEYWORD_FALSE() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserKEYWORD_FALSE, 0)
}

func (s *SpecialKeywordContext) KEYWORD_NIL() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserKEYWORD_NIL, 0)
}

func (s *SpecialKeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecialKeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecialKeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterSpecialKeyword(s)
	}
}

func (s *SpecialKeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitSpecialKeyword(s)
	}
}

func (p *HilcoPencilGrammarParser) SpecialKeyword() (localctx ISpecialKeywordContext) {
	localctx = NewSpecialKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, HilcoPencilGrammarParserRULE_specialKeyword)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1924145348608) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *HilcoPencilGrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *HilcoPencilGrammarParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 23)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 11)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
