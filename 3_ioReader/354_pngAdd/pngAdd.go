package main

import (
	"bytes"
	"encoding/binary"
	"hash/crc32"
	"io"
	"os"
)

/*
   PNG format
               | Chunk                          | Chunk | ...
   | Signature | Data len  | kind | data  | CRC |
   | 8 bytes   | 4 bytes   | 4    | <len> | 4   |
*/

func textChunk(text string) io.Reader {
	data := []byte(text)
	var chunk bytes.Buffer

	// Write Data len and Chunk kind
	binary.Write(&chunk, binary.BigEndian, int32(len(data)))
	chunk.WriteString("tEXt")
	chunk.Write(data)

	// Write CRC
	crc := crc32.NewIEEE()
	io.WriteString(crc, "tEXt")
	crc.Write(data)
	binary.Write(&chunk, binary.BigEndian, crc.Sum32())

	return &chunk
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
	file, err := os.Open("Lenna.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	newFile, err := os.Create("Lenna2.png")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	chunks := readChunks(file)
	// Write Signature and IHDR chunk
	io.WriteString(newFile, "\x89PNG\r\n\x1a\n")
	io.Copy(newFile, chunks[0])

	// Write my chunk
	io.Copy(newFile, textChunk("This is my tEXt chuunuuk!"))

	// Write remaining chunks
	for _, chunk := range chunks[1:] {
		io.Copy(newFile, chunk)
	}
}
