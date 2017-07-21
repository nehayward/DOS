package core

func (a Attacks) FindAttack(id int) Attack {
	for _, attack := range a {
		if attack.ID == id {
			return attack
		}
	}
	return Attack{}
}

func (a Attacks) GetNewID() (id int) {
	return len(a) + 1
}
