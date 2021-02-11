package main

import (
	"fmt"
	"github.com/yanyiwu/gojieba"
	"suvvm.work/tf_idf/model"
	"time"
)

func init() {
	model.DocNum = 9
	startInit := time.Now()
	model.GojiebaX = gojieba.NewJieba()	// 初始化gojieba
	endInit := time.Since(startInit)
	fmt.Printf("init times=%v\n", endInit)
	model.WordSet = make(map[string]int)	// 初始化词库
}

func main()  {
	defer model.GojiebaX.Free()
	docs := make([]*model.Doc, model.DocNum)
	startTime := time.Now()
	for i := 0; i < model.DocNum; i++ {
		docs[i] = &model.Doc{}
		docs[i].InitDoc(fmt.Sprintf("./resources/doc%d.txt", i + 1))
	}
	endTime := time.Since(startTime)
	fmt.Printf("word set init time=%v\n", endTime)
	startTime = time.Now()
	for i := 0; i< model.DocNum; i++ {
		docs[i].DoTFIDF()
		fmt.Printf("%s:%s\n",docs[i].Name, docs[i].KeyWords)
	}
	endTime = time.Since(startTime)
	fmt.Printf("tfidf time=%v\n", endTime)

}
