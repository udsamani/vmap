package vmap

import (
  "fmt",
  "binary"
)


// OMap stands for optimized map. This map utilizes the fact that GO's runtime does not kick in 
// garbage collector for maps with integer keys and values. Thus we create a wrapper around a amo
// with few more added utilities to optimize the map performance.
// 1) array here holds the actual data.
// 2) items is a map where the index of item's location in array is stored. It is reference by the 
//    a hashedKey
// 3) 
type OMap struct {
    items map[uint64]uint32
    array []byte
    tail int
    headerBuffer []byte
}


func (o *OMap) Add(hashedKey uint64, data []byte) {
     index := o.add(data)
     o.items[hashedKey] = uint32(index)
}

func (o *OMap) add(data []byte) int {
    dataLen := len(data)
    headerEntrySize := binary.PutUvarint(q.headerBuffer, uint64(dataLen))
    o.copy(q.headerBuffer, headerEntrySize)
    o.copy(data, dataLen)
    return o.tail
}

func (o *OMap) copy(data []byte, len int) {
    o.tail += copy(o.array[q.tail:], data[:len])
}

func (o *OMap) Get(hashedKey uint64) []byte {
    index := o.items[hashedKey]
    return o.get(index)
}

func (o *OMap) get(index int) []byte {
    blockSize, n := binary.Uvarint(o.array[index:])
    return o.array[index+n : index + n + int(blockSize)]
}
