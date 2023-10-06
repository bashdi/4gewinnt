package main

import (
	"fmt"
)

type VgPlayerConsoleInput struct {
	name string
}

func NewVgPlayerConsoleInput(name string) *VgPlayerConsoleInput {
	return &VgPlayerConsoleInput{name: name}
}

func (vgpc *VgPlayerConsoleInput) GetPlayerName() string {
	return vgpc.name
}

func (vgpc *VgPlayerConsoleInput) DoTurn(board [][]int) int {
	fmt.Println("Enter column:")
	var column int
	_, err := fmt.Scanf("%d", &column)
	if err != nil {
		return -1
	}
	return column
}
