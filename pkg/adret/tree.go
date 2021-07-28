package ubac

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

// Function wich aim to imitate the tree command output
// It prints the node name without prefix with the right indentation compare to its position in the Tree.
// If it is the last element of a directory it print it with a special character behind
// If it is in a directory which is "the last element" it print one "|" less
func specialPrint(name string, last bool, inlast bool) {
	// compute depth (using / counter)
	depth := strings.Count(name, "/")

	output := ""

	//determine the appropriate characters
	var lastCharacter string
	var inlastCharacter string

	tab := "   " //tabulate character
	if last {
		lastCharacter = "└── "
	} else {
		lastCharacter = "├── "
	}

	if inlast {
		inlastCharacter = ""
	} else {
		inlastCharacter = "|"
	}

	if depth != 0 {
		//if not root path
		if depth == 1 {
			output += lastCharacter
		}
		if depth == 2 {
			output += inlastCharacter + tab + lastCharacter
		}
		if depth > 2 {
			output += strings.Repeat("|"+tab, depth-2) + inlastCharacter + tab + lastCharacter
		}

	}
	output += filepath.Base(name) //whatever happens

	fmt.Println(output)
}

// Print the Tree struct in a fashion way (as tree command would do.. I hope)
func PrintTree(tree Tree, root string) {
	rootNode, err := getNodeByName(root, tree.Nodes)
	if err != nil {
		log.SetFlags(0)
		log.Fatal(err)
	}
	specialPrint(root, true, false)
	PrintNode(tree.Nodes, rootNode, false, false)
}

// (recursive) Print the tree under the Node (except the node itself)
// Retrieve all node under if it is a directory and print it, nothing if it is a file
func PrintNode(nodes []Node, node Node, last bool, inlast bool) {
	if node.Type == "file" {
		// the Node has already been printed
	} else if node.Type == "directory" {
		// Recursivity: print node under this node
		// Print all node with  a specific prefix ie node.Dir == prefix
		// Retrieve a list of all node
		// Then iterate over the list when we arrive at last PrintNode(node.Name,true)
		nodeWithPrefix := getNodeWithPrefix(node.Name, nodes)

		inlast = last //if we are in last we must now call PrintNode with inlast at true, and conversely

		//iterate to print the last one differently
		for i := range nodeWithPrefix {
			last := (len(nodeWithPrefix)-1 == i)
			specialPrint(nodeWithPrefix[i], last, inlast)
			//!recursivity
			node, err := getNodeByName(nodeWithPrefix[i], nodes)
			if err != nil {
				log.SetFlags(0)
				log.Fatal(err)
			}
			PrintNode(nodes, node, last, inlast)
		}
	} else {
		log.Fatal("Node/Resource with undefined type")
	}
}
