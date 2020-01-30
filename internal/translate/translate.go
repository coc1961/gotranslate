package translate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Translate interface {
	Translate() error
}

func New(language, sourceLanguage, inPath, outPath string) (Translate, error) {

	return (&translate{
		language:       language,
		sourceLanguage: sourceLanguage,
		inPath:         inPath,
		outPath:        outPath,
	}).validate()

}

type translate struct {
	sourceLanguage string
	language       string
	inPath         string
	outPath        string
}

func (t *translate) validate() (*translate, error) {
	if _, err := os.Stat(t.inPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("input dir does not exist")
	}
	if f, err := os.Stat(t.outPath); os.IsNotExist(err) || !f.IsDir() {
		return nil, fmt.Errorf("output dir does not exist")
	}
	if len(t.language) != 2 {
		return nil, fmt.Errorf("language is invalid")

	}
	return t, nil
}

func (t *translate) Translate() error {
	var err error
	if err = filepath.Walk(t.inPath, t.visit); err != nil {
		return err
	}
	return nil
}

// visita código fuente
func (t *translate) visit(path string, f os.FileInfo, err1 error) error {

	if err1 != nil {
		return err1
	}
	if f, err := os.Stat(path); err == nil && f.IsDir() {
		return nil
	}

	if ext := filepath.Ext(path); ext != ".go" {
		return nil
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("parseFile error:%+v", err)
	}

	for _, c := range node.Comments {
		for _, l := range c.List {
			if s, err := t.google_translate(l.Text, t.sourceLanguage, t.language); err == nil {
				l.Text = s
			} else {
				return err
			}
		}
	}

	pt := filepath.Join(t.outPath, path)
	if err := os.MkdirAll(filepath.Dir(pt), 0700); err != nil {
		return err
	}

	if fo, err := os.Create(pt); err == nil {
		printer.Fprint(fo, fset, node)
	} else {
		return err
	}

	return nil
}

type post struct {
	Text   string `json:"text"`
	Source string `json:"source"`
	Target string `json:"target"`
}

func (t *translate) google_translate(text, source, target string) (string, error) {
	endpoint := "https://script.google.com/macros/s/AKfycbywwDmlmQrNPYoxL90NCZYjoEzuzRcnRuUmFCPzEqG7VdWBAhU/exec"

	postData, err := json.Marshal(post{text, source, target})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer([]byte(postData)))

	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("translate error")
	}

	bod := strings.ToLower(string(body))

	if strings.Contains(bod, "<html>") || bod == "source is empty" {
		return "", fmt.Errorf("translate error")
	}

	return string(body), nil
}
