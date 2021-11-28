package scanner

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
	"github.com/wdevore/RISCV-Meta-Assembler/src/scanner/literals"
)

type Scanner struct {
	assembler api.IAssembler

	source string
	tokens []*Token

	start   int
	current int
	line    int
}

func NewScanner(assembler api.IAssembler) *Scanner {
	s := new(Scanner)
	s.start = 0
	s.current = 0
	s.line = 1
	s.assembler = assembler
	return s
}

func (s *Scanner) Scan(source string) error {
	s.source = source

	dataPath, err := filepath.Abs(s.assembler.ConfigRelPath())

	if err != nil {
		return err
	}

	// file, err := os.Open(dataPath + "/" + source)
	// if err != nil {
	// 	return err
	// }

	// defer file.Close()

	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadFile(dataPath + "/" + source)
	if err != nil {
		return err
	}
	s.source = string(bytes)
	s.scanTokens(s.source)

	// ioScanner := bufio.NewScanner(file)
	// ioScanner.Split(bufio.ScanLines)

	// for ioScanner.Scan() {
	// 	s.source = ioScanner.Text()

	// 	if len(s.source) == 0 {
	// 		continue
	// 	}

	// 	s.scanTokens(s.source)
	// }

	for _, token := range s.tokens {
		log.Println((*token).String())
	}

	return nil
}

func (s *Scanner) scanTokens(line string) {
	for !s.isAtEnd() {
		// We are at the beginning of the next lexeme.
		s.start = s.current
		s.scanToken()
	}

	t := NewToken(EOF, "", literals.NewNilLiteral(), 1 /*line*/)
	s.tokens = append(s.tokens, t)
}

func (s *Scanner) scanToken() {
	c := s.advance()
	// log.Println(c)
	// if c == "a" {
	// 	fmt.Println(c)
	// }
	switch c {
	case "(":
		s.addTokenNullLiteral(LEFT_PAREN)
	case ")":
		s.addTokenNullLiteral(RIGHT_PAREN)
	case "{":
		s.addTokenNullLiteral(LEFT_BRACE)
	case "}":
		s.addTokenNullLiteral(RIGHT_BRACE)
	case "[":
		s.addTokenNullLiteral(LEFT_BRACKET)
	case "]":
		s.addTokenNullLiteral(RIGHT_BRACKET)
	case ",":
		s.addTokenNullLiteral(COMMA)
	case ".":
		s.addTokenNullLiteral(DOT)
	case "-":
		s.addTokenNullLiteral(MINUS)
	case "+":
		s.addTokenNullLiteral(PLUS)
	case "*":
		s.addTokenNullLiteral(STAR)
	case "%":
		s.addTokenNullLiteral(PERCENT)
	case "!":
		match := s.match("=")
		if match {
			s.addTokenNullLiteral(BANG_EQUAL)
		} else {
			s.addTokenNullLiteral(BANG)
		}
	case "=":
		match := s.match("=")
		if match {
			s.addTokenNullLiteral(EQUAL_EQUAL)
		} else {
			s.addTokenNullLiteral(EQUAL)
		}
	case "<":
		match := s.match("=")
		if match {
			s.addTokenNullLiteral(LESS_EQUAL)
		} else {
			// It could be "<42>" example or just "<"
			if !s.isDigit(s.peek()) {
				s.addTokenNullLiteral(LESS)
			}
		}
	case ">":
		match := s.match("=")
		if match {
			s.addTokenNullLiteral(GREATER_EQUAL)
		} else {
			s.addTokenNullLiteral(GREATER)
		}

	case "/":
		match := s.match("/")

		if match {
			// A comment goes until the end of the line.
			for s.peek() != "\n" && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addTokenNullLiteral(SLASH)
		}
	case " ", "\r", "\t":
		// Ignore whitespace.
	case "\n":
		s.line++
	case "\"":
		s.string()
	case "'":
		s.char()
	default:
		if s.isDigit(c) {
			s.number()
		} else if s.isAlpha(c) {
			s.identifier()
		} else {
			s.assembler.ReportLine(s.line, "Unexpected character '"+c+"'")
		}
	}
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) advance() string {
	s.current++
	return string(s.source[s.current-1])
}

func (s *Scanner) addTokenNullLiteral(ttype TokenType) {
	s.addToken(ttype, literals.NewNilLiteral())
}

func (s *Scanner) addToken(ttype TokenType, literal api.ILiteral) {
	text := s.source[s.start:s.current]
	token := NewToken(ttype, text, literal, s.line)
	s.tokens = append(s.tokens, token)
}

// We only consume the current character if it’s
// what we’re looking for.
func (s *Scanner) match(expected string) bool {
	if s.isAtEnd() {
		return false
	}
	if string(s.source[s.current]) != expected {
		return false
	}

	s.current++
	return true
}

// It’s sort of like advance() , but doesn’t consume the character.
// This is called lookahead.
func (s *Scanner) peek() string {
	if s.isAtEnd() {
		return "\n"
	}
	return string(s.source[s.current])
}

func (s *Scanner) string() {
	for s.peek() != "\"" && !s.isAtEnd() {
		if s.peek() == "\n" {
			s.line++
		}
		s.advance()
	}
	if s.isAtEnd() {
		s.assembler.ReportLine(s.line, "Unterminated string.")
		return
	}
	// The closing " character
	s.advance()

	// Trim the surrounding quotes.
	value := s.source[s.start+1 : s.current-1]
	s.addToken(STRING, literals.NewStringLiteral(value))
}

func (s *Scanner) char() {
	for s.peek() != "'" && !s.isAtEnd() {
		if s.peek() == "\n" {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		s.assembler.ReportLine(s.line, "Unterminated char.")
		return
	}

	// The closing "'"
	s.advance()

	// Trim the surrounding quotes.
	value := s.source[s.start+1 : s.current-1]

	// -2 for right "'" and advanced position
	if s.current-s.start-2 > 1 {
		s.assembler.ReportLine(s.line, "To many characters for '"+value+"'")
		return
	}

	s.addToken(STRING, literals.NewCharLiteral(value))
}

func (s *Scanner) isDigit(c string) bool {
	return c >= "0" && c <= "9"
}

func (s *Scanner) isAlpha(c string) bool {
	return (c >= "a" && c <= "z") || (c >= "A" && c <= "Z") || c == "_"
}

func (s *Scanner) isAlphaNumeric(c string) bool {
	return s.isAlpha(c) || s.isDigit(c)
}

func (s *Scanner) number() {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	// Look for a fractional part of a decimal number
	if s.peek() == "." && s.isDigit(s.peekNext()) {
		// Consume the "."
		s.advance()

		for s.isDigit(s.peek()) {
			s.advance()
		}

		value := s.source[s.start:s.current]
		s.addToken(NUMBER, literals.NewNumberLiteral(value))
		return
	}

	// Look for a base specifier "x","b".
	if s.peek() == "x" && s.isAlphaNumeric(s.peekNext()) {
		// Consume the "x"
		s.advance()

		for s.isAlphaNumeric(s.peek()) {
			s.advance()
		}

		// Trim "0x"
		value := s.source[s.start+2 : s.current]
		s.addToken(NUMBER, literals.NewHexNumberLiteral(value))
		return
	}

	if s.peek() == "b" && s.isDigit(s.peekNext()) {
		// Consume the "b"
		s.advance()

		for s.isDigit(s.peek()) {
			s.advance()
		}

		// Trim "0b"
		value := s.source[s.start+2 : s.current]
		s.addToken(NUMBER, literals.NewBinaryNumberLiteral(value))
		return
	}

	value := s.source[s.start:s.current]
	s.addToken(NUMBER, literals.NewNumberLiteral(value))
}

func (s *Scanner) identifier() {
	for s.isAlphaNumeric(s.peek()) {
		s.advance()
	}

	text := s.source[s.start:s.current]
	ttype := keywords[text]
	if ttype == UNDEFINED {
		ttype = IDENTIFIER
	}

	s.addTokenNullLiteral(ttype)
}

func (s *Scanner) peekNext() string {
	if s.current+1 >= len(s.source) {
		return "" // "\0"
	}

	return string(s.source[s.current+1])
}
