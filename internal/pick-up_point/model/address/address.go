package address

type PPAddress struct {
	Region   string `valid:"required,length(3|50),matches(^[A-Z][a-z]+$)"`
	City     string `valid:"required,length(3|50),matches(^[A-Z][a-z]+$)"`
	Street   string `valid:"required,length(2|50),matches(^[A-Z][a-z]+$)"`
	HouseNum string `valid:"required,length(1|10),matches(^[a-zA-Z0-9]+$)"`
}
