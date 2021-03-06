package main

import (
	"fmt"
	"strconv"
)

type urldir struct {
	url string
	dir string
}

type seginfo struct {
	name   string
	site   string
	link   string
	sid    [16]byte
	status int
	size   size
	down   size
	speed  rate
	eta    string
	task   *taskinfo
}

type taskinfo struct {
	title   string
	suffix  string
	dir     string
	quality string
	play    string
	segs    []*seginfo
	tid     [16]byte
}

type size int64
type rate int64

const (
	_ = 1 << (10 * iota)
	K
	M
	G
	T
)

func (sz size) String() string {
	switch {
	case sz < 0:
		return "-"
	case sz < K:
		return strconv.FormatInt(int64(sz), 10) + "B"
	case sz < M:
		return strconv.FormatInt(int64(sz)/K, 10) + "K"
	case sz < G:
		return fmt.Sprintf("%.1f", float64(sz)/float64(M)) + "M"
	case sz < T:
		return fmt.Sprintf("%.2f", float64(sz)/float64(G)) + "G"
	default:
		return fmt.Sprintf("%.3f", float64(sz)/float64(T)) + "T"
	}
}

func (v rate) String() string {
	if v < 0 {
		return size(v).String()
	}
	return size(v).String() + "/s"
}

type DirArray [5]string

func (a *DirArray) Add(d string) {
	if a[0] == "" {
		a[0] = d
		return
	}
	if a[0] == d {
		return
	}
	for i := range a {
		if a[i] == d {
			copy(a[1:i+1], a[:i])
			a[0] = d
			return
		}
	}
	copy(a[1:len(a)], a[:len(a)-1])
	a[0] = d
}

func (a *DirArray) Len() int {
	i := 0
	for i = range a {
		if a[i] == "" {
			break
		}
	}
	return i
}

type Config struct {
	Xdown     bool
	Threads   int32
	Automerge bool
	Autodel   bool
	CIndex    int32
	Dirs      DirArray
}
