package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

/*
   PNG format
               | Chunk                          | Chunk | ...
   | Signature | Data len  | kind | data  | CRC |
   | 8 bytes   | 4 bytes   | 4    | <len> | 4   |
*/

func dumpChunk(chunk io.Reader) {
	var data_len int32
	binary.Read(chunk, binary.BigEndian, &data_len)
	chunk_kind := make([]byte, 4)
	chunk.Read(chunk_kind)
	fmt.Printf("chunk '%s' (%d bytes)\n", chunk_kind, data_len)

	if bytes.Equal(chunk_kind, []byte("tEXt")) {
		rawText := make([]byte, data_len)
		chunk.Read(rawText)
		fmt.Println(string(rawText))
	}
}

func readChunks(file *os.File) []io.Reader {
	var chunks []io.Reader
	var chunk_start int64 = 8

	// Skip PNG signature
	file.Seek(chunk_start, 0)

	for {
		var chunk_len int32
		var data_len int32
		// Data len is 4 bytes long
		err := binary.Read(file, binary.BigEndian, &data_len)
		if err == io.EOF {
			break
		}

		// Chunk is 1 data + 3 metadata(each is 4 bytes long)
		chunk_len = data_len + 12

		chunks = append(chunks,
			io.NewSectionReader(file, chunk_start, int64(chunk_len)))

		// Seek to next chunk
		// NOTE: We Read 4 bytes already.
		chunk_start, _ = file.Seek(int64(chunk_len)-4, 1)
	}

	return chunks
}

func main() {
	file, err := os.Open("Lenna2.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	chunks := readChunks(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}
