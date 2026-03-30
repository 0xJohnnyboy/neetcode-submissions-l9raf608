type DynamicArray struct {
    values []int
}

func NewDynamicArray(capacity int) *DynamicArray {
    values := make([]int, 0, capacity)
    return &DynamicArray{
        values,
    }
}

func (da *DynamicArray) Get(i int) int {
    return da.values[i]
}

func (da *DynamicArray) Set(i int, n int) {
    da.values[i] = n
}

func (da *DynamicArray) Pushback(n int) {
    da.values = append(da.values, n)
}

func (da *DynamicArray) Popback() int {
    idx := len(da.values) - 1
    value := da.values[idx]
    da.values = da.values[:idx]
    return value
}

func (da *DynamicArray) resize() {
    newSlice := make([]int, len(da.values), 2*cap(da.values))
    copy(newSlice, da.values)
    da.values = newSlice
}

func (da *DynamicArray) GetSize() int {
    return len(da.values)
}

func (da *DynamicArray) GetCapacity() int {
    return cap(da.values)
}
