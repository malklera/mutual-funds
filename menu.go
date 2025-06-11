package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func menu() {
	for {
		fmt.Println("\nMutual funds options, use numbers only(1, 2, 3, etc)")
		fmt.Println("1- My funds")
		fmt.Println("2- All funds")
		fmt.Println("3- Create base files")
		fmt.Println("4- Update values")
		fmt.Println("5- Exit")
		fmt.Print("> ")

		opt, err := reader.ReadString('\n')
		opt = strings.TrimSuffix(opt, "\n")
		if err == nil {
			switch opt {
			case "1":
				optionsMenu(myFundsFile)
			case "2":
				optionsMenu(fundsFile)
			case "3":
				path, err := os.Getwd()
				if err != nil {
					log.Printf("Error getting path: %v", err)
				}
				createBaseFile(filepath.Join(path, fundsFile), baseFundsJson)
				createBaseFile(filepath.Join(path, myFundsFile), baseMyFundsJson)
			case "4":
				updateValues()
			case "5":
				innerFor := true
				for innerFor {
					fmt.Println("You sure want to exit? y/n")
					fmt.Print("> ")

					opt, err := reader.ReadString('\n')
					opt = strings.TrimSuffix(opt, "\n")
					if err == nil {
						switch opt {
						case "y", "Y":
							os.Exit(0)
						case "n", "N":
							innerFor = false
						default:
							fmt.Println("Wrong option")
						}
					} else {
						fmt.Printf("Error reading input: %v", err)
					}
				}
			default:
				fmt.Println("Wrong option")
			}
		} else {
			fmt.Printf("Error reading input: %v", err)
		}
	}
}

func optionsMenu(context string) {
	for {
		fmt.Print("\nOperating over ")
		if context == myFundsFile {
			fmt.Println("my funds")
		} else {
			fmt.Println("all funds")
		}

		fmt.Println("1- Show data")
		fmt.Println("2- Export data")
		fmt.Println("3- Modify data")
		fmt.Println("4- Back")
		fmt.Print("> ")

		opt, err := reader.ReadString('\n')
		opt = strings.TrimSuffix(opt, "\n")
		if err == nil {
			switch opt {
			case "1":
				innerFor := true
				for innerFor {
					fmt.Println("Do you wish to operate over:")
					fmt.Println("1- All funds")
					fmt.Println("2- Back")
					fmt.Println("- Input the name of a single fund")
					fmt.Print("> ")

					opt2, err2 := reader.ReadString('\n')
					opt2 = strings.TrimSuffix(opt2, "\n")
					if err2 == nil {
						switch opt2 {
						case "1":
							menuShow(context, "allFunds")
						case "2":
							innerFor = false
						default:
							if fundExist(context, opt2) {
								menuShow(context, opt2)
							} else {
								fmt.Println("Invalid fund name.")
							}
						}
					} else {
						fmt.Printf("Error reading input: %v", err2)
					}
				}
			case "2":
				innerFor := true
				for innerFor {
					fmt.Println("Do you wish to operate over:")
					fmt.Println("1- All funds")
					fmt.Println("2- Back")
					fmt.Println("- Input the name of a single fund")
					fmt.Print("> ")

					opt2, err2 := reader.ReadString('\n')
					opt2 = strings.TrimSuffix(opt2, "\n")
					if err2 == nil {
						switch opt2 {
						case "1":
							menuExport(context, "allFunds")
						case "2":
							innerFor = false
						default:
							if fundExist(context, opt2) {
								menuExport(context, opt2)
							} else {
								fmt.Println("Invalid fund name.")
							}
						}
					} else {
						fmt.Printf("Error reading input: %v", err2)
					}
				}
			case "3":
				menuModify(context)
			case "4":
				return
			default:
				fmt.Println("Wrong option")
			}
		} else {
			fmt.Printf("Error reading input: %v", err)
		}
	}
}

func menuShow(context string, choosenFunds string) {
	for {
		fmt.Print("\nOperating over ")
		switch context {
		case myFundsFile:
			fmt.Println("my funds")
		case fundsFile:
			fmt.Println("all funds")
		default:
			fmt.Println(context)
		}
		// NOTE: I have this menu because in the future there will be more options
		// about how to show the data
		fmt.Println("1- Show data")
		fmt.Println("2- Back")
		fmt.Print("> ")

		opt, err := reader.ReadString('\n')
		opt = strings.TrimSuffix(opt, "\n")
		if err == nil {
			switch opt {
			case "1":
				switch choosenFunds {
				case "allFunds":
					showData(context, "allFunds")
				default:
					showData(context, choosenFunds)
				}
			case "2":
				return
			default:
				fmt.Println("Wrong option")
			}
		} else {
			fmt.Printf("Error reading input: %v", err)
		}
	}
}

func menuExport(context string, choosenFunds string) {
	// TODO: Give options of exporting to a place on the pc choosen by the user
	// first option will be a .json file, later on maybe other options

	for {
		fmt.Print("\nOperating over ")
		if context == myFundsFile {
			fmt.Println("my funds")
		} else {
			fmt.Println("all funds")
		}
		// NOTE: Leave the menu for future options about how to export it
		fmt.Println("1- Export data")
		fmt.Println("2- Back")
		fmt.Print("> ")

		opt, err := reader.ReadString('\n')
		opt = strings.TrimSuffix(opt, "\n")
		if err == nil {
			switch opt {
			case "1":
				fmt.Println("exporting data")
			case "2":
				return
			default:
				fmt.Println("Wrong option")
			}
		} else {
			fmt.Printf("Error reading input: %v", err)
		}
	}
}

func menuModify(context string) {
	for {
		fmt.Print("\nOperating over ")
		switch context {
		case myFundsFile:
			fmt.Println("my funds")
		case fundsFile:
			fmt.Println("all funds")
		default:
			log.Fatalf("Wrong context: %s", context)
		}

		fmt.Println("1- Modify fund")
		fmt.Println("2- Add fund")
		fmt.Println("3- Delete fund")
		fmt.Println("4- Back")
		fmt.Print("> ")

		opt, err := reader.ReadString('\n')
		opt = strings.TrimSuffix(opt, "\n")
		if err == nil {
			switch opt {
			case "1":
				subMenuModify(context)
			case "2":
				subMenuAdd(context)
			case "3":
				// TODO: do the subMenuDelete(context) function and the actual
				// deleting function on operations
				fmt.Println("deleting fund")
			case "4":
				return
			default:
				fmt.Println("Wrong option")
			}
		} else {
			fmt.Printf("Error reading input: %v", err)
		}
	}
}

func subMenuModify(context string) {
	fmt.Print("\nOperating over ")
	switch context {
	case myFundsFile:
		fmt.Println("my funds")
	case fundsFile:
		fmt.Println("all funds")
	default:
		fmt.Println(context)
	}

	for {
		fmt.Println("\nWhat do you wish to do?")
		fmt.Println("1- Back")
		fmt.Println("2- Show funds")
		fmt.Println("- Input fund name to modify")
		fmt.Print("> ")

		opt, err := reader.ReadString('\n')
		opt = strings.TrimSuffix(opt, "\n")
		if err == nil {
			switch opt {
			case "1":
				return
			case "2":
				showData(context, "allFunds")
			default:
				if fundExist(context, opt) {
					modifyData(context, opt)
				} else {
					fmt.Printf("Fund do not exist: %s", opt)
				}
			}
		} else {
			fmt.Printf("Error reading input: %v", err)
		}
	}
}

func subMenuAdd(context string) {
	fmt.Print("\nOperating over ")
	switch context {
	case myFundsFile:
		fmt.Println("my funds")
	case fundsFile:
		fmt.Println("all funds")
	default:
		fmt.Println(context)
	}
	for {
		fmt.Println("\nWhat do you wish to do?")
		fmt.Println("1- Back")
		fmt.Println("2- Show funds")
		fmt.Println("- Input fund name to add")
		fmt.Print("> ")

		opt, err := reader.ReadString('\n')
		if err == nil {
			opt = strings.TrimSuffix(opt, "\n")
			switch opt {
			case "1":
				return
			case "2":
				showData(context, "allFunds")
			default:
				if fundExist(context, opt) {
					fmt.Printf("The fund '%s' already exist", opt)
				} else {
					addData(context, opt)
				}
			}
		} else {
			fmt.Printf("Error reading input: %v", err)
		}
	}
}
