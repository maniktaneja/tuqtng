//line n1ql.y:2
package goyacc

import __yyfmt__ "fmt"

//line n1ql.y:2
import "github.com/couchbaselabs/clog"
import "github.com/couchbaselabs/tuqtng/parser"
import "github.com/couchbaselabs/tuqtng/ast"

func logDebugGrammar(format string, v ...interface{}) {
	clog.To(parser.PARSER_CHANNEL, format, v...)
}

//line n1ql.y:13
type yySymType struct {
	yys int
	s   string
	n   int
	f   float64
}

const ALTER = 57346
const BETWEEN = 57347
const BUCKET = 57348
const CAST = 57349
const COLLATE = 57350
const DATABASE = 57351
const DELETE = 57352
const EACH = 57353
const EXCEPT = 57354
const EXISTS = 57355
const IF = 57356
const INLINE = 57357
const INSERT = 57358
const INTERSECT = 57359
const INTO = 57360
const JOIN = 57361
const PATH = 57362
const UNION = 57363
const UPDATE = 57364
const POOL = 57365
const EXPLAIN = 57366
const CREATE = 57367
const DROP = 57368
const PRIMARY = 57369
const VIEW = 57370
const INDEX = 57371
const ON = 57372
const USING = 57373
const DISTINCT = 57374
const UNIQUE = 57375
const SELECT = 57376
const AS = 57377
const FROM = 57378
const WHERE = 57379
const KEY = 57380
const KEYS = 57381
const ORDER = 57382
const BY = 57383
const ASC = 57384
const DESC = 57385
const LIMIT = 57386
const OFFSET = 57387
const GROUP = 57388
const HAVING = 57389
const LBRACE = 57390
const RBRACE = 57391
const LBRACKET = 57392
const RBRACKET = 57393
const COMMA = 57394
const COLON = 57395
const TRUE = 57396
const FALSE = 57397
const NULL = 57398
const INT = 57399
const NUMBER = 57400
const IDENTIFIER = 57401
const STRING = 57402
const PLUS = 57403
const MINUS = 57404
const MULT = 57405
const DIV = 57406
const CONCAT = 57407
const AND = 57408
const OR = 57409
const NOT = 57410
const EQ = 57411
const NE = 57412
const GT = 57413
const GTE = 57414
const LT = 57415
const LTE = 57416
const LPAREN = 57417
const RPAREN = 57418
const LIKE = 57419
const IS = 57420
const VALUED = 57421
const MISSING = 57422
const DOT = 57423
const CASE = 57424
const WHEN = 57425
const THEN = 57426
const ELSE = 57427
const END = 57428
const ANY = 57429
const ALL = 57430
const FIRST = 57431
const ARRAY = 57432
const IN = 57433
const SATISFIES = 57434
const EVERY = 57435
const UNNEST = 57436
const FOR = 57437
const INNER = 57438
const LEFT = 57439
const OUTER = 57440
const MOD = 57441

var yyToknames = []string{
	"ALTER",
	"BETWEEN",
	"BUCKET",
	"CAST",
	"COLLATE",
	"DATABASE",
	"DELETE",
	"EACH",
	"EXCEPT",
	"EXISTS",
	"IF",
	"INLINE",
	"INSERT",
	"INTERSECT",
	"INTO",
	"JOIN",
	"PATH",
	"UNION",
	"UPDATE",
	"POOL",
	"EXPLAIN",
	"CREATE",
	"DROP",
	"PRIMARY",
	"VIEW",
	"INDEX",
	"ON",
	"USING",
	"DISTINCT",
	"UNIQUE",
	"SELECT",
	"AS",
	"FROM",
	"WHERE",
	"KEY",
	"KEYS",
	"ORDER",
	"BY",
	"ASC",
	"DESC",
	"LIMIT",
	"OFFSET",
	"GROUP",
	"HAVING",
	"LBRACE",
	"RBRACE",
	"LBRACKET",
	"RBRACKET",
	"COMMA",
	"COLON",
	"TRUE",
	"FALSE",
	"NULL",
	"INT",
	"NUMBER",
	"IDENTIFIER",
	"STRING",
	"PLUS",
	"MINUS",
	"MULT",
	"DIV",
	"CONCAT",
	"AND",
	"OR",
	"NOT",
	"EQ",
	"NE",
	"GT",
	"GTE",
	"LT",
	"LTE",
	"LPAREN",
	"RPAREN",
	"LIKE",
	"IS",
	"VALUED",
	"MISSING",
	"DOT",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"END",
	"ANY",
	"ALL",
	"FIRST",
	"ARRAY",
	"IN",
	"SATISFIES",
	"EVERY",
	"UNNEST",
	"FOR",
	"INNER",
	"LEFT",
	"OUTER",
	"MOD",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 183
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1631

var yyAct = []int{

	50, 292, 210, 141, 31, 83, 146, 101, 233, 89,
	127, 76, 34, 155, 316, 85, 145, 131, 201, 313,
	131, 110, 111, 112, 113, 115, 306, 4, 125, 81,
	45, 265, 133, 214, 49, 79, 213, 297, 128, 295,
	291, 126, 174, 164, 151, 96, 191, 359, 103, 328,
	303, 302, 259, 209, 343, 127, 247, 256, 131, 114,
	143, 134, 137, 138, 139, 132, 110, 111, 112, 113,
	115, 116, 117, 125, 118, 123, 121, 122, 119, 120,
	258, 257, 124, 128, 329, 192, 126, 327, 350, 305,
	84, 351, 87, 88, 140, 161, 162, 152, 153, 192,
	293, 157, 170, 227, 114, 143, 326, 324, 169, 281,
	278, 176, 177, 178, 179, 180, 181, 182, 183, 184,
	185, 186, 187, 188, 189, 190, 175, 196, 193, 127,
	46, 294, 208, 173, 211, 273, 35, 287, 206, 197,
	236, 237, 112, 113, 115, 271, 148, 125, 81, 85,
	199, 198, 92, 35, 79, 248, 228, 128, 231, 225,
	126, 286, 232, 127, 246, 229, 239, 238, 94, 95,
	149, 37, 255, 103, 244, 243, 226, 36, 114, 249,
	92, 125, 234, 93, 90, 236, 237, 94, 95, 230,
	32, 128, 252, 192, 126, 163, 35, 92, 17, 92,
	16, 160, 156, 208, 208, 107, 235, 97, 91, 206,
	206, 93, 82, 267, 268, 269, 270, 172, 272, 43,
	274, 260, 261, 171, 84, 275, 87, 88, 93, 276,
	93, 159, 242, 165, 300, 158, 279, 283, 284, 280,
	299, 277, 289, 100, 282, 254, 51, 285, 288, 240,
	222, 241, 262, 224, 221, 166, 147, 330, 298, 325,
	301, 290, 223, 208, 296, 220, 307, 308, 251, 206,
	48, 13, 99, 167, 168, 47, 40, 109, 41, 236,
	237, 304, 94, 95, 319, 21, 27, 25, 321, 17,
	320, 322, 29, 30, 362, 323, 73, 342, 74, 129,
	130, 195, 68, 69, 70, 194, 72, 56, 64, 341,
	53, 332, 333, 245, 334, 335, 52, 336, 337, 108,
	106, 105, 104, 58, 42, 22, 338, 23, 19, 339,
	59, 26, 211, 340, 344, 60, 154, 62, 63, 142,
	67, 61, 66, 354, 355, 2, 353, 65, 357, 18,
	205, 358, 12, 10, 127, 204, 44, 264, 57, 55,
	356, 17, 54, 16, 363, 110, 111, 112, 113, 115,
	116, 117, 125, 118, 123, 121, 122, 119, 120, 98,
	39, 124, 128, 102, 86, 126, 33, 347, 78, 77,
	348, 75, 28, 127, 15, 250, 14, 24, 38, 20,
	11, 7, 9, 114, 110, 111, 112, 113, 115, 116,
	117, 125, 118, 123, 121, 122, 119, 120, 8, 6,
	124, 128, 5, 1, 126, 0, 317, 0, 0, 318,
	0, 0, 127, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 114, 110, 111, 112, 113, 115, 116, 117,
	125, 118, 123, 121, 122, 119, 120, 0, 0, 124,
	128, 0, 0, 126, 0, 314, 0, 0, 315, 0,
	0, 127, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 114, 110, 111, 112, 113, 115, 116, 117, 125,
	118, 123, 121, 122, 119, 120, 0, 0, 124, 128,
	0, 0, 126, 3, 12, 10, 0, 0, 0, 0,
	127, 0, 219, 17, 0, 16, 218, 0, 0, 0,
	114, 110, 111, 112, 113, 115, 116, 117, 125, 118,
	123, 121, 122, 119, 120, 0, 0, 124, 128, 0,
	0, 126, 0, 0, 0, 0, 0, 0, 0, 127,
	0, 217, 0, 0, 0, 216, 0, 0, 0, 114,
	110, 111, 112, 113, 115, 116, 117, 125, 118, 123,
	121, 122, 119, 120, 0, 0, 124, 128, 0, 0,
	126, 0, 0, 0, 0, 361, 0, 0, 127, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 114, 110,
	111, 112, 113, 115, 116, 117, 125, 118, 123, 121,
	122, 119, 120, 0, 0, 124, 128, 0, 0, 126,
	0, 0, 0, 0, 360, 0, 0, 127, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 114, 110, 111,
	112, 113, 115, 116, 117, 125, 118, 123, 121, 122,
	119, 120, 0, 0, 124, 128, 0, 0, 126, 0,
	0, 0, 0, 352, 0, 0, 127, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 114, 110, 111, 112,
	113, 115, 116, 117, 125, 118, 123, 121, 122, 119,
	120, 0, 0, 124, 128, 0, 0, 126, 0, 0,
	0, 0, 349, 0, 0, 127, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 114, 110, 111, 112, 113,
	115, 116, 117, 125, 118, 123, 121, 122, 119, 120,
	0, 0, 124, 128, 0, 0, 126, 0, 0, 0,
	0, 346, 0, 0, 127, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 114, 110, 111, 112, 113, 115,
	116, 117, 125, 118, 123, 121, 122, 119, 120, 0,
	0, 124, 128, 0, 0, 126, 0, 0, 0, 0,
	345, 0, 0, 127, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 114, 110, 111, 112, 113, 115, 116,
	117, 125, 118, 123, 121, 122, 119, 120, 0, 0,
	124, 128, 0, 0, 126, 0, 331, 0, 0, 0,
	0, 0, 127, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 114, 110, 111, 112, 113, 115, 116, 117,
	125, 118, 123, 121, 122, 119, 120, 0, 0, 124,
	128, 0, 0, 126, 0, 0, 0, 0, 312, 0,
	0, 127, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 114, 110, 111, 112, 113, 115, 116, 117, 125,
	118, 123, 121, 122, 119, 120, 0, 0, 124, 128,
	0, 0, 126, 0, 0, 0, 0, 0, 0, 0,
	127, 0, 0, 311, 0, 0, 0, 0, 0, 0,
	114, 110, 111, 112, 113, 115, 116, 117, 125, 118,
	123, 121, 122, 119, 120, 0, 0, 124, 128, 0,
	0, 126, 0, 0, 0, 0, 0, 0, 0, 127,
	0, 0, 310, 0, 0, 0, 0, 0, 0, 114,
	110, 111, 112, 113, 115, 116, 117, 125, 118, 123,
	121, 122, 119, 120, 0, 0, 124, 128, 0, 0,
	126, 0, 0, 0, 0, 309, 0, 0, 127, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 114, 110,
	111, 112, 113, 115, 116, 117, 125, 118, 123, 121,
	122, 119, 120, 0, 0, 124, 128, 0, 0, 126,
	0, 0, 266, 0, 0, 0, 0, 127, 253, 0,
	0, 0, 0, 0, 0, 0, 0, 114, 110, 111,
	112, 113, 115, 116, 117, 125, 118, 123, 121, 122,
	119, 120, 0, 0, 124, 128, 0, 0, 126, 0,
	0, 0, 0, 0, 0, 0, 127, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 114, 110, 111, 112,
	113, 115, 116, 117, 125, 118, 123, 121, 122, 119,
	120, 0, 0, 124, 128, 0, 0, 126, 0, 0,
	0, 0, 0, 0, 0, 127, 0, 0, 215, 0,
	0, 0, 0, 0, 0, 114, 110, 111, 112, 113,
	115, 116, 117, 125, 118, 123, 121, 122, 119, 120,
	0, 0, 124, 128, 0, 0, 126, 0, 0, 0,
	0, 0, 0, 0, 127, 0, 0, 212, 0, 0,
	0, 0, 0, 0, 114, 110, 111, 112, 113, 115,
	116, 117, 125, 118, 123, 121, 122, 119, 120, 0,
	0, 124, 128, 0, 0, 126, 0, 0, 0, 0,
	0, 0, 0, 127, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 114, 110, 111, 112, 113, 115, 116,
	117, 125, 118, 123, 121, 122, 119, 120, 0, 0,
	124, 128, 0, 0, 263, 0, 0, 0, 0, 0,
	0, 0, 127, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 114, 110, 111, 112, 113, 115, 116, 117,
	125, 118, 123, 121, 122, 119, 120, 0, 0, 124,
	128, 0, 0, 150, 127, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 110, 111, 112, 113, 115,
	116, 114, 125, 118, 123, 121, 122, 119, 120, 0,
	0, 124, 128, 0, 0, 126, 127, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 110, 111, 112,
	113, 115, 0, 114, 125, 118, 123, 121, 122, 119,
	120, 0, 0, 124, 128, 202, 203, 126, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 73, 0, 74, 0, 114, 0, 68, 69, 70,
	71, 72, 56, 64, 0, 53, 207, 0, 0, 0,
	0, 52, 0, 0, 0, 0, 0, 0, 58, 200,
	0, 0, 0, 0, 0, 59, 0, 0, 0, 0,
	60, 0, 62, 63, 0, 73, 61, 74, 0, 0,
	0, 68, 69, 70, 71, 72, 56, 64, 0, 53,
	207, 0, 0, 0, 0, 52, 0, 0, 0, 0,
	0, 0, 58, 0, 0, 0, 0, 0, 0, 59,
	0, 0, 0, 0, 60, 0, 62, 63, 0, 73,
	61, 74, 0, 0, 0, 68, 69, 70, 71, 72,
	56, 64, 0, 53, 80, 0, 0, 0, 0, 52,
	0, 0, 0, 0, 0, 0, 58, 0, 0, 0,
	0, 0, 0, 59, 0, 0, 0, 0, 60, 0,
	62, 63, 0, 73, 61, 74, 144, 0, 0, 68,
	69, 70, 71, 72, 56, 64, 0, 53, 0, 0,
	0, 0, 0, 52, 0, 0, 0, 0, 0, 0,
	58, 0, 0, 0, 0, 0, 0, 59, 0, 0,
	0, 0, 60, 0, 62, 63, 0, 73, 61, 74,
	0, 0, 0, 68, 69, 70, 71, 72, 56, 64,
	0, 53, 0, 0, 0, 0, 0, 52, 0, 0,
	0, 0, 0, 0, 58, 0, 0, 0, 0, 0,
	0, 59, 0, 0, 0, 0, 60, 0, 62, 63,
	0, 73, 61, 74, 0, 0, 0, 68, 69, 70,
	71, 72, 136, 64, 0, 53, 0, 0, 0, 0,
	0, 52, 0, 0, 0, 0, 0, 0, 58, 0,
	0, 0, 0, 0, 0, 59, 0, 0, 0, 0,
	60, 0, 62, 63, 0, 73, 61, 74, 0, 0,
	0, 68, 69, 70, 71, 72, 135, 64, 0, 53,
	0, 0, 0, 0, 0, 52, 0, 0, 0, 0,
	0, 0, 58, 0, 0, 0, 0, 0, 0, 59,
	0, 0, 0, 0, 60, 0, 62, 63, 0, 0,
	61,
}
var yyPact = []int{

	479, -1000, -1000, 327, -1000, -1000, -1000, -1000, -1000, -1000,
	299, 245, 298, 251, 249, 260, 137, -1000, -1000, 118,
	232, 237, 295, 160, 249, 77, 224, 1449, 1361, -1000,
	-1000, -1000, 153, -4, 149, -1000, -36, 148, -1000, 227,
	186, 1449, 292, 291, 224, -1000, 146, 255, 236, -1000,
	1084, -1000, 1449, 1449, -1000, -1000, -17, -1000, 1449, -51,
	1537, 1493, 1449, 1449, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 45, 1405, -1000, -1000, 204, -1000, 111,
	-1000, 1162, -37, -1000, 94, 94, 317, -1000, -85, -1000,
	143, 244, 178, 142, 1449, 1449, 136, -38, -1000, 176,
	-1000, -1000, 203, 231, 49, 164, -1000, -39, -1000, 1449,
	1449, 1449, 1449, 1449, 1449, 1449, 1449, 1449, 1449, 1449,
	1449, 1449, 1449, 1449, 1449, -31, 134, 248, 71, -1000,
	-1000, 1273, -23, 1449, 1045, -55, -58, 1006, 460, 421,
	-1000, 216, 202, 197, -1000, 211, 201, 1361, 117, -1000,
	40, 94, 130, 147, 94, -1000, 244, -1000, 198, 175,
	-1000, 1084, 1084, -1000, 116, -1000, 1449, -1000, -1000, 282,
	105, -19, 96, -1000, 94, 221, 79, 79, 113, 113,
	113, 113, 1226, 1194, -40, -40, -40, -40, -40, -40,
	-40, 1449, -1000, 967, 192, 115, -1000, 1, -1000, -1000,
	-1000, -24, 1317, 1317, 200, -1000, -1000, -1000, 1123, -1000,
	-54, 928, 1449, 1449, 1449, 1449, 86, 1449, 76, 1449,
	-1000, 0, 1449, -1000, 1449, -1000, -1000, -1000, -1000, 51,
	-4, -1000, -1000, -4, 50, 241, 1449, 1449, 102, -1000,
	-1000, 191, 210, -41, -1000, 72, -42, 1449, -44, -1000,
	-1000, 1449, -40, -1000, 183, 209, -1000, -1000, -1000, -1000,
	-25, -26, 1317, 26, -60, 1449, 1449, 889, 850, 811,
	772, -72, 382, -77, 343, -1000, -1000, -1000, -4, -1000,
	-1000, 241, -4, 1084, 1084, -4, 241, 48, 208, -1000,
	-1000, 47, -1000, -1000, -1000, 28, -27, 25, -1000, 206,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 1084, 733, -1000,
	1449, 1449, -1000, 1449, 1449, -1000, 1449, 1449, -1000, -1000,
	-4, -1000, -1000, -4, 241, -1000, -1000, 278, 266, -21,
	-1000, 1449, 694, 655, 304, 616, 5, 577, -1000, -1000,
	-4, 72, 72, 1449, -1000, -1000, -1000, 1449, -1000, -1000,
	1449, -1000, -1000, -1000, -1000, -1000, -29, 538, 499, 263,
	-1000, -1000, 72, -1000,
}
var yyPgo = []int{

	0, 423, 345, 27, 422, 419, 418, 402, 1, 16,
	401, 400, 399, 398, 271, 397, 331, 275, 396, 395,
	6, 394, 392, 391, 11, 389, 388, 0, 4, 386,
	5, 12, 9, 8, 384, 7, 383, 380, 379, 246,
	362, 359, 358, 2, 357, 18, 355, 350, 347, 342,
	340, 3, 339,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 2, 4, 4, 6, 6,
	6, 6, 7, 7, 7, 7, 7, 8, 8, 5,
	5, 3, 10, 11, 11, 17, 17, 19, 19, 14,
	21, 22, 22, 22, 23, 24, 24, 25, 25, 25,
	25, 26, 26, 15, 15, 15, 18, 18, 28, 28,
	30, 30, 30, 30, 30, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 30, 33,
	33, 34, 34, 34, 29, 29, 29, 29, 29, 29,
	32, 32, 16, 16, 12, 12, 35, 35, 36, 36,
	36, 13, 13, 13, 37, 38, 20, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 39, 39, 39, 40, 41,
	41, 41, 41, 41, 41, 41, 41, 41, 41, 41,
	41, 41, 41, 41, 41, 41, 41, 41, 41, 43,
	43, 44, 44, 31, 31, 31, 31, 31, 31, 45,
	45, 46, 46, 47, 47, 42, 42, 42, 42, 42,
	42, 42, 48, 48, 49, 49, 51, 51, 52, 50,
	50, 9, 9,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 1, 1, 1, 5, 8,
	7, 10, 8, 11, 10, 13, 5, 1, 1, 5,
	8, 1, 3, 4, 4, 0, 4, 0, 2, 3,
	1, 0, 1, 1, 1, 1, 3, 1, 1, 3,
	2, 1, 3, 0, 2, 5, 2, 5, 1, 2,
	2, 4, 3, 3, 3, 5, 4, 3, 5, 4,
	4, 6, 5, 4, 5, 5, 6, 6, 7, 2,
	2, 1, 1, 2, 1, 2, 3, 2, 4, 3,
	2, 2, 0, 2, 0, 3, 1, 3, 1, 2,
	2, 0, 1, 2, 2, 2, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 4, 3, 4, 6, 5, 5, 3, 4,
	3, 4, 3, 4, 1, 2, 2, 1, 1, 1,
	1, 3, 5, 5, 7, 7, 5, 9, 7, 7,
	5, 9, 7, 7, 5, 3, 4, 5, 5, 3,
	5, 0, 2, 1, 4, 6, 5, 5, 3, 1,
	3, 1, 1, 1, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 3, 1, 3, 3, 2,
	3, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, 24, -3, -4, -5, -10, -6, -7,
	26, -11, 25, -14, -18, -21, 36, 34, -2, 29,
	-12, 40, 27, 29, -15, 36, -16, 37, -22, 32,
	33, -28, 53, -29, -31, 59, 59, 53, -13, -37,
	44, 41, 29, 59, -16, -28, 53, -17, 46, -20,
	-27, -39, 68, 62, -40, -41, 59, -42, 75, 82,
	87, 93, 89, 90, 60, -48, -49, -50, 54, 55,
	56, 57, 58, 48, 50, -23, -24, -25, -26, -20,
	63, -27, 59, -30, 94, 19, -34, 96, 97, -32,
	35, 59, 50, 81, 38, 39, 81, 59, -38, 45,
	57, -35, -36, -20, 30, 30, -17, 59, -14, 41,
	61, 62, 63, 64, 99, 65, 66, 67, 69, 73,
	74, 71, 72, 70, 77, 68, 81, 50, 78, -39,
	-39, 75, -20, 83, -27, 59, 59, -27, -27, -27,
	49, -51, -52, 60, 51, -9, -20, 52, 35, 59,
	81, 81, -31, -31, 19, 98, 59, -32, 57, 53,
	59, -27, -27, 59, 81, 57, 52, 42, 43, 59,
	53, 59, 53, -3, 81, -9, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, 77, 59, -27, 57, 53, 56, 68, 80, 79,
	76, -45, 32, 33, -46, -47, -20, 63, -27, 76,
	-43, -27, 92, 91, 91, 92, 95, 91, 95, 91,
	49, 52, 53, 51, 52, -24, 59, 63, -28, 35,
	59, -30, -32, -33, 35, 59, 38, 39, -31, -32,
	51, 53, 57, 59, -35, 31, 59, 75, 59, -28,
	-19, 47, -27, 51, 53, 57, 56, 80, 79, 76,
	-45, -45, 52, 81, -44, 85, 84, -27, -27, -27,
	-27, 59, -27, 59, -27, -51, -20, -9, 59, -30,
	-30, 59, -33, -27, -27, -33, 59, 35, 57, 51,
	51, 81, -8, 28, 59, 81, -9, 81, -20, 57,
	51, 51, 76, 76, -45, 63, 86, -27, -27, 86,
	92, 92, 86, 91, 83, 86, 91, 83, 86, -30,
	-33, -30, -30, -33, 59, 51, 59, 59, 76, 59,
	51, 83, -27, -27, -27, -27, -27, -27, -30, -30,
	-33, 31, 31, 75, -43, 86, 86, 83, 86, 86,
	83, 86, 86, -30, -8, -8, -9, -27, -27, 76,
	86, 86, 31, -8,
}
var yyDef = []int{

	0, -2, 1, 0, 3, 4, 5, 21, 6, 7,
	0, 84, 0, 43, 82, 31, 0, 30, 2, 0,
	91, 0, 0, 0, 82, 0, 25, 0, 0, 32,
	33, 46, 0, 48, 74, 153, 0, 0, 22, 92,
	0, 0, 0, 0, 25, 44, 0, 0, 0, 83,
	96, 124, 0, 0, 127, 128, 129, 130, 0, 0,
	0, 0, 0, 0, 165, 166, 167, 168, 169, 170,
	171, 172, 173, 0, 0, 29, 34, 35, 37, 38,
	41, 96, 0, 49, 0, 0, 0, 71, 72, 75,
	0, 77, 0, 0, 0, 0, 0, 0, 93, 0,
	94, 85, 86, 88, 0, 0, 23, 0, 24, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 125,
	126, 0, 0, 0, 0, 129, 129, 0, 0, 0,
	174, 0, 176, 0, 179, 0, 181, 0, 0, 40,
	0, 0, 50, 0, 0, 73, 76, 79, 0, 0,
	158, 80, 81, 19, 0, 95, 0, 89, 90, 8,
	0, 0, 0, 16, 0, 27, 97, 98, 99, 100,
	101, 102, 103, 104, 105, 106, 107, 108, 109, 110,
	111, 0, 113, 0, 172, 0, 118, 0, 120, 122,
	145, 0, 0, 0, 159, 161, 162, 163, 96, 131,
	151, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	175, 0, 0, 180, 0, 36, 39, 42, 47, 0,
	52, 53, 54, 57, 0, 0, 0, 0, 0, 78,
	154, 0, 0, 0, 87, 0, 0, 0, 0, 45,
	26, 0, 112, 114, 0, 0, 119, 121, 123, 146,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 177, 178, 182, 51, 56,
	60, 0, 59, 69, 70, 63, 0, 0, 0, 156,
	157, 0, 10, 17, 18, 0, 0, 0, 28, 0,
	116, 117, 147, 148, 160, 164, 132, 152, 149, 133,
	0, 0, 136, 0, 0, 140, 0, 0, 144, 55,
	58, 62, 64, 65, 0, 155, 20, 9, 12, 0,
	115, 0, 0, 0, 0, 0, 0, 0, 61, 66,
	67, 0, 0, 0, 150, 134, 135, 0, 139, 138,
	0, 143, 142, 68, 11, 14, 0, 0, 0, 13,
	137, 141, 0, 15,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77, 78, 79, 80, 81,
	82, 83, 84, 85, 86, 87, 88, 89, 90, 91,
	92, 93, 94, 95, 96, 97, 98, 99,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line n1ql.y:55
		{
			logDebugGrammar("INPUT")
		}
	case 2:
		//line n1ql.y:59
		{
			logDebugGrammar("INPUT - EXPLAIN")
			parsingStatement.SetExplainOnly(true)
		}
	case 3:
		//line n1ql.y:65
		{
			logDebugGrammar("STMT - SELECT")
		}
	case 4:
		//line n1ql.y:69
		{
		}
	case 5:
		//line n1ql.y:72
		{
			logDebugGrammar("STMT - DROP INDEX")
		}
	case 6:
		//line n1ql.y:79
		{
			logDebugGrammar("STMT - CREATE PRIMARY INDEX")
		}
	case 7:
		//line n1ql.y:83
		{
			logDebugGrammar("STMT - CREATE SECONDARY INDEX")
		}
	case 8:
		//line n1ql.y:89
		{
			bucket := yyS[yypt-0].s
			createIndexStmt := ast.NewCreateIndexStatement()
			createIndexStmt.Bucket = bucket
			createIndexStmt.Primary = true
			parsingStatement = createIndexStmt
		}
	case 9:
		//line n1ql.y:97
		{
			pool := yyS[yypt-2].s
			bucket := yyS[yypt-0].s
			createIndexStmt := ast.NewCreateIndexStatement()
			createIndexStmt.Pool = pool
			createIndexStmt.Bucket = bucket
			createIndexStmt.Primary = true
			parsingStatement = createIndexStmt
		}
	case 10:
		//line n1ql.y:107
		{
			method := parsingStack.Pop().(string)
			bucket := yyS[yypt-2].s
			createIndexStmt := ast.NewCreateIndexStatement()
			createIndexStmt.Bucket = bucket
			createIndexStmt.Method = method
			createIndexStmt.Primary = true
			parsingStatement = createIndexStmt
		}
	case 11:
		//line n1ql.y:117
		{
			method := parsingStack.Pop().(string)
			bucket := yyS[yypt-2].s
			pool := yyS[yypt-4].s
			createIndexStmt := ast.NewCreateIndexStatement()
			createIndexStmt.Pool = pool
			createIndexStmt.Bucket = bucket
			createIndexStmt.Method = method
			createIndexStmt.Primary = true
			parsingStatement = createIndexStmt
		}
	case 12:
		//line n1ql.y:131
		{
			on := parsingStack.Pop().(ast.ExpressionList)
			bucket := yyS[yypt-3].s
			name := yyS[yypt-5].s
			createIndexStmt := ast.NewCreateIndexStatement()
			createIndexStmt.On = on
			createIndexStmt.Bucket = bucket
			createIndexStmt.Name = name
			createIndexStmt.Primary = false
			parsingStatement = createIndexStmt
		}
	case 13:
		//line n1ql.y:143
		{
			on := parsingStack.Pop().(ast.ExpressionList)
			bucket := yyS[yypt-3].s
			pool := yyS[yypt-5].s
			name := yyS[yypt-8].s
			createIndexStmt := ast.NewCreateIndexStatement()
			createIndexStmt.On = on
			createIndexStmt.Pool = pool
			createIndexStmt.Bucket = bucket
			createIndexStmt.Name = name
			createIndexStmt.Primary = false
			parsingStatement = createIndexStmt
		}
	case 14:
		//line n1ql.y:157
		{
			method := parsingStack.Pop().(string)
			on := parsingStack.Pop().(ast.ExpressionList)
			bucket := yyS[yypt-5].s
			name := yyS[yypt-7].s
			createIndexStmt := ast.NewCreateIndexStatement()
			createIndexStmt.On = on
			createIndexStmt.Bucket = bucket
			createIndexStmt.Name = name
			createIndexStmt.Method = method
			createIndexStmt.Primary = false
			parsingStatement = createIndexStmt
		}
	case 15:
		//line n1ql.y:171
		{
			method := parsingStack.Pop().(string)
			on := parsingStack.Pop().(ast.ExpressionList)
			bucket := yyS[yypt-5].s
			pool := yyS[yypt-7].s
			name := yyS[yypt-10].s
			createIndexStmt := ast.NewCreateIndexStatement()
			createIndexStmt.On = on
			createIndexStmt.Pool = pool
			createIndexStmt.Bucket = bucket
			createIndexStmt.Name = name
			createIndexStmt.Method = method
			createIndexStmt.Primary = false
			parsingStatement = createIndexStmt
		}
	case 16:
		//line n1ql.y:188
		{

			logDebugGrammar("COMPOUND_CREATE_INDEX_STATEMENT")
			select_statement := parsingStatement.(*ast.SelectStatement)
			name := yyS[yypt-2].s
			createIndexStmt := ast.NewCreateIndexStatement()
			createIndexStmt.Select = select_statement
			createIndexStmt.Name = name
			createIndexStmt.Primary = false
			parsingStatement = createIndexStmt
		}
	case 17:
		//line n1ql.y:203
		{
			parsingStack.Push("view")
		}
	case 18:
		//line n1ql.y:207
		{
			parsingStack.Push(yyS[yypt-0].s)
		}
	case 19:
		//line n1ql.y:213
		{
			bucket := yyS[yypt-2].s
			name := yyS[yypt-0].s
			dropIndexStmt := ast.NewDropIndexStatement()
			dropIndexStmt.Bucket = bucket
			dropIndexStmt.Name = name
			parsingStatement = dropIndexStmt
		}
	case 20:
		//line n1ql.y:222
		{
			bucket := yyS[yypt-2].s
			pool := yyS[yypt-4].s
			name := yyS[yypt-0].s
			dropIndexStmt := ast.NewDropIndexStatement()
			dropIndexStmt.Pool = pool
			dropIndexStmt.Bucket = bucket
			dropIndexStmt.Name = name
			parsingStatement = dropIndexStmt
		}
	case 21:
		//line n1ql.y:236
		{
			logDebugGrammar("SELECT_STMT")
		}
	case 22:
		//line n1ql.y:242
		{
			// future extensibility for comining queries with UNION, etc
			logDebugGrammar("SELECT_COMPOUND")
		}
	case 23:
		//line n1ql.y:249
		{
			logDebugGrammar("SELECT_CORE")
		}
	case 24:
		//line n1ql.y:253
		{
			logDebugGrammar("SELECT_CORE")
		}
	case 25:
		//line n1ql.y:259
		{
		}
	case 26:
		//line n1ql.y:262
		{
			group_by := parsingStack.Pop().(ast.ExpressionList)
			switch parsingStatement := parsingStatement.(type) {
			case *ast.SelectStatement:
				parsingStatement.GroupBy = group_by
			default:
				logDebugGrammar("This statement does not support GROUP BY")
			}
		}
	case 27:
		//line n1ql.y:274
		{
		}
	case 28:
		//line n1ql.y:277
		{
			logDebugGrammar("SELECT HAVING - EXPR")
			having_part := parsingStack.Pop().(ast.Expression)
			switch parsingStatement := parsingStatement.(type) {
			case *ast.SelectStatement:
				parsingStatement.Having = having_part
			default:
				logDebugGrammar("This statement does not support HAVING")
			}
		}
	case 29:
		//line n1ql.y:290
		{
			logDebugGrammar("SELECT_SELECT")
		}
	case 30:
		//line n1ql.y:296
		{
			logDebugGrammar("SELECT_SELECT_HEAD")
		}
	case 31:
		//line n1ql.y:302
		{
		}
	case 32:
		//line n1ql.y:305
		{
			logDebugGrammar("SELECT_SELECT_QUALIFIER DISTINCT")
			switch parsingStatement := parsingStatement.(type) {
			case *ast.SelectStatement:
				parsingStatement.Distinct = true
			default:
				logDebugGrammar("This statement does not support WHERE")
			}
		}
	case 33:
		//line n1ql.y:315
		{
			logDebugGrammar("SELECT_SELECT_QUALIFIER UNIQUE")
			switch parsingStatement := parsingStatement.(type) {
			case *ast.SelectStatement:
				parsingStatement.Distinct = true
			default:
				logDebugGrammar("This statement does not support WHERE")
			}
		}
	case 34:
		//line n1ql.y:327
		{
			logDebugGrammar("SELECT SELECT TAIL - EXPR")
			result_expr_list := parsingStack.Pop().(ast.ResultExpressionList)
			switch parsingStatement := parsingStatement.(type) {
			case *ast.SelectStatement:
				parsingStatement.Select = result_expr_list
			default:
				logDebugGrammar("This statement does not support WHERE")
			}

		}
	case 35:
		//line n1ql.y:341
		{
			result_expr := parsingStack.Pop().(*ast.ResultExpression)
			parsingStack.Push(ast.ResultExpressionList{result_expr})
		}
	case 36:
		//line n1ql.y:346
		{
			result_expr_list := parsingStack.Pop().(ast.ResultExpressionList)
			result_expr := parsingStack.Pop().(*ast.ResultExpression)
			// list items pushed onto the stack end up in reverse order
			// this prepends items in the list to restore order
			new_list := ast.ResultExpressionList{result_expr}
			for _, v := range result_expr_list {
				new_list = append(new_list, v)
			}
			parsingStack.Push(new_list)
		}
	case 37:
		//line n1ql.y:359
		{
			logDebugGrammar("RESULT STAR")
		}
	case 38:
		//line n1ql.y:363
		{
			logDebugGrammar("RESULT EXPR")
			expr_part := parsingStack.Pop().(ast.Expression)
			result_expr := ast.NewResultExpression(expr_part)
			parsingStack.Push(result_expr)
		}
	case 39:
		//line n1ql.y:370
		{
			logDebugGrammar("RESULT EXPR AS ID")
			expr_part := parsingStack.Pop().(ast.Expression)
			result_expr := ast.NewResultExpressionWithAlias(expr_part, yyS[yypt-0].s)
			parsingStack.Push(result_expr)
		}
	case 40:
		//line n1ql.y:377
		{
			logDebugGrammar("RESULT EXPR ID")
			expr_part := parsingStack.Pop().(ast.Expression)
			result_expr := ast.NewResultExpressionWithAlias(expr_part, yyS[yypt-0].s)
			parsingStack.Push(result_expr)
		}
	case 41:
		//line n1ql.y:386
		{
			logDebugGrammar("STAR")
			result_expr := ast.NewStarResultExpression()
			parsingStack.Push(result_expr)
		}
	case 42:
		//line n1ql.y:392
		{
			logDebugGrammar("PATH DOT STAR")
			expr_part := parsingStack.Pop().(ast.Expression)
			result_expr := ast.NewDotStarResultExpression(expr_part)
			parsingStack.Push(result_expr)
		}
	case 43:
		//line n1ql.y:401
		{
			logDebugGrammar("SELECT FROM - EMPTY")
		}
	case 44:
		//line n1ql.y:405
		{
			logDebugGrammar("SELECT FROM - DATASOURCE")
			from := parsingStack.Pop().(*ast.From)
			switch parsingStatement := parsingStatement.(type) {
			case *ast.SelectStatement:
				parsingStatement.From = from
			default:
				logDebugGrammar("This statement does not support FROM")
			}
		}
	case 45:
		//line n1ql.y:416
		{
			logDebugGrammar("SELECT FROM - DATASOURCE WITH POOL")
			from := parsingStack.Pop().(*ast.From)
			from.Pool = yyS[yypt-2].s
			switch parsingStatement := parsingStatement.(type) {
			case *ast.SelectStatement:
				parsingStatement.From = from
			default:
				logDebugGrammar("This statement does not support FROM")
			}
		}
	case 46:
		//line n1ql.y:431
		{
			logDebugGrammar("SELECT FROM - DATASOURCE over here")
			from := parsingStack.Pop().(*ast.From)
			switch parsingStatement := parsingStatement.(type) {
			case *ast.SelectStatement:
				parsingStatement.From = from
			default:
				logDebugGrammar("This statement does not support FROM")
			}
		}
	case 47:
		//line n1ql.y:442
		{
			logDebugGrammar("SELECT FROM - DATASOURCE WITH POOL")
			from := parsingStack.Pop().(*ast.From)
			from.Pool = yyS[yypt-2].s
			switch parsingStatement := parsingStatement.(type) {
			case *ast.SelectStatement:
				parsingStatement.From = from
			default:
				logDebugGrammar("This statement does not support FROM")
			}
		}
	case 48:
		//line n1ql.y:456
		{
			logDebugGrammar("FROM DATASOURCE WITHOUT UNNEST")
		}
	case 49:
		//line n1ql.y:460
		{
			logDebugGrammar("FROM DATASOURCE WITH UNNEST")
			rest := parsingStack.Pop().(*ast.From)
			last := parsingStack.Pop().(*ast.From)
			last.Over = rest
			parsingStack.Push(last)
		}
	case 50:
		//line n1ql.y:471
		{
			logDebugGrammar("UNNEST")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: ""})
		}
	case 51:
		//line n1ql.y:478
		{
			logDebugGrammar("UNNEST AS")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
		}
	case 52:
		//line n1ql.y:485
		{
			logDebugGrammar("UNNEST AS")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
		}
	case 53:
		//line n1ql.y:492
		{
			logDebugGrammar("UNNEST nested")
			rest := parsingStack.Pop().(*ast.From)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Over: rest})
		}
	case 54:
		//line n1ql.y:499
		{
			logDebugGrammar("UNNEST KEY_EXPR")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Keys: key_expr})
		}
	case 55:
		//line n1ql.y:506
		{
			logDebugGrammar("UNNEST AS nested")
			rest := parsingStack.Pop().(*ast.From)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Over: rest})
		}
	case 56:
		//line n1ql.y:513
		{
			logDebugGrammar("UNNEST AS nested")
			rest := parsingStack.Pop().(*ast.From)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Over: rest})
		}
	case 57:
		//line n1ql.y:520
		{
			logDebugGrammar("JOIN KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Keys: key_expr})
		}
	case 58:
		//line n1ql.y:527
		{
			logDebugGrammar("JOIN AS KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Keys: key_expr})
		}
	case 59:
		//line n1ql.y:534
		{
			logDebugGrammar("JOIN AS KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Keys: key_expr})
		}
	case 60:
		//line n1ql.y:541
		{
			logDebugGrammar("JOIN KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Keys: key_expr, Over: rest})
		}
	case 61:
		//line n1ql.y:549
		{
			logDebugGrammar("JOIN AS KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Keys: key_expr, Over: rest})
		}
	case 62:
		//line n1ql.y:557
		{
			logDebugGrammar("JOIN AS KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Keys: key_expr, Over: rest})
		}
	case 63:
		//line n1ql.y:565
		{
			logDebugGrammar("TYPE JOIN KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Type: Type, Keys: key_expr})

		}
	case 64:
		//line n1ql.y:574
		{
			logDebugGrammar("TYPE JOIN KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Type: Type, Keys: key_expr, Over: rest})
		}
	case 65:
		//line n1ql.y:583
		{
			logDebugGrammar("TYPE JOIN KEY IDENTIFIER")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Type: Type, Keys: key_expr})

		}
	case 66:
		//line n1ql.y:592
		{
			logDebugGrammar("TYPE JOIN KEY IDENTIFIER NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Type: Type, Keys: key_expr, Over: rest})
		}
	case 67:
		//line n1ql.y:601
		{
			logDebugGrammar("TYPE JOIN KEY AS IDENTIFIER")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Type: Type, Keys: key_expr})
		}
	case 68:
		//line n1ql.y:609
		{
			logDebugGrammar("TYPE JOIN KEY AS IDENTIFIER NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Type: Type, Keys: key_expr, Over: rest})
		}
	case 69:
		//line n1ql.y:620
		{
			logDebugGrammar("FROM JOIN DATASOURCE with KEY")
			key := parsingStack.Pop().(ast.Expression)
			key_expr := ast.NewKeyExpression(key, "KEY")
			parsingStack.Push(key_expr)
		}
	case 70:
		//line n1ql.y:627
		{
			logDebugGrammar("FROM DATASOURCE with KEYS")
			keys := parsingStack.Pop().(ast.Expression)
			keys_expr := ast.NewKeyExpression(keys, "KEYS")
			parsingStack.Push(keys_expr)

		}
	case 71:
		//line n1ql.y:636
		{
			logDebugGrammar("INNER")
			parsingStack.Push("INNER")
		}
	case 72:
		//line n1ql.y:641
		{
			logDebugGrammar("OUTER")
			parsingStack.Push("LEFT")
		}
	case 73:
		//line n1ql.y:646
		{
			logDebugGrammar("LEFT OUTER")
			parsingStack.Push("LEFT")
		}
	case 74:
		//line n1ql.y:653
		{
			logDebugGrammar("FROM DATASOURCE")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj})
		}
	case 75:
		//line n1ql.y:659
		{
			logDebugGrammar("FROM KEY(S) DATASOURCE")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj})
		}
	case 76:
		//line n1ql.y:665
		{
			// fixme support over as
			logDebugGrammar("FROM DATASOURCE AS ID")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
		}
	case 77:
		//line n1ql.y:672
		{
			// fixme support over as
			logDebugGrammar("FROM DATASOURCE ID")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
		}
	case 78:
		//line n1ql.y:679
		{
			logDebugGrammar("FROM DATASOURCE AS ID KEY(S)")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s})

		}
	case 79:
		//line n1ql.y:686
		{
			logDebugGrammar("FROM DATASOURCE ID KEY(s)")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s})

		}
	case 80:
		//line n1ql.y:695
		{
			logDebugGrammar("FROM DATASOURCE with KEY")
			keys := parsingStack.Pop().(ast.Expression)
			switch parsingStatement := parsingStatement.(type) {
			case *ast.SelectStatement:
				parsingStatement.Keys = ast.NewKeyExpression(keys, "KEY")
			default:
				logDebugGrammar("This statement does not support KEY")
			}
		}
	case 81:
		//line n1ql.y:706
		{
			logDebugGrammar("FROM DATASOURCE with KEYS")
			keys := parsingStack.Pop().(ast.Expression)
			switch parsingStatement := parsingStatement.(type) {
			case *ast.SelectStatement:
				parsingStatement.Keys = ast.NewKeyExpression(keys, "KEYS")
			default:
				logDebugGrammar("This statement does not support KEYS")
			}
		}
	case 82:
		//line n1ql.y:720
		{
			logDebugGrammar("SELECT WHERE - EMPTY")
		}
	case 83:
		//line n1ql.y:724
		{
			logDebugGrammar("SELECT WHERE - EXPR")
			where_part := parsingStack.Pop().(ast.Expression)
			switch parsingStatement := parsingStatement.(type) {
			case *ast.SelectStatement:
				parsingStatement.Where = where_part
			default:
				logDebugGrammar("This statement does not support WHERE")
			}
		}
	case 85:
		//line n1ql.y:738
		{

		}
	case 86:
		//line n1ql.y:744
		{

		}
	case 87:
		//line n1ql.y:748
		{

		}
	case 88:
		//line n1ql.y:753
		{
			logDebugGrammar("SORT EXPR")
			expr := parsingStack.Pop()
			switch parsingStatement := parsingStatement.(type) {
			case *ast.SelectStatement:
				parsingStatement.OrderBy = append(parsingStatement.OrderBy, ast.NewSortExpression(expr.(ast.Expression), true))
			default:
				logDebugGrammar("This statement does not support ORDER BY")
			}
		}
	case 89:
		//line n1ql.y:764
		{
			logDebugGrammar("SORT EXPR ASC")
			expr := parsingStack.Pop()
			switch parsingStatement := parsingStatement.(type) {
			case *ast.SelectStatement:
				parsingStatement.OrderBy = append(parsingStatement.OrderBy, ast.NewSortExpression(expr.(ast.Expression), true))
			default:
				logDebugGrammar("This statement does not support ORDER BY")
			}
		}
	case 90:
		//line n1ql.y:775
		{
			logDebugGrammar("SORT EXPR DESC")
			expr := parsingStack.Pop()
			switch parsingStatement := parsingStatement.(type) {
			case *ast.SelectStatement:
				parsingStatement.OrderBy = append(parsingStatement.OrderBy, ast.NewSortExpression(expr.(ast.Expression), false))
			default:
				logDebugGrammar("This statement does not support ORDER BY")
			}
		}
	case 91:
		//line n1ql.y:787
		{

		}
	case 92:
		//line n1ql.y:791
		{

		}
	case 93:
		//line n1ql.y:795
		{

		}
	case 94:
		//line n1ql.y:801
		{
			logDebugGrammar("LIMIT %d", yyS[yypt-0].n)
			if yyS[yypt-0].n < 0 {
				panic("LIMIT cannot be negative")
			}
			switch parsingStatement := parsingStatement.(type) {
			case *ast.SelectStatement:
				parsingStatement.Limit = yyS[yypt-0].n
			default:
				logDebugGrammar("This statement does not support LIMIT")
			}
		}
	case 95:
		//line n1ql.y:815
		{
			logDebugGrammar("OFFSET %d", yyS[yypt-0].n)
			if yyS[yypt-0].n < 0 {
				panic("OFFSET cannot be negative")
			}
			switch parsingStatement := parsingStatement.(type) {
			case *ast.SelectStatement:
				parsingStatement.Offset = yyS[yypt-0].n
			default:
				logDebugGrammar("This statement does not support OFFSET")
			}
		}
	case 96:
		//line n1ql.y:832
		{
			logDebugGrammar("EXPRESSION")
		}
	case 97:
		//line n1ql.y:837
		{
			logDebugGrammar("EXPR - PLUS")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 98:
		//line n1ql.y:845
		{
			logDebugGrammar("EXPR - MINUS")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 99:
		//line n1ql.y:853
		{
			logDebugGrammar("EXPR - MULT")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 100:
		//line n1ql.y:861
		{
			logDebugGrammar("EXPR - DIV")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 101:
		//line n1ql.y:869
		{
			logDebugGrammar("EXPR - MOD")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 102:
		//line n1ql.y:877
		{
			logDebugGrammar("EXPR - CONCAT")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 103:
		//line n1ql.y:885
		{
			logDebugGrammar("EXPR - AND")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewAndOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
			parsingStack.Push(thisExpression)
		}
	case 104:
		//line n1ql.y:893
		{
			logDebugGrammar("EXPR - OR")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewOrOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
			parsingStack.Push(thisExpression)
		}
	case 105:
		//line n1ql.y:901
		{
			logDebugGrammar("EXPR - EQ")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 106:
		//line n1ql.y:909
		{
			logDebugGrammar("EXPR - LT")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 107:
		//line n1ql.y:917
		{
			logDebugGrammar("EXPR - LTE")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 108:
		//line n1ql.y:925
		{
			logDebugGrammar("EXPR - GT")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 109:
		//line n1ql.y:933
		{
			logDebugGrammar("EXPR - GTE")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 110:
		//line n1ql.y:941
		{
			logDebugGrammar("EXPR - NE")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 111:
		//line n1ql.y:949
		{
			logDebugGrammar("EXPR - LIKE")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 112:
		//line n1ql.y:957
		{
			logDebugGrammar("EXPR - NOT LIKE")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 113:
		//line n1ql.y:965
		{
			logDebugGrammar("EXPR DOT MEMBER")
			right := ast.NewProperty(yyS[yypt-0].s)
			left := parsingStack.Pop()
			thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
			parsingStack.Push(thisExpression)
		}
	case 114:
		//line n1ql.y:973
		{
			logDebugGrammar("EXPR BRACKET MEMBER")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 115:
		//line n1ql.y:981
		{
			logDebugGrammar("EXPR COLON EXPR SLICE BRACKET MEMBER")
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-3].n)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
			parsingStack.Push(thisExpression)
		}
	case 116:
		//line n1ql.y:988
		{
			logDebugGrammar("EXPR COLON SLICE BRACKET MEMBER")
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-2].n)), ast.NewLiteralNumber(float64(0)))
			parsingStack.Push(thisExpression)

		}
	case 117:
		//line n1ql.y:996
		{
			logDebugGrammar("COLON EXPR SLICE BRACKET MEMBER")
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(0)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
			parsingStack.Push(thisExpression)
		}
	case 118:
		//line n1ql.y:1003
		{
			logDebugGrammar("SUFFIX_EXPR IS NULL")
			operand := parsingStack.Pop()
			thisExpression := ast.NewIsNullOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 119:
		//line n1ql.y:1010
		{
			logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
			operand := parsingStack.Pop()
			thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 120:
		//line n1ql.y:1017
		{
			logDebugGrammar("SUFFIX_EXPR IS MISSING")
			operand := parsingStack.Pop()
			thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 121:
		//line n1ql.y:1024
		{
			logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
			operand := parsingStack.Pop()
			thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 122:
		//line n1ql.y:1031
		{
			logDebugGrammar("SUFFIX_EXPR IS VALUED")
			operand := parsingStack.Pop()
			thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 123:
		//line n1ql.y:1038
		{
			logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
			operand := parsingStack.Pop()
			thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 124:
		//line n1ql.y:1045
		{

		}
	case 125:
		//line n1ql.y:1051
		{
			logDebugGrammar("EXPR - NOT")
			operand := parsingStack.Pop()
			thisExpression := ast.NewNotOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 126:
		//line n1ql.y:1058
		{
			logDebugGrammar("EXPR - CHANGE SIGN")
			operand := parsingStack.Pop()
			thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 127:
		//line n1ql.y:1065
		{

		}
	case 128:
		//line n1ql.y:1070
		{
			logDebugGrammar("SUFFIX_EXPR")
		}
	case 129:
		//line n1ql.y:1076
		{
			logDebugGrammar("IDENTIFIER - %s", yyS[yypt-0].s)
			thisExpression := ast.NewProperty(yyS[yypt-0].s)
			parsingStack.Push(thisExpression)
		}
	case 130:
		//line n1ql.y:1082
		{
			logDebugGrammar("LITERAL")
		}
	case 131:
		//line n1ql.y:1086
		{
			logDebugGrammar("NESTED EXPR")
		}
	case 132:
		//line n1ql.y:1090
		{
			logDebugGrammar("CASE WHEN THEN ELSE END")
			cwtee := ast.NewCaseOperator()
			topStack := parsingStack.Pop()
			switch topStack := topStack.(type) {
			case ast.Expression:
				cwtee.Else = topStack
				// now look for whenthens
				nextStack := parsingStack.Pop().([]*ast.WhenThen)
				cwtee.WhenThens = nextStack
			case []*ast.WhenThen:
				// no else
				cwtee.WhenThens = topStack
			}
			parsingStack.Push(cwtee)
		}
	case 133:
		//line n1ql.y:1107
		{
			logDebugGrammar("ANY SATISFIES")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			collectionAny := ast.NewCollectionAnyOperator(condition, sub, "")
			parsingStack.Push(collectionAny)
		}
	case 134:
		//line n1ql.y:1115
		{
			logDebugGrammar("ANY IN SATISFIES")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			collectionAny := ast.NewCollectionAnyOperator(condition, sub, yyS[yypt-5].s)
			parsingStack.Push(collectionAny)
		}
	case 135:
		//line n1ql.y:1123
		{
			logDebugGrammar("ANY IN SATISFIES")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			collectionAny := ast.NewCollectionAllOperator(condition, sub, yyS[yypt-5].s)
			parsingStack.Push(collectionAny)
		}
	case 136:
		//line n1ql.y:1131
		{
			logDebugGrammar("ANY SATISFIES")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			collectionAny := ast.NewCollectionAllOperator(condition, sub, "")
			parsingStack.Push(collectionAny)
		}
	case 137:
		//line n1ql.y:1139
		{
			logDebugGrammar("FIRST FOR IN WHEN")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionFirst := ast.NewCollectionFirstOperator(condition, sub, yyS[yypt-5].s, output)
			parsingStack.Push(collectionFirst)
		}
	case 138:
		//line n1ql.y:1148
		{
			logDebugGrammar("FIRST IN WHEN")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionFirst := ast.NewCollectionFirstOperator(condition, sub, "", output)
			parsingStack.Push(collectionFirst)
		}
	case 139:
		//line n1ql.y:1157
		{
			logDebugGrammar("FIRST FOR IN")
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionFirst := ast.NewCollectionFirstOperator(nil, sub, yyS[yypt-3].s, output)
			parsingStack.Push(collectionFirst)
		}
	case 140:
		//line n1ql.y:1165
		{
			logDebugGrammar("FIRST IN")
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionFirst := ast.NewCollectionFirstOperator(nil, sub, "", output)
			parsingStack.Push(collectionFirst)
		}
	case 141:
		//line n1ql.y:1173
		{
			logDebugGrammar("ARRAY FOR IN WHEN")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionArray := ast.NewCollectionArrayOperator(condition, sub, yyS[yypt-5].s, output)
			parsingStack.Push(collectionArray)
		}
	case 142:
		//line n1ql.y:1182
		{
			logDebugGrammar("ARRAY IN WHEN")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionArray := ast.NewCollectionArrayOperator(condition, sub, "", output)
			parsingStack.Push(collectionArray)
		}
	case 143:
		//line n1ql.y:1191
		{
			logDebugGrammar("ARRAY FOR IN")
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionArray := ast.NewCollectionArrayOperator(nil, sub, yyS[yypt-3].s, output)
			parsingStack.Push(collectionArray)
		}
	case 144:
		//line n1ql.y:1199
		{
			logDebugGrammar("ARRAY IN")
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionArray := ast.NewCollectionArrayOperator(nil, sub, "", output)
			parsingStack.Push(collectionArray)
		}
	case 145:
		//line n1ql.y:1207
		{
			logDebugGrammar("FUNCTION EXPR NOPARAM")
			thisExpression := ast.NewFunctionCall(yyS[yypt-2].s, ast.FunctionArgExpressionList{})
			parsingStack.Push(thisExpression)
		}
	case 146:
		//line n1ql.y:1213
		{
			logDebugGrammar("FUNCTION EXPR PARAM")
			funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
			thisExpression := ast.NewFunctionCall(yyS[yypt-3].s, funarg_exp_list)
			parsingStack.Push(thisExpression)
		}
	case 147:
		//line n1ql.y:1220
		{
			logDebugGrammar("FUNCTION DISTINCT EXPR PARAM")
			funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
			function := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
			function.SetDistinct(true)
			parsingStack.Push(function)
		}
	case 148:
		//line n1ql.y:1228
		{
			logDebugGrammar("FUNCTION EXPR PARAM")
			funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
			thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
			parsingStack.Push(thisExpression)
		}
	case 149:
		//line n1ql.y:1237
		{
			logDebugGrammar("THEN_LIST - SINGLE")
			when_then_list := make([]*ast.WhenThen, 0)
			when_then := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
			when_then_list = append(when_then_list, &when_then)
			parsingStack.Push(when_then_list)
		}
	case 150:
		//line n1ql.y:1245
		{
			logDebugGrammar("THEN_LIST - COMPOUND")
			rest := parsingStack.Pop().([]*ast.WhenThen)
			last := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
			new_list := make([]*ast.WhenThen, 0, len(rest)+1)
			new_list = append(new_list, &last)
			for _, v := range rest {
				new_list = append(new_list, v)
			}
			parsingStack.Push(new_list)
		}
	case 151:
		//line n1ql.y:1259
		{
			logDebugGrammar("ELSE - EMPTY")
		}
	case 152:
		//line n1ql.y:1263
		{
			logDebugGrammar("ELSE - EXPR")
		}
	case 153:
		//line n1ql.y:1269
		{
			logDebugGrammar("PATH - %v", yyS[yypt-0].s)
			thisExpression := ast.NewProperty(yyS[yypt-0].s)
			parsingStack.Push(thisExpression)
		}
	case 154:
		//line n1ql.y:1275
		{
			logDebugGrammar("PATH BRACKET - %v[%v]", yyS[yypt-3].s, yyS[yypt-1].n)
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
			parsingStack.Push(thisExpression)
		}
	case 155:
		//line n1ql.y:1282
		{
			logDebugGrammar("PATH SLICE BRACKET MEMBER - %v[%v-%v]", yyS[yypt-5].s, yyS[yypt-3].n, yyS[yypt-1].n)
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-3].n)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
			parsingStack.Push(thisExpression)
		}
	case 156:
		//line n1ql.y:1289
		{
			logDebugGrammar("PATH SLICE BRACKET MEMBER - %v[%v:]", yyS[yypt-4].s, yyS[yypt-2].n)
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-2].n)), ast.NewLiteralNumber(float64(0)))
			parsingStack.Push(thisExpression)

		}
	case 157:
		//line n1ql.y:1297
		{
			logDebugGrammar("PATH SLICE BRACKET MEMBER -%v[:%v]", yyS[yypt-4].s, yyS[yypt-1].n)
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(0)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
			parsingStack.Push(thisExpression)
		}
	case 158:
		//line n1ql.y:1304
		{
			logDebugGrammar("PATH DOT PATH - $1.s")
			right := ast.NewProperty(yyS[yypt-0].s)
			left := parsingStack.Pop()
			thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
			parsingStack.Push(thisExpression)
		}
	case 159:
		//line n1ql.y:1315
		{
			funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
			parsingStack.Push(ast.FunctionArgExpressionList{funarg_expr})
		}
	case 160:
		//line n1ql.y:1320
		{
			funarg_expr_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
			funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
			// list items pushed onto the stack end up in reverse order
			// this prepends items in the list to restore order
			new_list := ast.FunctionArgExpressionList{funarg_expr}
			for _, v := range funarg_expr_list {
				new_list = append(new_list, v)
			}
			parsingStack.Push(new_list)
		}
	case 161:
		//line n1ql.y:1334
		{
			logDebugGrammar("FUNARG STAR")
		}
	case 162:
		//line n1ql.y:1338
		{
			logDebugGrammar("FUNARG EXPR")
			expr_part := parsingStack.Pop().(ast.Expression)
			funarg_expr := ast.NewFunctionArgExpression(expr_part)
			parsingStack.Push(funarg_expr)
		}
	case 163:
		//line n1ql.y:1347
		{
			logDebugGrammar("FUNSTAR")
			funarg_expr := ast.NewStarFunctionArgExpression()
			parsingStack.Push(funarg_expr)
		}
	case 164:
		//line n1ql.y:1353
		{
			logDebugGrammar("FUN PATH DOT STAR")
			expr_part := parsingStack.Pop().(ast.Expression)
			funarg_expr := ast.NewDotStarFunctionArgExpression(expr_part)
			parsingStack.Push(funarg_expr)
		}
	case 165:
		//line n1ql.y:1363
		{
			logDebugGrammar("STRING %s", yyS[yypt-0].s)
			thisExpression := ast.NewLiteralString(yyS[yypt-0].s)
			parsingStack.Push(thisExpression)
		}
	case 166:
		//line n1ql.y:1369
		{
			logDebugGrammar("NUMBER")
		}
	case 167:
		//line n1ql.y:1373
		{
			logDebugGrammar("OBJECT")
		}
	case 168:
		//line n1ql.y:1377
		{
			logDebugGrammar("ARRAY")
		}
	case 169:
		//line n1ql.y:1381
		{
			logDebugGrammar("TRUE")
			thisExpression := ast.NewLiteralBool(true)
			parsingStack.Push(thisExpression)
		}
	case 170:
		//line n1ql.y:1387
		{
			logDebugGrammar("FALSE")
			thisExpression := ast.NewLiteralBool(false)
			parsingStack.Push(thisExpression)
		}
	case 171:
		//line n1ql.y:1393
		{
			logDebugGrammar("NULL")
			thisExpression := ast.NewLiteralNull()
			parsingStack.Push(thisExpression)
		}
	case 172:
		//line n1ql.y:1401
		{
			logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
			thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
			parsingStack.Push(thisExpression)
		}
	case 173:
		//line n1ql.y:1407
		{
			logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
			thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
			parsingStack.Push(thisExpression)
		}
	case 174:
		//line n1ql.y:1415
		{
			logDebugGrammar("EMPTY OBJECT")
			emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
			parsingStack.Push(emptyObject)
		}
	case 175:
		//line n1ql.y:1421
		{
			logDebugGrammar("OBJECT")
		}
	case 176:
		//line n1ql.y:1427
		{
			logDebugGrammar("NAMED EXPR LIST SINGLE")
		}
	case 177:
		//line n1ql.y:1431
		{
			logDebugGrammar("NAMED EXPR LIST COMPOUND")
			last := parsingStack.Pop().(*ast.LiteralObject)
			rest := parsingStack.Pop().(*ast.LiteralObject)
			for k, v := range last.Val {
				rest.Val[k] = v
			}
			parsingStack.Push(rest)
		}
	case 178:
		//line n1ql.y:1443
		{
			logDebugGrammar("NAMED EXPR SINGLE")
			thisKey := yyS[yypt-2].s
			thisValue := parsingStack.Pop().(ast.Expression)
			thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
			parsingStack.Push(thisExpression)
		}
	case 179:
		//line n1ql.y:1453
		{
			logDebugGrammar("EMPTY ARRAY")
			thisExpression := ast.NewLiteralArray(ast.ExpressionList{})
			parsingStack.Push(thisExpression)
		}
	case 180:
		//line n1ql.y:1459
		{
			logDebugGrammar("ARRAY")
			exp_list := parsingStack.Pop().(ast.ExpressionList)
			thisExpression := ast.NewLiteralArray(exp_list)
			parsingStack.Push(thisExpression)
		}
	case 181:
		//line n1ql.y:1468
		{
			logDebugGrammar("EXPRESSION LIST SINGLE")
			exp_list := make(ast.ExpressionList, 0)
			exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
			parsingStack.Push(exp_list)
		}
	case 182:
		//line n1ql.y:1475
		{
			logDebugGrammar("EXPRESSION LIST COMPOUND")
			rest := parsingStack.Pop().(ast.ExpressionList)
			last := parsingStack.Pop()
			new_list := make(ast.ExpressionList, 0, len(rest)+1)
			new_list = append(new_list, last.(ast.Expression))
			for _, v := range rest {
				new_list = append(new_list, v)
			}
			parsingStack.Push(new_list)
		}
	}
	goto yystack /* stack new state and value */
}
