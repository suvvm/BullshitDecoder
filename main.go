package main

import (
	"fmt"
	"github.com/yanyiwu/gojieba"
	"suvvm.work/tf_idf/model"
)

func init() {
	model.DocNum = 9
	model.GojiebaX = gojieba.NewJieba()	// 初始化gojieba
	model.GojiebaX.AddWord("宋健")
	model.WordSet = make(map[string]int)	// 初始化词库
}

func main()  {
	defer model.GojiebaX.Free()
	docs := make([]*model.Doc, model.DocNum)
	for i := 0; i < model.DocNum; i++ {
		docs[i] = &model.Doc{}
		docs[i].InitDoc(fmt.Sprintf("./resources/doc%d.txt", i + 1))
	}
	for i := 0; i< model.DocNum; i++ {
		docs[i].DoTFIDF()
		fmt.Printf("%s:%s\n",docs[i].Name, docs[i].KeyWords)
	}
}
