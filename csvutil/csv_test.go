package csvutil

import "testing"

func TestWriteFile(t *testing.T) {
	filePath := "testdata.csv"
	head := []string{"序号", "姓名", "电话"}
	body := make([][]string, 0)
	body = append(body, []string{"1", "Mark", "18812345678"}, []string{"2", "马克", "18612345678"})
	err := WriteFile(filePath, body, head)
	t.Logf("WriteFile: %v", err)
}

func TestWriteBytes(t *testing.T) {
	head := []string{"序号", "姓名", "电话"}
	body := make([][]string, 0)
	body = append(body, []string{"1", "Mark", "18812345678"}, []string{"2", "马克", "18612345678"})
	b, _ := WriteBytes(body, head)
	t.Logf("WriteBytes: %s", b)
}

func TestReadFile(t *testing.T) {
	filePath := "testdata.csv"
	body, head, _ := ReadFile(filePath)
	t.Logf("body: %s", body)
	t.Logf("head: %s", head)
}

func TestReadFileOffset(t *testing.T) {
	filePath := "testdata.csv"
	body, head, err := ReadFileOffset(filePath, 4, 2)
	t.Logf("body: %s", body)
	t.Logf("head: %s", head)
	t.Logf("err: %v", err)
}
