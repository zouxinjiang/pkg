package arg_parser

import (
	"github.com/spf13/cobra"
	"reflect"
	"strconv"
	"strings"
)

type tag struct {
	data string
	temp map[string]string
}

func (t *tag) prepare() {
	if len(t.temp) != 0 {
		return
	}
	if t.temp == nil {
		t.temp = map[string]string{}
	}
	items := strings.Split(t.data, ";")
	for _, v := range items {
		kv := strings.SplitN(v, ":", 2)
		if len(kv) < 2 {
			continue
		}
		t.temp[strings.ToLower(kv[0])] = kv[1]
	}
}

func (t *tag) GetName() string {
	t.prepare()
	return t.temp["name"]
}

func (t tag) GetDefault() string {
	t.prepare()
	return t.temp["default"]
}

func (t *tag) GetUsage() string {
	t.prepare()
	return t.temp["usage"]
}

func InitCobraFlag(cmd *cobra.Command, args interface{}) {
	typ := reflect.TypeOf(args)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return
	}
	cnt := typ.NumField()
	for i := 0; i < cnt; i++ {
		f := typ.Field(i)
		tag := tag{
			data: f.Tag.Get("cmd"),
		}
		name := tag.GetName()
		if name == "" {
			name = strings.ToLower(f.Name)
		}
		switch f.Type.Kind() {
		case reflect.Int, reflect.Uint, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val, _ := strconv.ParseInt(tag.GetDefault(), 10, 64)
			cmd.Flags().Int(name, int(val), tag.GetUsage())
		case reflect.Bool:
			val, _ := strconv.ParseBool(tag.GetDefault())
			cmd.Flags().Bool(name, val, tag.GetUsage())
		case reflect.String:
			cmd.Flags().String(name, tag.GetDefault(), tag.GetUsage())
		}
	}
}

func ParseCobraFlag(cmd *cobra.Command, args interface{}) {
	typ := reflect.TypeOf(args)
	val := reflect.ValueOf(args)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return
	}
	cnt := typ.NumField()
	for i := 0; i < cnt; i++ {
		f := typ.Field(i)
		fv := val.Field(i)
		if !fv.CanSet() {
			continue
		}
		tag := tag{
			data: f.Tag.Get("cmd"),
		}
		name := tag.GetName()
		if name == "" {
			name = strings.ToLower(f.Name)
		}
		switch f.Type.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			cmdval, _ := cmd.Flags().GetInt64(name)
			val.SetInt(cmdval)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			cmdval, _ := cmd.Flags().GetUint64(name)
			val.SetUint(cmdval)
		case reflect.Bool:
			cmdval, _ := cmd.Flags().GetBool(name)
			val.SetBool(cmdval)
		case reflect.String:
			cmdval, _ := cmd.Flags().GetString(name)
			val.SetString(cmdval)
		}
	}
}
