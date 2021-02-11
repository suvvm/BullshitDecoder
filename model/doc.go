package model

import (
	"fmt"
	"github.com/yanyiwu/gojieba"
	"io/ioutil"
	"log"
	"math"
	"os"
)

var (
	DocNum int
	WordSet map[string]int
	GojiebaX *gojieba.Jieba
)

// 文档信息
type Doc struct {
	Path string					// 文档路径
	Name string					// 文档名
	Content string				// 文档内容
	Words map[string]int		// 文档分词结果
	WordsTotal int				// 文档总词数量
	KeyWords string				// 文档关键词
	WordsTF map[string]float64	// 关键词词频
	WordsIDF map[string]float64	// 关键词逆向文件频率
	TFIDF map[string]float64	// 词频-逆文档频率
}

// InitDoc read file from path
func (d *Doc) InitDoc(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("open file:%s error:%v", path, err)
		return
	}
	defer file.Close()
	d.Name = file.Name()
	d.Path = path
	d.Words = make(map[string]int)
	d.WordsTF = make(map[string]float64)
	d.WordsIDF = make(map[string]float64)
	d.TFIDF = make(map[string]float64)
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("read file:%s data error:%v", d.Name, err)
		return
	}
	d.Content = string(data)
	err = d.InitWord()
	if err != nil {
		log.Fatalf("read file:%s init word error:%v", d.Name, err)
		return
	}
}

func (d *Doc)InitWord() error {
	if d.Content == "" {
		return fmt.Errorf("doc contant is empty")
	}
	x := gojieba.NewJieba()
	defer x.Free()
	words := x.Cut(d.Content, true)
	d.WordsTotal = len(words)
	for _, word := range words {
		if _, ok := d.Words[word]; ok {
			d.Words[word]++
		} else {
			d.Words[word] = 1
		}
	}
	for key, _ := range d.Words {
		if _, ok := WordSet[key]; ok{
			WordSet[key]++
		} else {
			WordSet[key] = 1
		}
	}
	return nil
}

func (d *Doc) computeTF() {
	for key, val := range d.Words {
		d.WordsTF[key] = float64(val) / float64(d.WordsTotal)
	}
}

func (d *Doc) computeIDF() {
	for key, _ := range d.Words {
		d.WordsIDF[key] = math.Log10(float64(DocNum + 1) / float64(WordSet[key] + 1))
	}
}

func (d *Doc) computeTFIDF() {
	for key, val := range d.WordsTF {
		d.TFIDF[key] = val * d.WordsIDF[key]
		// fmt.Printf("%s: CNT=%v TF=%v IDF=%v TFIDF=%v\t", key, d.Words[key], d.WordsTF[key], d.WordsIDF[key], d.TFIDF[key])
	}
}

func (d *Doc) DoTFIDF() {
	d.computeTF()
	d.computeIDF()
	d.computeTFIDF()
	maxTFIDF := float64(0)
	for key, val := range d.TFIDF{
		if val > maxTFIDF {
			maxTFIDF = val
			d.KeyWords = key
		}
	}
}
