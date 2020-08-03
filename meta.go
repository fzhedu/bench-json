package json_bench
type ColumnInfo struct {
	ID                    int64       `json:"id"`
	Name                  string       `json:"name"`
	Offset                int         `json:"offset"`
	OriginDefaultValue    interface{} `json:"origin_default"`
	OriginDefaultValueBit []byte      `json:"origin_default_bit"`
	DefaultValue          interface{} `json:"default"`
	DefaultValueBit       []byte      `json:"default_bit"`
	// DefaultIsExpr is indicates the default value string is expr.
	DefaultIsExpr       bool                `json:"default_is_expr"`
	GeneratedExprString string              `json:"generated_expr_string"`
	GeneratedStored     bool                `json:"generated_stored"`
	Dependences         map[string]struct{} `json:"dependences"`
	Comment             string      `json:"comment"`
	// A hidden column is used internally(expression index) and are not accessible by users.
	Hidden bool `json:"hidden"`
	// Version means the version of the column info.
	// Version = 0: For OriginDefaultValue and DefaultValue of timestamp column will stores the default time in system time zone.
	//              That is a bug if multiple TiDB servers in different system time zone.
	// Version = 1: For OriginDefaultValue and DefaultValue of timestamp column will stores the default time in UTC time zone.
	//              This will fix bug in version 0. For compatibility with version 0, we add version field in column info struct.
	Version uint64 `json:"version"`
}

// IndexInfo provides meta data describing a DB index.
// It corresponds to the statement `CREATE INDEX Name ON Table (Column);`
// See https://dev.mysql.com/doc/refman/5.7/en/create-index.html
type IndexInfo struct {
	ID        int64          `json:"id"`
	Name      string          `json:"idx_name"` // Index name.
	Table     string          `json:"tbl_name"` // Table name.
	Columns   []string        `json:"idx_cols"` // Index columns.
	Comment   string         `json:"comment"`      // Comment
	Unique    bool           `json:"is_unique"`    // Whether the index is unique.
	Primary   bool           `json:"is_primary"`   // Whether the index is primary key.
	Invisible bool           `json:"is_invisible"` // Whether the index is invisible.
}
// TableInfo provides meta data describing a DB table.
type TableInfo struct {
	ID      int64  `json:"id"`
	Name    string  `json:"name"`
	Charset string `json:"charset"`
	Collate string `json:"collate"`
	// Columns are listed in the order in which they appear in the schema.
	Columns     []*ColumnInfo     `json:"cols"`
	Indices     []*IndexInfo      `json:"index_info"`
	// PKIsHandle is true when primary key is a single integer column.
	PKIsHandle bool `json:"pk_is_handle"`
	// IsCommonHandle is true when clustered index feature is
	// enabled and the primary key is not a single integer column.
	IsCommonHandle bool `json:"is_common_handle"`

	Comment         string `json:"comment"`
	AutoIncID       int64  `json:"auto_inc_id"`
	AutoIdCache     int64  `json:"auto_id_cache"`
	AutoRandID      int64  `json:"auto_rand_id"`
	MaxColumnID     int64  `json:"max_col_id"`
	MaxIndexID      int64  `json:"max_idx_id"`
	MaxConstraintID int64  `json:"max_cst_id"`
	// UpdateTS is used to record the timestamp of updating the table's schema information.
	// These changing schema operations don't include 'truncate table' and 'rename table'.
	UpdateTS uint64 `json:"update_timestamp"`
	// OldSchemaID :
	// Because auto increment ID has schemaID as prefix,
	// We need to save original schemaID to keep autoID unchanged
	// while renaming a table from one database to another.
	// TODO: Remove it.
	// Now it only uses for compatibility with the old version that already uses this field.
	OldSchemaID int64 `json:"old_schema_id,omitempty"`

	// ShardRowIDBits specify if the implicit row ID is sharded.
	ShardRowIDBits uint64
	// MaxShardRowIDBits uses to record the max ShardRowIDBits be used so far.
	MaxShardRowIDBits uint64 `json:"max_shard_row_id_bits"`
	// AutoRandomBits is used to set the bit number to shard automatically when PKIsHandle.
	AutoRandomBits uint64 `json:"auto_random_bits"`
	// PreSplitRegions specify the pre-split region when create table.
	// The pre-split region num is 2^(PreSplitRegions-1).
	// And the PreSplitRegions should less than or equal to ShardRowIDBits.
	PreSplitRegions uint64 `json:"pre_split_regions"`
	Compression string `json:"compression"`
	// Version means the version of the table info.
	Version uint16 `json:"version"`
}
func genTables(num int) []*TableInfo {
	tbl := make([]*TableInfo,num)
	ColNum := 10
	idxNum := 10
	cols := make([]*ColumnInfo,ColNum)
	for i :=0; i< ColNum; i++ {
		cols[i] = & ColumnInfo{
			ID:                    int64(i),
			Name:                  "colssssssssssss                 ssssssss",
			Offset:                0,
			OriginDefaultValue:    nil,
			OriginDefaultValueBit: nil,
			DefaultValue:          nil,
			DefaultValueBit:       nil,
			DefaultIsExpr:         false,
			GeneratedExprString:   "bcs                                  asdf",
			GeneratedStored:       false,
			Dependences:           nil,
			Comment:               "test test commewnts   atet           asd",
			Hidden:                false,
			Version:               0,
		}
	}
	idxes := make([]*IndexInfo,idxNum)
	for i :=0; i<idxNum; i++ {
		idxes[i] = &IndexInfo{
			ID:        int64(i),
			Name:      "n       a              m                     e",
			Table:     "t          a           b           ll        e",
			Columns:   nil,
			Comment:   "c     o        m          m      e    n      t",
			Unique:    false,
			Primary:   false,
			Invisible: false,
		}
	}
	for i:=0 ;i <num; i++ {
		tbl[i] = &TableInfo{
			ID:                int64(i),
			Name:              "na                                                                 me",
			Charset:           "c             h          a       r           s             e       t",
			Collate:           "c          o               l             l            a       t     e",
			Columns:           cols,
			Indices:           idxes,
			PKIsHandle:        false,
			IsCommonHandle:    false,
			Comment:           "c        o              m             m          e          n       t",
			AutoIncID:         0,
			AutoIdCache:       0,
			AutoRandID:        0,
			MaxColumnID:       0,
			MaxIndexID:        0,
			MaxConstraintID:   0,
			UpdateTS:          0,
			OldSchemaID:       0,
			ShardRowIDBits:    0,
			MaxShardRowIDBits: 0,
			AutoRandomBits:    0,
			PreSplitRegions:   0,
			Compression:       "c          o              m   ppp           r       e        s        s",
			Version:           0,
		}
	}
	return tbl
}