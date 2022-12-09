package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	workers := make([]worker, 0)

	leader := newManager("Kut Zhang", "ODC", "IT Manager")
	workers = append(workers, leader)

	m := newMember("Lucy Liu", "ODC", leader.id)
	workers = append(workers, m)

	for _, worker := range workers {
		worker.work()

		switch target := worker.(type) {
		case *member:
			fmt.Printf("member info: %q - %q\n", target.name, target.leaderId)

		case *manager:
			fmt.Printf("manager info: %q - %q\n", target.name, target.remark)

		default:
			fmt.Printf("unknown worker type: %T\n", target)
		}
	}

	fmt.Println("------------------------------------------------------------")

	for _, worker := range workers {
		if m, ok := worker.(*member); ok {
			fmt.Printf("member info: %q - %q\n", m.name, m.leaderId)
			continue
		}

		if ma, ok := worker.(*manager); ok {
			fmt.Printf("manager info: %q - %q\n", ma.name, ma.remark)
			continue
		}

		fmt.Printf("unknown worker type: %T\n", worker)
	}
}

// worker interface
type worker interface {
	// work
	work()
}

// staff
type staff struct {
	id         string
	name       string
	position   string
	department string
}

// member
type member struct {
	staff
	leaderId string
}

// new member
func newMember(name, department, leaderId string) *member {
	return &member{
		staff: staff{
			id:         uuid.NewString(),
			name:       name,
			position:   "member",
			department: department,
		},
		leaderId: leaderId,
	}
}

// work
func (m *member) work() {
	fmt.Printf("member %q working....done!!!\n", m.name)
}

// manager
type manager struct {
	staff
	remark string
}

// new manager
func newManager(name, department, remark string) *manager {
	return &manager{
		staff: staff{
			id:         uuid.NewString(),
			name:       name,
			position:   "manager",
			department: department,
		},
		remark: remark,
	}
}

// work
func (m *manager) work() {
	fmt.Printf("manager %q working......complete!!!!!\n", m.name)
}
