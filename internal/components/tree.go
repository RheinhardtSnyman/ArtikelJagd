package component

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type placedTree struct {
	img *ebiten.Image
	x   float64
	y   float64
}

type tree struct {
	treeRow []placedTree
	w       float64
	scale   float64
}

func NewTree(y int, density float64, scale float64, w float64, in float64) Component {
	imgUrls := []string{
		"./assets/images/Stall/tree_oak.png",
		"./assets/images/Stall/tree_pine.png"}

	var trees []placedTree
	for x := 0.0; x < w; x += in + (in * (1 - density)) {
		img, _, err := ebitenutil.NewImageFromFile(imgUrls[int(getRandom(0, len(imgUrls)))])
		if err != nil {
			log.Fatal(err)
		}
		trees = append(trees, placedTree{
			img: img,
			x:   getRandom(int(x-15), int(x+15)),
			y:   getRandom(y-25, y+25),
		})
	}
	return &tree{
		treeRow: trees,
		w:       w,
		scale:   scale,
	}
}

func (tree *tree) Draw(screen *ebiten.Image) error {

	for _, t := range tree.treeRow {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Scale(tree.scale, tree.scale)
		options.GeoM.Translate(t.x, t.y)
		screen.DrawImage(t.img, options)
	}

	return nil
}

func (tree *tree) Update() error {
	return nil
}

func (tree *tree) OnScreen() bool {
	return true
}
