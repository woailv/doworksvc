package dao

func limits(page, size int) (limit, start int) {
	return size, (page - 1) * size
}
