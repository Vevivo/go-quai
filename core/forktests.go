package core

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/spruce-solutions/go-quai/common"
	"github.com/spruce-solutions/go-quai/core/types"
)

// struct for notation
type blockGenSpec struct {
	numbers    [3]int    // prime idx, region idx, zone idx
	parentTags [3]string // (optionally) Override the parents to point to tagged blocks. Empty strings are ignored.
	tag        string    // (optionally) Give this block a named tag. Empty strings are ignored.
}

// Generate blocks to form a network of chains
func SendNetworkBlocksToNode(graph [3][3][]*blockGenSpec, blocks []*types.Block) error {
	return errors.New("Not implemented")
}

func GetNodeHeadAndTD() (common.Hash, [3]*big.Int, error) {
	return common.Hash{}, [3]*big.Int{}, errors.New("Not implemented")
}

// Example test for a fork choice scenario shown in slide00 (not a real slide)
func ForkChoiceTest_Slide05() {

	// Define the network graph here
	graph := [3][3][]*blockGenSpec{
		{ // Region1
			{ // Zone1
				&blockGenSpec{[3]int{-1, 1, 1}, [3]string{}, "z00_1"},
				&blockGenSpec{[3]int{-1, -1, 2}, [3]string{}, ""},
				// ...
				&blockGenSpec{[3]int{-1, 2, 2}, [3]string{"", "z00_1", "z00_1"}, ""}, //Fork at z00_1
				&blockGenSpec{[3]int{-1, -1, 3}, [3]string{}, ""},
				// ...
				&blockGenSpec{[3]int{-1, 3, 5}, [3]string{"", "z01_1", ""}, ""}, //Fork at z01_1
				&blockGenSpec{[3]int{-1, -1, 6}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, 4, 8}, [3]string{}, ""},
			},
			{
				&blockGenSpec{[3]int{-1, 2, 1}, [3]string{"", "z00_1", ""}, "z01_1"},
			},
			{}, // ... Zone3 omitted
		},
		{}, // ... Region2 omitted
		{}, // ... Region3 omitted
	}

	fmt.Println(graph)
	// Generate the blocks for this graph
	/*
		blocks, err := GenerateNetworkBlocks(graph)
		if err != nil {
			panic("Error generating blocks!")
		}

		// Send internal & external blocks to the node
		err = SendNetworkBlocksToNode(graph, blocks)
		if err != nil {
			panic("Failed to send all blocks to the node!")
		}*/

	// Confirm the node arrived at the expected head, with the expected total difficulty
	expectedHead := common.Hash{}
	expectedTD := [3]*big.Int{big.NewInt(0), big.NewInt(0), big.NewInt(0)}
	head, td, err := GetNodeHeadAndTD()
	if err != nil {
		panic("Error getting node head & total difficulty!")
	} else if head != expectedHead {
		panic("Node is on the wrong head!")
	} else if td != expectedTD {
		panic("Node has the wrong total difficulty!")
	}
}

// Example test for a fork choice scenario shown in slide00 (not a real slide)
func ForkChoiceTest_Slide06() {

	// Define the network graph here
	graph := [3][3][]*blockGenSpec{
		{}, // ... Region1 omitted
		{ // Region2
			{ //Zone1
				&blockGenSpec{[3]int{1, 1, 1}, [3]string{}, "z00_1"},
			},
			{}, //... Zone2 omitted
			{}, //... Zone3 omitted
		},
		{ //Region 3
			{ //Zone1
				&blockGenSpec{[3]int{2, 1, 1}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, -1, 5}, [3]string{}, ""},
			},
			{ //Zone2
				&blockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&blockGenSpec{[3]int{2, 1, 2}, [3]string{"z00_1", "", ""}, ""}, //Fork at z00_1
				&blockGenSpec{[3]int{-1, -1, 3}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, -1, 6}, [3]string{}, ""},
			},
			{ //Zone3
				&blockGenSpec{[3]int{-1, 2, 1}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, 3, 2}, [3]string{}, ""},
			},
		},
	}

	fmt.Println(graph)
	// Generate the blocks for this graph
	/*
		blocks, err := GenerateNetworkBlocks(graph)
		if err != nil {
			panic("Error generating blocks!")
		}

		// Send internal & external blocks to the node
		err = SendNetworkBlocksToNode(graph, blocks)
		if err != nil {
			panic("Failed to send all blocks to the node!")
		}*/

	// Confirm the node arrived at the expected head, with the expected total difficulty
	expectedHead := common.Hash{}
	expectedTD := [3]*big.Int{big.NewInt(0), big.NewInt(0), big.NewInt(0)}
	head, td, err := GetNodeHeadAndTD()
	if err != nil {
		panic("Error getting node head & total difficulty!")
	} else if head != expectedHead {
		panic("Node is on the wrong head!")
	} else if td != expectedTD {
		panic("Node has the wrong total difficulty!")
	}
}

// Example test for a fork choice scenario shown in slide08
func ForkChoiceTest_Slide08() {

	// Define the network graph here
	graph := [3][3][]*blockGenSpec{
		{ // Region1
			{
				&blockGenSpec{[3]int{-1, 1, 1}, [3]string{}, ""},
			},
			{
				&blockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, 2, 3}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, -1, 3}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, -1, 11}, [3]string{}, "z00_1"},
				&blockGenSpec{[3]int{-1, 3, 12}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, -1, 12}, [3]string{"", "", "z00_1"}, ""}, //Fork at z00_1
				&blockGenSpec{[3]int{-1, -1, 13}, [3]string{}, ""},
			},
			{}, // ... Zone3 omitted
		},
		{}, // ... Region2 omitted
		{}, // ... Region3 omitted
	}

	fmt.Println(graph)
	// Generate the blocks for this graph
	/*
		blocks, err := GenerateNetworkBlocks(graph)
		if err != nil {
			panic("Error generating blocks!")
		}

		// Send internal & external blocks to the node
		err = SendNetworkBlocksToNode(graph, blocks)
		if err != nil {
			panic("Failed to send all blocks to the node!")
		}*/

	// Confirm the node arrived at the expected head, with the expected total difficulty
	expectedHead := common.Hash{}
	expectedTD := [3]*big.Int{big.NewInt(0), big.NewInt(0), big.NewInt(0)}
	head, td, err := GetNodeHeadAndTD()
	if err != nil {
		panic("Error getting node head & total difficulty!")
	} else if head != expectedHead {
		panic("Node is on the wrong head!")
	} else if td != expectedTD {
		panic("Node has the wrong total difficulty!")
	}
}

// Example test for a fork choice scenario shown in slide11
func ForkChoiceTest_Slide11() {

	// Define the network graph here
	graph := [3][3][]*blockGenSpec{
		{ // Region1
			{
				&blockGenSpec{[3]int{-1, 1, 1}, [3]string{}, ""},
			},
			{
				&blockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, 2, 3}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, -1, 3}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, -1, 11}, [3]string{}, "z00_1"},
				&blockGenSpec{[3]int{-1, 3, 12}, [3]string{}, "z00_3"},
				&blockGenSpec{[3]int{-1, -1, 12}, [3]string{"", "", "z00_1"}, "z00_2"}, //Fork at z00_1
				&blockGenSpec{[3]int{-1, -1, 13}, [3]string{"", "", "z00_2"}, ""},
				&blockGenSpec{[3]int{-1, -1, 13}, [3]string{"", "", "z00_3"}, ""},
			},
			{}, // ... Zone3 omitted
		},
		{}, // ... Region2 omitted
		{}, // ... Region3 omitted
	}

	fmt.Println(graph)
	// Generate the blocks for this graph
	/*
		blocks, err := GenerateNetworkBlocks(graph)
		if err != nil {
			panic("Error generating blocks!")
		}

		// Send internal & external blocks to the node
		err = SendNetworkBlocksToNode(graph, blocks)
		if err != nil {
			panic("Failed to send all blocks to the node!")
		}*/

	// Confirm the node arrived at the expected head, with the expected total difficulty
	expectedHead := common.Hash{}
	expectedTD := [3]*big.Int{big.NewInt(0), big.NewInt(0), big.NewInt(0)}
	head, td, err := GetNodeHeadAndTD()
	if err != nil {
		panic("Error getting node head & total difficulty!")
	} else if head != expectedHead {
		panic("Node is on the wrong head!")
	} else if td != expectedTD {
		panic("Node has the wrong total difficulty!")
	}
}

// Example test for a fork choice scenario shown in slide13
func ForkChoiceTest_Slide13() {

	// Define the network graph here
	graph := [3][3][]*blockGenSpec{
		{ // Region1
			{
				&blockGenSpec{[3]int{-1, 1, 1}, [3]string{}, "z00_1"},
			},
			{
				&blockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, -1, 2}, [3]string{}, "z01_1"},
				&blockGenSpec{[3]int{-1, -1, 3}, [3]string{}, ""},
				//...
				&blockGenSpec{[3]int{-1, 2, 3}, [3]string{"", "z00_1", "z01_1"}, ""},
				&blockGenSpec{[3]int{-1, -1, 4}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, -1, 5}, [3]string{}, ""},
			},
			{}, // ... Zone3 omitted
		},
		{}, // ... Region2 omitted
		{}, // ... Region3 omitted
	}

	fmt.Println(graph)
	// Generate the blocks for this graph
	/*
		blocks, err := GenerateNetworkBlocks(graph)
		if err != nil {
			panic("Error generating blocks!")
		}

		// Send internal & external blocks to the node
		err = SendNetworkBlocksToNode(graph, blocks)
		if err != nil {
			panic("Failed to send all blocks to the node!")
		}*/

	// Confirm the node arrived at the expected head, with the expected total difficulty
	expectedHead := common.Hash{}
	expectedTD := [3]*big.Int{big.NewInt(0), big.NewInt(0), big.NewInt(0)}
	head, td, err := GetNodeHeadAndTD()
	if err != nil {
		panic("Error getting node head & total difficulty!")
	} else if head != expectedHead {
		panic("Node is on the wrong head!")
	} else if td != expectedTD {
		panic("Node has the wrong total difficulty!")
	}
}

// Example test for a fork choice scenario shown in slide15
func ForkChoiceTest_Slide15() {

	// Define the network graph here
	graph := [3][3][]*blockGenSpec{
		{ // Region1
			{ // Zone1
				&blockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, -1, 2}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, 3, 6}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, -1, 7}, [3]string{}, ""},
			},
			{ // Zone2
				&blockGenSpec{[3]int{1, 1, 1}, [3]string{}, "z12_1"},
				&blockGenSpec{[3]int{-1, -1, 2}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, 2, 3}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, -1, 4}, [3]string{}, ""},
				//...
				&blockGenSpec{[3]int{3, 4, 7}, [3]string{"z31_1", "", ""}, ""},
			},
			{}, // ... Zone3 omitted
		},
		{
			{}, // ... Zone1 omitted
			{ // Zone2
				&blockGenSpec{[3]int{2, 1, 1}, [3]string{}, ""},
			},
			{}, // ... Zone3 omitted
		},
		{
			{
				&blockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&blockGenSpec{[3]int{2, 2, 3}, [3]string{"z12_1", "", ""}, "z31_1"},
				&blockGenSpec{[3]int{-1, -1, 4}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, -1, 7}, [3]string{}, ""},
			},
			{}, // ... Zone2 omitted
			{
				&blockGenSpec{[3]int{-1, 1, 1}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, -1, 2}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, 3, 6}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, 4, 7}, [3]string{}, ""},
			},
		},
	}

	fmt.Println(graph)
	// Generate the blocks for this graph
	/*
		blocks, err := GenerateNetworkBlocks(graph)
		if err != nil {
			panic("Error generating blocks!")
		}

		// Send internal & external blocks to the node
		err = SendNetworkBlocksToNode(graph, blocks)
		if err != nil {
			panic("Failed to send all blocks to the node!")
		}*/

	// Confirm the node arrived at the expected head, with the expected total difficulty
	expectedHead := common.Hash{}
	expectedTD := [3]*big.Int{big.NewInt(0), big.NewInt(0), big.NewInt(0)}
	head, td, err := GetNodeHeadAndTD()
	if err != nil {
		panic("Error getting node head & total difficulty!")
	} else if head != expectedHead {
		panic("Node is on the wrong head!")
	} else if td != expectedTD {
		panic("Node has the wrong total difficulty!")
	}
}

// Example test for a fork choice scenario shown in slide17
func ForkChoiceTest_Slide17() {

	// Define the network graph here

	graph := [3][3][]*blockGenSpec{
		{ // Region1
			{
				&blockGenSpec{[3]int{-1, 1, 1}, [3]string{}, "z00_1"},
			},
			{
				&blockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, 2, 3}, [3]string{"", "z00_1", ""}, ""},
				&blockGenSpec{[3]int{-1, -1, 4}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, 3, 7}, [3]string{"", "z02_1", ""}, ""},
			},
			{
				&blockGenSpec{[3]int{-1, 2, 1}, [3]string{"", "z00_1", ""}, "z02_1"},
			},
		},
		{}, // ... Region2 omitted
		{}, // ... Region3 omitted
	}

	fmt.Println(graph)
	// Generate the blocks for this graph
	/*
		blocks, err := GenerateNetworkBlocks(graph)
		if err != nil {
			panic("Error generating blocks!")
		}

		// Send internal & external blocks to the node
		err = SendNetworkBlocksToNode(graph, blocks)
		if err != nil {
			panic("Failed to send all blocks to the node!")
		}*/

	// Confirm the node arrived at the expected head, with the expected total difficulty
	expectedHead := common.Hash{}
	expectedTD := [3]*big.Int{big.NewInt(0), big.NewInt(0), big.NewInt(0)}
	head, td, err := GetNodeHeadAndTD()
	if err != nil {
		panic("Error getting node head & total difficulty!")
	} else if head != expectedHead {
		panic("Node is on the wrong head!")
	} else if td != expectedTD {
		panic("Node has the wrong total difficulty!")
	}
}

// Example test for a fork choice scenario shown in slide18
func ForkChoiceTest_Slide18() {

	// Define the network graph here

	graph := [3][3][]*blockGenSpec{
		{ // Region1
			{
				&blockGenSpec{[3]int{-1, 1, 1}, [3]string{}, "z00_1"},
			},
			{
				&blockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, 2, 3}, [3]string{"", "z00_1", ""}, ""},
				&blockGenSpec{[3]int{1, 2, 4}, [3]string{"", "z00_1", ""}, ""},
			},
			{},
		},
		{}, // ... Region2 omitted
		{}, // ... Region3 omitted
	}

	fmt.Println(graph)
	// Generate the blocks for this graph
	/*
		blocks, err := GenerateNetworkBlocks(graph)
		if err != nil {
			panic("Error generating blocks!")
		}

		// Send internal & external blocks to the node
		err = SendNetworkBlocksToNode(graph, blocks)
		if err != nil {
			panic("Failed to send all blocks to the node!")
		}*/

	// Confirm the node arrived at the expected head, with the expected total difficulty
	expectedHead := common.Hash{}
	expectedTD := [3]*big.Int{big.NewInt(0), big.NewInt(0), big.NewInt(0)}
	head, td, err := GetNodeHeadAndTD()
	if err != nil {
		panic("Error getting node head & total difficulty!")
	} else if head != expectedHead {
		panic("Node is on the wrong head!")
	} else if td != expectedTD {
		panic("Node has the wrong total difficulty!")
	}
}

// Example test for a fork choice scenario shown in slide23
func ForkChoiceTest_Slide23() {

	// Define the network graph here
	graph := [3][3][]*blockGenSpec{
		{ // Region1
			{
				&blockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, 1, 2}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, -1, 3}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, -1, 4}, [3]string{}, "z00_1"},
				//...
				&blockGenSpec{[3]int{-1, 2, 5}, [3]string{"", "", "z00_1"}, ""},
				&blockGenSpec{[3]int{1, 3, 5}, [3]string{"", "", "z00_1"}, ""},
			},
			{},
			{},
		},
		{}, // ... Region2 omitted
		{}, // ... Region3 omitted
	}

	fmt.Println(graph)
	// Generate the blocks for this graph
	/*
		blocks, err := GenerateNetworkBlocks(graph)
		if err != nil {
			panic("Error generating blocks!")
		}

		// Send internal & external blocks to the node
		err = SendNetworkBlocksToNode(graph, blocks)
		if err != nil {
			panic("Failed to send all blocks to the node!")
		}*/

	// Confirm the node arrived at the expected head, with the expected total difficulty
	expectedHead := common.Hash{}
	expectedTD := [3]*big.Int{big.NewInt(0), big.NewInt(0), big.NewInt(0)}
	head, td, err := GetNodeHeadAndTD()
	if err != nil {
		panic("Error getting node head & total difficulty!")
	} else if head != expectedHead {
		panic("Node is on the wrong head!")
	} else if td != expectedTD {
		panic("Node has the wrong total difficulty!")
	}
}

// Example test for a fork choice scenario shown in slide29
func ForkChoiceTest_Slide29() {

	// Define the network graph here
	graph := [3][3][]*blockGenSpec{
		{ // Region1
			{
				&blockGenSpec{[3]int{1, 1, 1}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, -1, 2}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, 1, 3}, [3]string{}, "z00_1"},
				&blockGenSpec{[3]int{3, 3, 4}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, -1, 5}, [3]string{}, ""},
				// ...
				&blockGenSpec{[3]int{-1, -1, 4}, [3]string{}, "z00_1"},
			},
			{}, // ... Zone2 omitted
			{}, // ... Zone3 omitted
		},
		{
			{
				&blockGenSpec{[3]int{2, 1, 1}, [3]string{}, ""},
			},
			{}, // ... Zone2 omitted
			{}, // ... Zone3 omitted
		},
		{}, // ... Region3 omitted
	}

	fmt.Println(graph)
	// Generate the blocks for this graph
	/*
		blocks, err := GenerateNetworkBlocks(graph)
		if err != nil {
			panic("Error generating blocks!")
		}

		// Send internal & external blocks to the node
		err = SendNetworkBlocksToNode(graph, blocks)
		if err != nil {
			panic("Failed to send all blocks to the node!")
		}*/

	// Confirm the node arrived at the expected head, with the expected total difficulty
	expectedHead := common.Hash{}
	expectedTD := [3]*big.Int{big.NewInt(0), big.NewInt(0), big.NewInt(0)}
	head, td, err := GetNodeHeadAndTD()
	if err != nil {
		panic("Error getting node head & total difficulty!")
	} else if head != expectedHead {
		panic("Node is on the wrong head!")
	} else if td != expectedTD {
		panic("Node has the wrong total difficulty!")
	}
}

// Example test for a fork choice scenario shown in slide30
func ForkChoiceTest_Slide30() {

	// Define the network graph here
	graph := [3][3][]*blockGenSpec{
		{ // Region1
			{
				&blockGenSpec{[3]int{1, 1, 1}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, -1, 2}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, 1, 3}, [3]string{}, "z00_1"},
				&blockGenSpec{[3]int{3, 3, 4}, [3]string{}, ""},
				// ...
				&blockGenSpec{[3]int{-1, -1, 4}, [3]string{}, "z00_1"},
			},
			{}, // ... Zone2 omitted
			{}, // ... Zone3 omitted
		},
		{
			{
				&blockGenSpec{[3]int{2, 1, 1}, [3]string{}, ""},
			},
			{}, // ... Zone2 omitted
			{}, // ... Zone3 omitted
		},
		{}, // ... Region3 omitted
	}

	fmt.Println(graph)
	// Generate the blocks for this graph
	/*
		blocks, err := GenerateNetworkBlocks(graph)
		if err != nil {
			panic("Error generating blocks!")
		}

		// Send internal & external blocks to the node
		err = SendNetworkBlocksToNode(graph, blocks)
		if err != nil {
			panic("Failed to send all blocks to the node!")
		}*/

	// Confirm the node arrived at the expected head, with the expected total difficulty
	expectedHead := common.Hash{}
	expectedTD := [3]*big.Int{big.NewInt(0), big.NewInt(0), big.NewInt(0)}
	head, td, err := GetNodeHeadAndTD()
	if err != nil {
		panic("Error getting node head & total difficulty!")
	} else if head != expectedHead {
		panic("Node is on the wrong head!")
	} else if td != expectedTD {
		panic("Node has the wrong total difficulty!")
	}
}

// Example test for a fork choice scenario shown in slide31
func ForkChoiceTest_Slide31() {

	// Define the network graph here
	graph := [3][3][]*blockGenSpec{
		{ // Region1
			{
				&blockGenSpec{[3]int{1, 1, 1}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, -1, 2}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, 2, 3}, [3]string{}, "z00_1"},
				&blockGenSpec{[3]int{3, 3, 4}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, -1, 5}, [3]string{}, ""},
				// ...
				&blockGenSpec{[3]int{-1, -1, 4}, [3]string{"", "", "z00_1"}, ""},
				&blockGenSpec{[3]int{-1, -1, 5}, [3]string{}, ""},
				&blockGenSpec{[3]int{3, 3, 6}, [3]string{"z12_1", "z00_1", ""}, ""},
			},
			{}, // ... Zone2 omitted
			{}, // ... Zone3 omitted
		},
		{
			{}, // ... Zone1 omitted
			{}, // ... Zone2 omitted
			{
				&blockGenSpec{[3]int{2, 1, 1}, [3]string{}, "z12_1"},
			},
		},
		{}, // ... Region3 omitted
	}

	fmt.Println(graph)
	// Generate the blocks for this graph
	/*
		blocks, err := GenerateNetworkBlocks(graph)
		if err != nil {
			panic("Error generating blocks!")
		}

		// Send internal & external blocks to the node
		err = SendNetworkBlocksToNode(graph, blocks)
		if err != nil {
			panic("Failed to send all blocks to the node!")
		}*/

	// Confirm the node arrived at the expected head, with the expected total difficulty
	expectedHead := common.Hash{}
	expectedTD := [3]*big.Int{big.NewInt(0), big.NewInt(0), big.NewInt(0)}
	head, td, err := GetNodeHeadAndTD()
	if err != nil {
		panic("Error getting node head & total difficulty!")
	} else if head != expectedHead {
		panic("Node is on the wrong head!")
	} else if td != expectedTD {
		panic("Node has the wrong total difficulty!")
	}
}

// Example test for a fork choice scenario shown in slide32
func ForkChoiceTest_Slide32() {

	// Define the network graph here
	graph := [3][3][]*blockGenSpec{
		{ // Region1
			{
				&blockGenSpec{[3]int{1, 1, 1}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, -1, 2}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, 2, 3}, [3]string{}, "z00_1"},
				&blockGenSpec{[3]int{3, 3, 4}, [3]string{}, ""},
				&blockGenSpec{[3]int{-1, -1, 5}, [3]string{}, ""},
				// ...
				&blockGenSpec{[3]int{-1, -1, 4}, [3]string{"", "", "z00_1"}, ""},
				&blockGenSpec{[3]int{-1, -1, 5}, [3]string{}, ""},
				&blockGenSpec{[3]int{4, 3, 6}, [3]string{"", "z00_1", ""}, ""},
			},
			{},
			{},
		},
		{
			{},
			{},
			{
				&blockGenSpec{[3]int{2, 1, 1}, [3]string{}, "z12_1"},
			},
		},
		{
			{
				&blockGenSpec{[3]int{3, 1, 1}, [3]string{"z12_1", "", ""}, ""},
			},
			{},
			{},
		}, // ... Region3 omitted
	}

	fmt.Println(graph)
	// Generate the blocks for this graph
	/*
		blocks, err := GenerateNetworkBlocks(graph)
		if err != nil {
			panic("Error generating blocks!")
		}

		// Send internal & external blocks to the node
		err = SendNetworkBlocksToNode(graph, blocks)
		if err != nil {
			panic("Failed to send all blocks to the node!")
		}*/

	// Confirm the node arrived at the expected head, with the expected total difficulty
	expectedHead := common.Hash{}
	expectedTD := [3]*big.Int{big.NewInt(0), big.NewInt(0), big.NewInt(0)}
	head, td, err := GetNodeHeadAndTD()
	if err != nil {
		panic("Error getting node head & total difficulty!")
	} else if head != expectedHead {
		panic("Node is on the wrong head!")
	} else if td != expectedTD {
		panic("Node has the wrong total difficulty!")
	}
}

/*	Outline for Graph
graph := [3][3][]*blockGenSpec{
	{ // Region1
		{}, // ... Zone1 omitted
		{}, // ... Zone2 omitted
		{}, // ... Zone3 omitted
	},
	{}, // ... Region2 omitted
	{}, // ... Region3 omitted
}*/
