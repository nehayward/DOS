package core

// FindAttack checks attack list based on ID
func (a Attacks) FindAttack(id int) Attack {
	for _, attack := range a {
		if attack.ID == id {
			return attack
		}
	}
	return Attack{}
}

// RemoveAttack removes an attack after being stopped
func (a *Attacks) RemoveAttack(id int) {
	*a = (*a)[:id-1]
}

// GetNewID generates unique id for an attack
func (a Attacks) GetNewID() (id int) {
	return len(a) + 1
}
