package db

import (
	"time"
)

type CNCType struct {
	ID     int    `gorm:"column:Id;AUTO_INCREMENT;NOT NULL"`
	Maker  string `gorm:"column:Maker;NOT NULL"`
	Series string `gorm:"column:Series;NOT NULL"`
	Model  string `gorm:"column:Model;NOT NULL"`
}

func (m *CNCType) TableName() string {
	return "CNCType"
}

type CNCDescription struct {
	ID        int    `gorm:"column:Id;AUTO_INCREMENT;NOT NULL"`
	Message   string `gorm:"column:Message;NOT NULL"`
	CNCTypeID int    `gorm:"column:CNCType_Id;NOT NULL"`
}

func (m *CNCDescription) TableName() string {
	return "CNCDescription"
}

type CNCProcess struct {
	ID        int    `gorm:"column:Id;AUTO_INCREMENT;NOT NULL"`
	Name      string `gorm:"column:Name;NOT NULL"`
	CNCTypeID int    `gorm:"column:CNCType_Id;NOT NULL"`
}

func (m *CNCProcess) TableName() string {
	return "CNCProcess"
}

type CNCParam struct {
	ID           int     `gorm:"column:Id;AUTO_INCREMENT;NOT NULL"`
	Name         string  `gorm:"column:Name;NOT NULL"`
	Dimension    string  `gorm:"column:Dimension;NOT NULL"`
	Frequence    float64 `gorm:"column:Frequence;NOT NULL"`
	MinValue     float64 `gorm:"column:MinValue;NOT NULL"`
	MaxValue     float64 `gorm:"column:MaxValue;NOT NULL"`
	DefaultValue float64 `gorm:"column:DefaultValue;NOT NULL"`
	CNCProcessID int     `gorm:"column:CNCProcess_Id;NOT NULL"`
}

func (m *CNCParam) TableName() string {
	return "CNCParam"
}

type User struct {
	ID       int    `gorm:"column:Id;AUTO_INCREMENT;NOT NULL"`
	Login    string `gorm:"column:Login;NOT NULL"`
	Password string `gorm:"column:Password;NOT NULL"`
}

func (m *User) TableName() string {
	return "User"
}

type Log struct {
	ID      int       `gorm:"column:Id;AUTO_INCREMENT;NOT NULL"`
	Level   int       `gorm:"column:Level;NOT NULL"`
	Time    time.Time `gorm:"column:Time;NOT NULL"`
	Message string    `gorm:"column:Message;NOT NULL"`
	Source  string    `gorm:"column:Source;NOT NULL"`
	UserID  int       `gorm:"column:User_Id;NOT NULL"`
}

func (m *Log) TableName() string {
	return "Log"
}

type Context struct {
	ID        int `gorm:"column:Id;AUTO_INCREMENT;NOT NULL"`
	UserID    int `gorm:"column:User_Id;NOT NULL"`
	CNCTypeID int `gorm:"column:CNCType_Id"`
}

func (m *Context) TableName() string {
	return "Context"
}

type Backlog struct {
	ID        int    `gorm:"column:Id;AUTO_INCREMENT;NOT NULL"`
	Command   string `gorm:"column:Command;NOT NULL"`
	Answer    string `gorm:"column:Answer;NOT NULL"`
	ContextID int    `gorm:"column:Context_Id;NOT NULL"`
}

func (m *Backlog) TableName() string {
	return "Backlog"
}

type Message struct {
	ID        int       `gorm:"column:Id;AUTO_INCREMENT;NOT NULL"`
	Message   string    `gorm:"column:Message;NOT NULL"`
	Date      time.Time `gorm:"column:Date;NOT NULL"`
	ContextID int       `gorm:"column:Context_Id;NOT NULL"`
}

func (m *Message) TableName() string {
	return "Message"
}
