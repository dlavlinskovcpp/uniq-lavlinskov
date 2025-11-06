// options.go
package uniq

import (
	"errors"
	"flag"
	"io"
	"os"
)

type Options struct {
	Count      bool
	Dups       bool
	Uniq       bool
	IgnoreCase bool
	SkipFields int
	SkipChars  int
}

func (o *Options) Validate() error {
	if o.Count && o.Dups || o.Count && o.Uniq || o.Dups && o.Uniq {
		return errors.New("флаги -c, -d и -u нельзя использовать вместе")
	}
	if o.SkipFields < 0 {
		return errors.New("значение -f должно быть неотрицательным")
	}
	if o.SkipChars < 0 {
		return errors.New("значение -s должно быть неотрицательным")
	}
	return nil
}

func Usage() string {
	return `uniq [опции] [вход [выход]]
Фильтрует повторяющиеся строки

Опции:
  -c    добавить счётчик появлений перед строкой
  -d    вывести только повторяющиеся строки
  -u    вывести только уникальные строки  
  -i    сравнивать без учёта регистра
  -f N  пропустить первые N полей
  -s N  после полей пропустить N символов

Примеры:
  uniq file.txt
  uniq -c input.txt output.txt  
  cat file.txt | uniq -d
`
}

func ParseFlags() (*Options, []string, error) {
	o := &Options{}
	fs := flag.NewFlagSet("uniq", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	fs.BoolVar(&o.Count, "c", false, "добавить счётчик")
	fs.BoolVar(&o.Dups, "d", false, "только дубликаты")
	fs.BoolVar(&o.Uniq, "u", false, "только уникальные")
	fs.BoolVar(&o.IgnoreCase, "i", false, "игнорировать регистр")
	fs.IntVar(&o.SkipFields, "f", 0, "пропустить поля")
	fs.IntVar(&o.SkipChars, "s", 0, "пропустить символы")

	if err := fs.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	return o, fs.Args(), nil
}
