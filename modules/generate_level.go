package modules

import (
	"bufio"
	"fmt"
	"github.com/Estuardo2015/rogue_wizard/modules/commons"
	"github.com/Estuardo2015/rogue_wizard/modules/grid"
	"github.com/Estuardo2015/rogue_wizard/modules/logging"
	"github.com/Estuardo2015/rogue_wizard/modules/utils"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
)

func CreateLevelFromTxtFile(path string) (l *Level, err error) {
	f, err := OpenFile(path)
	if logging.Error(err, "unable to open file '%s'", path) {
		return
	}
	defer f.Close()

	g, p, err := GenerateGridAndPlayer(f)
	if logging.Error(err, "unable to generate tile map from file '%s'", path) {
		return
	}

	l = NewLevel(commons.TileWidth, g, p)

	return
}

func OpenFile(path string) (f *os.File, err error) {
	f, err = os.Open(path)
	if err != nil {
		return
	}
	return
}

func GenerateGridAndPlayer(f *os.File) (g *grid.Grid, p *Player, err error) {
	scanner := bufio.NewScanner(f)

	p, err = generatePlayer(scanner)
	if err != nil {
		return
	}

	charGrid := make([][]string, 0)
	for scanner.Scan() {
		log.Debug().Msg(scanner.Text())
		row := make([]string, 0)
		for _, c := range scanner.Text() {
			char := string(c)
			row = append(row, char)
		}

		charGrid = append(charGrid, row)
	}

	tg, err := TileGridFromChars(charGrid)
	if err != nil {
		return
	}

	g = grid.NewGrid(len(tg), len(tg[0]), tg)

	if err = scanner.Err(); err != nil {
		return
	}

	return
}

// Level must be generated this way otherwise it comes out flipped -_-
func TileGridFromChars(charGrid [][]string) (tg [][]*grid.Tile, err error) {
	tgLen := len(charGrid[0])
	tgWidth := len(charGrid)

	// Initialize tile grid
	tg = make([][]*grid.Tile, tgLen)
	for i, _ := range tg {
		tg[i] = make([]*grid.Tile, tgWidth)
	}

	// Loop through char grid
	for r, row := range charGrid {
		for i, char := range row {
			tg[i][r] = tileFromChar(char)
		}
	}

	return
}

func generatePlayer(scanner *bufio.Scanner) (*Player, error) {
	scanner.Scan()
	line := scanner.Text()

	log.Debug().Msg(line)

	// Read player coords from first line
	pCoords := strings.Split(line, ",")
	if len(pCoords) != 2 {
		return nil, fmt.Errorf("unable to parse player coords from string '%s'", line)
	}

	p := NewPlayer("Hiro", utils.AtoI(pCoords[0], -1), utils.AtoI(pCoords[1], -1), PlayerImg)

	return p, nil
}

func tileFromChar(char string) *grid.Tile {
	switch char {
	case "#":
		return grid.NewTile(WallImg, true)
	case "v":
		return grid.NewTile(GrassImg, false)
	default:
		return nil
	}

	return nil
}
