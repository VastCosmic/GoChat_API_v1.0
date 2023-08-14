package chatModel

import "time"

// Messages 定义一个结构体，用于存储聊天对话消息数组
type Messages []Message

// Message 定义一个结构体，用于存储聊天对话消息
type Message struct {
	Role    string `json:"role"`    // 角色
	Content string `json:"content"` // 内容
}

// ChatParams 定义一个结构体，用于存储聊天对话参数
type ChatParams struct {
	Model string `json:"model"` // 模型名称
	//ChatID      string    `json:"chat_id"`     // 聊天标识符
	Messages    []Message `json:"messages"`    // 聊天对话消息数组
	Temperature float64   `json:"temperature"` // 聊天对话温度
}

// ChatCompletion 定义一个结构体，用于存储聊天对话补全结果
type ChatCompletion struct {
	ID      string   `json:"id"`      // 标识符
	Object  string   `json:"object"`  // 对象类型
	Created int      `json:"created"` // 创建时间
	Choices []Choice `json:"choices"` // 补全选项数组
	Usage   Usage    `json:"usage"`   // 令牌使用情况
}

// Choice 定义一个结构体，用于存储补全选项
type Choice struct {
	Index        int     `json:"index"`         // 位置
	Message      Message `json:"message"`       // 消息
	FinishReason string  `json:"finish_reason"` // 结束原因
}

// Usage 定义一个结构体，用于存储令牌使用情况
type Usage struct {
	PromptTokens     uint `json:"prompt_tokens"`     // 参数令牌数
	CompletionTokens uint `json:"completion_tokens"` // 补全令牌数
	TotalTokens      uint `json:"total_tokens"`      // 总令牌数
}

// MessageTable 用于存储聊天对话相关内容到数据库中
type MessageTable struct {
	RecordID uint `db:"record_id" sql:"AUTO_INCREMENT" gorm:"primaryKey"` // 单条对话记录id，设置为主键和自增
	ChatID   uint `db:"chat_id"`                                          // 对话id
	UserID   uint `db:"user_id"`                                          // 用户id
	//CreatedTime      time.Time `db:"created_time"`      // 创建时间
	Role             string `db:"role"`              // 对话角色
	Content          string `db:"content"`           // 对话文本内容
	PromptTokens     uint   `db:"prompt_tokens"`     // prompt令牌数
	CompletionTokens uint   `db:"completion_tokens"` // 回复令牌数
	TotalTokens      uint   `db:"total_tokens"`      // 总令牌数
}

// TableName returns the table name for MessageHistory struct
func (MessageHistory) TableName() string {
	return "message_table"
}

// TableName returns the table name for MessageTable struct
func (MessageTable) TableName() string {
	return "message_table"
}

// MessageHistory 查询数据库后存储的对话结构体
type MessageHistory struct {
	RecordID         int       `gorm:"column:record_id"`
	ChatID           int       `gorm:"column:chat_id"`
	UserID           int       `gorm:"column:user_id"`
	CreatedTime      time.Time `gorm:"column:created_time"`
	Role             string    `gorm:"column:role"`
	Content          string    `gorm:"column:content"`
	PromptTokens     int       `gorm:"column:prompt_tokens"`
	CompletionTokens int       `gorm:"column:completion_tokens"`
	TotalTokens      int       `gorm:"column:total_tokens"`
}

// MessageHistoryReturn 用于返回给前端的对话历史记录
type MessageHistoryReturn struct {
	ChatID        int             `json:"chat_id"`
	UserID        int             `json:"user_id"`
	MessageReturn []MessageReturn `json:"message_return"`
}

// MessageReturn 对话历史记录子结构体
type MessageReturn struct {
	CreatedTime time.Time `json:"created_time"`
	Role        string    `json:"role"`
	Content     string    `json:"content"`
}
