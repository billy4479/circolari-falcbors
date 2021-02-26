package main

type diff struct {
	HasDiff   bool        `json:"HasDiff"`
	Circolari []circolare `json:"Circolari"`
}

func getDiff(new, old []circolare) diff {
	d := diff{}

	for _, n := range new {
		for _, o := range old {
			if n.Title != o.Title {
				d.HasDiff = true
				d.Circolari = append(d.Circolari, n)
			}
		}
	}

	return d
}
