package ubac_test

import (
	"strings"
	"testing"

	"github.com/ariary/AravisFS/pkg/ubac"
)

func TestTree(t *testing.T) {

	// Construct tree test
	tree := ubac.GetTreeFromFS("../../../test/arafs/encrypted.arafs")

	resources := [8]string{
		"AAAAAAAAAAAAAAAA6ihdrw4ttG+sj+eQMnlA235ttcqjgwlrfQDRy+r2o07a",
		"AAAAAAAAAAAAAAAA6ihdrw4ttG+sj+eQMnlA237KVk6le4+fERV81xq4dTDo0PnkM3M=",
		"AAAAAAAAAAAAAAAA6ihdrw4ttG+sj+eQMnlA237KVk6lex5C5vNk45UL+SkLdoTh7qyXvwxYug==",
		"AAAAAAAAAAAAAAAA6ihdrw4ttG+sj+eQMnlA237KVk6lex5B+/xxlkV968vlzkYBq0N5t83gW3p22zDwZ9mv",
		"AAAAAAAAAAAAAAAA6ihdrw4ttG+sj+eQMnlA237KVkalc/4wmSpT0hUWczNuTuK7owM=",
		"AAAAAAAAAAAAAAAA6ihdrw4ttG+sj+eQMnlA237KVkCldR9X6ubVaNeqN7l/zuVswoxLF6J6",
		"AAAAAAAAAAAAAAAA6ihdrw4ttG+sj+eQMnlA237KVkalcx5X5+Zlykha96qBXPlK2ErQStsX9XptnT8=",
		"AAAAAAAAAAAAAAAA6ihdrw4ttG+sj+eQMnlA237KVkalcx5W5udkykha97Z2/re2f44mqN5WoU9tLZw=",
	}

	for i := 0; i < len(resources); i++ {
		if !strings.Contains(tree, resources[i]) {
			t.Errorf("Tree construction was incorrect,  %v was not in  %v.", resources[i], tree)
		}
	}

}
