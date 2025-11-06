package uniq

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type processor struct {
	opts     *Options
	lastKey  string
	lastLine string
	count    int
	first    bool
}

func Run(r io.Reader, w io.Writer, o *Options) error {
	p := &processor{opts: o, first: true}
	scanner := bufio.NewScanner(r)

	const maxCapacity = 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		line := strings.TrimRightFunc(scanner.Text(), unicode.IsSpace)
		p.processLine(line, w)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("чтение входных данных: %w", err)
	}

	p.flush(w)
	return nil
}

func (p *processor) processLine(line string, w io.Writer) {
	key := p.normalize(line)

	if p.first {
		p.first = false
		p.lastKey = key
		p.lastLine = line
		p.count = 1
		return
	}

	if key == p.lastKey {
		p.count++
		return
	}

	p.flush(w)
	p.lastKey = key
	p.lastLine = line
	p.count = 1
}

func (p *processor) flush(w io.Writer) {
	if p.first {
		return
	}

	if p.opts.Dups && p.count < 2 {
		return
	}
	if p.opts.Uniq && p.count > 1 {
		return
	}

	if p.opts.Count {
		fmt.Fprintf(w, "%d %s\n", p.count, p.lastLine)
	} else {
		fmt.Fprintf(w, "%s\n", p.lastLine)
	}
}

func (p *processor) normalize(s string) string {
	if p.opts.IgnoreCase {
		s = strings.ToLower(s)
	}

	if p.opts.SkipFields > 0 {
		fields := strings.Fields(s)
		if len(fields) > p.opts.SkipFields {
			s = strings.Join(fields[p.opts.SkipFields:], " ")
		} else {
			s = ""
		}
	}

	if p.opts.SkipChars > 0 {
		runes := []rune(s)
		if len(runes) > p.opts.SkipChars {
			s = string(runes[p.opts.SkipChars:])
		} else {
			s = ""
		}
	}

	return s
}
