package mem

import (
	"github.com/cihub/seelog"
	. "innovi-event-server/model"
	"time"
)

type ChannelType chan InnoviEvent

type Stage struct{
	Name string
	InCh, OutCh ChannelType
}

func  NewStage( name string ) *Stage   {
	return  &Stage{
		Name:name,
		InCh:make( chan InnoviEvent, 10 ),
		OutCh:make( chan InnoviEvent, 10 ),
	}
}


func (s *Stage ) Pump( in ChannelType ) ChannelType {

	go func() {

		seelog.Debugf("inside %s Pump ", s.Name)
		for{
			select {
			case ev := <- in:
				s.process( ev )
				s.OutCh <- ev
			}
		}

	}()
	return s.OutCh
}

func (s *Stage ) process( ev InnoviEvent ){
	seelog.Debugf("inside %s process", s.Name)
	time.Sleep( time.Millisecond * 50 )
}