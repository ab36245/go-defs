package model

import "github.com/ab36245/go-writer"

func String(m any) string {
	return writer.Value(m)
}
