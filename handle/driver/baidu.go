package driver

type BaiDu struct {
	Content string
}

func (baidu *BaiDu) Output() (string, error) {
	return baidu.Content, nil
}
