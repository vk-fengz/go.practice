// counters包 提供告警计数器功能
package counters


//alterCounter是一个未公开的类型
// 用于保存告警计数器
type alterCounter int

// New创建并返回一个未公开的
// alterCounter类型的值
func New(value int) alterCounter {
    return alterCounter(value)
}