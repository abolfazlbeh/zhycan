package utils

import (
	"go/ast"
	"reflect"
	"time"
)

// GenerateStructType - convert ast.StructType to reflect.Type
func GenerateStructType(structType *ast.StructType) reflect.Type {
	fields := make([]reflect.StructField, 0)
	for _, field := range structType.Fields.List {
		for _, fieldName := range field.Names {
			fieldType := field.Type
			fieldTag := reflect.StructTag("")
			if field.Tag != nil {
				fieldTag = reflect.StructTag(field.Tag.Value)
			}

			fields = append(fields, reflect.StructField{
				Name: fieldName.Name,
				Type: ParseFieldType(fieldType),
				Tag:  fieldTag,
			})
		}
	}

	return reflect.StructOf(fields)
}

// ParseFieldType - parse ast.Expr to reflect.Type
func ParseFieldType(expr ast.Expr) reflect.Type {
	// Implement parsing of various field types
	// For simplicity, let's assume all fields are of type string for now
	switch fieldType := expr.(type) {
	case *ast.Ident:
		switch fieldType.Name {
		case "int":
			return reflect.TypeOf(int(0))
		case "int8":
			return reflect.TypeOf(int8(0))
		case "int64":
			return reflect.TypeOf(int64(0))
		case "int32":
			return reflect.TypeOf(int32(0))
		case "string":
			return reflect.TypeOf("")
		case "float32":
			return reflect.TypeOf(float32(0.0))
		case "float64":
			return reflect.TypeOf(float64(0.0))
		case "uint":
			return reflect.TypeOf(uint(0))
		case "uint32":
			return reflect.TypeOf(uint32(0))
		case "uint64":
			return reflect.TypeOf(uint64(0))
		case "uint8":
			return reflect.TypeOf(uint8(0))
		case "time.Time":
			return reflect.TypeOf(time.Now())
		default:
			return reflect.TypeOf(nil)
		}
	case *ast.StarExpr:
		// Pointer types
		elemType := ParseFieldType(fieldType.X)
		return reflect.PtrTo(elemType)
	case *ast.ArrayType:
		// Slice types
		elemType := ParseFieldType(fieldType.Elt)
		return reflect.SliceOf(elemType)
	}
	return reflect.TypeOf(nil)
}
