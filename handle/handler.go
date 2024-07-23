package handle

// +----------------------------------------------------------------------
// | 兔兔答题考试系统
// +----------------------------------------------------------------------
// | 感谢使用兔兔答题考试系统
// | 本系统经过商业授权，不能转售、开源或者以其他方式进行使用，只可在购买协议范围内容使用
// | 若发现未被版权协议，将追究相应的公司、团队、企业等
// | 访问官网：https://www.tutudati.com
// | 开发者：Mandy
// | 创建时间：2024/7/24 00:16
// | 兔兔答题考试系统开发者版权所有 拥有最终解释权
// +----------------------------------------------------------------------

// InputInterface define a unified input and output interface to receive translated content and translated content
type InputInterface interface {
	Input(content string)
	Output() string
}
