package main

import (
	_ "embed"
	"fmt"
	"log"
	"sort"
	"text/template"
)

var (
	//go:embed column.tpl
	columnSrc string
	//go:embed array.tpl
	arraySrc string
	//go:embed dynamic.tpl
	dynamicSrc string
)
var (
	types            []_type
	supportedGoTypes []string
	dynamicTypes     []_type
)

const (
	typeBFloat16 = "BFloat16"
)

type _type struct {
	Size int

	ChType string
	GoType string

	SkipArray bool
}

func init() {
	for _, size := range []int{8, 16, 32, 64} {
		types = append(types, _type{
			Size:   size,
			ChType: fmt.Sprintf("Int%d", size),
			GoType: fmt.Sprintf("int%d", size),
		}, _type{
			Size:   size,
			ChType: fmt.Sprintf("UInt%d", size),
			GoType: fmt.Sprintf("uint%d", size),
		})
	}
	for _, size := range []int{32, 64} {
		types = append(types, _type{
			Size:   size,
			ChType: fmt.Sprintf("Float%d", size),
			GoType: fmt.Sprintf("float%d", size),
		})
	}
	// BFloat16 - Brain Floating Point 16
	types = append(types, _type{
		Size:      16,
		ChType:    typeBFloat16,
		GoType:    "float32", // BFloat16 uses float32 in user-facing apis.
		SkipArray: true,
	})
	sort.Slice(types, func(i, j int) bool {
		return sequenceKey(types[i].ChType) < sequenceKey(types[j].ChType)
	})

	for _, typ := range types {
		// Skip BFloat16 from supportedGoTypes to avoid conflict with float32
		if typ.ChType == typeBFloat16 {
			continue
		}
		supportedGoTypes = append(supportedGoTypes, typ.GoType)
	}

	supportedGoTypes = append(supportedGoTypes,
		"string", "[]byte", "sql.NullString",
		"int", "uint", "big.Int", "decimal.Decimal",
		"bool", "sql.NullBool",
		"time.Time", "sql.NullTime",
		"uuid.UUID",
		"netip.Addr", "net.IP", "proto.IPv6", "[16]byte",
		"orb.LineString", "orb.MultiLineString",
		"orb.MultiPolygon", "orb.Point", "orb.Polygon", "orb.Ring",
	)

	dynamicTypes = make([]_type, 0, len(types))
	for _, typ := range types {

		if typ.GoType == "uint8" {
			// Prevent conflict with []byte and []uint8
			typ.SkipArray = true
			dynamicTypes = append(dynamicTypes, typ)
			continue
		}

		if typ.ChType == typeBFloat16 {
			// Skip BFloat16 from dynamic types to prevent conflict with Float32.
			// Both Float32 and BFloat16 uses float32 Go type.
			continue
		}

		dynamicTypes = append(dynamicTypes, typ)
	}

	// Best-effort type matching for Dynamic inference
	dynamicTypes = append(dynamicTypes, []_type{
		{ChType: "String", GoType: "string"},
		{ChType: "String", GoType: "json.RawMessage"},
		{ChType: "String", GoType: "sql.NullString"},
		{ChType: "Bool", GoType: "bool"},
		{ChType: "Bool", GoType: "sql.NullBool"},
		{ChType: "DateTime64(3)", GoType: "time.Time"},
		{ChType: "DateTime64(3)", GoType: "sql.NullTime"},
		{ChType: "UUID", GoType: "uuid.UUID"},
		{ChType: "IPv6", GoType: "proto.IPv6"},
		{ChType: "LineString", GoType: "orb.LineString"},
		{ChType: "MultiLineString", GoType: "orb.MultiLineString"},
		{ChType: "MultiPolygon", GoType: "orb.MultiPolygon"},
		{ChType: "Point", GoType: "orb.Point"},
		{ChType: "Polygon", GoType: "orb.Polygon"},
		{ChType: "Ring", GoType: "orb.Ring"},
	}...)

}
func write(name string, v any, t *template.Template) error { _ = "STUB: not implemented"; return nil }

//	fmt.Println(out.String())

func main() {
	for name, tpl := range map[string]struct {
		template *template.Template
		args     any
	}{
		"column_gen":  {template.Must(template.New("column").Parse(columnSrc)), types},
		"array_gen":   {template.Must(template.New("array").Parse(arraySrc)), supportedGoTypes},
		"dynamic_gen": {template.Must(template.New("dynamic").Parse(dynamicSrc)), dynamicTypes},
	} {
		if err := write(name, tpl.args, tpl.template); err != nil {
			log.Fatal(err)
		}
	}
}

const maxByte = 1<<8 - 1

func isDigit(d byte) bool { _ = "STUB: not implemented"; return false }

func sequenceKey(key string) string { _ = "STUB: not implemented"; return "" }
