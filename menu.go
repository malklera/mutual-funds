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
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
		} else {
			opt = strings.TrimSuffix(opt, "\n")
			switch opt {
			case "1":
				optionsMenu(myFundsFile)
			case "2":
				optionsMenu(fundsFile)
			case "3":
				path, err := os.Getwd()
				if err != nil {
					fmt.Printf("Error getting path: %v\n", err)
				}
				createBaseFile(filepath.Join(path, fundsFile), &baseFundsJSON)
				createBaseFile(filepath.Join(path, myFundsFile), &baseMyFundsJSON)
			case "4":
				err := updateValues()
				if err != nil {
					fmt.Printf("Error with updateValues() : %v\n", err)
				}
			case "5":
				innerFor := true
				for innerFor {
					fmt.Println("You sure want to exit? y/n")
					fmt.Print("> ")

					opt, err := reader.ReadString('\n')
					if err != nil {
						fmt.Printf("Error reading input: %v\n", err)
					} else {
						opt = strings.TrimSuffix(opt, "\n")
						switch opt {
						case "y", "Y":
							os.Exit(0)
						case "n", "N":
							innerFor = false
						default:
							fmt.Println("Wrong option")
						}
					}
				}
			default:
				fmt.Println("Wrong option")
			}
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
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
		} else {
			opt = strings.TrimSuffix(opt, "\n")
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
					if err2 != nil {
						fmt.Printf("Error reading input: %v\n", err2)
					} else {
						opt2 = strings.TrimSuffix(opt2, "\n")
						switch opt2 {
						case "1":
							menuShow(context, "allFunds")
						case "2":
							innerFor = false
						default:
							exists, err := fundExist(context, opt2)
							if err != nil {
								fmt.Printf("Error checking the existence of %s : %v\n", opt2, err)
							} else {
								if exists {
									menuShow(context, opt2)
								} else {
									fmt.Println("Invalid fund name.")
								}
							}
						}
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
					if err2 != nil {
						fmt.Printf("Error reading input: %v\n", err2)
					} else {
						opt2 = strings.TrimSuffix(opt2, "\n")
						switch opt2 {
						case "1":
							menuExport(context, "allFunds")
						case "2":
							innerFor = false
						default:
							exists, err := fundExist(context, opt2)
							if err != nil {
								fmt.Printf("Error checking the existence of %s : %v\n", opt2, err)
							} else {
								if exists {
									menuExport(context, opt2)
								} else {
									fmt.Println("Invalid fund name.")
								}
							}
						}
					}
				}
			case "3":
				menuModify(context)
			case "4":
				return
			default:
				fmt.Println("Wrong option.")
			}
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
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
		} else {
			opt = strings.TrimSuffix(opt, "\n")
			switch opt {
			case "1":
				switch choosenFunds {
				case "allFunds":
					err := showData(context, "allFunds")
					if err != nil {
						fmt.Printf("Error in showData(%s, %s) : %v", context, choosenFunds, err)
					}
				default:
					err := showData(context, choosenFunds)
					if err != nil {
						fmt.Printf("Error in showData(%s, %s) : %v", context, choosenFunds, err)
					}
				}
			case "2":
				return
			default:
				fmt.Println("Wrong option.")
			}
		}
	}
}

func menuExport(context string, choosenFunds string) {
	// TODO: In the future i may allow other format apart from .json for export
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

		fmt.Println("1- Export data")
		fmt.Println("2- Back")
		fmt.Print("> ")

		opt, err := reader.ReadString('\n')
		opt = strings.TrimSuffix(opt, "\n")
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
		} else {
			switch opt {
			case "1":
				subMenuExport(context, choosenFunds)
			case "2":
				return
			default:
				fmt.Println("Wrong option.")
			}
		}
	}
}

func subMenuExport(context string, choosenFunds string) {
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

		fmt.Println("1- Back")
		fmt.Println("Input the path where to export")
		fmt.Print("> ")
		path, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
		} else {
			path = strings.TrimSuffix(path, "\n")
			if path == "1" {
				return
			} else {
				dir, err := os.Stat(path)
				if err != nil {
					log.Printf("Error getting the stats of %s : %v\n", path, err)
				} else {
					if dir.IsDir() {
						err := exportData(context, path, choosenFunds)
						if err != nil {
							fmt.Printf("Error with exportData(%s, %s, %s) : %v", context, path, choosenFunds, err)
						}
					} else {
						log.Printf("Error, '%s' is not a valid path\n", path)
					}
				}
			}
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
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
		} else {
			switch opt {
			case "1":
				subMenuModify(context)
			case "2":
				subMenuAdd(context)
			case "3":
				subMenuDelete(context)
			case "4":
				return
			default:
				fmt.Println("Wrong option.")
			}
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
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
		} else {
			opt = strings.TrimSuffix(opt, "\n")
			switch opt {
			case "1":
				return
			case "2":
				err := showData(context, "allFunds")
				if err != nil {
					fmt.Printf("Error on showData(%s, allFunds) : %v\n", context, err)
				}
			default:
				exists, err := fundExist(context, opt)
				if err != nil {
					fmt.Printf("Error checking existence of %s : %v", opt, err)
				} else {
					if exists {
						err := modifyData(context, opt)
						if err != nil {
							fmt.Printf("Error with modifyData(%s, %s) : %v", context, opt, err)
						}
					} else {
						fmt.Printf("Fund do not exist: %s\n", opt)
					}
				}
			}
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
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
		} else {
			opt = strings.TrimSuffix(opt, "\n")
			switch opt {
			case "1":
				return
			case "2":
				err := showData(context, "allFunds")
				if err != nil {
					fmt.Printf("Error with showData(%s, allFunds) : %v", context, err)
				}
			default:
				exists, err := fundExist(context, opt)
				if err != nil {
					fmt.Printf("Error checking existence of %s : %v\n", opt, err)
				} else {
					if exists {
						fmt.Printf("The fund '%s' already exist\n", opt)
					} else {
						err := addData(context, opt)
						if err != nil {
							fmt.Printf("Error with addData(%s, %s) : %v", context, opt, err)
						}
					}
				}
			}
		}
	}
}

func subMenuDelete(context string) {
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
		fmt.Println("- Input fund name to delete")
		fmt.Print("> ")

		opt, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
		} else {
			opt = strings.TrimSuffix(opt, "\n")
			switch opt {
			case "1":
				return
			case "2":
				err := showData(context, "allFunds")
				if err != nil {
					fmt.Printf("Error with showData(%s, allFunds) : %v", context, err)
				}
			default:
				exists, err := fundExist(context, opt)
				if err != nil {
					fmt.Printf("Error checking existence of %s : %v\n", opt, err)
				} else {
					if exists {
						err := deleteData(context, opt)
						if err != nil {
							fmt.Printf("Error with deleteData(%s, %s) : %v", context, opt, err)
						}
					} else {
						fmt.Printf("The fund '%s' do not exist\n", opt)
					}
				}
			}
		}
	}
}
