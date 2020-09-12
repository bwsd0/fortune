// fortune samples lines from a file
package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
)

var (
	fortunes = filepath.Join(os.Getenv("HOME"), os.Getenv("FORTUNES"))

	fInfo, ixInfo os.FileInfo
	fbuf          *buf
	off           = make([]byte, 4)
	choice        []byte
)

const (
	startBufSize = 2048 // Size of initial allocation for buffer
)

type buf struct {
	rdline int   // number of bytes after rdline
	offset int64 // offset of buffer in file
	bbuf   []byte
	rdr    *bufio.Reader
	fd     *os.File
}

// Binit initializes a standard size buffer with the open file descriptor passed
// by the caller.
func Binit(fd *os.File) *buf {
	return &buf{
		rdline: 0,
		offset: 0,
		bbuf:   make([]byte, startBufSize),
		rdr:    bufio.NewReaderSize(fd, startBufSize),
		fd:     fd,
	}
}

// Boffset returns the current file offset.
func (b *buf) Boffset() uint32 {
	off, err := b.fd.Seek(0, 1)
	if err != nil {
		log.Fatal(err)
	}
	off += int64(len(b.bbuf))
	if int64(uint32(off)) != off {
		log.Fatal("fortune: index >= 4GB")
	}

	return uint32(off)
}

// Brdline reads a string from the file associated with b up to and including
// the first delim character.
// The delimiter character at the end of the line is not altered. Blinelen
// returns the length (including the delimiter) of the most recent string
// returned by Brdline.
func (b *buf) Brdline(delim byte) ([]byte, int) {
	p, err := b.rdr.ReadBytes(delim)
	if err != nil {
		if err == io.EOF {
			return nil, 0
		}
		log.Fatal(err)
	}
	n := copy(b.bbuf, p)
	b.rdline = n
	b.offset += int64(n)
	return p, n
}

// Blinelen returns the length (including the delimiter) of the most recent
// string returned by Brdline.
func (b *buf) Blinelen() int {
	return b.rdline
}

// puint32 returns the byte sequence of a little-endian 32-bit unsigned integer
func puint32(b []byte, v uint32) []byte {
	_ = b[3] // bounds check
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	return b
}

func main() {
	log.SetFlags(0)
	flag.Parse()
	args := flag.Args()

	if len(args) > 0 {
		log.Printf("Misfortune!\n")
		os.Exit(2)
	}

	f, err := os.Open(fortunes)
	if err != nil {
		log.Fatalf("Misfortune!\n%s", err)
	}
	defer f.Close()

	fbuf = Binit(f)
	rand.Seed(int64(os.Getpid()))
	for i := 1; ; i++ {
		offs := fbuf.Boffset()
		p, n := fbuf.Brdline('\n')
		if n == 0 {
			break
		}
		p = p[:fbuf.Blinelen()-1]
		off = puint32(off, offs)
		if rand.Int31()%int32(i) == 0 {
			choice = p
		}
	}

	log.Printf("%s", choice)
	os.Exit(0)
}
