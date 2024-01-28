package hashcash

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/bits"
	"strconv"
	"strings"
	"time"
)

const version = 1
const dateFormat = "060102" // no RFC for this one
const segmentsLen = 7
const lifetimeHours = 48

const DefaultZeroBits = 20
const DefaultSaltLen = 8

type Storage interface {
	Exists(header string) bool
	Push(header string)
}

type Hash struct {
	zeroBits int
	saltLen  int
	ext      string
}

func New(zeroBits int, saltLen int, ext string) Hash {
	return Hash{
		zeroBits: zeroBits,
		saltLen:  saltLen,
		ext:      ext,
	}
}

func Make() Hash {
	return New(DefaultZeroBits, DefaultSaltLen, "")
}

func (h Hash) Mint(resource string) (string, error) {
	dateString := time.Now().Format(dateFormat)
	saltString, err := salt(h.saltLen)

	if err != nil {
		return "", err
	}

	counter := 0

	for {
		header := fmt.Sprintf(
			"%d:%d:%s:%s:%s:%s:%s",
			version,
			h.zeroBits,
			dateString,
			resource,
			h.ext,
			saltString,
			intToBase64(counter),
		)

		hashSum := sha1.Sum(([]byte)(header))
		if containsLeadZeros(h.zeroBits, hashSum) {
			return header, nil
		}

		counter += 1
	}
}

func Verify(header string) bool {
	segments := strings.Split(header, ":")

	if len(segments) != segmentsLen {
		return false
	}

	date, err := time.Parse(dateFormat, segments[2])
	if err != nil {
		return false
	}

	diff := time.Since(date)
	if diff.Hours() >= lifetimeHours {
		return false
	}

	zeroBits, err := strconv.ParseInt(segments[1], 10, 32)
	if err != nil {
		return false
	}

	hashSum := sha1.Sum(([]byte)(header))
	if !containsLeadZeros(int(zeroBits), hashSum) {
		return false
	}

	return true
}

func VerifyWithStorage(header string, storage Storage) bool {
	if !Verify(header) {
		return false
	}

	if storage.Exists(header) {
		return false
	}

	storage.Push(header)
	return true
}

func containsLeadZeros(zeros int, s [20]byte) bool {
	count := 0
	for _, b := range s {
		tmp := bits.LeadingZeros8(b)
		count += tmp
		if tmp < 8 {
			break
		}
	}
	return count >= zeros
}

func salt(size int) (string, error) {
	random, err := randomBytes(size)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(random), nil
}

func randomBytes(size int) (buf []byte, err error) {
	buf = make([]byte, size)
	_, err = rand.Read(buf)

	return buf, err
}

func intToBase64(n int) string {
	return base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(n)))
}
