package toss

// PatternOddEven 奇偶模式
func PatternOddEven(value int) Schema {
	if value%2 == 1 {
		return SchemaPositive
	}

	return SchemaNegative
}
