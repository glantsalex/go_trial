package main

import (
	"github.com/cihub/seelog"
	"innovi-event-server/model"
	"learn_nlp/mem"
	"sort"
	"time"
)

type Pair struct {
	Key string
	Value int
}

type PairList []Pair

func main(){

	defer seelog.Flush()

	stageA := mem.NewStage("stage a")
	stageB := mem.NewStage("stage b")
	stageC := mem.NewStage("stage c")

	ch := make( chan model.InnoviEvent, 10 )

	out:= stageC.Pump( stageB.Pump( stageA.Pump( ch ) ) )

	for {

		ev := getInnoviEventFromFile( "test_data/event_with_clip.json")
		ch <- *ev

		<-out

		time.Sleep( time.Millisecond*100)
	}
}

func rankByWordCount(wordFrequencies map[string]int) PairList{
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }



/*
	myself, _:= user.Current()

	homeDir := myself.HomeDir
	pathToFile := path.Join( homeDir, "/Downloads/Limonov_-_The_Other_Russia.txt")
	data, err := ioutil.ReadFile( pathToFile )

	if err != nil{
		seelog.Errorf("Error read file: %s", err )
		return
	}

	doc, err := prose.NewDocument( string( data ), prose.WithSegmentation(false)  )

	freqOfWords:= make( map[string]int, 0, )

	for _, ent := range doc.Entities() {
		if  ent.Label != "GPE"{
			continue
		}

		if  _, ok := freqOfWords[ent.Text];ok{
			freqOfWords[ent.Text] = freqOfWords[ent.Text] + 1
			continue
		}else{
			freqOfWords[ent.Text] = 1
		}
		//fmt.Println(ent.Text, ent.Label)
	}

	wordPairs := rankByWordCount(freqOfWords)

	for _, wp := range wordPairs {
		fmt.Printf("%.30s  %d\n", wp.Key,wp.Value)
	}
*/