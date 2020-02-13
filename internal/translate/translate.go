package translate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
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

	if x, err := filepath.Abs(inPath); err == nil {
		inPath = x
	}

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
	files          []source
}

type source struct {
	t       *translate
	fileSet *token.FileSet
	node    *ast.File
	path    string
	comment []*ast.Comment
}

func (s source) save() error {
	b := bytes.Buffer{}
	for _, s := range s.comment {
		b.Write([]byte(s.Text + "\n\r\r\n"))
	}
	str := b.String()

	if txt, err := google_translate(str, s.t.sourceLanguage, s.t.language); err != nil {
		return err
	} else {
		arr := strings.Split(txt, "\n\r\r\n")
		if len(arr)-1 == len(s.comment) {
			for i := 0; i < len(s.comment); i++ {
				s.comment[i].Text = strings.ReplaceAll(strings.ReplaceAll(arr[i], "* /", "*/"), "/ *", "/*")
			}
		}
	}

	if err := os.MkdirAll(filepath.Dir(s.path), 0700); err != nil {
		return err
	}

	if fo, err := os.Create(s.path); err == nil {
		printer.Fprint(fo, s.fileSet, s.node)
		defer fo.Close()
	} else {
		return err
	}
	return nil
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
	t.files = make([]source, 0)
	var err error

	if err = filepath.Walk(t.inPath, t.visit); err != nil {
		return err
	}

	for _, f := range t.files {
		fmt.Println("Processing file " + f.path)
		if err := f.save(); err != nil {
			fmt.Println("file " + f.path + " Error")
		}
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

	var src source
	src.comment = make([]*ast.Comment, 0)

	for _, c := range node.Comments {
		for _, l := range c.List {
			if len(l.Text) > 0 {
				src.comment = append(src.comment, l)
			}
		}
	}

	if len(src.comment) > 0 {

		tmp, err := filepath.Rel(t.inPath, path)

		if err != nil {
			return nil
		}

		pt := filepath.Join(t.outPath, tmp)

		src.t = t
		src.path = pt
		src.fileSet = fset
		src.node = node
		t.files = append(t.files, src)
	}

	return nil
}

type post struct {
	Text   string `json:"text"`
	Source string `json:"source"`
	Target string `json:"target"`
}

func google_translate(text, source, target string) (string, error) {
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
