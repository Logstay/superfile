package model

import (
	"fmt"
	"log"
	"strings"

	"github.com/helloticket/superfile/field"
)

type Linha struct {
	Block        string
	Name         string
	Start        int
	End          int
	Size         int
	Picture      string
	Format       string
	Options      map[string]string
	DefaultValue interface{}
	Description  string
	Value        interface{}
}

type LinhaSorted []Linha

func (a LinhaSorted) Len() int           { return len(a) }
func (a LinhaSorted) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a LinhaSorted) Less(i, j int) bool { return a[i].Start < a[j].Start }

func (l *Linha) Clone(value interface{}) Linha {
	return Linha{
		Block:        l.Block,
		DefaultValue: l.DefaultValue,
		Description:  l.Description,
		End:          l.End,
		Format:       l.Format,
		Name:         l.Name,
		Picture:      l.Picture,
		Size:         l.Size,
		Start:        l.Start,
		Value:        value,
	}
}

func (l *Linha) Decode(d field.Picture) (interface{}, error) {
	opt := field.Options{}
	if l.Format != "" {
		opt["data_format"] = l.Format
	}

	if len(l.Options) != 0 {
		for k, v := range l.Options {
			opt[k] = v
		}
	}

	decoded, err := d.Decode(l.value(), l.Picture, opt)
	if err != nil {
		log.Printf("Parser error: %s - %v", l, err)
		return nil, fmt.Errorf("Parser error: %s - %v", l, err)
	}

	return decoded, nil
}

func (l *Linha) Encode(d field.Picture) string {
	data := l.DefaultValue
	if l.Value != nil {
		data = l.Value
	}

	opt := field.Options{}
	if l.Format != "" {
		opt["data_format"] = l.Format
	}

	if len(l.Options) != 0 {
		for k, v := range l.Options {
			opt[k] = v
		}
	}

	output, err := d.Encode(data, l.Picture, opt)
	if err != nil {
		log.Printf("Parser error: %s - %v", l, err)
	}

	return output
}

func (l *Linha) value() string {
	content, _ := l.Value.(string)
	if len(content) >= l.Start {
		end := l.End
		if end >= len(content) {
			end = len(content)
		}
		return content[l.Start-1 : end]
	}
	return ""
}

func (l *Linha) String() string {
	b := strings.Builder{}
	b.WriteString("[")
	b.WriteString("Block: ")
	b.WriteString(l.Block)
	b.WriteString(", ")
	b.WriteString("Name: ")
	b.WriteString(l.Name)
	b.WriteString(", ")
	b.WriteString("Pos: ")
	b.WriteString(fmt.Sprintf("(%d,%d)", l.Start, l.End))
	b.WriteString(", ")
	b.WriteString("Picture: ")
	b.WriteString(l.Picture)
	b.WriteString(", ")
	b.WriteString("Format: ")
	b.WriteString(l.Format)
	b.WriteString(", ")
	b.WriteString("Default: ")
	b.WriteString(fmt.Sprintf("%v", l.DefaultValue))
	b.WriteString(", ")
	b.WriteString("Value: ")
	b.WriteString(fmt.Sprintf("%v", l.Value))
	b.WriteString("]")
	return b.String()
}
