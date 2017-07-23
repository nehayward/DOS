package core

import (
	"fmt"
	"testing"
)

func TestRemoveAttack(t *testing.T) {
	var attacks Attacks

	// Removing nil
	attacks.RemoveAttack(1)
	if len(attacks) > 0 {
		t.Error("Expected 0, got", len(attacks))
	}

	firstAttack := Attack{ID: 1}
	secondAttack := Attack{ID: 2}
	attacks.AddAttack(firstAttack)
	attacks.AddAttack(secondAttack)

	// Removing item that doesn't exist
	attacks.RemoveAttack(5)
	if len(attacks) > 2 {
		t.Error("Expected 2, got", len(attacks))
	}

	// Removing item 1
	attacks.RemoveAttack(1)
	if len(attacks) != 1 {
		t.Error("Expected 1, got", len(attacks))
	}

	// Removing item 2
	attacks.RemoveAttack(2)
	if len(attacks) != 0 {
		t.Error("Expected 1, got", len(attacks))
	}
}

func TestAddAttack(t *testing.T) {
	var attacks Attacks

	firstAttack := Attack{ID: 1}
	secondAttack := Attack{ID: 2}
	attacks.AddAttack(firstAttack)

	fmt.Println(len(attacks))
	if len(attacks) != 1 {
		t.Error("Expected 1, got", len(attacks))
	}

	attacks.AddAttack(secondAttack)
	if len(attacks) != 2 {
		t.Error("Expected 1, got", len(attacks))
	}
}

// // FindAttack checks attack list based on ID
// func (a Attacks) FindAttack(id int) Attack {
// 	for _, attack := range a {
// 		if attack.ID == id {
// 			return attack
// 		}
// 	}
// 	return Attack{}
// }
//

// // GetNewID generates unique id for an attack
// func (a Attacks) GetNewID() (id int) {
// 	return len(a) + 1
// }
