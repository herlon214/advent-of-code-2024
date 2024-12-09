package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	data, err := os.ReadFile("day09/example.txt")
	if err != nil {
		panic(err)
	}

	// Part 1
	// disk := NewDisk(data)
	// disk.Optimize(false)
	// fmt.Println("Part 1:", disk.Checksum())

	disk := NewDisk(data)
	disk.Optimize(true)
	fmt.Println("Part 2:", disk.Checksum())
}

type Data struct {
	size int
	id   int
	tile string
}

func (d Data) Size() int {
	return d.size
}

func (d *Data) SetSize(size int) {
	d.size = size
}

func (d Data) ID() int {
	return d.id
}

func (d Data) String() string {
	out := ""
	for i := 0; i < d.Size(); i++ {
		out += d.tile
	}

	return out
}

type Block struct {
	Data
}

func (b Block) Clone() DiskData {
	return &Block{
		Data: Data{
			size: b.Data.size,
			id:   b.Data.id,
			tile: b.Data.tile,
		},
	}
}

type Space struct {
	Data
}

func (e Space) Clone() DiskData {
	return &Space{
		Data: Data{
			size: e.Data.size,
			id:   e.Data.id,
			tile: e.Data.tile,
		},
	}
}

type DiskData interface {
	ID() int
	Size() int
	SetSize(size int)
	Clone() DiskData
	String() string
}

type Disk struct {
	data []DiskData
}

func NewDisk(data []byte) *Disk {
	diskData := make([]DiskData, 0)

	blockId, emptyId := 0, 0
	for i, char := range string(data) {
		val, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}

		if i%2 == 0 {
			if val > 0 {
				diskData = append(diskData, &Block{
					Data: Data{
						size: val,
						id:   blockId,
						tile: fmt.Sprint(blockId),
					},
				})
				blockId++
			}
		} else {
			if val > 0 {
				diskData = append(diskData, &Space{
					Data: Data{
						size: val,
						id:   emptyId,
						tile: ".",
					},
				})
				emptyId++
			}
		}
	}

	return &Disk{
		data: diskData,
	}
}

func (d *Disk) LeftMostSpace(start int, minSize int) int {
	i := start
	for {
		if i >= len(d.data) {
			break
		}

		current := d.data[i]

		switch current.(type) {
		case *Space:
			if current.Size() < minSize {
				i += 1
				continue
			}
		default:
			i += 1
			continue
		}

		return i
	}

	return -1
}

func (d *Disk) RightMostBlock(start int) int {
	i := start
	for {
		if i < 0 || i > len(d.data) {
			break
		}

		current := d.data[i]

		switch current.(type) {
		case *Block:
		default:
			i -= 1
			continue
		}

		return i
	}

	return -1
}

func (d *Disk) Optimize(fitWholeBlock bool) {
	blockStart := len(d.data) - 1
	spaceStart := 0

	i := 0
	for {
		if i > len(d.data) {
			break
		}

		// Get right-most block
		blockIdx := d.RightMostBlock(blockStart)

		// Get left-most empty space
		var spaceIdx int
		if !fitWholeBlock {
			spaceIdx = d.LeftMostSpace(spaceStart, 1)
		} else {
			spaceIdx = d.LeftMostSpace(spaceStart, d.data[blockIdx].Size())

			// Not able to move this block, skip
			if spaceIdx == -1 {
				blockStart -= 1
				continue
			}
		}

		// Check boundaries
		if spaceIdx >= len(d.data) || blockIdx == -1 || spaceIdx == -1 || spaceIdx > blockIdx {
			fmt.Println("stop")
			break
		}

		block := d.data[blockIdx]
		blockStart = blockIdx
		space := d.data[spaceIdx]
		spaceStart = spaceIdx

		targetSize := space.Size() - block.Size()

		// Not enough space
		if targetSize < 0 {
			// Create a new block with the same size as the empty space
			newBlock := block.Clone()
			newBlock.SetSize(space.Size())

			// Replace the empty space with the new block
			d.data[spaceIdx] = newBlock

			// Reduce the size from the original block
			block.SetSize(block.Size() - space.Size())

			// Add space to the end
			newSpace := space.Clone()
			newSpace.SetSize(newBlock.Size())
			d.data = append(d.data, newSpace)
		} else if targetSize == 0 {
			// Swap block and space
			d.data[spaceIdx], d.data[blockIdx] = d.data[blockIdx], d.data[spaceIdx]
		} else { // Enough space
			// Create a new data with everything until the current space
			newData := slices.Clone(d.data)[:spaceIdx]

			// Add the block
			newData = append(newData, block)

			// Shrink space
			space.SetSize(space.Size() - block.Size())
			newData = append(newData, space)

			// Copy the rest
			newData = append(newData, d.data[spaceIdx+1:blockIdx]...)

			// Add moved space
			newSpace := space.Clone()
			newSpace.SetSize(block.Size())
			newData = append(newData, newSpace)

			d.data = newData

		}

		i += 1
		d.Print()
	}
}

func (d *Disk) Checksum() int64 {
	i := 0
	total := int64(0)

	idx := 0
	for {
		if i >= len(d.data) {
			break
		}

		item := d.data[i]

		switch item.(type) {
		case *Block:
		default:
			i++
			continue
		}

		for j := 0; j < item.Size(); j++ {
			total += int64(idx * item.ID())
			// fmt.Printf("%d * %d = %d\n", idx, item.ID(), int64(idx*item.ID()))
			idx += 1
		}

		i++
	}

	return total
}

func (d *Disk) Print() {
	for _, item := range d.data {
		fmt.Printf("%s", item.String())

	}

	fmt.Printf("\n")
}
