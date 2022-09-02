package csvutil

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
)

// WriteFile 追加写文件
func WriteFile(fileName string, body [][]string, head []string) error {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.Comma = ','
	w.UseCRLF = true
	if len(head) > 0 {
		r := csv.NewReader(f)
		var record []string
		if record, err = r.Read(); err != nil || len(record) == 0 {
			f.WriteString("\xEF\xBB\xBF")
			if err = w.Write(head); err != nil {
				return err
			}
		}
	}

	err = w.WriteAll(body)
	if err != nil {
		return err
	}
	w.Flush()
	return nil
}

// WriteBytes 写内存
func WriteBytes(body [][]string, head []string) ([]byte, error) {
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	w.Comma = ','
	w.UseCRLF = true
	if len(head) > 0 {
		if buf.Len() == 0 {
			buf.WriteString("\xEF\xBB\xBF")
			if err := w.Write(head); err != nil {
				return nil, err
			}
		}
	}
	if err := w.WriteAll(body); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// ReadFile 读文件
func ReadFile(fileName string) ([][]string, []string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	content, err := r.ReadAll()
	if err != nil {
		return nil, nil, err
	}
	return content[1:], content[0], nil
}

func ReadFileOffset(fileName string, offset, limit int) ([][]string, []string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	body := make([][]string, 0)
	head := make([]string, 0)
	r := csv.NewReader(f)
	counter := 1
	for {
		var row []string
		if row, err = r.Read(); err != nil && err != io.EOF {
			return nil, nil, err
		}
		if err == io.EOF {
			break
		}
		if counter == 1 {
			head = row
		} else if counter >= offset && counter < offset+limit {
			body = append(body, row)
		} else if counter >= offset+limit {
			break
		}
		counter++
	}
	return body, head, nil
}
