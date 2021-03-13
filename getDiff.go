package main

type diff struct {
	HasDiff   bool        `json:"hasDiff"`
	Circolari []circolare `json:"circolari"`
}

func getDiff(new, old []circolare) diff {
	d := diff{}

	for _, n := range new {
		found := false
		for _, o := range old {
			if n.Title == o.Title {
				found = true
				break
			}
		}
		if !found {
			d.HasDiff = true
			d.Circolari = append(d.Circolari, n)
		}
	}

	return d
}
