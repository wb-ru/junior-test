package factory

type (
	byGround struct {
		from map[int]string
		to   map[int]string
	}
	bySea struct {
		from map[int]string
		to   map[int]string
	}
	byAir struct {
		from map[int]string
		to   map[int]string
	}
	taskManager interface {
		getInfo(int) (string, string)
		putInfo(int, string, string)
	}
)

func (truck byGround) getInfo(id int) (string, string) {
	if truck.from[id] == "" || truck.to[id] == "" {
		return "", ""
	}

	return truck.from[id], truck.to[id]
}

func (ship bySea) getInfo(id int) (string, string) {
	if ship.from[id] == "" || ship.to[id] == "" {
		return "", ""
	}

	return ship.from[id], ship.to[id]
}

func (plane byAir) getInfo(id int) (string, string) {
	if plane.from[id] == "" || plane.to[id] == "" {
		return "", ""
	}
	return plane.from[id], plane.to[id]
}

func (truck byGround) putInfo(id int, From string, To string) {
	truck.from[id] = From
	truck.to[id] = To
}

func (ship bySea) putInfo(id int, From string, To string) {
	ship.from[id] = From
	ship.to[id] = To
}

func (plane byAir) putInfo(id int, From string, To string) {
	plane.from[id] = From
	plane.to[id] = To
}

func logistic(way string) taskManager {
	switch way {
	case "Ground":
		return byGround{
			from: make(map[int]string),
			to:   make(map[int]string),
		}
	case "Sea":
		return bySea{
			from: make(map[int]string),
			to:   make(map[int]string),
		}
	case "Air":
		return byAir{
			from: make(map[int]string),
			to:   make(map[int]string),
		}
	default:
		return nil
	}
}
