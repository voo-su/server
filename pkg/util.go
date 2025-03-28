package pkg

import (
	"bytes"
	"fmt"
	"html/template"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"math/rand"
	"reflect"
	"runtime"
	"time"
)

type ImageMeta struct {
	Width  int
	Height int
}

func ReadImageMeta(r io.Reader) *ImageMeta {
	c, _, _ := image.DecodeConfig(r)

	return &ImageMeta{c.Width, c.Height}
}

func Retry(num int, sleep time.Duration, fn func() error) error {
	var err error
	for i := 0; i < num; i++ {
		if err = fn(); err == nil {
			return nil
		}
		time.Sleep(sleep)
	}

	return err
}

func RenderTemplate(text []byte, data any) (string, error) {
	tmpl, _ := template.New("tmpl").Parse(string(text))
	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return "", err
	}

	return body.String(), nil
}

func MtRand(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min+1) + min
}

func PanicTrace(err interface{}) string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%v\n", err)
	for i := 2; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
	}

	return buf.String()
}

func IsZeroValue(val reflect.Value) bool {
	switch val.Kind() {
	case reflect.Ptr, reflect.Interface:
		return val.IsNil()
	case reflect.String:
		return val.Len() == 0
	case reflect.Array, reflect.Slice, reflect.Map:
		return val.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return val.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return val.Float() == 0
	case reflect.Bool:
		return !val.Bool()
	}

	return false
}
