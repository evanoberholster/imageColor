package main

import (
	"fmt"
	"log"

	"github.com/lucasb-eyer/go-colorful"
)

func transformMaterialColors() {
	mcolors900 := map[string]string{
		"materialRed":        "#B71C1C", // red
		"materialPink":       "#880E4F",
		"materialPurple":     "#4A148C",
		"materialDeepPurple": "#311B92",
		"materialIndigo":     "#1A237E",
		"materialBlue":       "#0D47A1",
		"materialLightBlue":  "#01579B",
		"materialCyan":       "#006064",
		"materialTeal":       "#004D40",
		"materialGreen":      "#1B5E20",
		"materialLightGreen": "#33691E",
		"materialLime":       "#827717",
		"materialYellow":     "#F57F17",
		"materialAmber":      "#FF6F00",
		"materialOrange":     "#E65100",
		"materialDeepOrange": "#BF360C",
		"materialBrown":      "#3E2723",
		"materialGrey":       "#212121",
		"materialBlueGrey":   "#263238",
	}
	mcolors700 := map[string]string{
		"materialRed":        "#D32F2F",
		"materialPink":       "#C2185B",
		"materialPurple":     "#7B1FA2",
		"materialDeepPurple": "#512DA8",
		"materialIndigo":     "#303F9F",
		"materialBlue":       "#1976D2",
		"materialLightBlue":  "#0288D1",
		"materialCyan":       "#0097A7",
		"materialTeal":       "#00796B",
		"materialGreen":      "#388E3C",
		"materialLightGreen": "#689F38",
		"materialLime":       "#AFB42B",
		"materialYellow":     "#FBC02D",
		"materialAmber":      "#FFA000",
		"materialOrange":     "#F57C00",
		"materialDeepOrange": "#E64A19",
		"materialBrown":      "#5D4037",
		"materialGrey":       "#616161",
		"materialBlueGrey":   "#455A64",
	}

	mcolors500 := map[string]string{
		"materialRed":        "#F44336", // red #F44336
		"materialPink":       "#E91E63", // pink
		"materialPurple":     "#9C27B0", // purple
		"materialDeepPurple": "#673AB7", // deep purple
		"materialIndigo":     "#3F51B5", // indigo
		"materialBlue":       "#2196F3", // blue
		"materialLightBlue":  "#03A9F4", // light blue
		"materialCyan":       "#00BCD4", // cyan
		"materialTeal":       "#009688", // teal
		"materialGreen":      "#4CAF50", // green
		"materialLightGreen": "#8BC34A", // lightgreen
		"materialLime":       "#CDDC39", // lime
		"materialYellow":     "#FFEB3B", // yellow
		"materialAmber":      "#FFC107", // amber
		"materialOrange":     "#FF9800", // orange
		"materialDeepOrange": "#FF5722", // deep orange
		"materialBrown":      "#795548", // brown
		"materialGrey":       "#9E9E9E", // grey
		"materialBlueGrey":   "#607D8B", // blue grey
	}
	mcolors300 := map[string]string{
		"materialRed":        "#E57373",
		"materialPink":       "#F06292",
		"materialPurple":     "#BA68C8",
		"materialDeepPurple": "#9575CD",
		"materialIndigo":     "#7986CB",
		"materialBlue":       "#64B5F6",
		"materialLightBlue":  "#4FC3F7",
		"materialCyan":       "#4DD0E1",
		"materialTeal":       "#4DB6AC",
		"materialGreen":      "#81C784",
		"materialLightGreen": "#AED581",
		"materialLime":       "#DCE775",
		"materialYellow":     "#FFF176",
		"materialAmber":      "#FFD54F",
		"materialOrange":     "#FFB74D",
		"materialDeepOrange": "#FF8A65",
		"materialBrown":      "#A1887F",
		"materialGrey":       "#E0E0E0",
		"materialBlueGrey":   "#90A4AE",
	}

	mcolors100 := map[string]string{
		"materialRed":        "#FFCDD2",
		"materialPink":       "#F8BBD0",
		"materialPurple":     "#E1BEE7",
		"materialDeepPurple": "#D1C4E9",
		"materialIndigo":     "#C5CAE9",
		"materialBlue":       "#BBDEFB",
		"materialLightBlue":  "#B3E5FC",
		"materialCyan":       "#B2EBF2",
		"materialTeal":       "#B2DFDB",
		"materialGreen":      "#C8E6C9",
		"materialLightGreen": "#DCEDC8",
		"materialLime":       "#F0F4C3",
		"materialYellow":     "#FFF9C4",
		"materialAmber":      "#FFECB3",
		"materialOrange":     "#FFE0B2",
		"materialDeepOrange": "#FFCCBC",
		"materialBrown":      "#D7CCC8",
		"materialGrey":       "#F5F5F5",
		"materialBlueGrey":   "#CFD8DC",
	}

	mcolors200 := []string{
		"#EF9A9A",
		"#F48FB1",
		"#CE93D8",
		"#B39DDB",
		"#9FA8DA",
		"#90CAF9",
		"#81D4FA",
		"#80DEEA",
		"#80CBC4",
		"#A5D6A7",
		"#C5E1A5",
		"#E6EE9C",
		"#FFF59D",
		"#FFE082",
		"#FFCC80",
		"#FFAB91",
		"#BCAAA4",
		"#EEEEEE",
		"#B0BEC5",
	}

	fmt.Println(len(mcolors900), len(mcolors700), len(mcolors500), len(mcolors300), len(mcolors200), len(mcolors100))

	fmt.Printf("map[materialColor]ColorHCL{\n")
	for name, m := range mcolors100 {
		col, err := colorful.Hex(m)
		if err != nil {
			log.Fatal(err)
		}
		h, c, l := col.Hsl()
		fmt.Printf("%v: ColorHCL{%.15f, %.15f, %.15f},\n", name, h, c, l)
	}
	fmt.Printf("}\n")
}
