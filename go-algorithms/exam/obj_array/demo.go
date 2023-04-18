package obj_array

type Obj struct {
	ID int
}

func Creator1(n int) []*Obj {
	objs := make([]*Obj, 0)
	for i := 0; i < n; i++ {
		obj := &Obj{ID: 1}
		objs = append(objs, obj)
	}
	return objs
}

func Creator2(n int) []*Obj {
	objs := make([]*Obj, n) //这里一次性创建出长度为n的指针数组，减少切片扩容的拷贝操作
	for i := 0; i < n; i++ {
		obj := &Obj{ID: 1}
		objs = append(objs, obj)
	}
	return objs
}

func Creator3(n int) []*Obj {
	res := make([]*Obj, n)
	objs := make([]Obj, n) //1. 这里一次性创建出长度为n的数组，连续空间，减少切片扩容的拷贝操作
	for i := 0; i < n; i++ {
		// 给已经创建好的元素赋值即可
		objs[i].ID = i
		res = append(res, &objs[i])
	}
	// 因为元素在连续的内存空间上，指针访问的时候速度快一点
	return res
}
