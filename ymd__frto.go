package cmmn

type Ymd__frto struct {
	fr Ymd
	to Ymd
}

func (y *Ymd__frto) Get__days() int {
	days := 0
	t_curr := y.Get__fr().Get__time()
	for t_curr.Before(y.to.Get__time()) {
		days++
		t_curr = t_curr.AddDate(0, 0, 1)
	}
	return days + 1
}

func (y *Ymd__frto) Get__fr() Ymd {
	return y.fr
}

func (y *Ymd__frto) Get__to() Ymd {
	return y.to
}

func Ymd__frto__new(fr, to Ymd) Ymd__frto {
	if fr.Get__time().After(to.Get__time()) {
		tmp := fr
		fr = to
		to = tmp
	}
	return Ymd__frto{
		fr: fr,
		to: to,
	}
}
