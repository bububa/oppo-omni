package enum

type ClueTransformType int

const (
	ClueTransformType_SUBMIT_FORM  ClueTransformType = 1
	ClueTransformType_PHONE_CALL   ClueTransformType = 3
	ClueTransformType_CONSULT      ClueTransformType = 4
	ClueTransformType_BUY_PRODUCT  ClueTransformType = 5
	ClueTransformType_VIEW_CONTENT ClueTransformType = 6
	ClueTransformType_COPY_WECHAT  ClueTransformType = 7
	ClueTransformType_PAY          ClueTransformType = 8
	ClueTransformType_OTHERS       ClueTransformType = 9
)
