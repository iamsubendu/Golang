package utility

//making public
func GetName() string {
	return getPackageName()
}

// we are having functions in different files but go treat them 
//as a single file --- like classes