package swgen

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
)

var camelingRegex = regexp.MustCompile("[0-9A-Za-z]+")

// ReflectTypeHash returns private (unexported) `hash` field of the Golang internal reflect.rtype struct for a given reflect.Type
// This hash is used to (quasi-)uniquely identify a reflect.Type value
func ReflectTypeHash(t reflect.Type) uint32 {
	return uint32(reflect.Indirect(reflect.ValueOf(t)).FieldByName("hash").Uint())
}

// ReflectTypeReliableName returns real name of given reflect.Type, if it is non-empty, or auto-generates "anon_*"]
// name for anonymous structs
func ReflectTypeReliableName(t reflect.Type) string {
	if t.Name() != "" {
		byteSrc := []byte(t.PkgPath())
		parts := camelingRegex.FindAll(byteSrc, -1)
		if len(parts) > 2 {
			chunks := parts[len(parts)-2 : len(parts)]
			for idx, val := range chunks {
				chunks[idx] = bytes.Title(val)
			}
			return string(bytes.Join(chunks, nil)) + t.Name()
		}
		return t.Name()
	}
	return fmt.Sprintf("anon_%08x", ReflectTypeHash(t))
}
