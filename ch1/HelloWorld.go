/**
	Packages need to be clean
	Should express all the requirement of build
	No Cyclic dependencies
	Package name and path are not the same.
	Visibility is package level and not type level.
	G is meant for server side development.
	language for cloud infrastructure

 */
package main // the main package is the root of the initialisation tree
//Compiler loads bits of code and calls the inits recursively


/**
	Imports are...
	Implemented by the compiler
	Imports a package instead of a set of identifiers
 */
import "fmt"

/**
 */
func main()  {
	/**
	there is no semicolons ; it is injected by the compiler
	Every Identifier is Go is either a local identifier to the package or qualified byt type or import
	 */
	fmt.Println("Hellow, 世界!!")
}