package encode

import "golang.org/x/crypto/bcrypt"

// Encode 加密密码
func Encode(pwdStr string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwdStr), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	pwdHash := string(hash)
	return pwdHash, nil
}

// ComparePasswords 验证密码
func ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePwd := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		return false
	}
	return true
}
