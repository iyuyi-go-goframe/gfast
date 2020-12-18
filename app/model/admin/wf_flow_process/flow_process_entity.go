// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package wf_flow_process

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
)

// Entity is the golang structure for table wf_flow_process.
type Entity struct {
	Id              int    `orm:"id,primary"        json:"id"`                //
	FlowId          uint   `orm:"flow_id"           json:"flow_id"`           // 流程ID
	ProcessName     string `orm:"process_name"      json:"process_name"`      // 步骤名称
	ProcessType     string `orm:"process_type"      json:"process_type"`      // 步骤类型
	ProcessTo       string `orm:"process_to"        json:"process_to"`        // 转交下一步骤号
	AutoPerson      uint   `orm:"auto_person"       json:"auto_person"`       // 3自由选择|4指定人员|5指定角色|6事务接受
	AutoSponsorIds  string `orm:"auto_sponsor_ids"  json:"auto_sponsor_ids"`  // 4指定步骤主办人ids
	AutoSponsorText string `orm:"auto_sponsor_text" json:"auto_sponsor_text"` // 4指定步骤主办人text
	WorkIds         string `orm:"work_ids"          json:"work_ids"`          // 6事务接受
	WorkText        string `orm:"work_text"         json:"work_text"`         // 6事务接受
	AutoRoleIds     string `orm:"auto_role_ids"     json:"auto_role_ids"`     // 5角色ids
	AutoRoleText    string `orm:"auto_role_text"    json:"auto_role_text"`    // 5角色 text
	RangeUserIds    string `orm:"range_user_ids"    json:"range_user_ids"`    // 3自由选择IDS
	RangeUserText   string `orm:"range_user_text"   json:"range_user_text"`   // 3自由选择用户ID
	IsSing          uint   `orm:"is_sing"           json:"is_sing"`           // 1允许|2不允许
	IsBack          uint   `orm:"is_back"           json:"is_back"`           // 1允许|2不允许
	OutCondition    string `orm:"out_condition"     json:"out_condition"`     // 转出条件
	Setleft         uint   `orm:"setleft"           json:"setleft"`           // 左 坐标
	Settop          uint   `orm:"settop"            json:"settop"`            // 上 坐标
	Style           string `orm:"style"             json:"style"`             // 样式 序列化
	IsDel           uint   `orm:"is_del"            json:"is_del"`            //
	Updatetime      uint   `orm:"updatetime"        json:"updatetime"`        // 更新时间
	Dateline        uint   `orm:"dateline"          json:"dateline"`          //
	WfMode          uint   `orm:"wf_mode"           json:"wf_mode"`           // 0 单一线性，1，转出条件 2，同步模式
	WfAction        string `orm:"wf_action"         json:"wf_action"`         // 对应方法
	WorkSql         string `orm:"work_sql"          json:"work_sql"`          //
	WorkMsg         string `orm:"work_msg"          json:"work_msg"`          //
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// InsertIgnore does "INSERT IGNORE INTO ..." statement for inserting current object into table.
func (r *Entity) InsertIgnore() (result sql.Result, err error) {
	return Model.Data(r).InsertIgnore()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Update() (result sql.Result, err error) {
	return Model.Data(r).Where(gdb.GetWhereConditionOfStruct(r)).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
func (r *Entity) Delete() (result sql.Result, err error) {
	return Model.Where(gdb.GetWhereConditionOfStruct(r)).Delete()
}