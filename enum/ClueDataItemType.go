package enum

type ClueDataItemType int

const (
	ClueDataItemType_TEXT_INPUT       ClueDataItemType = 1
	ClueDataItemType_NUMBER_INPUT     ClueDataItemType = 2
	ClueDataItemType_PHONE            ClueDataItemType = 3
	ClueDataItemType_EMAIL            ClueDataItemType = 4
	ClueDataItemType_SINGLE_SELECTION ClueDataItemType = 5
	ClueDataItemType_MULTI_SELECTION  ClueDataItemType = 6
	ClueDataItemType_NAME             ClueDataItemType = 7
	ClueDataItemType_GENDER           ClueDataItemType = 8
	ClueDataItemType_CITY             ClueDataItemType = 9
	ClueDataItemType_DATE             ClueDataItemType = 10
	ClueDataItemType_SINGLE_DROPDOWN  ClueDataItemType = 11
	ClueDataItemType_MULTI_DROPDOWN   ClueDataItemType = 12
)
