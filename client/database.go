package client

import (
	"fmt"
	"strconv"

	pb "github.com/yaoyaochil/wcf-client-go/proto"
	"google.golang.org/protobuf/proto"
)

// GetDBNames 获取数据库名列表
func (c *ClientWCF) GetDBNames() ([]string, error) {
	req := &pb.Request{
		Func: pb.Functions_FUNC_GET_DB_NAMES,
		Msg: &pb.Request_Empty{
			Empty: &pb.Empty{},
		},
	}

	data, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := c.call(data)
	if err != nil {
		return nil, err
	}

	var response pb.Response
	if err := proto.Unmarshal(resp, &response); err != nil {
		return nil, err
	}

	dbs, ok := response.Msg.(*pb.Response_Dbs)
	if !ok {
		return nil, fmt.Errorf("unexpected response type")
	}

	return dbs.Dbs.Names, nil
}

// GetDbTables 获取数据库表
func (c *ClientWCF) GetDbTables(dbName string) ([]*pb.DbTable, error) {
	req := &pb.Request{
		Func: pb.Functions_FUNC_GET_DB_TABLES,
		Msg: &pb.Request_Str{
			Str: dbName,
		},
	}

	data, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := c.call(data)
	if err != nil {
		return nil, err
	}

	var response pb.Response
	if err := proto.Unmarshal(resp, &response); err != nil {
		return nil, err
	}

	tables, ok := response.Msg.(*pb.Response_Tables)
	if !ok {
		return nil, fmt.Errorf("unexpected response type")
	}

	return tables.Tables.Tables, nil
}

// ExecDbQuery 执行数据库查询
func (c *ClientWCF) ExecDbQuery(dbName, sql string) ([]map[string]any, error) {
	req := &pb.Request{
		Func: pb.Functions_FUNC_EXEC_DB_QUERY,
		Msg: &pb.Request_Query{
			Query: &pb.DbQuery{
				Db:  dbName,
				Sql: sql,
			},
		},
	}

	data, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := c.call(data)
	if err != nil {
		return nil, err
	}

	var response pb.Response
	if err := proto.Unmarshal(resp, &response); err != nil {
		return nil, err
	}

	rows := []map[string]any{}
	// 将字段名和值转换为map
	for _, row := range response.GetRows().GetRows() {
		fields := map[string]any{}
		// 将字段名和值转换为map
		for _, field := range row.Fields {
			fields[field.Column] = c.ParseDbField(field)
		}
		rows = append(rows, fields)
	}

	return rows, nil
}

// ParseDbField 解析数据库字段
func (c *ClientWCF) ParseDbField(field *pb.DbField) any {
	str := string(field.Content)
	switch field.Type {
	case 1:
		n, _ := strconv.ParseInt(str, 10, 64)
		return n
	case 2:
		n, _ := strconv.ParseFloat(str, 64)
		return n
	case 4:
		return field.Content
	case 5:
		return nil
	default:
		return str
	}
}
