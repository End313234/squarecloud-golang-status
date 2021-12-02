// Package squrecloud provides some useful functions to
// get the current bot data, such as the avaliable ram,
// the current ram usage and the ssd usage.
//
// This documentation is written only in English (EN-US).
// If you want access to the documentation written in
// Portuguese (PT-BR), access the documentation at
// https://github.com/End313234/squarecloud-golang-status
package squarecloud

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	kb = 1 << (10 * iota)
	mb
	gb
)

const (
	RAM_USAGE_PATH = "/sys/fs/cgroup/memory/memory.usage_in_bytes"
	TOTAL_RAM_PATH = "/sys/fs/cgroup/memory/memory.usage_in_bytes"
	BOT_PATH       = "/application"
)

var sizes map[string]float64 = map[string]float64{
	"KB": float64(kb),
	"MB": float64(mb),
	"GB": float64(gb),
}

func convert(originalSize float64, size string) (float64, error) {
	bytes, exists := sizes[strings.ToUpper(size)]
	if !exists {
		return 0.0, fmt.Errorf("size '%s' does not exist", size)
	}
	return originalSize / bytes, nil
}

func readFile(path string, size string) (float64, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, errors.New("could not open file")
	}

	parsedData, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, fmt.Errorf("can not convert '%s' to int", string(data))
	}

	return convert(float64(parsedData), size)
}

// UsedRam returns the current ram usage.
// This method has support for some multiples of the unity
// byte, for example, kB (kilobyte) MB (megabyte) and GB
// (gigabyte). You only have to specify the multiple like that:
//
//		 squarecloud.RamUsage("MB") // 23
//		 squarecloud.RamUsage("KB") // 23000
func RamUsage(size string) (float64, error) {
	bytes, err := readFile(RAM_USAGE_PATH, size)
	return bytes, err
}

// TotalRam returns the avaliable ram.
// This method also have support for the same multiples
// of the unity byte as RamUsage.
func TotalRam(size string) (float64, error) {
	bytes, err := readFile(TOTAL_RAM_PATH, size)
	return bytes, err
}

// SSD returns the ssd usage.
// This method was made based in
// https://stackoverflow.com/a/32482941.
// It also have support for multiples of byte
//
func SSD(size string) (float64, error) {
	var byteSize int64

	err := filepath.Walk(BOT_PATH, func(_ string, data os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !data.IsDir() {
			byteSize += data.Size()
		}

		return nil
	})
	if err != nil {
		return 0.0, err
	}

	return convert(float64(byteSize), size)
}
