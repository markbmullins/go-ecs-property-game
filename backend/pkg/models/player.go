package models

type Player struct {
    ID        int
    Funds     float64
    Properties []*Property
}
