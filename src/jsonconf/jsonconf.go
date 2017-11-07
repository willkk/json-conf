package jsonconf

import (
	"os"
	"bufio"
	"log"
	"strings"
	"encoding/json"
	"io"
	"errors"
)

func Unmarshal(filepath string, v interface{}) error  {
	file, err :=  os.Open(filepath)
	if err != nil {
		log.Printf("Open config file %s failed.", filepath)
		return err
	}

	reader := bufio.NewReader(file)
	multi_comment := 0
	var target_file []string
	for {
		skip_line := true
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatalf("Read line failed. err=%s.", err)
				return err
			}
			break
		}
		line = strings.TrimSpace(line)
		if strings.Contains(line, "/*") {
			var offset, offset2 int = -1, -1
			var part1, part2 string
			offset = strings.Index(line, "/*")
			rs := []rune(line)
			part1 = string(rs[:offset])

			if multi_comment == 0 {
				skip_line = false
			}
			multi_comment ++

			if strings.Contains(line, "*/") {
				multi_comment --
				// Single line, find the last "*/"
				offset2 = strings.Index(line, "*/")
				part2 = string(rs[offset2+2:])
			}
			line = part1 + part2
		}
		if strings.Contains(line, "*/") {
			if multi_comment == 0 {
				log.Printf("No matched openning /*")
				return errors.New("No matched openning /*")
			}
			multi_comment --
			offset := strings.Index(line, "*/")
			rs := []rune(line)
			line = string(rs[offset+2:])
			if multi_comment == 0 {
				skip_line = false
			}
		}

		if skip_line && multi_comment > 0 {
			continue
		}

		offset := strings.Index(line, "//")
		if offset >= 0 {
			rs := []rune(line)
			line = string(rs[0:offset])
		}

		// Skip empty line
		if len(strings.TrimSpace(line)) > 0 {
			target_file = append(target_file, line)
		}
	}
	if multi_comment > 0 {
		return errors.New("No matched closing */")
	}

	target_file_str := strings.Join(target_file, "\n")
	return json.Unmarshal([]byte(target_file_str), v)
}
