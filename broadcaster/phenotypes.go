package broadcaster

import "github.com/dustin/go-broadcast"

var phenotypeChannels = make(map[string]broadcast.Broadcaster)

func openListener(phenotypeid string) chan interface{} {
	listener := make(chan interface{})
	phenotype(phenotypeid).Register(listener)
	return listener
}

func closeListener(phenotypeid string, listener chan interface{}) {
	phenotype(phenotypeid).Unregister(listener)
	close(listener)
}

func deleteBroadcast(phenotypeid string) {
	b, ok := phenotypeChannels[phenotypeid]
	if ok {
		b.Close()
		delete(phenotypeChannels, phenotypeid)
	}
}

func phenotype(phenotypeid string) broadcast.Broadcaster {
	b, ok := phenotypeChannels[phenotypeid]
	if !ok {
		b = broadcast.NewBroadcaster(10)
		phenotypeChannels[phenotypeid] = b
	}
	return b
}
