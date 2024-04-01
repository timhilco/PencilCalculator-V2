// Generated from /Users/timothyhilgenberg/go/src/github.com/timhilco/go-PencilCalculator/HilcoPencilGrammarLexer.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class HilcoPencilGrammarLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		ASSIGNMENT=1, CASE=2, END_CASE=3, IS=4, SWITCH=5, END_SWITCH=6, IF=7, 
		THEN=8, ELSE=9, DOUBLE_COLON=10, SEMI_COLON=11, NOT=12, EXPONENTIAL=13, 
		MULTIPLY=14, DIVIDE=15, ADD=16, MINUS=17, GREATER_THAN=18, GREATER_THAN_EQUAL=19, 
		LESS_THAN=20, LESS_THAN_EQUAL=21, EQUAL=22, NOT_EQUAL=23, AND=24, OR=25, 
		COLON=26, LBRACKET=27, RBRACKET=28, CURLYLBRACKET=29, CURLYRBRACKET=30, 
		LPAREN=31, RPAREN=32, UNDERSCORE=33, PERCENT_SIGN=34, AT_SIGN=35, COMMA=36, 
		DOT=37, KEYWORD_TRUE=38, KEYWORD_FALSE=39, KEYWORD_NIL=40, CLASSNAME=41, 
		ID=42, ATFUNCTION=43, INT=44, STRING=45, FLOAT=46, PERCENT=47, DATE=48, 
		NEWLINE=49, WS=50;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"ASSIGNMENT", "CASE", "END_CASE", "IS", "SWITCH", "END_SWITCH", "IF", 
			"THEN", "ELSE", "DOUBLE_COLON", "SEMI_COLON", "NOT", "EXPONENTIAL", "MULTIPLY", 
			"DIVIDE", "ADD", "MINUS", "GREATER_THAN", "GREATER_THAN_EQUAL", "LESS_THAN", 
			"LESS_THAN_EQUAL", "EQUAL", "NOT_EQUAL", "AND", "OR", "COLON", "LBRACKET", 
			"RBRACKET", "CURLYLBRACKET", "CURLYRBRACKET", "LPAREN", "RPAREN", "UNDERSCORE", 
			"PERCENT_SIGN", "AT_SIGN", "COMMA", "DOT", "KEYWORD_TRUE", "KEYWORD_FALSE", 
			"KEYWORD_NIL", "DIGIT", "LOWERCASELETTERS", "UPPERCASELETTERS", "CLASSNAME", 
			"ID", "ATFUNCTION", "INT", "STRING", "FLOAT", "PERCENT", "DATE", "EXPONENT", 
			"HEX_DIGIT", "ESC_SEQ", "OCTAL_ESC", "UNICODE_ESC", "NEWLINE", "WS"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "':='", "'case'", "'endcase'", "'is'", "'switch'", "'endswitch'", 
			"'if'", "'then'", "'else'", "'::'", "';'", "'NOT'", "'^'", "'*'", "'/'", 
			"'+'", "'-'", "'>'", "'>='", "'<'", "'<='", "'='", "'~='", "'AND'", "'OR'", 
			"':'", "'['", "']'", "'{'", "'}'", "'('", "')'", "'_'", "'%'", "'@'", 
			"','", "'.'", "'true'", "'false'", "'nil'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "ASSIGNMENT", "CASE", "END_CASE", "IS", "SWITCH", "END_SWITCH", 
			"IF", "THEN", "ELSE", "DOUBLE_COLON", "SEMI_COLON", "NOT", "EXPONENTIAL", 
			"MULTIPLY", "DIVIDE", "ADD", "MINUS", "GREATER_THAN", "GREATER_THAN_EQUAL", 
			"LESS_THAN", "LESS_THAN_EQUAL", "EQUAL", "NOT_EQUAL", "AND", "OR", "COLON", 
			"LBRACKET", "RBRACKET", "CURLYLBRACKET", "CURLYRBRACKET", "LPAREN", "RPAREN", 
			"UNDERSCORE", "PERCENT_SIGN", "AT_SIGN", "COMMA", "DOT", "KEYWORD_TRUE", 
			"KEYWORD_FALSE", "KEYWORD_NIL", "CLASSNAME", "ID", "ATFUNCTION", "INT", 
			"STRING", "FLOAT", "PERCENT", "DATE", "NEWLINE", "WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public HilcoPencilGrammarLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "HilcoPencilGrammarLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\64\u0193\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64"+
		"\t\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\3\2\3\2\3"+
		"\2\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3"+
		"\b\3\b\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\f\3\f"+
		"\3\r\3\r\3\r\3\r\3\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\22\3\22\3\23"+
		"\3\23\3\24\3\24\3\24\3\25\3\25\3\26\3\26\3\26\3\27\3\27\3\30\3\30\3\30"+
		"\3\31\3\31\3\31\3\31\3\32\3\32\3\32\3\33\3\33\3\34\3\34\3\35\3\35\3\36"+
		"\3\36\3\37\3\37\3 \3 \3!\3!\3\"\3\"\3#\3#\3$\3$\3%\3%\3&\3&\3\'\3\'\3"+
		"\'\3\'\3\'\3(\3(\3(\3(\3(\3(\3)\3)\3)\3)\3*\3*\3+\3+\3,\3,\3-\3-\3-\3"+
		"-\3-\7-\u0104\n-\f-\16-\u0107\13-\3.\3.\3.\3.\3.\3.\3.\7.\u0110\n.\f."+
		"\16.\u0113\13.\3/\3/\3/\5/\u0118\n/\3/\3/\3/\7/\u011d\n/\f/\16/\u0120"+
		"\13/\3\60\6\60\u0123\n\60\r\60\16\60\u0124\3\61\3\61\3\61\7\61\u012a\n"+
		"\61\f\61\16\61\u012d\13\61\3\61\3\61\3\62\6\62\u0132\n\62\r\62\16\62\u0133"+
		"\3\62\3\62\7\62\u0138\n\62\f\62\16\62\u013b\13\62\3\63\3\63\5\63\u013f"+
		"\n\63\3\63\3\63\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\5\64\u014c"+
		"\n\64\3\64\3\64\5\64\u0150\n\64\3\64\3\64\5\64\u0154\n\64\5\64\u0156\n"+
		"\64\3\64\3\64\3\64\3\64\3\64\3\64\5\64\u015e\n\64\3\64\3\64\3\65\3\65"+
		"\5\65\u0164\n\65\3\65\6\65\u0167\n\65\r\65\16\65\u0168\3\66\3\66\3\67"+
		"\3\67\3\67\3\67\5\67\u0171\n\67\38\38\38\38\38\38\38\38\38\58\u017c\n"+
		"8\39\39\39\39\39\39\39\3:\5:\u0186\n:\3:\3:\3:\3:\3;\3;\3;\3;\5;\u0190"+
		"\n;\3;\3;\2\2<\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16"+
		"\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34"+
		"\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S\2U\2W\2Y+[,]-_.a/c\60e\61g\62"+
		"i\2k\2m\2o\2q\2s\63u\64\3\2\t\4\2))^^\4\2//\61\61\4\2GGgg\4\2--//\5\2"+
		"\62;CHch\n\2$$))^^ddhhppttvv\5\2\13\13\16\16\"\"\2\u01ac\2\3\3\2\2\2\2"+
		"\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2"+
		"\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2"+
		"\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2"+
		"\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2"+
		"\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2"+
		"\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2"+
		"K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3"+
		"\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2s\3\2\2"+
		"\2\2u\3\2\2\2\3w\3\2\2\2\5z\3\2\2\2\7\177\3\2\2\2\t\u0087\3\2\2\2\13\u008a"+
		"\3\2\2\2\r\u0091\3\2\2\2\17\u009b\3\2\2\2\21\u009e\3\2\2\2\23\u00a3\3"+
		"\2\2\2\25\u00a8\3\2\2\2\27\u00ab\3\2\2\2\31\u00ad\3\2\2\2\33\u00b1\3\2"+
		"\2\2\35\u00b3\3\2\2\2\37\u00b5\3\2\2\2!\u00b7\3\2\2\2#\u00b9\3\2\2\2%"+
		"\u00bb\3\2\2\2\'\u00bd\3\2\2\2)\u00c0\3\2\2\2+\u00c2\3\2\2\2-\u00c5\3"+
		"\2\2\2/\u00c7\3\2\2\2\61\u00ca\3\2\2\2\63\u00ce\3\2\2\2\65\u00d1\3\2\2"+
		"\2\67\u00d3\3\2\2\29\u00d5\3\2\2\2;\u00d7\3\2\2\2=\u00d9\3\2\2\2?\u00db"+
		"\3\2\2\2A\u00dd\3\2\2\2C\u00df\3\2\2\2E\u00e1\3\2\2\2G\u00e3\3\2\2\2I"+
		"\u00e5\3\2\2\2K\u00e7\3\2\2\2M\u00e9\3\2\2\2O\u00ee\3\2\2\2Q\u00f4\3\2"+
		"\2\2S\u00f8\3\2\2\2U\u00fa\3\2\2\2W\u00fc\3\2\2\2Y\u00fe\3\2\2\2[\u0108"+
		"\3\2\2\2]\u0114\3\2\2\2_\u0122\3\2\2\2a\u0126\3\2\2\2c\u0131\3\2\2\2e"+
		"\u013e\3\2\2\2g\u0142\3\2\2\2i\u0161\3\2\2\2k\u016a\3\2\2\2m\u0170\3\2"+
		"\2\2o\u017b\3\2\2\2q\u017d\3\2\2\2s\u0185\3\2\2\2u\u018f\3\2\2\2wx\7<"+
		"\2\2xy\7?\2\2y\4\3\2\2\2z{\7e\2\2{|\7c\2\2|}\7u\2\2}~\7g\2\2~\6\3\2\2"+
		"\2\177\u0080\7g\2\2\u0080\u0081\7p\2\2\u0081\u0082\7f\2\2\u0082\u0083"+
		"\7e\2\2\u0083\u0084\7c\2\2\u0084\u0085\7u\2\2\u0085\u0086\7g\2\2\u0086"+
		"\b\3\2\2\2\u0087\u0088\7k\2\2\u0088\u0089\7u\2\2\u0089\n\3\2\2\2\u008a"+
		"\u008b\7u\2\2\u008b\u008c\7y\2\2\u008c\u008d\7k\2\2\u008d\u008e\7v\2\2"+
		"\u008e\u008f\7e\2\2\u008f\u0090\7j\2\2\u0090\f\3\2\2\2\u0091\u0092\7g"+
		"\2\2\u0092\u0093\7p\2\2\u0093\u0094\7f\2\2\u0094\u0095\7u\2\2\u0095\u0096"+
		"\7y\2\2\u0096\u0097\7k\2\2\u0097\u0098\7v\2\2\u0098\u0099\7e\2\2\u0099"+
		"\u009a\7j\2\2\u009a\16\3\2\2\2\u009b\u009c\7k\2\2\u009c\u009d\7h\2\2\u009d"+
		"\20\3\2\2\2\u009e\u009f\7v\2\2\u009f\u00a0\7j\2\2\u00a0\u00a1\7g\2\2\u00a1"+
		"\u00a2\7p\2\2\u00a2\22\3\2\2\2\u00a3\u00a4\7g\2\2\u00a4\u00a5\7n\2\2\u00a5"+
		"\u00a6\7u\2\2\u00a6\u00a7\7g\2\2\u00a7\24\3\2\2\2\u00a8\u00a9\7<\2\2\u00a9"+
		"\u00aa\7<\2\2\u00aa\26\3\2\2\2\u00ab\u00ac\7=\2\2\u00ac\30\3\2\2\2\u00ad"+
		"\u00ae\7P\2\2\u00ae\u00af\7Q\2\2\u00af\u00b0\7V\2\2\u00b0\32\3\2\2\2\u00b1"+
		"\u00b2\7`\2\2\u00b2\34\3\2\2\2\u00b3\u00b4\7,\2\2\u00b4\36\3\2\2\2\u00b5"+
		"\u00b6\7\61\2\2\u00b6 \3\2\2\2\u00b7\u00b8\7-\2\2\u00b8\"\3\2\2\2\u00b9"+
		"\u00ba\7/\2\2\u00ba$\3\2\2\2\u00bb\u00bc\7@\2\2\u00bc&\3\2\2\2\u00bd\u00be"+
		"\7@\2\2\u00be\u00bf\7?\2\2\u00bf(\3\2\2\2\u00c0\u00c1\7>\2\2\u00c1*\3"+
		"\2\2\2\u00c2\u00c3\7>\2\2\u00c3\u00c4\7?\2\2\u00c4,\3\2\2\2\u00c5\u00c6"+
		"\7?\2\2\u00c6.\3\2\2\2\u00c7\u00c8\7\u0080\2\2\u00c8\u00c9\7?\2\2\u00c9"+
		"\60\3\2\2\2\u00ca\u00cb\7C\2\2\u00cb\u00cc\7P\2\2\u00cc\u00cd\7F\2\2\u00cd"+
		"\62\3\2\2\2\u00ce\u00cf\7Q\2\2\u00cf\u00d0\7T\2\2\u00d0\64\3\2\2\2\u00d1"+
		"\u00d2\7<\2\2\u00d2\66\3\2\2\2\u00d3\u00d4\7]\2\2\u00d48\3\2\2\2\u00d5"+
		"\u00d6\7_\2\2\u00d6:\3\2\2\2\u00d7\u00d8\7}\2\2\u00d8<\3\2\2\2\u00d9\u00da"+
		"\7\177\2\2\u00da>\3\2\2\2\u00db\u00dc\7*\2\2\u00dc@\3\2\2\2\u00dd\u00de"+
		"\7+\2\2\u00deB\3\2\2\2\u00df\u00e0\7a\2\2\u00e0D\3\2\2\2\u00e1\u00e2\7"+
		"\'\2\2\u00e2F\3\2\2\2\u00e3\u00e4\7B\2\2\u00e4H\3\2\2\2\u00e5\u00e6\7"+
		".\2\2\u00e6J\3\2\2\2\u00e7\u00e8\7\60\2\2\u00e8L\3\2\2\2\u00e9\u00ea\7"+
		"v\2\2\u00ea\u00eb\7t\2\2\u00eb\u00ec\7w\2\2\u00ec\u00ed\7g\2\2\u00edN"+
		"\3\2\2\2\u00ee\u00ef\7h\2\2\u00ef\u00f0\7c\2\2\u00f0\u00f1\7n\2\2\u00f1"+
		"\u00f2\7u\2\2\u00f2\u00f3\7g\2\2\u00f3P\3\2\2\2\u00f4\u00f5\7p\2\2\u00f5"+
		"\u00f6\7k\2\2\u00f6\u00f7\7n\2\2\u00f7R\3\2\2\2\u00f8\u00f9\4\62;\2\u00f9"+
		"T\3\2\2\2\u00fa\u00fb\4c|\2\u00fbV\3\2\2\2\u00fc\u00fd\4C\\\2\u00fdX\3"+
		"\2\2\2\u00fe\u0105\5W,\2\u00ff\u0104\5U+\2\u0100\u0104\5W,\2\u0101\u0104"+
		"\5S*\2\u0102\u0104\5C\"\2\u0103\u00ff\3\2\2\2\u0103\u0100\3\2\2\2\u0103"+
		"\u0101\3\2\2\2\u0103\u0102\3\2\2\2\u0104\u0107\3\2\2\2\u0105\u0103\3\2"+
		"\2\2\u0105\u0106\3\2\2\2\u0106Z\3\2\2\2\u0107\u0105\3\2\2\2\u0108\u0111"+
		"\5U+\2\u0109\u0110\5U+\2\u010a\u0110\5W,\2\u010b\u0110\5S*\2\u010c\u0110"+
		"\5C\"\2\u010d\u0110\5\65\33\2\u010e\u0110\7^\2\2\u010f\u0109\3\2\2\2\u010f"+
		"\u010a\3\2\2\2\u010f\u010b\3\2\2\2\u010f\u010c\3\2\2\2\u010f\u010d\3\2"+
		"\2\2\u010f\u010e\3\2\2\2\u0110\u0113\3\2\2\2\u0111\u010f\3\2\2\2\u0111"+
		"\u0112\3\2\2\2\u0112\\\3\2\2\2\u0113\u0111\3\2\2\2\u0114\u0117\5G$\2\u0115"+
		"\u0118\5U+\2\u0116\u0118\5W,\2\u0117\u0115\3\2\2\2\u0117\u0116\3\2\2\2"+
		"\u0118\u011e\3\2\2\2\u0119\u011d\5U+\2\u011a\u011d\5W,\2\u011b\u011d\5"+
		"S*\2\u011c\u0119\3\2\2\2\u011c\u011a\3\2\2\2\u011c\u011b\3\2\2\2\u011d"+
		"\u0120\3\2\2\2\u011e\u011c\3\2\2\2\u011e\u011f\3\2\2\2\u011f^\3\2\2\2"+
		"\u0120\u011e\3\2\2\2\u0121\u0123\5S*\2\u0122\u0121\3\2\2\2\u0123\u0124"+
		"\3\2\2\2\u0124\u0122\3\2\2\2\u0124\u0125\3\2\2\2\u0125`\3\2\2\2\u0126"+
		"\u012b\7)\2\2\u0127\u012a\5m\67\2\u0128\u012a\n\2\2\2\u0129\u0127\3\2"+
		"\2\2\u0129\u0128\3\2\2\2\u012a\u012d\3\2\2\2\u012b\u0129\3\2\2\2\u012b"+
		"\u012c\3\2\2\2\u012c\u012e\3\2\2\2\u012d\u012b\3\2\2\2\u012e\u012f\7)"+
		"\2\2\u012fb\3\2\2\2\u0130\u0132\5S*\2\u0131\u0130\3\2\2\2\u0132\u0133"+
		"\3\2\2\2\u0133\u0131\3\2\2\2\u0133\u0134\3\2\2\2\u0134\u0135\3\2\2\2\u0135"+
		"\u0139\7\60\2\2\u0136\u0138\5S*\2\u0137\u0136\3\2\2\2\u0138\u013b\3\2"+
		"\2\2\u0139\u0137\3\2\2\2\u0139\u013a\3\2\2\2\u013ad\3\2\2\2\u013b\u0139"+
		"\3\2\2\2\u013c\u013f\5_\60\2\u013d\u013f\5c\62\2\u013e\u013c\3\2\2\2\u013e"+
		"\u013d\3\2\2\2\u013f\u0140\3\2\2\2\u0140\u0141\5E#\2\u0141f\3\2\2\2\u0142"+
		"\u0143\5;\36\2\u0143\u0144\5S*\2\u0144\u0145\5S*\2\u0145\u0155\t\3\2\2"+
		"\u0146\u0147\5S*\2\u0147\u0148\5S*\2\u0148\u0156\3\2\2\2\u0149\u014c\5"+
		"U+\2\u014a\u014c\5W,\2\u014b\u0149\3\2\2\2\u014b\u014a\3\2\2\2\u014c\u014f"+
		"\3\2\2\2\u014d\u0150\5U+\2\u014e\u0150\5W,\2\u014f\u014d\3\2\2\2\u014f"+
		"\u014e\3\2\2\2\u0150\u0153\3\2\2\2\u0151\u0154\5U+\2\u0152\u0154\5W,\2"+
		"\u0153\u0151\3\2\2\2\u0153\u0152\3\2\2\2\u0154\u0156\3\2\2\2\u0155\u0146"+
		"\3\2\2\2\u0155\u014b\3\2\2\2\u0156\u0157\3\2\2\2\u0157\u0158\t\3\2\2\u0158"+
		"\u0159\5S*\2\u0159\u015d\5S*\2\u015a\u015b\5S*\2\u015b\u015c\5S*\2\u015c"+
		"\u015e\3\2\2\2\u015d\u015a\3\2\2\2\u015d\u015e\3\2\2\2\u015e\u015f\3\2"+
		"\2\2\u015f\u0160\5=\37\2\u0160h\3\2\2\2\u0161\u0163\t\4\2\2\u0162\u0164"+
		"\t\5\2\2\u0163\u0162\3\2\2\2\u0163\u0164\3\2\2\2\u0164\u0166\3\2\2\2\u0165"+
		"\u0167\4\62;\2\u0166\u0165\3\2\2\2\u0167\u0168\3\2\2\2\u0168\u0166\3\2"+
		"\2\2\u0168\u0169\3\2\2\2\u0169j\3\2\2\2\u016a\u016b\t\6\2\2\u016bl\3\2"+
		"\2\2\u016c\u016d\7^\2\2\u016d\u0171\t\7\2\2\u016e\u0171\5q9\2\u016f\u0171"+
		"\5o8\2\u0170\u016c\3\2\2\2\u0170\u016e\3\2\2\2\u0170\u016f\3\2\2\2\u0171"+
		"n\3\2\2\2\u0172\u0173\7^\2\2\u0173\u0174\4\62\65\2\u0174\u0175\4\629\2"+
		"\u0175\u017c\4\629\2\u0176\u0177\7^\2\2\u0177\u0178\4\629\2\u0178\u017c"+
		"\4\629\2\u0179\u017a\7^\2\2\u017a\u017c\4\629\2\u017b\u0172\3\2\2\2\u017b"+
		"\u0176\3\2\2\2\u017b\u0179\3\2\2\2\u017cp\3\2\2\2\u017d\u017e\7^\2\2\u017e"+
		"\u017f\7w\2\2\u017f\u0180\5k\66\2\u0180\u0181\5k\66\2\u0181\u0182\5k\66"+
		"\2\u0182\u0183\5k\66\2\u0183r\3\2\2\2\u0184\u0186\7\17\2\2\u0185\u0184"+
		"\3\2\2\2\u0185\u0186\3\2\2\2\u0186\u0187\3\2\2\2\u0187\u0188\7\f\2\2\u0188"+
		"\u0189\3\2\2\2\u0189\u018a\b:\2\2\u018at\3\2\2\2\u018b\u0190\t\b\2\2\u018c"+
		"\u018d\7\17\2\2\u018d\u0190\7\f\2\2\u018e\u0190\7\5\2\2\u018f\u018b\3"+
		"\2\2\2\u018f\u018c\3\2\2\2\u018f\u018e\3\2\2\2\u0190\u0191\3\2\2\2\u0191"+
		"\u0192\b;\2\2\u0192v\3\2\2\2\33\2\u0103\u0105\u010f\u0111\u0117\u011c"+
		"\u011e\u0124\u0129\u012b\u0133\u0139\u013e\u014b\u014f\u0153\u0155\u015d"+
		"\u0163\u0168\u0170\u017b\u0185\u018f\3\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}