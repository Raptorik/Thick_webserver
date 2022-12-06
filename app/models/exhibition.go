package models

type Exhibition struct {
	Name    string    `json:"name"`
	Artists []*Artist `json:"artists"`
}

func (e Exhibition) AddArtist(a *Artist) {
	e.Artists = append(e.Artists, a)
}

func (e *Exhibition) DeleteArtist(artist *Artist) []*Artist {
	for i, artist := range e.Artists {
		if artist == artist {
			if len(e.Artists) == 1 {
				e.Artists = []*Artist{}
			} else {
				e.Artists = e.Artists[i:len(e.Artists)]
			}
			return nil
		}
	}
	return e.Artists
}
