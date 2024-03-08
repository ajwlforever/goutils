package main

type Cache struct {
	Cap   int
	List  []int
	Value map[int]int
}

func New(Cap int) Cache {
	return Cache{
		Cap:   Cap,
		List:  make([]int, Cap),
		Value: make(map[int]int, 0),
	}
}

func (c *Cache) Get(k int) (res int) {
	if c.Value[k] == 0 {
		res = -1
	} else {
		res = c.Value[k]
	}
	return
}

func (c *Cache) Put(k int, v int) {
	if c.Get(k) != -1 {
		// already exists
		index := -1
		for idx, v := range c.List {
			if v == k {
				index = idx
				break
			}
		}
		//
		c.List = append(v, c.List[0:index], c.List[index+1:len(c.List)-1])

	} else {
		//
		if len(c.Value) == Cap {
			//
			delete(c.Value, c.List[len(c.List)-1])
			c.List[len(c.List)-1] = k
			c.Value[k] = v
			delete(c.Value, k)

		} else {
			// Cap enough
			c.List = append(k, c.List)
			c.Value[k] = v
		}
	}
}

func main() {
	c := New(2)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Get(1)
}
