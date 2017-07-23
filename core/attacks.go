package core

// FindAttack checks attack list based on ID
func (a Attacks) FindAttack(id int) (Attack, bool, int) {
	for index, attack := range a {
		if attack.ID == id {
			return attack, true, index
		}
	}
	return Attack{}, false, -0
}

// RemoveAttack removes an attack after being stopped
func (a *Attacks) RemoveAttack(id int) {
	_, exists, index := a.FindAttack(id)
	if exists {
		*a = append((*a)[:index], (*a)[index+1:]...)
	}
}

// AddAttack removes an attack after being stopped
func (a *Attacks) AddAttack(attack Attack) {
	*a = append((*a), attack)
}

// GetNewID generates unique id for an attack
func (a Attacks) GetNewID() (id int) {
	return len(a) + 1
}
