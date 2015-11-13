package main 

import(
	"fmt"
	"UJSPackage/file"
)



func main() {
	var filename string
	// Get filename
	fmt.Println("Enter file to be read")
    fmt.Scanf("%s\n", &filename)
    fmt.Println()
	// Read in file to slice
	aRelation := file.IntLines(filename)
	printRel(aRelation)
	// Test for Reflexivity
	reflexive := isReflexive(aRelation)
	if reflexive {
		fmt.Println("The relation is reflexive\n")
	}
	// Test for Symmetry
	symmetric := isSymmetric(aRelation)
	if symmetric {
		fmt.Println("The relation is symmetric\n")
	}
	// Test for Transitivity
	transitive := isTransitive(aRelation)
	if transitive {
		fmt.Println("The relation is transitive\n")
	}
	// Output Results
	if reflexive && symmetric && transitive {
		fmt.Println("The relation is an equivalence relation.\n")
		fmt.Println("The equivalence classes are:")
		eClasses(aRelation)
	} else {
		fmt.Println("The relation is not an equivalence relation.")
	}
}

// This function returns true if the given relation is reflexive
func isReflexive(aRelation [][]int) bool {
	var reflex bool = true
	// For each line i, set reflex to false if i,i is not 1
	for i, _ := range aRelation {
		reflex = reflex && aRelation[i][i] == 1
	}
	return reflex
}

// This function returns true is the given relation is symmetric
func isSymmetric(aRelation [][]int) bool {
	symmetry := true
	// on each line i, for each row j < i, symmetry is false if
	//    i,j = 1 && j,i = 0
	for i, _ := range aRelation {
		for j := 0; j < i; j++ {
			if aRelation[i][j] == 1 && aRelation[j][i] == 0 {
				symmetry = false
			}
		}
	}
	return symmetry
}

// This function returns true if the given function is transitive
func isTransitive(aRelation [][]int) bool{
	transitivity := true
	// For each line i, if i,j = 1 then scan line j and
	//    check that for all j,k = 1, i,k = 1
	for i, _ := range aRelation {
		for j, _ := range aRelation[i] {
			if aRelation[i][j] == 1 {
				for k, _ := range aRelation[j] {
					if aRelation[j][k] == 1 && aRelation[i][k] == 0 {
						transitivity = false
					}
				}
			}
		}
	}
	return transitivity
}

// Function to print out the equivalence classes of given relation,
// assumes the input is an equivalence relation
func eClasses(aRelation [][]int) {
	// Create a slice of slices to hold our classes
	classList := make([][]int,0)
	// Go down each row in the relation and see if they are in a class yet
	for i, _ := range aRelation {
		if !inClass(i,classList){
			// If the element is not in a class, make a new class for it
			newClass := make([]int,1)
			newClass[0] = i
			// Now that we've gotten the first element,
			// Build out the rest of the class and add it to the list
			newClass = buildClass(newClass,aRelation)
			classList = append(classList,newClass)
		}
	}
	// Just print it out raw. Why not?
	fmt.Println(classList)
}

// given a partial relation class, and the relation it is from
// will find and add the rest of the elements in the class
func buildClass(newClass []int, aRelation [][]int) []int {
	// go slices are fun. I can iterate over this slice as I build it,
	// and check each element for more relations
	for _, w := range newClass {
		// w is the current 'member' of the class we are iterating over
		// We will look at each element of row and see if they are already
		// in the relation class
		for j, v := range aRelation[w] {
			// If any column represents a related element,
			//  add it to the class
			if v == 1 && !contains(newClass, j) {
				newClass = append(newClass, j)
			}
		}
	}
	return newClass
}

// returns true if the given integer slice contains the given integer
func contains(slice []int, n int) bool {
	for i, _ := range slice {
		if slice[i] == n {
			return true
		}
	}
	return false
}

// returns true if integer i is found in the slice of integer slices
func inClass(member int, classList [][]int) bool {
	for i, _ := range classList {
		for j, _ := range classList[i] {
			if classList[i][j] == member {
				return true
			}
		}
	}
	return false
}

//!!! Testing only!
// I have this in here just so that it is easier to see the relation the
// program is workin on
func printRel(aRelation [][]int) {
	for i, _ := range aRelation {
		for j, _ := range aRelation[i] {
			fmt.Printf("%d ", aRelation[i][j])
		}
		fmt.Println()
		fmt.Println()
	}
}