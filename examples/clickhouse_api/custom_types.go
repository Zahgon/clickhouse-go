package clickhouse_api

type customStr string

func (s *customStr) Scan(src any) error { _ = "STUB: not implemented"; return nil }

func (s customStr) String() string { _ = "STUB: not implemented"; return "" }

func CustomTypes() error { _ = "STUB: not implemented"; return nil }
