package main

import (
	"fmt"
	"os"
)

func createBaseFile(filePath string, baseContentJSON *[]byte) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		err := os.WriteFile(filePath, *baseContentJSON, 0644)
		if err != nil {
			fmt.Printf("Failed to create base %s: %v\n", filePath, err)
		}
	} else {
		fmt.Printf("File:\n%s\n Already exist\n", filePath)
	}
}

var baseMyFundsJSON = []byte(`[]`)

var baseFundsJSON = []byte(`[
  {
    "name": "Superfondo Acciones",
    "url": "https://www.santander.com.ar/personas/inversiones/informacion-fondos#/detail/1"
  },
  {
    "name": "Superfondo Renta $",
    "url": "https://www.santander.com.ar/personas/inversiones/informacion-fondos#/detail/2"
  },
  {
    "name": "Super Ahorro $",
    "url": "https://www.santander.com.ar/personas/inversiones/informacion-fondos#/detail/6"
  },
  {
    "name": "Superfondo Acciones Brasil cuota C",
    "url": "https://www.santander.com.ar/personas/inversiones/informacion-fondos#/detail/9"
  },
  {
    "name": "Superfondo Renta Variable",
    "url": "https://www.santander.com.ar/personas/inversiones/informacion-fondos#/detail/12"
  },
  {
    "name": "Super Bonos",
    "url": "https://www.santander.com.ar/personas/inversiones/informacion-fondos#/detail/16"
  },
  {
    "name": "Supergestión cuota C",
    "url": "https://www.santander.com.ar/personas/inversiones/informacion-fondos#/detail/17"
  },
  {
    "name": "Superfondo Renta Fija",
    "url": "https://www.santander.com.ar/personas/inversiones/informacion-fondos#/detail/18"
  },
  {
    "name": "Superfondo Latinoamérica cuota C",
    "url": "https://www.santander.com.ar/personas/inversiones/informacion-fondos#/detail/19"
  },
  {
    "name": "Supergestión Balanceado cuota C",
    "url": "https://www.santander.com.ar/personas/inversiones/informacion-fondos#/detail/27"
  },
  {
    "name": "Supergestión MIX VI",
    "url": "https://www.santander.com.ar/personas/inversiones/informacion-fondos#/detail/44"
  },
  {
    "name": "Super Ahorro Plus",
    "url": "https://www.santander.com.ar/personas/inversiones/informacion-fondos#/detail/49"
  },
  {
    "name": "Superfondo Combinado",
    "url": "https://www.santander.com.ar/personas/inversiones/informacion-fondos#/detail/59"
  },
  {
    "name": "Supergestión",
    "url": "https://www.santander.com.ar/personas/inversiones/informacion-fondos#/detail/64"
  }
]`)
