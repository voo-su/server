// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package encoding

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() any {
		return &bytes.Buffer{}
	},
}

func NewEncode(data []byte) ([]byte, error) {
	buf := bufferPool.Get().(*bytes.Buffer)
	var length = int32(len(string(data)))
	if err := binary.Write(buf, binary.LittleEndian, length); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.LittleEndian, data); err != nil {
		return nil, err
	}

	buffer := buf.Bytes()
	buf.Reset()
	bufferPool.Put(buf)

	return buffer, nil
}

func NewDecode(r io.Reader) ([]byte, error) {
	var length int32
	if err := binary.Read(r, binary.LittleEndian, &length); err != nil {
		return nil, err
	}
	if length < 0 {
		return nil, fmt.Errorf("размер ответного сообщения отрицателен: %v", length)
	}

	buf := make([]byte, length)
	if _, err := io.ReadFull(r, buf); err != nil {
		return nil, err
	}

	return buf, nil
}
