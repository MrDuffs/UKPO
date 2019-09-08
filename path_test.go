package main

import "testing"

func testPath(worldInput string, t *testing.T, expectedDist float64) {
	world := ParseWorld(worldInput)
	t.Logf("Input world\n%s", world.RenderPath([]Pather{}))
	p, dist, found := Path(world.From(), world.To())
	if !found {
		t.Log("Could not find a path")
	} else {
		t.Logf("Resulting path\n%s", world.RenderPath(p))
	}
	if !found && expectedDist >= 0 {
		t.Fatal("Could not find a path")
	}
	if found && dist != expectedDist {
		t.Fatalf("Expected dist to be %v but got %v", expectedDist, dist)
	}
}

// Проверка прямого пути, если нет препятствий
func TestStraightLine(t *testing.T) {
	testPath(`
.....~......
.....MM.....
.F........T.
....MMM.....
............
`, t, 9)
}

// Проверка обхода горы
func TestPathAroundMountain(t *testing.T) {
	testPath(`
.....~......
.....MM.....
.F..MMMM..T.
....MMM.....
............
`, t, 13)
}

// Проверка отсутствия пути
func TestBlocked(t *testing.T) {
	testPath(`
............
.........XXX
.F.......XTX
.........XXX
............
`, t, -1)
}

// Длинный путь
func TestMaze(t *testing.T) {
	testPath(`
FX.X........
.X...XXXX.X.
.X.X.X....X.
...X.X.XXXXX
.XX..X.....T
`, t, 27)
}

// Переход через гору(3), если путь более затратный
func TestMountainClimber(t *testing.T) {
	testPath(`
..F..M......
.....MM.....
....MMMM..T.
....MMM.....
............
`, t, 12)
}

// Путь через реку(2), если путь более затратный
func TestRiverSwimmer(t *testing.T) {
	testPath(`
.....~......
.....~......
.F...X...T..
.....M......
.....M......
`, t, 11)
}
