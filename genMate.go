package json_bench

func genPTables(num int) []*PTableInfo {
	tbl := make([]*PTableInfo,num)
	ColNum := 10
	idxNum := 10
	cols := make([]*PColumnInfo,ColNum)
	for i :=0; i< ColNum; i++ {
		cols[i] = & PColumnInfo{
			ID:                    int64(i),
			Name:                  "colssssssssssss                 ssssssss",
			Offset:                0,
			OriginDefaultValue:    0,
			OriginDefaultValueBit: nil,
			DefaultValue:          0,
			DefaultValueBit:       nil,
			DefaultIsExpr:         false,
			GeneratedExprString:   "bcs                                  asdf",
			GeneratedStored:       false,
			Dependences:           0,
			Comment:               "test test commewnts   atet           asd",
			Hidden:                false,
			Version:               0,
		}
	}
	idxes := make([]*PIndexInfo,idxNum)
	for i :=0; i<idxNum; i++ {
		idxes[i] = &PIndexInfo{
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
		tbl[i] = &PTableInfo{
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