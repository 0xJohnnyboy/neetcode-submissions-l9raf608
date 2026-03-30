public class DynamicArray {
    private int[] _Array;
    private int _Capacity;
    private int _Size;

    public DynamicArray(int capacity) {
        if (capacity <= 0){
            throw new ArgumentException("capacity must be over 0");
        }

        this._Capacity = capacity;
        this._Size = 0;
        this._Array = new int[capacity];
    }

    public int Get(int i) {
        return this._Array[i];
    }

    public void Set(int i, int n) {
        this._Array[i] = n;
    }

    public void PushBack(int n) {
        if (this._Size == this._Capacity){
            this. Resize();
        }

        this._Array[this._Size] = n;
        this._Size++;
    }

    public int PopBack() {
        if (this._Size > 0){
            this._Size--;
        }
        return this._Array[this._Size];
    }

    private void Resize() {
        this._Capacity *= 2;

        int[] newArray = new int[this._Capacity];

        for (int i = 0; i < this._Size; i++){
            newArray[i] = this._Array[i];
        }

        this._Array = newArray;
    }

    public int GetSize() {
        return this._Size;
    }

    public int GetCapacity() {
        return this._Capacity;
    }
}
