package util

type Bitmap struct {
	size int
	data []uint64
}

func NewBitmap(size int) *Bitmap {
	return &Bitmap{
		size: size,
		data: make([]uint64, (size+63)/64),
	}
}

func (b *Bitmap) Set(bit int) {
	index := bit / 64
	offset := bit % 64
	b.data[index] |= 1 << offset
}

func (b *Bitmap) Get(bit int) bool {
	index := bit / 64
	offset := bit % 64
	return (b.data[index] & (1 << offset)) != 0
}

func (b *Bitmap) Delete(bit int) {
	index := bit / 64
	offset := bit % 64
	b.data[index] &= ^(1 << offset)
}

func (b *Bitmap) Capacity() int {
	return len(b.data)
}

func (b *Bitmap) Size() int {
	return b.size
}

func (b *Bitmap) Clone() *Bitmap {
	copyBm := NewBitmap(b.size)
	for i, v := range b.data {
		copyBm.data[i] = v
	}

	return copyBm
}
