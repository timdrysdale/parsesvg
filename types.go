package parsesvg

import "github.com/timdrysdale/geo"

type TextField struct {
	Rect        geo.Rect //Corner.X/Y, Dim.W/H
	ID          string
	Prefill     string
	TabSequence int64
}

type Ladder struct {
	Anchor     geo.Point //X,Y
	Dim        geo.Dim   //W,H
	ID         string
	TextFields []TextField
}
